// Copyright 2021 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package iterator implements some feature of C++ STL iterators
package iterator

import "github.com/duke-git/lancet/v2/lancetconstraints"

// Iterator is used to iterate over a slice of data.
type Iterator[T any] interface {
	HasNext() bool
	Next() (T, bool)
}

// FromSlice returns an iterator over a slice of data.
func FromSlice[T any](data []T) Iterator[T] {
	// return &sliceIterator[T]{slice: data, index: 0, step: 1}
	return &sliceIterator[T]{slice: data}
}

type sliceIterator[T any] struct {
	slice []T
	// index int
	// step  int
}

func (iter *sliceIterator[T]) HasNext() bool {
	if len(iter.slice) == 0 {
		return false
	}
	return true
}

func (iter *sliceIterator[T]) Next() (T, bool) {
	if len(iter.slice) == 0 {
		var zero T
		return zero, false
	}
	item := iter.slice[0]
	iter.slice = iter.slice[1:]
	return item, true
}

// Count
func (iter *sliceIterator[T]) Count() int {
	count := len(iter.slice)
	iter.slice = []T{}
	return count
}

// func (iter *sliceIterator[T]) Begin(data []T) Iterator[T] {
// 	iter.data = data
// 	iter.index = 0
// 	iter.step = 1

// 	return iter
// }

// func (iter *sliceIterator[T]) End(data []T) Iterator[T] {
// 	iter.data = data
// 	iter.index = len(data)
// 	iter.step = 1

// 	return iter
// }

// FromRange creates a iterator which returns the numeric range between start inclusive and end
// exclusive by the step size. start should be less than end, step shoud be positive.
func FromRange[T lancetconstraints.Number](start, end, step T) Iterator[T] {
	if end < start {
		panic("RangeIterator: start should be before end")
	} else if step <= 0 {
		panic("RangeIterator: step should be positive")
	}

	return &rangeIterator[T]{start: start, end: end, step: step}
}

type rangeIterator[T lancetconstraints.Number] struct {
	start, end, step T
}

func (iter *rangeIterator[T]) HasNext() bool {
	if iter.start >= iter.end {
		return false
	}
	return true
}

func (iter *rangeIterator[T]) Next() (T, bool) {
	if iter.start >= iter.end {
		var zero T
		return zero, false
	}
	num := iter.start
	iter.start += iter.step
	return num, true
}

func (iter *rangeIterator[T]) Count() int {
	count := (iter.end - iter.start) / iter.step
	iter.start = iter.end
	return int(count)
}
