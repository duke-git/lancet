package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func TestBSTree_PreOrderTraverse(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBSTree_PreOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	acturl := bstree.PreOrderTraverse()
	assert.Equal([]int{6, 5, 2, 4, 7}, acturl)
}

func TestBSTree_PostOrderTraverse(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBSTree_PostOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	acturl := bstree.PostOrderTraverse()
	assert.Equal([]int{5, 2, 4, 7, 6}, acturl)
}

func TestBSTree_InOrderTraverse(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBSTree_InOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	acturl := bstree.InOrderTraverse()
	assert.Equal([]int{2, 4, 5, 6, 7}, acturl)
}

func TestBSTree_LevelOrderTraverse(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBSTree_LevelOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	acturl := bstree.LevelOrderTraverse()
	assert.Equal([]int{6, 5, 7, 2, 4}, acturl)
}

func TestBSTree_Delete(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBSTree_Delete")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	bstree.Delete(4)

	acturl1 := bstree.InOrderTraverse()
	assert.Equal([]int{2, 5, 6, 7}, acturl1)
}

func TestBSTree_Depth(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBSTree_Depth")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	assert.Equal(bstree.Depth(), 4)
}

func TestBSTree_IsSubTree(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBSTree_IsSubTree")

	superTree := NewBSTree(8, &intComparator{})
	superTree.Insert(4)
	superTree.Insert(5)
	superTree.Insert(6)
	superTree.Insert(9)
	superTree.Insert(4)

	superTree.Print()

	subTree := NewBSTree(5, &intComparator{})
	subTree.Insert(4)
	subTree.Insert(6)

	assert.Equal(true, superTree.HasSubTree(subTree))
	assert.Equal(false, subTree.HasSubTree(superTree))
}
