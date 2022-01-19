package system

import (
	"strings"
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestOsDetection(t *testing.T) {
	assert := internal.NewAssert(t, "TestOsJudgment")

	osType, _, _ := ExecCommand("echo $OSTYPE")
	if strings.Index(osType, "linux") != -1 {
		assert.Equal(true, IsLinux())
	}
	if strings.Index(osType, "darwin") != -1 {
		assert.Equal(true, IsMac())
	}
}

func TestOsEnvOperation(t *testing.T) {
	assert := internal.NewAssert(t, "TestOsEnvOperation")

	envNotExist := GetOsEnv("foo")
	assert.Equal("", envNotExist)

	SetOsEnv("foo", "foo_value")
	envExist := GetOsEnv("foo")
	assert.Equal("foo_value", envExist)

	assert.Equal(true, CompareOsEnv("foo", "foo_value"))
	assert.Equal(false, CompareOsEnv("foo", "abc"))
	assert.Equal(false, CompareOsEnv("abc", "abc"))
	assert.Equal(false, CompareOsEnv("abc", "abc"))

	err := RemoveOsEnv("foo")
	if err != nil {
		t.Fail()
	}
	assert.Equal(false, CompareOsEnv("foo", "foo_value"))
}

func TestExecCommand(t *testing.T) {
	assert := internal.NewAssert(t, "TestExecCommand")

	out, errout, err := ExecCommand("ls")
	t.Log("std out: ", out)
	t.Log("std err: ", errout)
	assert.IsNil(err)

	out, errout, err = ExecCommand("abc")
	t.Log("std out: ", out)
	t.Log("std err: ", errout)
	if err != nil {
		t.Logf("error: %v\n", err)
	}

	assert.IsNotNil(err)
}
