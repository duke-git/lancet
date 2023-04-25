package compare

import (
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestEqual(t *testing.T) {
	assert := internal.NewAssert(t, "TestEqual")

	assert.Equal(true, Equal(1, 1))
	assert.Equal(true, Equal(int64(1), int64(1)))
	assert.Equal(true, Equal("a", "a"))
	assert.Equal(true, Equal(true, true))
	assert.Equal(true, Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.Equal(true, Equal(map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 2: "b"}))

	assert.Equal(false, Equal(1, 2))
	assert.Equal(false, Equal(1, int64(1)))
	assert.Equal(false, Equal("a", "b"))
	assert.Equal(false, Equal(true, false))
	assert.Equal(false, Equal([]int{1, 2}, []int{1, 2, 3}))
	assert.Equal(false, Equal(map[int]string{1: "a", 2: "b"}, map[int]string{1: "a"}))

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	time3 := time1.Add(time.Second)

	assert.Equal(false, Equal(time1, time2))
	assert.Equal(true, Equal(time2, time3))

	st1 := struct {
		A string
		B string
	}{
		A: "a",
		B: "b",
	}

	st2 := struct {
		A string
		B string
	}{
		A: "a",
		B: "b",
	}

	st3 := struct {
		A string
		B string
	}{
		A: "a1",
		B: "b",
	}

	assert.Equal(true, Equal(st1, st2))
	assert.Equal(false, Equal(st1, st3))
}

func TestEqualValue(t *testing.T) {
	assert := internal.NewAssert(t, "TestEqualValue")

	assert.Equal(true, EqualValue(1, 1))
	assert.Equal(true, EqualValue(int(1), int64(1)))
	assert.Equal(true, EqualValue(1, "1"))

	assert.Equal(false, EqualValue(1, "2"))
}
