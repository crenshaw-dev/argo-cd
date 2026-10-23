package commands

import (
	"os"
	"time"
)

// defaultPodTerminationGrace is the Kubernetes default terminationGracePeriodSeconds.
// argocd-repo-server does not set its own.
const defaultPodTerminationGrace = 30 * time.Second

// defaultShutdownDrainTimeout is the pod grace period minus ARGOCD_EXEC_FATAL_TIMEOUT.
// Cancelling then still leaves the fatal timeout for a command to exit before the
// kubelet sends SIGKILL.
func defaultShutdownDrainTimeout() time.Duration {
	fatal := 10 * time.Second
	if raw := os.Getenv("ARGOCD_EXEC_FATAL_TIMEOUT"); raw != "" {
		if parsed, err := time.ParseDuration(raw); err == nil {
			fatal = parsed
		}
	}
	if fatal >= defaultPodTerminationGrace {
		return 0
	}
	return defaultPodTerminationGrace - fatal
}

// drainThenCancel lets in-flight work finish for up to drain after shutdown is
// closed. When the deadline passes, cancel is called so those requests stop,
// and then gracefulStop is waited on until it returns.
func drainThenCancel(shutdown <-chan struct{}, gracefulStop func(), cancel func(), drain time.Duration) {
	<-shutdown
	stopped := make(chan struct{})
	go func() {
		gracefulStop()
		close(stopped)
	}()
	timer := time.NewTimer(drain)
	defer timer.Stop()
	select {
	case <-stopped:
	case <-timer.C:
		cancel()
		<-stopped
	}
}
