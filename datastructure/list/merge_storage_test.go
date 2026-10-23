package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestMergePreservesBackingStorage(t *testing.T) {
	assert := internal.NewAssert(t, "TestMergePreservesBackingStorage")
	backing := []int{1, 99, 100}
	left := NewList(backing[:1])
	merged := left.Merge(NewList([]int{2, 3}))
	assert.Equal([]int{1, 2, 3}, merged.Data())
	assert.Equal([]int{1, 99, 100}, backing)
	merged.UpdateAt(0, 42)
	assert.Equal([]int{1}, left.Data())
}
