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
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestMapIterator(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMapIterator")

	iter := FromSlice([]int{1, 2, 3, 4})

	iter = Map(iter, func(n int) int { return n / 2 })

	result := ToSlice(iter)
	assert.Equal([]int{0, 1, 1, 2}, result)
}

func TestFilterIterator(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFilterIterator")

	iter := FromSlice([]int{1, 2, 3, 4})

	iter = Filter(iter, func(n int) bool { return n < 3 })

	result := ToSlice(iter)
	assert.Equal([]int{1, 2}, result)
}

func TestJoinIterator(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestJoinIterator")

	iter1 := FromSlice([]int{1, 2})
	iter2 := FromSlice([]int{3, 4})

	iter := Join(iter1, iter2)

	item, ok := iter.Next()
	assert.Equal(1, item)
	assert.Equal(true, ok)

	assert.Equal([]int{2, 3, 4}, ToSlice(iter))
}

func TestReduce(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestReduce")

	iter := FromSlice([]int{1, 2, 3, 4})
	sum := Reduce(iter, 0, func(a, b int) int { return a + b })
	assert.Equal(10, sum)
}

func TestTakeIterator(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTakeIterator")

	iter := FromSlice([]int{1, 2, 3, 4, 5})

	iter = Take(iter, 3)

	result := ToSlice(iter)
	assert.Equal([]int{1, 2, 3}, result)
}
