package context

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// WithSigtermCancel returns a context that is canceled when the process
// receives a SIGTERM or SIGINT signal.
func WithSigtermCancel(ctx context.Context, onCancel func()) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-interruptCh
		onCancel()
		cancel()
	}()
	return ctx
}
