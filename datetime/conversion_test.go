package datetime

import (
	"testing"
	"time"

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

	tm1, err := NewFormat("2022/03/18 17:04:05")
	assert.IsNotNil(err)

	tm1, err = NewFormat("2022-03-18 17:04:05")
	assert.IsNil(err)

	// res := tm1.ToFormat()
	t1, err1 := time.Parse("2006-01-02 15:04:05", tm1.ToFormat())
	if err1 != nil {
		t.FailNow()
	}

	t2, err2 := time.Parse("2006-01-02 15:04:05", "2022-03-18 17:04:05")
	if err2 != nil {
		t.FailNow()
	}

	assert.Equal(t2, t1)
	// assert.Equal("2022-03-18 17:04:05", res)
}

func TestToFormatForTpl(t *testing.T) {
	assert := internal.NewAssert(t, "TestToFormatForTpl")

	tm1, err := NewFormat("2022/03/18 17:04:05")
	assert.IsNotNil(err)

	tm1, err = NewFormat("2022-03-18 17:04:05")
	assert.IsNil(err)

	res := tm1.ToFormatForTpl("2006/01/02 15:04:05")
	assert.Equal("2022/03/18 17:04:05", res)
}

func TestToIso8601(t *testing.T) {
	assert := internal.NewAssert(t, "TestToIso8601")

	tm1, err := NewISO8601("2022-03-18 17:04:05")
	assert.IsNotNil(err)

	tm1, err = NewISO8601("2006-01-02T15:04:05.999Z")
	assert.IsNil(err)

	res := tm1.ToIso8601()
	assert.Equal("2006-01-02T23:04:05+08:00", res)
}
