//go:build darwin

package system

import (
	"os/exec"
)

func WithForeground() Option {
	return func(c *exec.Cmd) {
		// if c.SysProcAttr == nil {
		// 	c.SysProcAttr = &syscall.SysProcAttr{
		// 		Foreground: true,
		// 	}
		// } else {
		// 	c.SysProcAttr.Foreground = true
		// }
	}
}

func WithWinHide() Option {
	return func(c *exec.Cmd) {

	}
}
