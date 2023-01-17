// Copyright 2023 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package stream implements a sequence of elements supporting sequential and parallel aggregate operations.
// this package is a experiment to explore if stream in go can work as the way java does. it's complete, but not
// powerful like other libs
package stream

import "golang.org/x/exp/constraints"

// A stream should implements methods:
// type StreamI[T any] interface {

// 	// part methods of Java Stream Specification.
// 	Distinct() StreamI[T]
// 	Filter(predicate func(item T) bool) StreamI[T]
// 	FlatMap(mapper func(item T) StreamI[T]) StreamI[T]
// 	Map(mapper func(item T) T) StreamI[T]
// 	Peek(consumer func(item T)) StreamI[T]

// 	Sort(less func(a, b T) bool) StreamI[T]
// 	Max(less func(a, b T) bool) (T, bool)
// 	Min(less func(a, b T) bool) (T, bool)

// 	Limit(maxSize int) StreamI[T]
// 	Skip(n int64) StreamI[T]

// 	AllMatch(predicate func(item T) bool) bool
// 	AnyMatch(predicate func(item T) bool) bool
// 	NoneMatch(predicate func(item T) bool) bool
// 	ForEach(consumer func(item T))
// 	Reduce(accumulator func(a, b T) T) (T, bool)
// 	Count() int

// 	FindFirst() (T, bool)

// 	ToSlice() []T

// 	// part of methods custom extension
// 	Reverse() StreamI[T]
// 	Range(start, end int64) StreamI[T]
// 	Concat(streams ...StreamI[T]) StreamI[T]
// }

type stream[T any] struct {
	source []T
}

// FromSlice create stream from slice.
func FromSlice[T any](source []T) stream[T] {
	return stream[T]{source: source}
}

// FromRange create a number stream from start to end. both start and end are included. [start, end]
func FromRange[T constraints.Integer | constraints.Float](start, end, step T) stream[T] {
	if end < start {
		panic("stream.FromRange: param start should be before param end")
	} else if step <= 0 {
		panic("stream.FromRange: param step should be positive")
	}

	l := int((end-start)/step) + 1
	source := make([]T, l, l)

	for i := 0; i < l; i++ {
		source[i] = start + (T(i) * step)
	}

	return stream[T]{source: source}
}

// ToSlice return the elements in the stream.
func (s stream[T]) ToSlice() []T {
	return s.source
}
