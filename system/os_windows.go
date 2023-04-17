//go:build windows

package system

import (
	"os/exec"
	"syscall"
)

func WithWinHide() Option {
	return func(c *exec.Cmd) {
		if c.SysProcAttr == nil {
			c.SysProcAttr = &syscall.SysProcAttr{
				HideWindow: true,
			}
		} else {
			c.SysProcAttr.HideWindow = true
		}
	}
}
