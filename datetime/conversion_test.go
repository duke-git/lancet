package datetime

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestToUnix(t *testing.T) {
	assert := internal.NewAssert(t, "TestToUnix")

	tm1 := NewUnixNow()
	unixTimestamp := tm1.ToUnix()
	tm2 := NewUnix(unixTimestamp)

	assert.Equal(tm1, tm2)
}

func TestToFormat(t *testing.T) {
	assert := internal.NewAssert(t, "TestToFormat")

	_, err := NewFormat("2022/03/18 17:04:05")
	assert.IsNotNil(err)

	tm, err := NewFormat("2022-03-18 17:04:05")
	assert.IsNil(err)

	t.Log("ToFormat -> ", tm.ToFormat())
}

func TestToFormatForTpl(t *testing.T) {
	assert := internal.NewAssert(t, "TestToFormatForTpl")

	_, err := NewFormat("2022/03/18 17:04:05")
	assert.IsNotNil(err)

	tm, err := NewFormat("2022-03-18 17:04:05")
	assert.IsNil(err)

	t.Log("ToFormatForTpl -> ", tm.ToFormatForTpl("2006/01/02 15:04:05"))
}

func TestToIso8601(t *testing.T) {
	assert := internal.NewAssert(t, "TestToIso8601")

	_, err := NewISO8601("2022-03-18 17:04:05")
	assert.IsNotNil(err)

	tm, err := NewISO8601("2006-01-02T15:04:05.999Z")
	assert.IsNil(err)

	t.Log("ToIso8601 -> ", tm.ToIso8601())
}
