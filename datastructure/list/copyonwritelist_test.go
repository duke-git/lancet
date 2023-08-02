package datastructure

import (
	"github.com/duke-git/lancet/v2/internal"
	"testing"
)

func TestCopyOnWriteList_ValueOf(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})

	assert := internal.NewAssert(t, "CopyOnWriteList_IndexOf")
	of, ok := list.ValueOf(3)
	assert.Equal(4, *of)
	assert.Equal(true, ok)
}

func TestCopyOnWriteList_Contains(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_Contains")
	assert.Equal(true, list.Contain(3))
}

func TestCopyOnWriteList_IsEmpty(t *testing.T) {
	list := NewCopyOnWriteList([]int{})
	assert := internal.NewAssert(t, "CopyOnWriteList_IsEmpty")
	assert.Equal(true, list.IsEmpty())
}

func TestCopyOnWriteList_size(t *testing.T) {
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
	i, _ := list.Get(2)
	assert.Equal(3, i)
}

func TestCopyOnWriteList_Set(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_Set")
	list.Set(2, 6)
	assert.Equal(6, list.getList()[2])
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
}

func TestCopyOnWriteList_RemoveByIndex(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_RemoveByIndex")
	list.DeleteAt(2)
	assert.Equal([]int{1, 2, 4, 5}, list.getList())
}

func TestCopyOnWriteList_RemoveByValue(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_RemoveByValue")
	list.DeleteBy(3)
	assert.Equal([]int{1, 2, 4, 5}, list.getList())
}

func TestCopyOnWriteList_RemoveRange(t *testing.T) {
	list := NewCopyOnWriteList([]int{1, 2, 3, 4, 5})
	assert := internal.NewAssert(t, "CopyOnWriteList_RemoveRange")
	list.DeleteRange(1, 3)
	assert.Equal([]int{1, 4, 5}, list.getList())
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
