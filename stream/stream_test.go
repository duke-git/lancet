package stream

import (
	"fmt"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestOf(t *testing.T) {
	assert := internal.NewAssert(t, "TestFromSlice")

	stream := Of(1, 2, 3)
	assert.Equal([]int{1, 2, 3}, stream.ToSlice())
}

func TestGenerate(t *testing.T) {
	assert := internal.NewAssert(t, "TestFromSlice")

	n := 0
	max := 4

	generator := func() func() (int, bool) {
		return func() (int, bool) {
			n++
			return n, n < max
		}
	}

	stream := Generate(generator)

	assert.Equal([]int{1, 2, 3}, stream.ToSlice())
}

func TestFromSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestFromSlice")

	stream := FromSlice([]int{1, 2, 3})

	assert.Equal([]int{1, 2, 3}, stream.ToSlice())
}

func TestFromChannel(t *testing.T) {
	assert := internal.NewAssert(t, "TestFromChannel")

	ch := make(chan int)
	go func() {
		for i := 1; i < 4; i++ {
			ch <- i
		}
		close(ch)
	}()

	stream := FromChannel(ch)

	assert.Equal([]int{1, 2, 3}, stream.ToSlice())
}

func TestFromRange(t *testing.T) {
	assert := internal.NewAssert(t, "TestFromRange")

	s1 := FromRange(1, 5, 1)
	s2 := FromRange(1.1, 5.0, 1.0)

	assert.Equal([]int{1, 2, 3, 4, 5}, s1.ToSlice())
	assert.Equal([]float64{1.1, 2.1, 3.1, 4.1}, s2.ToSlice())
}

func TestStream_Distinct(t *testing.T) {
	assert := internal.NewAssert(t, "TestStream_Distinct")

	nums := FromSlice([]int{1, 2, 2, 3, 3, 3})
	distinctNums := nums.Distinct()

	assert.Equal([]int{1, 2, 2, 3, 3, 3}, nums.ToSlice())
	assert.Equal([]int{1, 2, 3}, distinctNums.ToSlice())

	type Person struct {
		Id   string
		Name string
		Age  uint
	}

	people := []Person{
		{Id: "001", Name: "Tom", Age: 10},
		{Id: "001", Name: "Tom", Age: 10},
		{Id: "002", Name: "Jim", Age: 20},
		{Id: "003", Name: "Mike", Age: 30},
	}

	stream := FromSlice(people)
	distinctStream := stream.Distinct()

	// {[{001 Tom 10} {001 Tom 10} {002 Jim 20} {003 Mike 30}]}
	t.Log(stream)

	// {[{001 Tom 10} {002 Jim 20} {003 Mike 30}]}
	t.Log(distinctStream)
}

func TestStream_Filter(t *testing.T) {
	assert := internal.NewAssert(t, "TestStream_Filter")

	stream := FromSlice([]int{1, 2, 3, 4, 5})

	isEven := func(n int) bool {
		return n%2 == 0
	}

	even := stream.Filter(isEven)

	assert.Equal([]int{2, 4}, even.ToSlice())
}

func TestStream_Map(t *testing.T) {
	assert := internal.NewAssert(t, "TestStream_Map")

	stream := FromSlice([]int{1, 2, 3})

	addOne := func(n int) int {
		return n + 1
	}

	s := stream.Map(addOne)

	assert.Equal([]int{2, 3, 4}, s.ToSlice())
}

func TestStream_Peek(t *testing.T) {
	assert := internal.NewAssert(t, "TestStream_Peek")

	stream := FromSlice([]int{1, 2, 3, 4, 5, 6})

	result := []string{}
	stream = stream.Filter(func(n int) bool {
		return n <= 3
	}).Peek(func(n int) {
		result = append(result, fmt.Sprint("current: ", n))
	})

	assert.Equal([]int{1, 2, 3}, stream.ToSlice())
	assert.Equal([]string{
		"current: 1", "current: 2", "current: 3",
	}, result)
}

func TestStream_Skip(t *testing.T) {
	assert := internal.NewAssert(t, "TestStream_Peek")

	stream := FromSlice([]int{1, 2, 3, 4, 5, 6})

	s1 := stream.Skip(-1)
	s2 := stream.Skip(0)

	assert.Equal([]int{1, 2, 3, 4, 5, 6}, s1.ToSlice())
	assert.Equal([]int{1, 2, 3, 4, 5, 6}, s2.ToSlice())
}

func TestStream_Limit(t *testing.T) {
	assert := internal.NewAssert(t, "TestStream_Limit")

	stream := FromSlice([]int{1, 2, 3, 4, 5, 6})

	s1 := stream.Limit(-1)
	s2 := stream.Limit(0)
	s3 := stream.Limit(1)
	s4 := stream.Limit(6)

	assert.Equal([]int{}, s1.ToSlice())
	assert.Equal([]int{}, s2.ToSlice())
	assert.Equal([]int{1}, s3.ToSlice())
	assert.Equal([]int{1, 2, 3, 4, 5, 6}, s4.ToSlice())
}

func TestStream_AllMatch(t *testing.T) {
	assert := internal.NewAssert(t, "TestStream_AllMatch")

	stream := FromSlice([]int{1, 2, 3, 4, 5, 6})

	result1 := stream.AllMatch(func(item int) bool {
		return item > 0
	})

	result2 := stream.AllMatch(func(item int) bool {
		return item > 1
	})

	assert.Equal(true, result1)
	assert.Equal(false, result2)
}

func TestStream_AnyMatch(t *testing.T) {
	assert := internal.NewAssert(t, "TestStream_AnyMatch")

	stream := FromSlice([]int{1, 2, 3})

	result1 := stream.AnyMatch(func(item int) bool {
		return item > 3
	})

	result2 := stream.AnyMatch(func(item int) bool {
		return item > 1
	})

	assert.Equal(false, result1)
	assert.Equal(true, result2)
}

func TestStream_NoneMatch(t *testing.T) {
	assert := internal.NewAssert(t, "TestStream_NoneMatch")

	stream := FromSlice([]int{1, 2, 3})

	result1 := stream.NoneMatch(func(item int) bool {
		return item > 3
	})

	result2 := stream.NoneMatch(func(item int) bool {
		return item > 1
	})

	assert.Equal(true, result1)
	assert.Equal(false, result2)
}
