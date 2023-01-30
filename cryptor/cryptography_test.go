package cryptor

import (
	"testing"
)

func TestName(t *testing.T) {
	t.Log(DefaultFastCryptography("hhh"))
}
