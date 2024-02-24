package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestCopyOnWriteList_ValueOf(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})

	assert := internal.NewAssert(t, "CopyOnWriteList_IndexOf")
	of, ok := list.ValueOf(3)
	assert.Equal(4, *of)
	assert.Equal(true, ok)

	_, ok = list.ValueOf(6)
	assert.Equal(false, ok)
}

func TestCopyOnWriteList_Contain(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_Contains")
	assert.Equal(true, list.Contain(3))
}

func TestCopyOnWriteList_IsEmpty(t *testing.T) {
	list := NewCopyOnWriteList([]int{})
	assert := internal.NewAssert(t, "CopyOnWriteList_IsEmpty")
	assert.Equal(true, list.IsEmpty())
}

func TestCopyOnWriteList_Size(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_size")
	assert.Equal(5, list.Size())
}

func TestCopyOnWriteList_GetList(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_GetList")
	assert.Equal([]int{1, 2, 3, 4, 5}, list.getList())
}

func TestCopyOnWriteList_Get(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_Get")
	i := list.Get(2)
	assert.Equal(3, *i)
}

func TestCopyOnWriteList_Set(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_Set")
	list.Set(2, 6)
	assert.Equal(6, list.getList()[2])

	list = NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	list.Set(0, 6)
	assert.Equal(6, list.getList()[0])

	list = NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	list.Set(0, 1)
	assert.Equal(1, list.getList()[0])
}

func TestCopyOnWriteList_Add(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_Add")
	list.Add(6)
	assert.Equal([]int{1, 2, 3, 4, 5, 6}, list.getList())
}

func TestCopyOnWriteList_AddAll(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_AddAll")
	list.AddAll([]int{6, 7, 8})
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8}, list.getList())
}

func TestCopyOnWriteList_AddByIndex(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_AddByIndex")
	list.AddByIndex(2, 6)
	assert.Equal([]int{1, 2, 6, 3, 4, 5}, list.getList())

	list = NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	list.AddByIndex(0, 6)
	assert.Equal([]int{6, 1, 2, 3, 4, 5}, list.getList())

	list = NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	list.AddByIndex(5, 6)
	assert.Equal([]int{1, 2, 3, 4, 5, 6}, list.getList())
}

func TestCopyOnWriteList_DeleteAt2(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_RemoveByIndex")
	list.DeleteAt(2)
	assert.Equal([]int{1, 2, 4, 5}, list.getList())

	list = NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	list.DeleteAt(4)
	assert.Equal([]int{1, 2, 3, 4}, list.getList())
}

func TestCopyOnWriteList_RemoveByValue(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_RemoveByValue")
	list.DeleteBy(3)
	assert.Equal([]int{1, 2, 4, 5}, list.getList())
}

func TestCopyOnWriteList_DeleteRange(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_RemoveRange")
	list.DeleteRange(1, 3)
	assert.Equal([]int{1, 4, 5}, list.getList())

	list = NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	list.DeleteRange(0, 5)
	assert.Equal([]int{}, list.getList())
}

func TestCopyOnWriteList_LastIndexOf(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5, 3})
	assert := internal.NewAssert(t, "CopyOnWriteList_LastIndexOf")
	assert.Equal(5, list.LastIndexOf(3))
}

func TestCopyOnWriteList_DeleteAt(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5, 3})
	assert := internal.NewAssert(t, "CopyOnWriteList_DeleteAt")
	list.DeleteAt(2)
	assert.Equal([]int{1, 2, 4, 5, 3}, list.getList())
}

func TestCopyOnWriteList_DeleteBy(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5, 3})
	assert := internal.NewAssert(t, "CopyOnWriteList_DeleteBy")
	list.DeleteBy(3)
	assert.Equal([]int{1, 2, 4, 5, 3}, list.getList())
}

func TestCopyOnWriteList_DeleteIf(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5, 3, 6})
	assert := internal.NewAssert(t, "CopyOnWriteList_DeleteIf")

	list.DeleteIf(func(i int) bool {
		return i%2 == 0
	})

	assert.Equal([]int{1, 3, 5, 3}, list.getList())
}

func TestCopyOnWriteList_Equal(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5, 3, 6})
	assert := internal.NewAssert(t, "CopyOnWriteList_Equal")

	assert.Equal(true, list.Equal(&[]int{1, 2, 3, 4, 5, 3, 6}))
}

func TestCopyOnWriteList_ForEach(t *testing.T) {
	testList := make([]int, 0)
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5, 3, 6})
	assert := internal.NewAssert(t, "CopyOnWriteList_ForEach")

	list.ForEach(func(i int) {
		testList = append(testList, i)
	})
	assert.Equal([]int{1, 2, 3, 4, 5, 3, 6}, testList)

	list.ForEach(func(i int) {
		list.Add(i)
	})
	assert.Equal([]int{1, 2, 3, 4, 5, 3, 6, 1, 2, 3, 4, 5, 3, 6}, list.getList())
}

func TestCopyOnWriteList_Clear(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5, 3, 6})
	assert := internal.NewAssert(t, "CopyOnWriteList_Clear")

	list.Clear()
	assert.Equal([]int{}, list.getList())
}

func TestCopyOnWriteList_Merge(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 3, 5, 7, 9})
	assert := internal.NewAssert(t, "CopyOnWriteList_Merge")

	list.Merge([]int{2, 4, 6, 8, 10})
	assert.Equal([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}, list.getList())

}

func TestCopyOnWriteList_Sort(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10})
	assert := internal.NewAssert(t, "CopyOnWriteList_Sort")

	list.Sort(func(i, j int) bool {
		return i < j
	})
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, list.getList())
}

func TestCopyOnWriteList_IndexOf(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10})
	assert := internal.NewAssert(t, "CopyOnWriteList_IndexOf")

	assert.Equal(0, list.IndexOf(1))
	assert.Equal(9, list.IndexOf(10))
	assert.Equal(-1, list.IndexOf(11))
}

func TestCopyOnWriteList_SubList(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10})

	assert := internal.NewAssert(t, "CopyOnWriteList_SubList")

	list = NewCopyOnWriteList([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10})
	subList := list.SubList(1, 3)
	assert.Equal([]int{3, 5}, subList)

	list = NewCopyOnWriteList([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10})
	subList = list.SubList(1, 1)
	assert.Equal([]int{}, subList)

	list = NewCopyOnWriteList([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10})
	assert.Equal(10, list.Size())
	subList = list.SubList(1, 10)
	assert.Equal([]int{3, 5, 7, 9, 2, 4, 6, 8, 10}, subList)

	list = NewCopyOnWriteList([]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10})
	subList = list.SubList(11, 1)
	assert.Equal([]int{}, subList)
}

func TestCopyOnWriteListIndexOfFunc(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIndexOfFunc")

	list := NewCopyOnWriteList([]int{1, 2, 3})
	i := list.IndexOfFunc(func(a int) bool { return a == 1 })
	assert.Equal(0, i)

	i = list.IndexOfFunc(func(a int) bool { return a == 4 })
	assert.Equal(-1, i)
}

func TestNewCopyOnWriteListLastIndexOfFunc(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLastIndexOfFunc")

	list := NewCopyOnWriteList([]int{1, 2, 3, 3, 3, 3, 4, 5, 6, 9})
	i := list.LastIndexOfFunc(func(a int) bool { return a == 3 })
	assert.Equal(5, i)

	i = list.LastIndexOfFunc(func(a int) bool { return a == 10 })
	assert.Equal(-1, i)

	i = list.LastIndexOfFunc(func(a int) bool { return a == 4 })
	assert.Equal(6, i)

	i = list.LastIndexOfFunc(func(a int) bool { return a == 1 })
	assert.Equal(0, i)
}
