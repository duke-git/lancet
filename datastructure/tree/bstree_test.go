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

func TestBSTree_Insert(t *testing.T) {
	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	bstree.Print()
}

func TestBSTree_PreOrderTraverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_PreOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	acturl := bstree.PreOrderTraverse()
	t.Log(acturl)
	assert.Equal([]int{6, 5, 2, 4, 7}, acturl)
}

func TestBSTree_PostOrderTraverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_PostOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	acturl := bstree.PostOrderTraverse()
	t.Log(acturl)
	assert.Equal([]int{5, 2, 4, 7, 6}, acturl)
}

func TestBSTree_InOrderTraverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_InOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	acturl := bstree.InOrderTraverse()
	t.Log(acturl)
	assert.Equal([]int{2, 4, 5, 6, 7}, acturl)
}

func TestBSTree_LevelOrderTraverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_LevelOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	bstree.Print()

	acturl := bstree.LevelOrderTraverse()
	t.Log(acturl)
	assert.Equal([]int{6, 5, 7, 2, 4}, acturl)
}

func TestBSTree_Delete(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_Delete")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	bstree.Print()

	bstree.Delete(4)
	bstree.Print()
	acturl1 := bstree.InOrderTraverse()
	t.Log(acturl1)
	assert.Equal([]int{2, 5, 6, 7}, acturl1)

	//todo
	// bstree.DeletetNode(6, comparator)
	// bstree.Print()
	// acturl2 := bstree.InOrderTraverse()
	// t.Log(acturl2)
	// assert.Equal([]int{2, 5, 7}, acturl2)
}

func TestBSTree_Depth(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_Depth")

	bstree := NewBSTree(6, &intComparator{})

	bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

	bstree.Print()

	assert.Equal(bstree.Depth(), 4)
}

func TestBSTree_IsSubTree(t *testing.T) {
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

	subTree.Print()

	assert.Equal(true, superTree.HasSubTree(subTree))
	assert.Equal(false, subTree.HasSubTree(superTree))
}
