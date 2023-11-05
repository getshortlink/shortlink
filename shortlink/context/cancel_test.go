package context

import (
	"context"
	"syscall"
	"testing"
)

func TestWithSigtermCancel(t *testing.T) {
	t.Run("cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		ctx = WithSigtermCancel(ctx, cancel)

		_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
		<-ctx.Done()
	})
}
