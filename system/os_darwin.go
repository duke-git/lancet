//go:build darwin

package system

import (
	"os/exec"
)

func WithForeground() Option {
	return func(c *exec.Cmd) {

	}
}

func WithWinHide() Option {
	return func(c *exec.Cmd) {

	}
}
