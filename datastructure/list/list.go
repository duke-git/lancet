// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package datastructure implements some data structure. eg. list, linklist, stack, queue, tree, graph.
package datastructure

import (
	"reflect"

	"github.com/duke-git/lancet/v2/iterator"
)

// List is a linear table, implemented with slice.
type List[T any] struct {
	data []T
}

// NewList return a pointer of List.
func NewList[T any](data []T) *List[T] {
	return &List[T]{data: data}
}

// Data return list data.
func (l *List[T]) Data() []T {
	return l.data
}

// ValueOf return the value pointer at index of list data.
func (l *List[T]) ValueOf(index int) (*T, bool) {
	if index < 0 || index >= len(l.data) {
		return nil, false
	}
	return &l.data[index], true
}

// IndexOf returns the index of value. if not found return -1.
func (l *List[T]) IndexOf(value T) int {
	index := -1
	data := l.data
	for i, v := range data {
		if reflect.DeepEqual(v, value) {
			index = i
			break
		}
	}
	return index
}

// LastIndexOf returns the index of the last occurrence of the value in this list.
// if not found return -1.
func (l *List[T]) LastIndexOf(value T) int {
	index := -1
	data := l.data
	for i := len(data) - 1; i >= 0; i-- {
		if reflect.DeepEqual(data[i], value) {
			index = i
			break
		}
	}
	return index
}

// IndexOfFunc returns the first index satisfying f(v)
// if not found return -1.
func (l *List[T]) IndexOfFunc(f func(T) bool) int {
	index := -1
	data := l.data
	for i, v := range data {
		if f(v) {
			index = i
			break
		}
	}
	return index
}

// LastIndexOfFunc returns the index of the last occurrence of the value in this list satisfying f(data[i])
// if not found return -1.
func (l *List[T]) LastIndexOfFunc(f func(T) bool) int {
	index := -1
	data := l.data
	for i := len(data) - 1; i >= 0; i-- {
		if f(data[i]) {
			index = i
			break
		}
	}
	return index
}

// Contain checks if the value in the list or not.
func (l *List[T]) Contain(value T) bool {
	data := l.data
	for _, v := range data {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Push append value to the list data.
func (l *List[T]) Push(value T) {
	l.data = append(l.data, value)
}

// InsertAtFirst insert value into list at first index.
func (l *List[T]) InsertAtFirst(value T) {
	l.InsertAt(0, value)
}

// InsertAtLast insert value into list at last index.
func (l *List[T]) InsertAtLast(value T) {
	l.InsertAt(len(l.data), value)
}

// InsertAt insert value into list at index.
func (l *List[T]) InsertAt(index int, value T) {
	data := l.data
	size := len(data)

	if index < 0 || index > size {
		return
	}
	l.data = append(data[:index], append([]T{value}, data[index:]...)...)
}

// PopFirst delete the first value of list and return it.
func (l *List[T]) PopFirst() (*T, bool) {
	if len(l.data) == 0 {
		return nil, false
	}

	v := l.data[0]
	l.DeleteAt(0)

	return &v, true
}

// PopLast delete the last value of list and return it.
func (l *List[T]) PopLast() (*T, bool) {
	size := len(l.data)
	if size == 0 {
		return nil, false
	}

	v := l.data[size-1]
	l.DeleteAt(size - 1)

	return &v, true
}

// DeleteAt delete the value of list at index.
func (l *List[T]) DeleteAt(index int) {
	data := l.data
	size := len(data)
	if index < 0 || index > size-1 {
		return
	}
	if index == size-1 {
		data = data[:index]
	} else {
		data = append(data[:index], data[index+1:]...)
	}
	l.data = data
}

// DeleteIf delete all satisfying f(data[i]), returns count of removed elements
func (l *List[T]) DeleteIf(f func(T) bool) int {
	data := l.data
	size := len(data)

	var c int
	for index := 0; index < len(data); index++ {
		if !f(data[index]) {
			continue
		}
		if index == size-1 {
			data = data[:index]
		} else {
			data = append(data[:index], data[index+1:]...)
			index--
		}
		c++
	}

	if c > 0 {
		l.data = data
	}
	return c
}

// UpdateAt update value of list at index, index shoud between 0 and list size -1
func (l *List[T]) UpdateAt(index int, value T) {
	data := l.data
	size := len(data)

	if index < 0 || index >= size {
		return
	}
	l.data = append(data[:index], append([]T{value}, data[index+1:]...)...)
}

// Equal compare list to other list, use reflect.DeepEqual.
func (l *List[T]) Equal(other *List[T]) bool {
	if len(l.data) != len(other.data) {
		return false
	}

	for i := 0; i < len(l.data); i++ {
		if !reflect.DeepEqual(l.data[i], other.data[i]) {
			return false
		}
	}

	return true
}

// IsEmpty check if the list is empty or not.
func (l *List[T]) IsEmpty() bool {
	return len(l.data) == 0
}

// Clear the data of list.
func (l *List[T]) Clear() {
	l.data = make([]T, 0)
}

// Clone return a copy of list.
func (l *List[T]) Clone() *List[T] {
	cl := NewList(make([]T, len(l.data)))
	copy(cl.data, l.data)

	return cl
}

// Merge two list, return new list, don't change original list.
func (l *List[T]) Merge(other *List[T]) *List[T] {
	l1, l2 := len(l.data), len(other.data)
	ml := NewList(make([]T, l1+l2))

	data := append([]T{}, append(l.data, other.data...)...)
	ml.data = data

	return ml
}

// Size return number of list data items.
func (l *List[T]) Size() int {
	return len(l.data)
}

// Cap return cap of the inner data.
func (l *List[T]) Cap() int {
	return cap(l.data)
}

// Swap the value of index i and j in list.
func (l *List[T]) Swap(i, j int) {
	size := len(l.data)
	if i < 0 || i >= size || j < 0 || j >= size {
		return
	}
	l.data[i], l.data[j] = l.data[j], l.data[i]
}

// Reverse the item order of list.
func (l *List[T]) Reverse() {
	for i, j := 0, len(l.data)-1; i < j; i, j = i+1, j-1 {
		l.data[i], l.data[j] = l.data[j], l.data[i]
	}
}

// Unique remove duplicate items in list.
func (l *List[T]) Unique() {
	data := l.data
	size := len(data)

	uniqueData := make([]T, 0)
	for i := 0; i < size; i++ {
		value := data[i]
		skip := true
		for _, v := range uniqueData {
			if reflect.DeepEqual(value, v) {
				skip = false
				break
			}
		}
		if skip {
			uniqueData = append(uniqueData, value)
		}
	}

	l.data = uniqueData
}

// Union creates a new list contain all element in list l and other, remove duplicate element.
func (l *List[T]) Union(other *List[T]) *List[T] {
	result := NewList([]T{})

	result.data = append(result.data, l.data...)
	result.data = append(result.data, other.data...)
	result.Unique()

	return result
}

// Intersection creates a new list whose element both be contained in list l and other.
func (l *List[T]) Intersection(other *List[T]) *List[T] {
	result := NewList(make([]T, 0))

	for _, v := range l.data {
		if other.Contain(v) {
			result.data = append(result.data, v)
		}
	}

	return result
}

// Difference returns the difference between two collections.
// return a list whose element in the original list, not in the given list.
func (l *List[T]) Difference(other *List[T]) *List[T] {
	result := NewList(make([]T, 0))

	intersectList := l.Intersection(other)

	for _, v := range l.data {
		if !intersectList.Contain(v) {
			result.data = append(result.data, v)
		}
	}

	return result
}

// SymmetricDifference oppoiste operation of intersection function.
func (l *List[T]) SymmetricDifference(other *List[T]) *List[T] {
	result := NewList(make([]T, 0))

	intersectList := l.Intersection(other)

	for _, v := range l.data {
		if !intersectList.Contain(v) {
			result.data = append(result.data, v)
		}
	}

	for _, v := range other.data {
		if !intersectList.Contain(v) {
			result.data = append(result.data, v)
		}
	}

	return result
}

// SubList returns a sub list of the original list between the specified fromIndex, inclusive, and toIndex, exclusive.
func (l *List[T]) SubList(fromIndex, toIndex int) *List[T] {
	data := l.data[fromIndex:toIndex]
	subList := make([]T, len(data))
	copy(subList, data)
	return NewList(subList)
}

// ForEach performs the given action for each element of the list.
func (l *List[T]) ForEach(consumer func(T)) {
	for _, it := range l.data {
		consumer(it)
	}
}

// RetainAll retains only the elements in this list that are contained in the given list.
func (l *List[T]) RetainAll(list *List[T]) bool {
	return l.batchRemove(list, true)
}

// DeleteAll removes from this list all of its elements that are contained in the given list.
func (l *List[T]) DeleteAll(list *List[T]) bool {
	return l.batchRemove(list, false)
}

func (l *List[T]) batchRemove(list *List[T], complement bool) bool {
	var (
		w    = 0
		data = l.data
		size = len(data)
	)

	for i := 0; i < size; i++ {
		if list.Contain(data[i]) == complement {
			data[w] = data[i]
			w++
		}
	}

	if w != size {
		l.data = data[:w]
		return true
	}
	return false
}

// Iterator returns an iterator over the elements in this list in proper sequence.
func (l *List[T]) Iterator() iterator.Iterator[T] {
	return iterator.FromSlice(l.data)
}

// ListToMap convert a list to a map based on iteratee function.
func ListToMap[T any, K comparable, V any](list *List[T], iteratee func(T) (K, V)) map[K]V {
	result := make(map[K]V, list.Size())
	for _, item := range list.data {
		k, v := iteratee(item)
		result[k] = v
	}

	return result
}
