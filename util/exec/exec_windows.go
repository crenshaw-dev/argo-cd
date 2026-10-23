//go:build windows

package exec

import (
	"os/exec"
	"syscall"
)

func applyIsolation(_ *exec.Cmd, _ Policy) {}

func setProcessGroup(_ *exec.Cmd) {}

func signalCmd(cmd *exec.Cmd, sig syscall.Signal) error {
	if cmd.Process == nil {
		return nil
	}
	return cmd.Process.Signal(sig)
}
