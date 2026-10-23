package exec

import (
	"context"
	"os/exec"
	"syscall"
	"time"
)

// Policy is how a child process is timed out and torn down.
// Zero values are not defaults: use DefaultPolicy.
type Policy struct {
	// Signal is the first signal sent when the timeout or the context fires.
	Signal syscall.Signal
	// GracePeriod is how long the command has to exit after Signal before SIGKILL.
	// It is also the bound on waiting for the output pipes to close.
	GracePeriod time.Duration
	// Timeout is how long the command may run. Zero means no exec timeout.
	Timeout time.Duration
	// Isolate starts the command in its own process group so a signal reaches
	// its children. The argocd CLI turns this off.
	Isolate bool
}

// DefaultPolicy matches the historical RunCommandExt defaults: SIGKILL on
// timeout, the fatal timeout as the pipe bound, and process-group isolation.
func DefaultPolicy() Policy {
	return Policy{
		Signal:      syscall.SIGKILL,
		GracePeriod: fatalTimeout,
		Timeout:     timeout,
		Isolate:     true,
	}
}

// CommandWithPolicy builds a command the same way exec.CommandContext does,
// and records whether it should run in its own process group.
func CommandWithPolicy(ctx context.Context, p Policy, name string, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, args...)
	applyIsolation(cmd, p)
	cmd.Cancel = func() error {
		LogTermination(ctx, cmd, TerminationCause(ctx))
		return Signal(cmd, p.Signal)
	}
	if p.GracePeriod > 0 {
		cmd.WaitDelay = p.GracePeriod
	}
	return cmd
}
