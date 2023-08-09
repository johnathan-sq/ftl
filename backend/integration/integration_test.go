//go:build integration

package integration

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/TBD54566975/ftl/backend/common/exec"
	"github.com/TBD54566975/ftl/backend/common/log"
	"github.com/TBD54566975/ftl/backend/common/rpc"
	"github.com/TBD54566975/ftl/protos/xyz/block/ftl/v1/ftlv1connect"
)

var binaries = []string{"ftl-controller", "ftl-runner"}

type assertion func(client ftlv1connect.ControllerServiceClient)
type asserts []assertion

type fixture interface {
	up() error
	down() error
}
type fixtures []fixture

func TestIntegration(t *testing.T) {
	binDir := t.TempDir()
	logger := log.Configure(os.Stderr, log.Config{Level: log.Warn})
	ctx := log.ContextWithLogger(context.Background(), logger)
	for _, binary := range binaries {
		t.Logf("Building %s", binary)
		err := exec.Command(ctx, "..", "go", "build", "-trimpath", "-ldflags=-s -w -buildid=", "-o", filepath.Join(binDir, binary), "./cmd/"+binary).Run()
		assert.NoError(t, err)
	}
	tests := []struct {
		name         string
		extraRunners int
		fixtures     fixtures
		assertions   asserts
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run(t, ctx, "ftl-controller")
			run(t, ctx, "ftl-runner", "--language=go")
			client := rpc.Dial(ftlv1connect.NewControllerServiceClient, "localhost:8893", log.Warn)
			for i := 0; i < tt.extraRunners; i++ {
				run(t, ctx, "ftl-runner", "--language=go", "--endpoint=http://localhost:"+strconv.Itoa(8893+i))
			}
			for _, fixture := range tt.fixtures {
				err := fixture.up()
				assert.NoError(t, err, "fixture failed")
			}
			for _, assertion := range tt.assertions {
				assertion(client)
			}
			for _, fixture := range tt.fixtures {
				err := fixture.down()
				assert.NoError(t, err, "fixture failed")
			}
		})
	}
}

func run(t testing.TB, ctx context.Context, args ...string) {
	t.Helper()
	binDir := t.TempDir()
	cmd := exec.Command(ctx, "..", filepath.Join(binDir, args[0]), args...)
	err := cmd.Start()
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := cmd.Kill(syscall.SIGTERM)
		assert.NoError(t, err)
	})
}