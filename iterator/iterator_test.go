// Copyright 2021 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package iterator implements some feature of C++ STL iterators
package iterator

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestSliceIterator(t *testing.T) {
	assert := internal.NewAssert(t, "TestSliceIterator")

	t.Run("slice iterator has next", func(t *testing.T) {
		iter := FromSlice([]int{1, 2, 3, 4})
		for {
			item, _ := iter.Next()
			if item < 4 {
				assert.Equal(true, iter.HasNext())
			} else {
				assert.Equal(false, iter.HasNext())
				break
			}
		}
	})

	t.Run("slice iterator next", func(t *testing.T) {
		iter := FromSlice([]int{1, 2, 3, 4})
		for i := 0; i < 4; i++ {
			item, ok := iter.Next()
			if !ok {
				break
			}
			assert.Equal(i+1, item)
		}
	})
}
