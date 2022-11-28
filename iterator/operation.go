// Copyright 2021 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package iterator implements some feature of C++ STL iterators
package iterator

// Map creates a new iterator which applies a function to all items of input iterator.
func Map[T any, U any](iter Iterator[T], iteratee func(item T) U) Iterator[U] {
	return &mapIterator[T, U]{
		iter:     iter,
		iteratee: iteratee,
	}
}

type mapIterator[T any, U any] struct {
	iter     Iterator[T]
	iteratee func(T) U
}

func (mr *mapIterator[T, U]) HasNext() bool {
	return mr.iter.HasNext()
}

func (mr *mapIterator[T, U]) Next() (U, bool) {
	var zero U
	item, ok := mr.iter.Next()
	if !ok {
		return zero, false
	}
	return mr.iteratee(item), true
}

// Filter creates a new iterator that returns only the items that pass specified predicate function.
func Filter[T any](iter Iterator[T], predicateFunc func(item T) bool) Iterator[T] {
	return &filterIterator[T]{iter: iter, predicateFunc: predicateFunc}
}

type filterIterator[T any] struct {
	iter          Iterator[T]
	predicateFunc func(T) bool
}

func (fr *filterIterator[T]) Next() (T, bool) {
	for item, ok := fr.iter.Next(); ok; item, ok = fr.iter.Next() {
		if fr.predicateFunc(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

func (fr *filterIterator[T]) HasNext() bool {
	return fr.iter.HasNext()
}
