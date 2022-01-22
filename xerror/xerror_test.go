package xerror

import (
	"github.com/duke-git/lancet/internal"
	"strconv"
	"testing"
)

func TestUnwrap(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnwrap")
	assert.Equal(42, Unwrap(strconv.Atoi("42")))
}

func TestUnwrapFail(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnwrapFail")

	_, err := strconv.Atoi("4o2")
	defer func() {
		v := recover()
		assert.Equal(err.Error(), v.(*strconv.NumError).Error())
	}()

	Unwrap(strconv.Atoi("4o2"))
}
