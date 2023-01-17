package stream

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestFromSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestFromSlice")

	stream := FromSlice([]int{1, 2, 3})

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
