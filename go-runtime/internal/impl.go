package internal

import (
	"context"
	"fmt"
	"reflect"

	"connectrpc.com/connect"

	ftlv1 "github.com/TBD54566975/ftl/backend/protos/xyz/block/ftl/v1"
	"github.com/TBD54566975/ftl/backend/protos/xyz/block/ftl/v1/ftlv1connect"
	schemapb "github.com/TBD54566975/ftl/backend/protos/xyz/block/ftl/v1/schema"
	"github.com/TBD54566975/ftl/backend/schema"
	"github.com/TBD54566975/ftl/go-runtime/encoding"
	"github.com/TBD54566975/ftl/go-runtime/ftl/reflection"
	"github.com/TBD54566975/ftl/internal/rpc"
)

// RealFTL is the real implementation of the [internal.FTL] interface using the Controller.
type RealFTL struct{}

// New creates a new [RealFTL]
func New() *RealFTL { return &RealFTL{} }

var _ FTL = &RealFTL{}

func (r *RealFTL) FSMSend(ctx context.Context, fsm, instance string, event any) error {
	client := rpc.ClientFromContext[ftlv1connect.VerbServiceClient](ctx)
	body, err := encoding.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	_, err = client.SendFSMEvent(ctx, connect.NewRequest(&ftlv1.SendFSMEventRequest{
		Fsm:      &schemapb.Ref{Module: reflection.Module(), Name: fsm},
		Instance: instance,
		Event:    schema.TypeToProto(reflection.ReflectTypeToSchemaType(reflect.TypeOf(event))),
		Body:     body,
	}))
	if err != nil {
		return fmt.Errorf("failed to send event: %w", err)
	}
	return nil
}