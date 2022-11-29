// Copyright 2022 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package iterator implements some feature of C++ STL iterators
package iterator

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestSliceIterator(t *testing.T) {
	assert := internal.NewAssert(t, "TestSliceIterator")

	// HashNext
	t.Run("slice iterator HasNext: ", func(t *testing.T) {
		iter1 := FromSlice([]int{1, 2, 3, 4})
		for {
			item, _ := iter1.Next()

			if item == 4 {
				assert.Equal(false, iter1.HasNext())
				break
			} else {
				assert.Equal(true, iter1.HasNext())
			}
		}

		iter2 := FromSlice([]int{})
		assert.Equal(false, iter2.HasNext())
	})

	//Next
	t.Run("slice iterator Next: ", func(t *testing.T) {
		iter1 := FromSlice([]int{1, 2, 3, 4})
		for i := 0; i < 4; i++ {
			item, ok := iter1.Next()
			if !ok {
				break
			}
			assert.Equal(i+1, item)
		}

		iter2 := FromSlice([]int{})
		_, ok := iter2.Next()
		assert.Equal(false, ok)
	})

}
