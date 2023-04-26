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

func TestLessThan(t *testing.T) {
	assert := internal.NewAssert(t, "TestLessThan")

	assert.Equal(true, LessThan(1, 2))
	assert.Equal(true, LessThan(1.1, 2.2))
	assert.Equal(true, LessThan("a", "b"))

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	assert.Equal(true, LessThan(time1, time2))

	assert.Equal(false, LessThan(1, 1))
	assert.Equal(false, LessThan(1, int64(1)))
}

func TestGreaterThan(t *testing.T) {
	assert := internal.NewAssert(t, "TestGreaterThan")

	assert.Equal(true, GreaterThan(2, 1))
	assert.Equal(true, GreaterThan(2.2, 1.1))
	assert.Equal(true, GreaterThan("b", "a"))

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	assert.Equal(true, GreaterThan(time2, time1))

	assert.Equal(false, GreaterThan(1, 2))
	assert.Equal(false, GreaterThan(int64(2), 1))
	assert.Equal(false, GreaterThan("b", "c"))
}

func TestLessOrEqual(t *testing.T) {
	assert := internal.NewAssert(t, "TestLessOrEqual")

	assert.Equal(true, LessOrEqual(1, 2))
	assert.Equal(true, LessOrEqual(1, 1))
	assert.Equal(true, LessOrEqual(1.1, 2.2))
	assert.Equal(true, LessOrEqual("a", "b"))

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	assert.Equal(true, LessOrEqual(time1, time2))

	assert.Equal(false, LessOrEqual(2, 1))
	assert.Equal(false, LessOrEqual(1, int64(2)))
}

func TestGreaterOrEqual(t *testing.T) {
	assert := internal.NewAssert(t, "TestGreaterThan")

	assert.Equal(true, GreaterOrEqual(2, 1))
	assert.Equal(true, GreaterOrEqual(1, 1))
	assert.Equal(true, GreaterOrEqual(2.2, 1.1))
	assert.Equal(true, GreaterOrEqual("b", "b"))

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	assert.Equal(true, GreaterOrEqual(time2, time1))

	assert.Equal(false, GreaterOrEqual(1, 2))
	assert.Equal(false, GreaterOrEqual(int64(2), 1))
	assert.Equal(false, GreaterOrEqual("b", "c"))
}
