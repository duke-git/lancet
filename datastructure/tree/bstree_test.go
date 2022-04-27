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

func TestBSTree_InsertNode(t *testing.T) {
	bstree := NewBSTree(6, &intComparator{})

	bstree.InsertNode(7)
	bstree.InsertNode(5)
	bstree.InsertNode(2)
	bstree.InsertNode(4)

	bstree.Print()
}

func TestBSTree_PreOrderTraverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_PreOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.InsertNode(7)
	bstree.InsertNode(5)
	bstree.InsertNode(2)
	bstree.InsertNode(4)

	acturl := bstree.PreOrderTraverse()
	t.Log(acturl)
	assert.Equal([]int{6, 5, 2, 4, 7}, acturl)
}

func TestBSTree_PostOrderTraverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_PostOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.InsertNode(7)
	bstree.InsertNode(5)
	bstree.InsertNode(2)
	bstree.InsertNode(4)

	acturl := bstree.PostOrderTraverse()
	t.Log(acturl)
	assert.Equal([]int{5, 2, 4, 7, 6}, acturl)
}

func TestBSTree_InOrderTraverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_InOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.InsertNode(7)
	bstree.InsertNode(5)
	bstree.InsertNode(2)
	bstree.InsertNode(4)

	acturl := bstree.InOrderTraverse()
	t.Log(acturl)
	assert.Equal([]int{2, 4, 5, 6, 7}, acturl)
}

func TestBSTree_LevelOrderTraverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_LevelOrderTraverse")

	bstree := NewBSTree(6, &intComparator{})

	bstree.InsertNode(7)
	bstree.InsertNode(5)
	bstree.InsertNode(2)
	bstree.InsertNode(4)

	bstree.Print()

	acturl := bstree.LevelOrderTraverse()
	t.Log(acturl)
	assert.Equal([]int{6, 5, 7, 2, 4}, acturl)
}

func TestBSTree_DeletetNode(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_DeletetNode")

	bstree := NewBSTree(6, &intComparator{})

	bstree.InsertNode(7)
	bstree.InsertNode(5)
	bstree.InsertNode(2)
	bstree.InsertNode(4)

	bstree.Print()

	bstree.DeletetNode(4)
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

	bstree.InsertNode(7)
	bstree.InsertNode(5)
	bstree.InsertNode(2)
	bstree.InsertNode(4)

	bstree.Print()

	assert.Equal(bstree.Depth(), 4)
}

func TestBSTree_IsSubTree(t *testing.T) {
	assert := internal.NewAssert(t, "TestBSTree_IsSubTree")

	superTree := NewBSTree(8, &intComparator{})
	superTree.InsertNode(4)
	superTree.InsertNode(5)
	superTree.InsertNode(6)
	superTree.InsertNode(9)
	superTree.InsertNode(4)

	superTree.Print()

	subTree := NewBSTree(5, &intComparator{})
	subTree.InsertNode(4)
	subTree.InsertNode(6)

	subTree.Print()

	assert.Equal(true, superTree.HasSubTree(subTree))
	assert.Equal(false, subTree.HasSubTree(superTree))
}
