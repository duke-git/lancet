package datastructure

import (
	"reflect"
	"sync"
)

type CopyOnWriteList[T any] struct {
	data []T
	lock sync.Locker
}

// NewCopyOnWriteList Creates an empty list.
func NewCopyOnWriteList[T any](data []T) *CopyOnWriteList[T] {
	return &CopyOnWriteList[T]{data: data, lock: &sync.RWMutex{}}
}

func (c *CopyOnWriteList[T]) getList() []T {
	return c.data
}
func (c *CopyOnWriteList[T]) setList(data []T) {
	c.data = data
}

// Size returns the number of elements in this list.
func (c *CopyOnWriteList[T]) Size() int {
	return len(c.getList())
}

// IsEmpty returns true if this list contains no elements.
func (c *CopyOnWriteList[T]) IsEmpty() bool {
	return c.Size() == 0
}

// Contain returns true if this list contains the specified element.
func (c *CopyOnWriteList[T]) Contain(e T) bool {
	list := c.getList()
	return indexOf(e, list, 0, c.Size()) >= 0
}

// ValueOf returns the index of the first occurrence of the specified element in this list, or null if this list does not contain the element.
func (c *CopyOnWriteList[T]) ValueOf(index int) *T {
	list := c.getList()
	if index < 0 || index >= len(c.data) {
		return nil
	}
	return get(list, index)
}

// IndexOf returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (c *CopyOnWriteList[T]) IndexOf(e T) int {
	list := c.getList()
	return indexOf(e, list, 0, c.Size())
}

// indexOf returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
// start the start position of the search (inclusive)
// end the end position of the search (exclusive)
func indexOf[T any](o T, e []T, start int, end int) int {
	if start >= end {
		return -1
	}
	for i := start; i < end; i++ {
		if reflect.DeepEqual(e[i], o) {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (c *CopyOnWriteList[T]) LastIndexOf(e T) int {
	list := c.getList()
	return lastIndexOf(e, list, 0, c.Size())
}

// lastIndexOf returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
// start the start position of the search (inclusive)
// end the end position of the search (exclusive)
func lastIndexOf[T any](o T, e []T, start int, end int) int {
	if start >= end {
		return -1
	}
	for i := end - 1; i >= start; i-- {
		if reflect.DeepEqual(e[i], o) {
			return i
		}
	}
	return -1
}

// get returns the element at the specified position in this list.
func get[T any](o []T, index int) *T {
	return &o[index]
}

// Get returns the element at the specified position in this list.
func (c *CopyOnWriteList[T]) Get(index int) *T {
	list := c.getList()
	if index < 0 || index >= len(list) {
		return nil
	}
	return get(list, index)
}

func (c *CopyOnWriteList[T]) set(index int, e T) (oldValue *T) {
	lock := c.lock
	lock.Lock()
	defer lock.Unlock()

	list := c.getList()
	oldValue = get(list, index)

	if reflect.DeepEqual(oldValue, e) {
		c.setList(list)
	} else {
		newList := make([]T, len(list))
		copy(newList, list)
		newList[index] = e
		c.setList(newList)
	}
	return
}

// Set replaces the element at the specified position in this list with the specified element.
func (c *CopyOnWriteList[T]) Set(index int, e T) (oldValue *T, ok bool) {
	list := c.getList()
	if index < 0 || index >= len(list) {
		return oldValue, false
	}
	return c.set(index, e), true
}

// Add appends the specified element to the end of this list.
func (c *CopyOnWriteList[T]) Add(e T) bool {
	lock := c.lock
	lock.Lock()
	defer lock.Unlock()

	list := c.getList()
	newList := make([]T, len(list)+1)
	copy(newList, list)
	newList[len(list)] = e
	c.setList(newList)
	return true
}

// AddAll appends all the elements in the specified collection to the end of this list
func (c *CopyOnWriteList[T]) AddAll(e []T) bool {
	lock := c.lock
	lock.Lock()
	defer lock.Unlock()

	list := c.getList()
	newList := make([]T, len(list)+len(e))
	copy(newList, list)
	copy(newList[len(list):], e)
	c.setList(newList)
	return true
}

// AddByIndex inserts the specified element at the specified position in this list.
func (c *CopyOnWriteList[T]) AddByIndex(index int, e T) bool {
	lock := c.lock
	lock.Lock()
	defer lock.Unlock()

	list := c.getList()
	length := len(list)
	if index < 0 || index > length {
		return false
	}
	var newList []T
	var numMove = length - index
	if numMove == 0 {
		newList = make([]T, length+1)
		copy(newList, list)
	} else {
		newList = make([]T, length+1)
		copy(newList, list[:index])
		copy(newList[index+1:], list[index:])
	}
	newList[index] = e
	c.setList(newList)
	return true
}

// delete removes the element at the specified position in this list.
func (c *CopyOnWriteList[T]) delete(index int) *T {
	lock := c.lock
	lock.Lock()
	defer lock.Unlock()

	list := c.getList()
	length := len(list)

	oldValue := get(list, index)
	numMove := length - index - 1
	var newList []T
	if numMove == 0 {
		newList = make([]T, length-1)
		copy(newList, list[:index])
	} else {
		newList = make([]T, length-1)
		copy(newList, list[:index])
		copy(newList[index:], list[index+1:])
	}

	c.setList(newList)
	return oldValue
}

// DeleteAt removes the element at the specified position in this list.
func (c *CopyOnWriteList[T]) DeleteAt(index int) (*T, bool) {
	list := c.getList()
	if index < 0 || index >= len(list) {
		return nil, false
	}
	return c.delete(index), true
}

// DeleteBy removes the first occurrence of the specified element from this list, if it is present.
func (c *CopyOnWriteList[T]) DeleteBy(o T) (*T, bool) {
	list := c.getList()
	index := indexOf(o, list, 0, len(list))
	if index == -1 {
		return nil, false
	}
	return c.delete(index), true
}

// DeleteRange removes from this list all the elements whose index is between fromIndex, inclusive, and toIndex, exclusive.
// left close and right open
func (c *CopyOnWriteList[T]) DeleteRange(start int, end int) {
	lock := c.lock
	lock.Lock()
	defer lock.Unlock()

	list := c.getList()
	length := len(list)
	if start < 0 || end > length || start > end {
		return
	}
	var newList []T
	numMove := length - end
	if numMove == 0 {
		newList = make([]T, length-(end-start))
		copy(newList, list[:start])
	} else {
		newList = make([]T, length-(end-start))
		copy(newList, list[:start])
		copy(newList[start:], list[end:])
	}
	c.setList(newList)
}

// DeleteIf removes all the elements of this collection that satisfy the given predicate.
func (c *CopyOnWriteList[T]) DeleteIf(f func(T) bool) {
	lock := c.lock
	lock.Lock()
	defer lock.Unlock()

	list := c.getList()
	length := len(list)
	var newList []T
	for i := 0; i < length; i++ {
		if !f(list[i]) {
			newList = append(newList, list[i])
		}
	}
	c.setList(newList)
}

// Equal returns true if the specified object is equal to this list.
func (c *CopyOnWriteList[T]) Equal(other *[]T) bool {
	if other == nil {
		return false
	}
	if c.Size() != len(*other) {
		return false
	}
	list := c.getList()
	otherList := NewCopyOnWriteList(*other).getList()
	for i := 0; i < len(list); i++ {
		if !reflect.DeepEqual(list[i], otherList[i]) {
			return false
		}
	}
	return true
}
