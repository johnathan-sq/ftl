package ftl

// RegisterTopic declares a topic
//
// Topics publish events, and subscriptions can listen to them.
func RegisterTopic[E any](name string) Topic[E] {
	return Topic[E]{name: name}
}

type Topic[E any] struct {
	name string
}

// Publish publishes an event to a topic
func (t Topic[E]) Publish(event E) error {
	panic("not implemented")
}

// RegisterSubscription declares a subscription to a topic
//
// Sinks can consume events from the subscription by including a "ftl:subscibe <subscription_name>" directive
func RegisterSubscription[E any](topic Topic[E], name string) Subscription[E] {
	return Subscription[E]{name: name}
}

type Subscription[E any] struct {
	name string
}