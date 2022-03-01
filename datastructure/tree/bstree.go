package datastructure

import (
	"math"

	"github.com/duke-git/lancet/datastructure"
	"github.com/duke-git/lancet/lancetconstraints"
)

// BSTree is a binary search tree data structure in which each node has at most two children,
// which are referred to as the left child and the right child.
// In BSTree: leftNode < rootNode < rightNode
// type T should implements Compare function in lancetconstraints.Comparator interface.
type BSTree[T any] struct {
	root *datastructure.TreeNode[T]
}

// NewBSTree create a BSTree pointer
func NewBSTree[T any](rootData T) *BSTree[T] {
	root := datastructure.NewTreeNode(rootData)
	return &BSTree[T]{root}
}

// InsertNode insert data into BSTree
func (t *BSTree[T]) InsertNode(data T, comparator lancetconstraints.Comparator) {
	root := t.root
	newNode := datastructure.NewTreeNode(data)
	if root == nil {
		t.root = newNode
	} else {
		insertTreeNode(root, newNode, comparator)
	}
}

// DeletetNode delete data into BSTree
func (t *BSTree[T]) DeletetNode(data T, comparator lancetconstraints.Comparator) {
	deleteTreeNode(t.root, data, comparator)
}

// NodeLevel get node level in BSTree
func (t *BSTree[T]) NodeLevel(node *datastructure.TreeNode[T]) int {
	if node == nil {
		return 0
	}
	left := float64(t.NodeLevel(node.Left))
	right := float64(t.NodeLevel(node.Right))

	return int(math.Max(left, right)) + 1
}

// PreOrderTraverse traverse tree node in pre order
func (t *BSTree[T]) PreOrderTraverse() []T {
	return preOrderTraverse(t.root)
}

// PostOrderTraverse traverse tree node in post order
func (t *BSTree[T]) PostOrderTraverse() []T {
	return postOrderTraverse(t.root)
}

// InOrderTraverse traverse tree node in mid order
func (t *BSTree[T]) InOrderTraverse() []T {
	return inOrderTraverse(t.root)
}

// Depth returns the calculated depth of a binary saerch tree
func (t *BSTree[T]) Depth() int {
	return calculateDepth(t.root, 0)
}

// Print the bstree structure
func (t *BSTree[T]) Print() {
	maxLevel := t.NodeLevel(t.root)
	nodes := []*datastructure.TreeNode[T]{t.root}
	printTreeNodes(nodes, 1, maxLevel)
}
