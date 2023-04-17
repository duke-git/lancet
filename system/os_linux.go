package system

import (
	"os/exec"
	"syscall"
)

func WithForeground() Option {
	return func(c *exec.Cmd) {
		if c.SysProcAttr == nil {
			c.SysProcAttr = &syscall.SysProcAttr{
				Foreground: true,
			}
		} else {
			c.SysProcAttr.Foreground = true
		}
	}
}
