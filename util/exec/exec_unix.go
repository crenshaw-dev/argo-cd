//go:build unix

package exec

import (
	"os/exec"
	"syscall"
)

func setProcessGroup(cmd *exec.Cmd) {
	if processGroupIsolationDisabled.Load() || cmd.SysProcAttr != nil {
		return
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func signalCmd(cmd *exec.Cmd, sig syscall.Signal) error {
	if cmd.Process == nil {
		return nil
	}
	if processGroupIsolationDisabled.Load() || cmd.SysProcAttr == nil || !cmd.SysProcAttr.Setpgid {
		return cmd.Process.Signal(sig)
	}
	if err := syscall.Kill(-cmd.Process.Pid, sig); err != nil {
		return cmd.Process.Signal(sig)
	}
	return nil
}
