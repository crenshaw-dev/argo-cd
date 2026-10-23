//go:build windows

package exec

import (
	"os/exec"
	"syscall"
)

func setProcessGroup(_ *exec.Cmd) {}

func signalCmd(cmd *exec.Cmd, sig syscall.Signal) error {
	if cmd.Process == nil {
		return nil
	}
	return cmd.Process.Signal(sig)
}
