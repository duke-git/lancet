// Copyright 2022 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package iterator provides a way to iterate over values stored in containers.
// note:
// 1. Full feature iterator is complicated, this package is just a experiment to explore how iterators could work in Go.
// 2. The functionality of this package is very simple and limited, may not meet the actual dev needs.
// 3. It is currently under development, unstable, and will not be completed for some time in the future.
// So, based on above factors, you may not use it in production. but, anyone is welcome to improve it.
// Hope that Go can support iterator in future. see https://github.com/golang/go/discussions/54245 and https://github.com/golang/go/discussions/56413
package iterator

import (
	"context"

	"golang.org/x/exp/constraints"
)

// Iterator supports iterating over a sequence of values of type `E`.
type Iterator[T any] interface {
	// Next checks if there is a next value in the iteration or not
	HasNext() bool
	// Next returns the next value in the iteration if there is one,
	// and reports whether the returned value is valid.
	// Once Next returns ok==false, the iteration is over,
	// and all subsequent calls will return ok==false.
	Next() (item T, ok bool)
}

// ResettableIterator supports to reset the iterator
type ResettableIterator[T any] interface {
	Iterator[T]
	// Reset allows for the iteration process over a sequence to be restarted from the beginning.
	// It enables reusing the iterator for multiple traversals without needing to recreate it.
	Reset()
}

// StopIterator is an interface for stopping Iterator.
type StopIterator[T any] interface {
	Iterator[T]

	// Stop indicates that the iterator will no longer be used.
	// After a call to Stop, future calls to Next may panic.
	// Stop may be called multiple times;
	// all calls after the first will have no effect.
	Stop()
}

// DeleteIter is an Iter that implements a Delete method.
type DeleteIterator[T any] interface {
	Iterator[T]

	// Delete deletes the current iterator element;
	// that is, the one returned by the last call to Next.
	// Delete should panic if called before Next or after
	// Next returns false.
	Delete()
}

// SetIterator is an Iter that implements a Set method.
type SetIterator[T any] interface {
	Iterator[T]

	// Set replaces the current iterator element with v.
	// Set should panic if called before Next or after
	// Next returns false.
	Set(v T)
}

// PrevIterator is an iterator with a Prev method.
type PrevIterator[T any] interface {
	Iterator[T]

	// Prev moves the iterator to the previous position.
	// After calling Prev, Next will return the value at
	// that position in the container. For example, after
	//   it.Next() returning (v, true)
	//   it.Prev()
	// another call to it.Next will again return (v, true).
	// Calling Prev before calling Next may panic.
	// Calling Prev after Next returns false will move
	// to the last element, or, if there are no elements,
	// to the iterator's initial state.
	Prev()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Functions that create an Iterator from some other type.                                       //
////////////////////////////////////////////////////////////////////////////////////////////////////

// FromSlice returns an iterator over a slice of data.
func FromSlice[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{slice: slice, index: -1}
}

func ToSlice[T any](iter Iterator[T]) []T {
	result := []T{}
	for item, ok := iter.Next(); ok; item, ok = iter.Next() {
		result = append(result, item)
	}
	return result
}

type SliceIterator[T any] struct {
	slice []T
	index int
}

func (iter *SliceIterator[T]) HasNext() bool {
	return iter.index < len(iter.slice)-1
}

func (iter *SliceIterator[T]) Next() (T, bool) {
	iter.index++

	ok := iter.index >= 0 && iter.index < len(iter.slice)

	var item T
	if ok {
		item = iter.slice[iter.index]
	}

	return item, ok
}

// Prev implements PrevIterator.
func (iter *SliceIterator[T]) Prev() {
	if iter.index == -1 {
		panic("Next function should be called Prev")
	}
	if iter.HasNext() {
		iter.index--
	} else {
		iter.index = len(iter.slice) - 1
	}
}

// Set implements SetIterator.
func (iter *SliceIterator[T]) Set(value T) {
	if iter.index == -1 {
		panic("Next function should be called Set")
	}
	if iter.index >= len(iter.slice) || len(iter.slice) == 0 {
		panic("No element in current iterator")
	}
	iter.slice[iter.index] = value
}

func (iter *SliceIterator[T]) Reset() {
	iter.index = -1
}

// FromRange creates a iterator which returns the numeric range between start inclusive and end
// exclusive by the step size. start should be less than end, step shoud be positive.
func FromRange[T constraints.Integer | constraints.Float](start, end, step T) *RangeIterator[T] {
	if end < start {
		panic("RangeIterator: start should be before end")
	} else if step <= 0 {
		panic("RangeIterator: step should be positive")
	}

	return &RangeIterator[T]{start: start, end: end, step: step, current: start}
}

type RangeIterator[T constraints.Integer | constraints.Float] struct {
	start, end, step, current T
}

func (iter *RangeIterator[T]) HasNext() bool {
	return iter.current < iter.end
}

func (iter *RangeIterator[T]) Next() (T, bool) {
	if iter.current >= iter.end {
		var zero T
		return zero, false
	}
	num := iter.current
	iter.current += iter.step
	return num, true
}

func (iter *RangeIterator[T]) Reset() {
	iter.current = iter.start
}

// FromChannel creates an iterator which returns items received from the provided channel.
// The iteration continues until the channel is closed.
func FromChannel[T any](channel <-chan T) *ChannelIterator[T] {
	return &ChannelIterator[T]{channel: channel}
}

type ChannelIterator[T any] struct {
	channel <-chan T
}

func (iter *ChannelIterator[T]) Next() (T, bool) {
	item, ok := <-iter.channel
	return item, ok
}

func (iter *ChannelIterator[T]) HasNext() bool {
	return len(iter.channel) == 0
}

// ToChannel create a new goroutine to pull items from the channel iterator to the returned channel.
func ToChannel[T any](ctx context.Context, iter Iterator[T], buffer int) <-chan T {
	result := make(chan T, buffer)

	go func() {
		defer close(result)

		for item, ok := iter.Next(); ok; item, ok = iter.Next() {
			select {
			case result <- item:
			case <-ctx.Done():
				return
			}
		}
	}()

	return result
}
