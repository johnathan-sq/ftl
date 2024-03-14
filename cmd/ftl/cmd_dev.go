package main

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/TBD54566975/ftl/backend/protos/xyz/block/ftl/v1/ftlv1connect"
	"github.com/TBD54566975/ftl/buildengine"
	"github.com/TBD54566975/ftl/internal/rpc"
)

type devCmd struct {
	Parallelism int           `short:"j" help:"Number of modules to build in parallel." default:"${numcpu}"`
	Dirs        []string      `arg:"" help:"Base directories containing modules." type:"existingdir" required:""`
	Watch       time.Duration `help:"Watch template directory at this frequency and regenerate on change." default:"500ms"`
	NoServe     bool          `help:"Do not start the FTL server." default:"false"`
	ServeCmd    serveCmd      `embed:""`
}

func (d *devCmd) Run(ctx context.Context) error {
	client := rpc.ClientFromContext[ftlv1connect.ControllerServiceClient](ctx)

	g, ctx := errgroup.WithContext(ctx)

	if !d.NoServe {
		if d.ServeCmd.Stop {
			err := d.ServeCmd.Run(ctx)
			if err != nil {
				return err
			}
			d.ServeCmd.Stop = false
		}
		if d.ServeCmd.isRunning(ctx, client) {
			return errors.New("FTL is already running")
		}

		g.Go(func() error {
			return d.ServeCmd.Run(ctx)
		})
	}

	err := d.ServeCmd.pollControllerOnine(ctx, client)
	if err != nil {
		return err
	}

	g.Go(func() error {
		engine, err := buildengine.New(ctx, client, d.Dirs, buildengine.Parallelism(d.Parallelism))
		if err != nil {
			return err
		}
		return engine.Dev(ctx, d.Watch)
	})

	return g.Wait()
}
