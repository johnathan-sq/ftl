package dal

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/types/optional"

	"github.com/TBD54566975/ftl/backend/controller/observability"
	"github.com/TBD54566975/ftl/backend/controller/sql"
	"github.com/TBD54566975/ftl/backend/controller/sql/sqltypes"
	dalerrs "github.com/TBD54566975/ftl/backend/dal"
	"github.com/TBD54566975/ftl/backend/schema"
	"github.com/TBD54566975/ftl/internal/encryption"
	"github.com/TBD54566975/ftl/internal/log"
	"github.com/TBD54566975/ftl/internal/model"
	"github.com/TBD54566975/ftl/internal/rpc"
	"github.com/TBD54566975/ftl/internal/slices"
)

func (d *DAL) PublishEventForTopic(ctx context.Context, module, topic, caller string, payload []byte) error {
	encryptedPayload, err := d.encrypt(encryption.AsyncSubKey, payload)
	if err != nil {
		return fmt.Errorf("failed to encrypt payload: %w", err)
	}

	// Store the current otel context with the event
	jsonOc, err := observability.RetrieveTraceContextFromContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve trace context: %w", err)
	}

	// Store the request key that initiated this publish, this will eventually
	// become the parent request key of the subscriber call
	requestKey := ""
	if rk, err := rpc.RequestKeyFromContext(ctx); err == nil {
		if rk, ok := rk.Get(); ok {
			requestKey = rk.String()
		}
	} else {
		return fmt.Errorf("failed to get request key: %w", err)
	}

	err = d.db.PublishEventForTopic(ctx, sql.PublishEventForTopicParams{
		Key:          model.NewTopicEventKey(module, topic),
		Module:       module,
		Topic:        topic,
		Caller:       caller,
		Payload:      encryptedPayload,
		RequestKey:   requestKey,
		TraceContext: jsonOc,
	})
	observability.PubSub.Published(ctx, module, topic, caller, err)
	if err != nil {
		return dalerrs.TranslatePGError(err)
	}
	return nil
}

func (d *DAL) GetSubscriptionsNeedingUpdate(ctx context.Context) ([]model.Subscription, error) {
	rows, err := d.db.GetSubscriptionsNeedingUpdate(ctx)
	if err != nil {
		return nil, dalerrs.TranslatePGError(err)
	}
	return slices.Map(rows, func(row sql.GetSubscriptionsNeedingUpdateRow) model.Subscription {
		return model.Subscription{
			Name:   row.Name,
			Key:    row.Key,
			Topic:  row.Topic,
			Cursor: row.Cursor,
		}
	}), nil
}

func (d *DAL) ProgressSubscriptions(ctx context.Context, eventConsumptionDelay time.Duration) (count int, err error) {
	tx, err := d.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.CommitOrRollback(ctx, &err)

	logger := log.FromContext(ctx)

	// get subscriptions needing update
	// also gets a lock on the subscription, and skips any subscriptions locked by others
	subs, err := tx.db.GetSubscriptionsNeedingUpdate(ctx)
	if err != nil {
		return 0, fmt.Errorf("could not get subscriptions to progress: %w", dalerrs.TranslatePGError(err))
	}

	successful := 0
	for _, subscription := range subs {
		nextCursor, err := tx.db.GetNextEventForSubscription(ctx, sqltypes.Duration(eventConsumptionDelay), subscription.Topic, subscription.Cursor)
		if err != nil {
			observability.PubSub.PropagationFailed(ctx, "GetNextEventForSubscription", subscription.Topic.Payload, nextCursor.Caller, subscriptionRef(subscription), optional.None[schema.RefKey]())
			return 0, fmt.Errorf("failed to get next cursor: %w", dalerrs.TranslatePGError(err))
		}
		nextCursorKey, ok := nextCursor.Event.Get()
		if !ok {
			observability.PubSub.PropagationFailed(ctx, "GetNextEventForSubscription-->Event.Get", subscription.Topic.Payload, nextCursor.Caller, subscriptionRef(subscription), optional.None[schema.RefKey]())
			return 0, fmt.Errorf("could not find event to progress subscription: %w", dalerrs.TranslatePGError(err))
		}
		if !nextCursor.Ready {
			logger.Tracef("Skipping subscription %s because event is too new", subscription.Key)
			continue
		}

		subscriber, err := tx.db.GetRandomSubscriber(ctx, subscription.Key)
		if err != nil {
			logger.Tracef("no subscriber for subscription %s", subscription.Key)
			continue
		}

		err = tx.db.BeginConsumingTopicEvent(ctx, subscription.Key, nextCursorKey)
		if err != nil {
			observability.PubSub.PropagationFailed(ctx, "BeginConsumingTopicEvent", subscription.Topic.Payload, nextCursor.Caller, subscriptionRef(subscription), optional.Some(subscriber.Sink))
			return 0, fmt.Errorf("failed to progress subscription: %w", dalerrs.TranslatePGError(err))
		}

		origin := AsyncOriginPubSub{
			Subscription: schema.RefKey{
				Module: subscription.Key.Payload.Module,
				Name:   subscription.Key.Payload.Name,
			},
		}

		_, err = tx.db.CreateAsyncCall(ctx, sql.CreateAsyncCallParams{
			Verb:              subscriber.Sink,
			Origin:            origin.String(),
			Request:           nextCursor.Payload, // already encrypted
			RemainingAttempts: subscriber.RetryAttempts,
			Backoff:           subscriber.Backoff,
			MaxBackoff:        subscriber.MaxBackoff,
			ParentRequestKey:  nextCursor.RequestKey,
			TraceContext:      nextCursor.TraceContext.RawMessage,
			CatchVerb:         subscriber.CatchVerb,
		})
		observability.AsyncCalls.Created(ctx, subscriber.Sink, subscriber.CatchVerb, origin.String(), int64(subscriber.RetryAttempts), err)
		if err != nil {
			observability.PubSub.PropagationFailed(ctx, "CreateAsyncCall", subscription.Topic.Payload, nextCursor.Caller, subscriptionRef(subscription), optional.Some(subscriber.Sink))
			return 0, fmt.Errorf("failed to schedule async task for subscription: %w", dalerrs.TranslatePGError(err))
		}

		observability.PubSub.SinkCalled(ctx, subscription.Topic.Payload, nextCursor.Caller, subscriptionRef(subscription), subscriber.Sink)
		successful++
	}

	if successful > 0 {
		// If no async calls were successfully created, then there is no need to
		// potentially increment the queue depth gauge.
		queueDepth, err := tx.db.AsyncCallQueueDepth(ctx)
		if err == nil {
			// Don't error out of progressing subscriptions just over a queue depth
			// retrieval error because this is only used for an observability gauge.
			observability.AsyncCalls.RecordQueueDepth(ctx, queueDepth)
		}
	}

	return successful, nil
}

func subscriptionRef(subscription sql.GetSubscriptionsNeedingUpdateRow) schema.RefKey {
	return schema.RefKey{Module: subscription.Key.Payload.Module, Name: subscription.Name}
}

func (d *DAL) CompleteEventForSubscription(ctx context.Context, module, name string) error {
	err := d.db.CompleteEventForSubscription(ctx, name, module)
	if err != nil {
		return fmt.Errorf("failed to complete event for subscription: %w", dalerrs.TranslatePGError(err))
	}
	return nil
}

// ResetSubscription resets the subscription cursor to the topic's head.
func (d *DAL) ResetSubscription(ctx context.Context, module, name string) (err error) {
	tx, err := d.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not start transaction: %w", err)
	}
	defer tx.CommitOrRollback(ctx, &err)

	subscription, err := tx.GetSubscription(ctx, name, module)
	if err != nil {
		if dalerrs.IsNotFound(err) {
			return fmt.Errorf("subscription %s.%s not found", module, name)
		}
		return fmt.Errorf("could not fetch subscription: %w", dalerrs.TranslatePGError(err))
	}

	topic, err := tx.GetTopic(ctx, subscription.TopicID)
	if err != nil {
		return fmt.Errorf("could not fetch topic: %w", dalerrs.TranslatePGError(err))
	}

	headEventID, ok := topic.Head.Get()
	if !ok {
		return fmt.Errorf("no events published to topic %s", topic.Name)
	}

	headEvent, err := tx.GetTopicEvent(ctx, headEventID)
	if err != nil {
		return fmt.Errorf("could not fetch topic head: %w", dalerrs.TranslatePGError(err))
	}

	err = tx.SetSubscriptionCursor(ctx, subscription.Key, headEvent.Key)
	if err != nil {
		return fmt.Errorf("failed to reset subscription: %w", dalerrs.TranslatePGError(err))
	}

	return nil
}

func (d *DAL) createSubscriptions(ctx context.Context, tx *sql.Tx, key model.DeploymentKey, module *schema.Module) error {
	logger := log.FromContext(ctx)

	for _, decl := range module.Decls {
		s, ok := decl.(*schema.Subscription)
		if !ok {
			continue
		}
		if !hasSubscribers(s, module.Decls) {
			// Ignore subscriptions without subscribers
			// This ensures that controllers don't endlessly try to progress subscriptions without subscribers
			// https://github.com/TBD54566975/ftl/issues/1685
			//
			// It does mean that a subscription will reset to the topic's head if all subscribers are removed and then later re-added
			logger.Debugf("Skipping upsert of subscription %s for %s due to lack of subscribers", s.Name, key)
			continue
		}
		subscriptionKey := model.NewSubscriptionKey(module.Name, s.Name)
		result, err := tx.UpsertSubscription(ctx, sql.UpsertSubscriptionParams{
			Key:         subscriptionKey,
			Module:      module.Name,
			Deployment:  key,
			TopicModule: s.Topic.Module,
			TopicName:   s.Topic.Name,
			Name:        s.Name,
		})
		if err != nil {
			return fmt.Errorf("could not insert subscription: %w", dalerrs.TranslatePGError(err))
		}
		if result.Inserted {
			logger.Debugf("Inserted subscription %s for %s", subscriptionKey, key)
		} else {
			logger.Debugf("Updated subscription %s to %s", subscriptionKey, key)
		}
	}
	return nil
}

func hasSubscribers(subscription *schema.Subscription, decls []schema.Decl) bool {
	for _, d := range decls {
		verb, ok := d.(*schema.Verb)
		if !ok {
			continue
		}
		for _, md := range verb.Metadata {
			subscriber, ok := md.(*schema.MetadataSubscriber)
			if !ok {
				continue
			}
			if subscriber.Name == subscription.Name {
				return true
			}
		}
	}
	return false
}

func (d *DAL) createSubscribers(ctx context.Context, tx *sql.Tx, key model.DeploymentKey, module *schema.Module) error {
	logger := log.FromContext(ctx)
	for _, decl := range module.Decls {
		v, ok := decl.(*schema.Verb)
		if !ok {
			continue
		}
		for _, md := range v.Metadata {
			s, ok := md.(*schema.MetadataSubscriber)
			if !ok {
				continue
			}
			sinkRef := schema.RefKey{
				Module: module.Name,
				Name:   v.Name,
			}
			retryParams := schema.RetryParams{}
			var err error
			if retryMd, ok := slices.FindVariant[*schema.MetadataRetry](v.Metadata); ok {
				retryParams, err = retryMd.RetryParams()
				if err != nil {
					return fmt.Errorf("could not parse retry parameters for %q: %w", v.Name, err)
				}
			}
			subscriberKey := model.NewSubscriberKey(module.Name, s.Name, v.Name)
			err = tx.InsertSubscriber(ctx, sql.InsertSubscriberParams{
				Key:              subscriberKey,
				Module:           module.Name,
				SubscriptionName: s.Name,
				Deployment:       key,
				Sink:             sinkRef,
				RetryAttempts:    int32(retryParams.Count),
				Backoff:          sqltypes.Duration(retryParams.MinBackoff),
				MaxBackoff:       sqltypes.Duration(retryParams.MaxBackoff),
				CatchVerb:        retryParams.Catch,
			})
			if err != nil {
				return fmt.Errorf("could not insert subscriber: %w", dalerrs.TranslatePGError(err))
			}
			logger.Debugf("Inserted subscriber %s for %s", subscriberKey, key)
		}
	}
	return nil
}

func (d *DAL) removeSubscriptionsAndSubscribers(ctx context.Context, tx *sql.Tx, key model.DeploymentKey) error {
	logger := log.FromContext(ctx)

	subscribers, err := tx.DeleteSubscribers(ctx, key)
	if err != nil {
		return fmt.Errorf("could not delete old subscribers: %w", dalerrs.TranslatePGError(err))
	}
	if len(subscribers) > 0 {
		logger.Debugf("Deleted subscribers for %s: %s", key, strings.Join(slices.Map(subscribers, func(key model.SubscriberKey) string { return key.String() }), ", "))
	}

	subscriptions, err := tx.DeleteSubscriptions(ctx, key)
	if err != nil {
		return fmt.Errorf("could not delete old subscriptions: %w", dalerrs.TranslatePGError(err))
	}
	if len(subscriptions) > 0 {
		logger.Debugf("Deleted subscriptions for %s: %s", key, strings.Join(slices.Map(subscriptions, func(key model.SubscriptionKey) string { return key.String() }), ", "))
	}

	return nil
}
