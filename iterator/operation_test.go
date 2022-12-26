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
	assert := internal.NewAssert(t, "TestMapIterator")

	iter := FromSlice([]int{1, 2, 3, 4})

	iter = Map(iter, func(n int) int { return n / 2 })

	result := ToSlice(iter)
	assert.Equal([]int{0, 1, 1, 2}, result)
}

func TestFilterIterator(t *testing.T) {
	assert := internal.NewAssert(t, "TestFilterIterator")

	iter := FromSlice([]int{1, 2, 3, 4})

	iter = Filter(iter, func(n int) bool { return n < 3 })

	result := ToSlice(iter)
	assert.Equal([]int{1, 2}, result)
}
