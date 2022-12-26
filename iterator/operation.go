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

// Join creates an iterator that join all elements of iters[0], then all elements of iters[1] and so on.
func Join[T any](iters ...Iterator[T]) Iterator[T] {
	return &joinIterator[T]{
		iters: iters,
	}
}

type joinIterator[T any] struct {
	iters []Iterator[T]
}

func (iter *joinIterator[T]) Next() (T, bool) {
	for len(iter.iters) > 0 {
		item, ok := iter.iters[0].Next()
		if ok {
			return item, true
		}
		iter.iters = iter.iters[1:]
	}
	var zero T
	return zero, false
}

func (iter *joinIterator[T]) HasNext() bool {
	if len(iter.iters) == 0 {
		return false
	}
	if len(iter.iters) == 1 {
		return iter.iters[0].HasNext()
	}

	result := iter.iters[0].HasNext()

	for i := 1; i < len(iter.iters); i++ {
		it := iter.iters[i]
		hasNext := it.HasNext()
		result = result || hasNext
	}

	return result
}
