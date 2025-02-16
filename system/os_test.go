package system

import (
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestOsDetection(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestOsJudgment")

	osType, _, _ := ExecCommand("echo $OSTYPE")
	if strings.Contains(osType, "linux") {
		assert.Equal(true, IsLinux())
	}
	if strings.Contains(osType, "darwin") {
		assert.Equal(true, IsMac())
	}
}

func TestOsEnvOperation(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestOsEnvOperation")

	envNotExist := GetOsEnv("foo")
	assert.Equal("", envNotExist)

	err := SetOsEnv("foo", "foo_value")
	assert.IsNil(err)

	envExist := GetOsEnv("foo")
	assert.Equal("foo_value", envExist)

	assert.Equal(true, CompareOsEnv("foo", "foo_value"))
	assert.Equal(false, CompareOsEnv("foo", "abc"))
	assert.Equal(false, CompareOsEnv("abc", "abc"))
	assert.Equal(false, CompareOsEnv("abc", "abc"))

	err = RemoveOsEnv("foo")
	if err != nil {
		t.Fail()
	}
	assert.Equal(false, CompareOsEnv("foo", "foo_value"))
}

func TestExecCommand(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestExecCommand")

	// linux or mac
	stdout, stderr, err := ExecCommand("ls", func(cmd *exec.Cmd) {
		cmd.Dir = "/"
	})
	t.Log("std out: ", stdout)
	t.Log("std err: ", stderr)
	assert.Equal("", stderr)
	assert.IsNil(err)

	// windows
	stdout, stderr, err = ExecCommand("dir")
	t.Log("std out: ", stdout)
	t.Log("std err: ", stderr)
	if IsWindows() {
		assert.IsNil(err)
	}

	// error command
	stdout, stderr, err = ExecCommand("abc")
	t.Log("std out: ", stdout)
	t.Log("std err: ", stderr)
	assert.IsNotNil(err)
}

func TestGetOsBits(t *testing.T) {
	t.Parallel()

	osBits := GetOsBits()
	switch osBits {
	case 32, 64:
		t.Logf("os is %d", osBits)
	default:
		t.Error("os is not 32 or 64 bits")
	}
}

func TestStartProcess(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestStartProcess")

	pid, err := StartProcess("ls", "-a")

	assert.IsNil(err)
	assert.Equal(true, pid > 0)
}

func TestKillProcess(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestKillProcess")

	pid, err := StartProcess("ls")
	assert.IsNil(err)
	assert.Equal(true, pid > 0)

	err = KillProcess(pid)
	assert.IsNil(err)
}

func TestStopProcess(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestStopProcess")

	pid, err := StartProcess("ls")
	assert.IsNil(err)
	assert.Equal(true, pid > 0)

	err = StopProcess(pid)
	assert.IsNil(err)
}

func TestGetProcessInfo(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestGetProcessInfo")

	pid, err := StartProcess("ls", "-a")
	assert.IsNil(err)

	time.Sleep(1 * time.Second)

	processInfo, err := GetProcessInfo(pid)
	assert.IsNil(err)

	t.Log(processInfo)
}
