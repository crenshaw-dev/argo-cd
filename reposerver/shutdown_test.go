package reposerver

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/argoproj/argo-cd/v3/reposerver/repository"
)

func TestCancelOnShutdownUnary_propagatesCause(t *testing.T) {
	shutdownCtx, shutdownCancel := context.WithCancelCause(context.Background())
	entered := make(chan context.Context, 1)
	errCh := make(chan error, 1)
	go func() {
		_, err := cancelOnShutdownUnary(shutdownCtx)(context.Background(), nil, nil, func(ctx context.Context, _ any) (any, error) {
			entered <- ctx
			<-ctx.Done()
			return nil, ctx.Err()
		})
		errCh <- err
	}()

	ctx := <-entered
	require.NoError(t, ctx.Err())
	shutdownCancel(repository.ErrShuttingDown)

	select {
	case err := <-errCh:
		require.ErrorIs(t, err, context.Canceled)
	case <-time.After(time.Second):
		t.Fatal("handler was not cancelled")
	}
	require.ErrorIs(t, context.Cause(ctx), repository.ErrShuttingDown)
}
