package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestListData(t *testing.T) {
	assert := internal.NewAssert(t, "TestListData")

	list := NewList([]int{1, 2, 3})
	assert.Equal([]int{1, 2, 3}, list.Data())
}

func TestValueOf(t *testing.T) {
	assert := internal.NewAssert(t, "TestValueOf")

	list := NewList([]int{1, 2, 3})
	v, ok := list.ValueOf(0)
	assert.Equal(1, *v)
	assert.Equal(true, ok)

	_, ok = list.ValueOf(3)
	assert.Equal(false, ok)
}

func TestIndexOf(t *testing.T) {
	assert := internal.NewAssert(t, "TestIndexOf")

	list := NewList([]int{1, 2, 3})
	i := list.IndexOf(1)
	assert.Equal(0, i)

	i = list.IndexOf(4)
	assert.Equal(-1, i)
}

func TestIndexOfFunc(t *testing.T) {
	assert := internal.NewAssert(t, "TestIndexOf")

	list := NewList([]int{1, 2, 3})
	i := list.IndexOfFunc(func(a int) bool { return a == 1 })
	assert.Equal(0, i)

	i = list.IndexOfFunc(func(a int) bool { return a == 4 })
	assert.Equal(-1, i)
}

func TestLastIndexOf(t *testing.T) {
	assert := internal.NewAssert(t, "TestIndexOf")

	list := NewList([]int{1, 2, 3, 3, 3, 3, 4, 5, 6, 9})
	i := list.LastIndexOf(3)
	assert.Equal(5, i)

	i = list.LastIndexOf(10)
	assert.Equal(-1, i)

	i = list.LastIndexOf(4)
	assert.Equal(6, i)

	i = list.LastIndexOf(1)
	assert.Equal(0, i)
}

func TestLastIndexOfFunc(t *testing.T) {
	assert := internal.NewAssert(t, "TestIndexOf")

	list := NewList([]int{1, 2, 3, 3, 3, 3, 4, 5, 6, 9})
	i := list.LastIndexOfFunc(func(a int) bool { return a == 3 })
	assert.Equal(5, i)

	i = list.LastIndexOfFunc(func(a int) bool { return a == 10 })
	assert.Equal(-1, i)

	i = list.LastIndexOfFunc(func(a int) bool { return a == 4 })
	assert.Equal(6, i)

	i = list.LastIndexOfFunc(func(a int) bool { return a == 1 })
	assert.Equal(0, i)
}

func TestContain(t *testing.T) {
	assert := internal.NewAssert(t, "TestContain")

	list := NewList([]int{1, 2, 3})
	assert.Equal(true, list.Contain(1))
	assert.Equal(false, list.Contain(0))
}

func TestPush(t *testing.T) {
	assert := internal.NewAssert(t, "TestPush")

	list := NewList([]int{1, 2, 3})
	list.Push(4)

	assert.Equal([]int{1, 2, 3, 4}, list.Data())
}

func TestInsertAtFirst(t *testing.T) {
	assert := internal.NewAssert(t, "TestInsertAtFirst")

	list := NewList([]int{1, 2, 3})
	list.InsertAtFirst(0)

	assert.Equal([]int{0, 1, 2, 3}, list.Data())
}

func TestInsertAtLast(t *testing.T) {
	assert := internal.NewAssert(t, "TestInsertAtLast")

	list := NewList([]int{1, 2, 3})
	list.InsertAtLast(4)

	assert.Equal([]int{1, 2, 3, 4}, list.Data())
}

func TestInsertAt(t *testing.T) {
	assert := internal.NewAssert(t, "TestInsertAt")

	list := NewList([]int{1, 2, 3})

	list.InsertAt(-1, 0)
	assert.Equal([]int{1, 2, 3}, list.Data())

	list.InsertAt(4, 0)
	assert.Equal([]int{1, 2, 3}, list.Data())

	list.InsertAt(0, 0)
	assert.Equal([]int{0, 1, 2, 3}, list.Data())

	list.InsertAt(4, 4)
	assert.Equal([]int{0, 1, 2, 3, 4}, list.Data())
}

func TestPopFirst(t *testing.T) {
	assert := internal.NewAssert(t, "TestPopFirst")

	list := NewList([]int{1, 2, 3})
	v, ok := list.PopFirst()
	assert.Equal(1, *v)
	assert.Equal(true, ok)
	assert.Equal([]int{2, 3}, list.Data())

	list2 := NewList([]int{})
	_, ok = list2.PopFirst()
	assert.Equal(false, ok)
	assert.Equal([]int{}, list2.Data())
}

func TestPopLast(t *testing.T) {
	assert := internal.NewAssert(t, "TestPopLast")

	list := NewList([]int{1, 2, 3})
	v, ok := list.PopLast()
	assert.Equal(3, *v)
	assert.Equal(true, ok)
	assert.Equal([]int{1, 2}, list.Data())

	list2 := NewList([]int{})
	_, ok = list2.PopLast()
	assert.Equal(false, ok)
	assert.Equal([]int{}, list2.Data())
}

func TestDeleteAt(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteAt")

	list := NewList([]int{1, 2, 3, 4})

	list.DeleteAt(-1)
	assert.Equal([]int{1, 2, 3, 4}, list.Data())

	list.DeleteAt(4)
	assert.Equal([]int{1, 2, 3, 4}, list.Data())

	list.DeleteAt(0)
	assert.Equal([]int{2, 3, 4}, list.Data())

	list.DeleteAt(2)
	assert.Equal([]int{2, 3}, list.Data())
}

func TestUpdateAt(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateAt")

	list := NewList([]int{1, 2, 3, 4})

	list.UpdateAt(-1, 0)
	assert.Equal([]int{1, 2, 3, 4}, list.Data())

	list.UpdateAt(4, 0)
	assert.Equal([]int{1, 2, 3, 4}, list.Data())

	list.UpdateAt(0, 5)
	assert.Equal([]int{5, 2, 3, 4}, list.Data())

	list.UpdateAt(3, 1)
	assert.Equal([]int{5, 2, 3, 1}, list.Data())
}

func TestEqual(t *testing.T) {
	assert := internal.NewAssert(t, "TestEqual")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{1, 2, 3, 4})
	list3 := NewList([]int{1, 2, 3})

	assert.Equal(true, list1.Equal(list2))
	assert.Equal(false, list1.Equal(list3))
}

func TestIsEmpty(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsEmpty")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{})

	assert.Equal(false, list1.IsEmpty())
	assert.Equal(true, list2.IsEmpty())
}

func TestIsClear(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsClear")

	list1 := NewList([]int{1, 2, 3, 4})
	list1.Clear()
	empty := NewList([]int{})

	assert.Equal(empty, list1)
}

func TestClone(t *testing.T) {
	assert := internal.NewAssert(t, "TestClone")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := list1.Clone()

	assert.Equal(true, list1.Equal(list2))
}

func TestMerge(t *testing.T) {
	assert := internal.NewAssert(t, "TestMerge")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{4, 5, 6})
	expected := NewList([]int{1, 2, 3, 4, 4, 5, 6})

	list3 := list1.Merge(list2)
	assert.Equal(true, expected.Equal(list3))
}

func TestSize(t *testing.T) {
	assert := internal.NewAssert(t, "TestSize")

	list := NewList([]int{1, 2, 3, 4})
	empty := NewList([]int{})

	assert.Equal(4, list.Size())
	assert.Equal(0, empty.Size())
}

func TestCap(t *testing.T) {
	assert := internal.NewAssert(t, "TestCap")

	data := make([]int, 0, 100)
	list := NewList(data)
	assert.Equal(100, list.Cap())

	data = make([]int, 0)
	list = NewList(data)
	assert.Equal(0, list.Cap())
}

func TestSwap(t *testing.T) {
	assert := internal.NewAssert(t, "TestSwap")

	list := NewList([]int{1, 2, 3, 4})
	expected := NewList([]int{4, 2, 3, 1})

	list.Swap(0, 3)

	assert.Equal(true, expected.Equal(list))
}

func TestReverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverse")

	list := NewList([]int{1, 2, 3, 4})
	expected := NewList([]int{4, 3, 2, 1})

	list.Reverse()

	assert.Equal(true, expected.Equal(list))
}

func TestUnique(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnique")

	list := NewList([]int{1, 2, 2, 3, 4})
	expected := NewList([]int{1, 2, 3, 4})

	list.Unique()

	assert.Equal(true, expected.Equal(list))
}

func TestUnion(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnion")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{4, 5, 6})
	expected := NewList([]int{1, 2, 3, 4, 5, 6})

	list3 := list1.Union(list2)
	assert.Equal(true, expected.Equal(list3))
}

func TestIntersection(t *testing.T) {
	assert := internal.NewAssert(t, "TestIntersection")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{4, 5, 6})
	expected := NewList([]int{4})

	list3 := list1.Intersection(list2)
	assert.Equal(true, expected.Equal(list3))
}

func TestDifference(t *testing.T) {
	assert := internal.NewAssert(t, "TestDifference")

	list1 := NewList([]int{1, 2, 3})
	list2 := NewList([]int{1, 2, 4})
	expected := NewList([]int{3})

	list3 := list1.Difference(list2)
	assert.Equal(true, expected.Equal(list3))
}

func TestSymmetricDifference(t *testing.T) {
	assert := internal.NewAssert(t, "TestSymmetricDifference")

	list1 := NewList([]int{1, 2, 3})
	list2 := NewList([]int{1, 2, 4})
	expected := NewList([]int{3, 4})

	list3 := list1.SymmetricDifference(list2)
	assert.Equal(true, expected.Equal(list3))
}

func TestSubSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestSubSlice")

	list := NewList([]int{1, 2, 3, 4, 5, 8})
	subList := list.SubList(2, 5)

	assert.Equal([]int{3, 4, 5}, subList.Data())
}

func BenchmarkSubSlice(b *testing.B) {
	list := NewList([]int{1, 2, 3, 4, 5, 8})
	for n := 0; n < b.N; n++ {
		list.SubList(2, 5)
	}
}

func TestDeleteIf(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteIf")

	list := NewList([]int{1, 1, 1, 1, 2, 3, 1, 1, 4, 1, 1, 1, 1, 1, 1})

	count := list.DeleteIf(func(a int) bool { return a == 1 })
	assert.Equal([]int{2, 3, 4}, list.Data())
	assert.Equal(12, count)

	count = list.DeleteIf(func(a int) bool { return a == 5 })
	assert.Equal([]int{2, 3, 4}, list.Data())
	assert.Equal(0, count)
}

func TestForEach(t *testing.T) {
	assert := internal.NewAssert(t, "TestForEach")

	list := NewList([]int{1, 2, 3, 4})

	rs := make([]int, 0)
	list.ForEach(func(i int) {
		rs = append(rs, i)
	})

	assert.Equal([]int{1, 2, 3, 4}, rs)
}

func TestRetainAll(t *testing.T) {
	assert := internal.NewAssert(t, "TestRetainAll")

	list := NewList([]int{1, 2, 3, 4})
	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{1, 2, 3, 4})

	retain := NewList([]int{1, 2})
	retain1 := NewList([]int{2, 3})
	retain2 := NewList([]int{1, 2, 5})

	list.RetainAll(retain)
	list1.RetainAll(retain1)
	list2.RetainAll(retain2)

	assert.Equal([]int{1, 2}, list.Data())
	assert.Equal([]int{2, 3}, list1.Data())
	assert.Equal([]int{1, 2}, list2.Data())
}

func TestDeleteAll(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteAll")

	list := NewList([]int{1, 2, 3, 4})
	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{1, 2, 3, 4})

	del := NewList([]int{1})
	del1 := NewList([]int{2, 3})
	del2 := NewList([]int{1, 2, 5})

	list.DeleteAll(del)
	list1.DeleteAll(del1)
	list2.DeleteAll(del2)
	assert.Equal([]int{2, 3, 4}, list.Data())
	assert.Equal([]int{1, 4}, list1.Data())
	assert.Equal([]int{3, 4}, list2.Data())
}

func TestIterator(t *testing.T) {
	assert := internal.NewAssert(t, "TestIterator")

	list := NewList([]int{1, 2, 3, 4})

	iterator := list.Iterator()

	rs := make([]int, 0)
	for iterator.HasNext() {
		item, _ := iterator.Next()
		rs = append(rs, item)
	}

	assert.Equal([]int{1, 2, 3, 4}, rs)
}

func TestListToMap(t *testing.T) {
	assert := internal.NewAssert(t, "ListToMap")

	list := NewList([]int{1, 2, 3, 4})

	result := ListToMap(list, func(n int) (int, bool) {
		return n, n > 1
	})
	expected := map[int]bool{1: false, 2: true, 3: true, 4: true}

	assert.Equal(expected, result)
}
