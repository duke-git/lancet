//go:build windows

package fileutil

import "testing"

func TestGetExeOrDllVersion(t *testing.T) {
	v, err := GetExeOrDllVersion(`C:\Windows\System32\cmd.exe`)
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
}
