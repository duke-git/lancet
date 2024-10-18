//go:build windows

package fileutil

import "testing"

func TestGetExeDllVersion(t *testing.T) {
	v, err := GetExeDllVersion(`C:\Windows\System32\cmd.exe`)
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
}
