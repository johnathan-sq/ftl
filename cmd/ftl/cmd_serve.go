package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/alecthomas/errors"
	"github.com/alecthomas/kong"
	"golang.org/x/sync/errgroup"

	"github.com/TBD54566975/ftl/backend/common/bind"
	"github.com/TBD54566975/ftl/backend/common/exec"
	"github.com/TBD54566975/ftl/backend/common/log"
	"github.com/TBD54566975/ftl/backend/controller"
	"github.com/TBD54566975/ftl/backend/controller/scaling/localscaling"
	"github.com/TBD54566975/ftl/backend/controller/sql/databasetesting"
)

type serveCmd struct {
	Bind        *url.URL `help:"Starting endpoint to bind to and advertise to. Each controller and runner will increment the port by 1" default:"http://localhost:8892"`
	DBPort      int      `help:"Port to use for the database." default:"5433"`
	Recreate    bool     `help:"Recreate the database even if it already exists." default:"false"`
	Controllers int      `short:"c" help:"Number of controllers to start." default:"1"`
	Runners     int      `short:"r" help:"Number of runners to start." default:"0"`
}

const ftlContainerName = "ftl-db"

func (s *serveCmd) Run(ctx context.Context) error {
	logger := log.FromContext(ctx)

	dsn, err := s.setupDB(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	logger.Infof("Starting %d controller(s) and %d runner(s)", s.Controllers, s.Runners)

	wg, ctx := errgroup.WithContext(ctx)

	bindAllocator, err := bind.NewBindAllocator(s.Bind)
	if err != nil {
		return errors.WithStack(err)
	}

	controllerAddresses := make([]*url.URL, 0, s.Controllers)
	for i := 0; i < s.Controllers; i++ {
		controllerAddresses = append(controllerAddresses, bindAllocator.Next())
	}

	runnerScaling, err := localscaling.NewLocalScaling(bindAllocator, controllerAddresses)
	if err != nil {
		return errors.WithStack(err)
	}

	for i := 0; i < s.Controllers; i++ {
		i := i
		config := controller.Config{
			Bind: controllerAddresses[i],
			DSN:  dsn,
		}
		if err := kong.ApplyDefaults(&config); err != nil {
			return errors.WithStack(err)
		}

		scope := fmt.Sprintf("controller%d", i)
		controllerCtx := log.ContextWithLogger(ctx, logger.Scope(scope))

		wg.Go(func() error {
			return errors.Wrapf(controller.Start(controllerCtx, config, runnerScaling), "controller%d failed", i)
		})
	}

	err = runnerScaling.SetReplicas(ctx, s.Runners, nil)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := wg.Wait(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s *serveCmd) setupDB(ctx context.Context) (string, error) {
	logger := log.FromContext(ctx)

	nameFlag := fmt.Sprintf("name=^/%s$", ftlContainerName)
	output, err := exec.Capture(ctx, ".", "docker", "ps", "-a", "--filter", nameFlag, "--format", "{{.Names}}")
	if err != nil {
		logger.Errorf(err, "%s", output)
		return "", errors.WithStack(err)
	}

	recreate := s.Recreate
	port := ""

	if len(output) == 0 {
		logger.Infof("Creating docker container '%s' for postgres db", ftlContainerName)

		// check if port s.DBPort is already in use
		_, err := exec.Capture(ctx, ".", "sh", "-c", fmt.Sprintf("lsof -i:%d", s.DBPort))
		if err == nil {
			return "", errors.Errorf("port %d is already in use", s.DBPort)
		}

		err = exec.Command(ctx, logger.GetLevel(), "./", "docker", "run",
			"-d", // run detached so we can follow with other commands
			"--name", ftlContainerName,
			"--user", "postgres",
			"--restart", "always",
			"-e", "POSTGRES_PASSWORD=secret",
			"-p", fmt.Sprintf("%d:5432", s.DBPort),
			"--health-cmd=pg_isready",
			"--health-interval=1s",
			"--health-timeout=60s",
			"--health-retries=60",
			"--health-start-period=80s",
			"postgres:latest", "postgres",
		).Run()

		if err != nil {
			return "", errors.WithStack(err)
		}

		err = pollContainerHealth(ctx, ftlContainerName, 10*time.Second)
		if err != nil {
			return "", err
		}

		recreate = true
	} else {
		// Grab the port from the existing container
		cmdStr := fmt.Sprintf("docker port %s 5432/tcp | grep -v '\\[::\\]' | awk -F: '{print $NF}'", ftlContainerName)
		portOutput, err := exec.Capture(ctx, ".", "sh", "-c", cmdStr)
		if err != nil {
			logger.Errorf(err, "%s", portOutput)
			return "", errors.WithStack(err)
		}

		port = strings.TrimSpace(string(portOutput))
		logger.Infof("Using docker container '%s' for postgres db", ftlContainerName)
	}

	dsn := fmt.Sprintf("postgres://postgres:secret@localhost:%s/%s?sslmode=disable", port, ftlContainerName)
	logger.Infof("Postgres DSN: %s", dsn)

	_, err = databasetesting.CreateForDevel(ctx, dsn, recreate)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return dsn, nil
}

func pollContainerHealth(ctx context.Context, containerName string, timeout time.Duration) error {
	logger := log.FromContext(ctx)
	logger.Infof("Waiting for %s to be healthy", containerName)

	pollCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		select {
		case <-pollCtx.Done():
			return errors.New("timed out waiting for container to be healthy")
		case <-time.After(1 * time.Millisecond):
			output, err := exec.Capture(pollCtx, ".", "docker", "inspect", "--format", "{{.State.Health.Status}}", containerName)
			if err != nil {
				return errors.WithStack(err)
			}

			status := strings.TrimSpace(string(output))
			if status == "healthy" {
				return nil
			}
		}
	}
}
