package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/internal"
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
	v, ok = list2.PopFirst()
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
	v, ok = list2.PopLast()
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

func TestEqutalTo(t *testing.T) {
	assert := internal.NewAssert(t, "TestEqutalTo")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{1, 2, 3, 4})
	list3 := NewList([]int{1, 2, 3})

	assert.Equal(true, list1.EqutalTo(list2))
	assert.Equal(false, list1.EqutalTo(list3))
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

	assert.Equal(true, list1.EqutalTo(list2))
}

func TestMerge(t *testing.T) {
	assert := internal.NewAssert(t, "TestMerge")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{4, 5, 6})
	expected := NewList([]int{1, 2, 3, 4, 4, 5, 6})

	list3 := list1.Merge(list2)
	assert.Equal(true, expected.EqutalTo(list3))
}

func TestSize(t *testing.T) {
	assert := internal.NewAssert(t, "TestSize")

	list := NewList([]int{1, 2, 3, 4})
	empty := NewList([]int{})

	assert.Equal(4, list.Size())
	assert.Equal(0, empty.Size())
}

func TestSwap(t *testing.T) {
	assert := internal.NewAssert(t, "TestSwap")

	list := NewList([]int{1, 2, 3, 4})
	expected := NewList([]int{4, 2, 3, 1})

	list.Swap(0, 3)

	assert.Equal(true, expected.EqutalTo(list))
}

func TestReverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverse")

	list := NewList([]int{1, 2, 3, 4})
	expected := NewList([]int{4, 3, 2, 1})

	list.Reverse()

	assert.Equal(true, expected.EqutalTo(list))
}

func TestUnique(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnique")

	list := NewList([]int{1, 2, 2, 3, 4})
	expected := NewList([]int{1, 2, 3, 4})

	list.Unique()

	assert.Equal(true, expected.EqutalTo(list))
}

func TestUnion(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnion")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{4, 5, 6})
	expected := NewList([]int{1, 2, 3, 4, 5, 6})

	list3 := list1.Union(list2)
	assert.Equal(true, expected.EqutalTo(list3))
}

func TestIntersection(t *testing.T) {
	assert := internal.NewAssert(t, "TestIntersection")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{4, 5, 6})
	expected := NewList([]int{4})

	list3 := list1.Intersection(list2)
	assert.Equal(true, expected.EqutalTo(list3))
}
