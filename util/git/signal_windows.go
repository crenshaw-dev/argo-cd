//go:build windows

package git

import "syscall"

func signalProcessGroup(int, syscall.Signal) error {
	return syscall.EINVAL
}
