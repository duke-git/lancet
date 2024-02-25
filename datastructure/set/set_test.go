package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestSet_FromSlice(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_FromSlice")

	s1 := FromSlice([]int{1, 2, 2, 3})
	assert.Equal(3, s1.Size())
	assert.Equal(true, s1.Contain(1))
	assert.Equal(true, s1.Contain(2))
	assert.Equal(true, s1.Contain(3))

	s2 := FromSlice([]int{})
	assert.Equal(0, s2.Size())
}

func TestSet_Add(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_Add")

	set := New[int]()
	set.Add(1, 2, 3)

	cmpSet := New(1, 2, 3)

	assert.Equal(true, set.Equal(cmpSet))
}

func TestSet_AddIfNotExist(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_AddIfNotExist")

	set := New[int]()
	set.Add(1, 2, 3)

	assert.Equal(false, set.AddIfNotExist(1))
	assert.Equal(true, set.AddIfNotExist(4))
	assert.Equal(New(1, 2, 3, 4), set)
}

func TestSet_AddIfNotExistBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_AddIfNotExistBy")

	set := New[int]()
	set.Add(1, 2)

	ok := set.AddIfNotExistBy(3, func(val int) bool {
		return val%2 != 0
	})

	notOk := set.AddIfNotExistBy(4, func(val int) bool {
		return val%2 != 0
	})

	assert.Equal(true, ok)
	assert.Equal(false, notOk)

	assert.Equal(true, set.Contain(3))
	assert.Equal(false, set.Contain(4))
}

func TestSet_Contain(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_Contain")

	set := New[int]()
	set.Add(1, 2, 3)

	assert.Equal(true, set.Contain(1))
	assert.Equal(false, set.Contain(4))
}

func TestSet_ContainAll(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_ContainAll")

	set1 := New(1, 2, 3)
	set2 := New(1, 2)
	set3 := New(1, 2, 3, 4)

	assert.Equal(true, set1.ContainAll(set2))
	assert.Equal(false, set1.ContainAll(set3))
}

func TestSet_Clone(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_Clone")

	set1 := New(1, 2, 3)
	set2 := set1.Clone()

	assert.Equal(true, set1.Size() == set2.Size())
	assert.Equal(true, set1.ContainAll(set2))
}

func TestSet_Delete(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_Delete")

	set := New[int]()
	set.Add(1, 2, 3)
	set.Delete(3)

	assert.Equal(true, set.Equal(New(1, 2)))
}

func TestSet_Equal(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_Equal")

	set1 := New(1, 2, 3)
	set2 := New(1, 2, 3)
	set3 := New(1, 2, 3, 4)

	assert.Equal(true, set1.Equal(set2))
	assert.Equal(false, set1.Equal(set3))
}

func TestSet_Iterate(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_Iterate")

	set := New(1, 2, 3)
	arr := []int{}
	set.Iterate(func(value int) {
		arr = append(arr, value)
	})

	assert.Equal(3, len(arr))
}

func TestSet_IsEmpty(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_IsEmpty")

	set := New[int]()
	assert.Equal(true, set.IsEmpty())
}

func TestSet_Size(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_Size")

	set := New(1, 2, 3)
	assert.Equal(3, set.Size())
}

func TestSet_Values(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_Values")

	set := New(1, 2, 3)
	values := set.Values()

	assert.Equal(3, len(values))
}

func TestSet_Union(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_Union")

	set1 := New(1, 2, 3)
	set2 := New(2, 3, 4, 5)

	unionSet := set1.Union(set2)

	assert.Equal(New(1, 2, 3, 4, 5), unionSet)
}

func TestSet_Intersection(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_Intersection")

	set1 := New(1, 2, 3)
	set2 := New(2, 3, 4, 5)
	intersectionSet := set1.Intersection(set2)

	assert.Equal(New(2, 3), intersectionSet)
}

func TestSet_SymmetricDifference(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_SymmetricDifference")

	set1 := New(1, 2, 3)
	set2 := New(2, 3, 4, 5)

	assert.Equal(New(1, 4, 5), set1.SymmetricDifference(set2))
}

func TestSet_Minus(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSet_Minus")

	set1 := New(1, 2, 3)
	set2 := New(2, 3, 4, 5)
	set3 := New(2, 3)

	assert.Equal(New(1), set1.Minus(set2))
	assert.Equal(New(4, 5), set2.Minus(set3))
}

func TestEachWithBreak(t *testing.T) {
	// s := New(1, 2, 3, 4, 5)

	// var sum int

	// s.EachWithBreak(func(n int) bool {
	// 	if n > 3 {
	// 		return false
	// 	}
	// 	sum += n
	// 	return true
	// })

	// assert := internal.NewAssert(t, "TestEachWithBreak")
	// assert.Equal(6, sum)
}

// func TestPop(t *testing.T) {
// 	assert := internal.NewAssert(t, "TestPop")

// 	s := New[int]()

// 	val, ok := s.Pop()
// 	assert.Equal(0, val)
// 	assert.Equal(false, ok)

// 	s.Add(1)
// 	s.Add(2)
// 	s.Add(3)

// 	// s = New(1, 2, 3, 4, 5)

// 	val, ok = s.Pop()
// 	assert.Equal(3, val)
// 	assert.Equal(true, ok)
// }
