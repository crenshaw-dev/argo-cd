package commands

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func closedShutdown() <-chan struct{} {
	ch := make(chan struct{})
	close(ch)
	return ch
}

func TestDrainThenCancel_cancelsAfterDrain(t *testing.T) {
	handlerCtx, handlerCancel := context.WithCancel(context.Background())
	t.Cleanup(handlerCancel)
	release := make(chan struct{})

	started := time.Now()
	var cancelledAt time.Time
	done := make(chan struct{})
	go func() {
		drainThenCancel(closedShutdown(), func() { <-release }, func() {
			cancelledAt = time.Now()
			close(release)
			handlerCancel()
		}, 150*time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("shutdown waited for the handler instead of cancelling it after the drain timeout")
	}
	require.Error(t, handlerCtx.Err())
	require.GreaterOrEqual(t, cancelledAt.Sub(started), 150*time.Millisecond)
}

func TestDrainThenCancel_handlerFinishesWithinDrain(t *testing.T) {
	cancelled := false
	done := make(chan struct{})
	go func() {
		drainThenCancel(closedShutdown(), func() {}, func() { cancelled = true }, time.Second)
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("shutdown did not return when the handler finished")
	}
	require.False(t, cancelled, "a handler that finished during the drain was cancelled")
}
