package xerror

import (
	"errors"
	"strconv"
	"strings"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestTryUnwrap(t *testing.T) {
	assert := internal.NewAssert(t, "TestTryUnwrap")
	assert.Equal(42, TryUnwrap(strconv.Atoi("42")))
}

func TestTryUnwrapFail(t *testing.T) {
	assert := internal.NewAssert(t, "TestTryUnwrapFail")

	_, err := strconv.Atoi("4o2")
	defer func() {
		v := recover()
		assert.Equal(err.Error(), v.(*strconv.NumError).Error())
	}()

	TryUnwrap(strconv.Atoi("4o2"))
}

func TestNew(t *testing.T) {
	assert := internal.NewAssert(t, "TestNew")

	err := New("error occurs")
	assert.Equal("error occurs", err.Error())
}

func TestWrap(t *testing.T) {
	assert := internal.NewAssert(t, "TestWrap")

	err := New("wrong password")
	wrapErr := Wrap(err, "error")

	assert.Equal("error: wrong password", wrapErr.Error())
}

func TestXError_Wrap(t *testing.T) {
	assert := internal.NewAssert(t, "TestXError_Wrap")

	err1 := New("error").With("level", "high")
	err2 := err1.Wrap(errors.New("bad"))

	assert.Equal("error: bad", err2.Error())
}

func TestXError_Unwrap(t *testing.T) {
	assert := internal.NewAssert(t, "TestXError_Unwrap")

	err1 := New("error").With("level", "high")

	err2 := err1.Wrap(errors.New("bad"))

	err := err2.Unwrap()

	assert.Equal("bad", err.Error())
}

func TestXError_StackTrace(t *testing.T) {
	assert := internal.NewAssert(t, "TestXError_StackTrace")

	err := New("error")

	stacks := err.Stacks()

	assert.Equal(3, len(stacks))
	assert.Equal("github.com/duke-git/lancet/v2/xerror.TestXError_StackTrace", stacks[0].Func)
	assert.Equal(69, stacks[0].Line)
	assert.Equal(true, strings.Contains(stacks[0].File, "xerror_test.go"))
}

func TestXError_With_Id_Is_Values(t *testing.T) {
	assert := internal.NewAssert(t, "TestXError_With_Id_Is_Values")

	baseErr := New("baseError")
	err1 := New("error1").Id("e001").With("level", "high")
	err2 := New("error2").Id("e002").With("level", "low")

	err := err1.Wrap(baseErr).With("v", "1.0")

	assert.Equal(true, errors.Is(err, baseErr))
	assert.NotEqual(err, err1)
	assert.IsNotNil(err.Values()["v"])
	assert.IsNil(err1.Values()["v"])

	assert.Equal(false, errors.Is(err, err2))
}

func TestXError_Info(t *testing.T) {
	assert := internal.NewAssert(t, "TestXError_Info")

	cause := errors.New("error")
	err := Wrap(cause, "invalid username").Id("e001").With("level", "high")

	errInfo := err.Info()
	assert.Equal("invalid username", errInfo.Message)
	assert.Equal("e001", errInfo.Id)
	assert.Equal(cause, errInfo.Cause)
	assert.Equal("high", errInfo.Values["level"])
}
