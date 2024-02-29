// Copyright 2022 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package iterator implements some feature of C++ STL iterators
package iterator

import (
	"context"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestSliceIterator(t *testing.T) {
	t.Parallel()

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

	// Reset
	t.Run("slice iterator Reset: ", func(t *testing.T) {
		iter1 := FromSlice([]int{1, 2, 3, 4})
		for i := 0; i < 4; i++ {
			item, ok := iter1.Next()
			if !ok {
				break
			}
			assert.Equal(i+1, item)
		}

		iter1.Reset()

		for i := 0; i < 4; i++ {
			item, ok := iter1.Next()
			if !ok {
				break
			}
			assert.Equal(i+1, item)
		}
	})

	t.Run("slice iterator ToSlice: ", func(t *testing.T) {
		iter := FromSlice([]int{1, 2, 3, 4})
		item, _ := iter.Next()
		assert.Equal(1, item)

		data := ToSlice[int](iter)
		assert.Equal([]int{2, 3, 4}, data)
	})
}

func TestRangeIterator(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRangeIterator")

	t.Run("range iterator: ", func(t *testing.T) {
		iter := FromRange(1, 4, 1)

		item, ok := iter.Next()
		assert.Equal(1, item)
		assert.Equal(true, ok)

		item, ok = iter.Next()
		assert.Equal(2, item)
		assert.Equal(true, ok)

		item, ok = iter.Next()
		assert.Equal(3, item)
		assert.Equal(true, ok)

		_, ok = iter.Next()
		assert.Equal(false, ok)
		assert.Equal(false, iter.HasNext())

		iter.Reset()

		item, ok = iter.Next()
		assert.Equal(1, item)
		assert.Equal(true, ok)

		item, ok = iter.Next()
		assert.Equal(2, item)
		assert.Equal(true, ok)

		item, ok = iter.Next()
		assert.Equal(3, item)
		assert.Equal(true, ok)

		_, ok = iter.Next()
		assert.Equal(false, ok)
		assert.Equal(false, iter.HasNext())
	})

	t.Run("range iterator reset: ", func(t *testing.T) {
		iter := FromRange(1, 4, 1)

		item, ok := iter.Next()
		assert.Equal(1, item)
		assert.Equal(true, ok)

		item, ok = iter.Next()
		assert.Equal(2, item)
		assert.Equal(true, ok)

		iter.Reset()

		item, ok = iter.Next()
		assert.Equal(1, item)
		assert.Equal(true, ok)

		item, ok = iter.Next()
		assert.Equal(2, item)
		assert.Equal(true, ok)

		item, ok = iter.Next()
		assert.Equal(3, item)
		assert.Equal(true, ok)

		_, ok = iter.Next()
		assert.Equal(false, ok)
		assert.Equal(false, iter.HasNext())
	})

}

func TestChannelIterator(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRangeIterator")

	var iter Iterator[int] = FromSlice([]int{1, 2, 3, 4})

	ctx, cancel := context.WithCancel(context.Background())
	iter = FromChannel(ToChannel(ctx, iter, 0))
	item, ok := iter.Next()
	assert.Equal(1, item)
	assert.Equal(true, ok)
	assert.Equal(true, iter.HasNext())

	cancel()

	_, ok = iter.Next()
	assert.Equal(false, ok)
}
