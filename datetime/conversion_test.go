package datetime

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestToUnix(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToUnix")

	tm1 := NewUnixNow()
	unixTimestamp := tm1.ToUnix()
	tm2 := NewUnix(unixTimestamp)

	assert.Equal(tm1, tm2)
}

func TestToFormat(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToFormat")

	tm, err := NewFormat("2022-03-18 17:04:05")
	t.Log("TestToFormat", tm.ToFormat())
	assert.IsNil(err)
}

func TestToFormatForTpl(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToFormatForTpl")

	_, err := NewFormat("2022/03/18 17:04:05")
	assert.IsNotNil(err)

	tm, err := NewFormat("2022-03-18 17:04:05")
	t.Log("TestToFormatForTpl", tm.ToFormatForTpl("2006/01/02 15:04:05"))
	assert.IsNil(err)
}

func TestToIso8601(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToIso8601")

	_, err := NewISO8601("2022-03-18 17:04:05")
	assert.IsNotNil(err)

	tm, err := NewISO8601("2006-01-02T15:04:05.999Z")
	t.Log("TestToIso8601", tm.ToIso8601())
	assert.IsNil(err)
}
