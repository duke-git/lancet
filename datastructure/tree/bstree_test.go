package datastructure

import "testing"

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
	bstree := NewBSTree(6)

	comparator := &intComparator{}
	bstree.Insert(7, comparator)
	bstree.Insert(5, comparator)
	bstree.Insert(2, comparator)
	bstree.Insert(4, comparator)

	bstree.Print()
}
