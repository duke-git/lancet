package datastructure_test

import (
	"reflect"
	"testing"

	tree "github.com/duke-git/lancet/v2/datastructure/tree"
)

type postorderComparator struct{}

func (postorderComparator) Compare(a, b any) int {
	x, y := a.(int), b.(int)
	if x < y {
		return -1
	}
	if x > y {
		return 1
	}
	return 0
}
func TestBSTreePostOrderVisitsDescendantsBeforeParents(t *testing.T) {
	for _, tc := range []struct {
		name        string
		input, want []int
	}{
		{"balanced", []int{4, 2, 6, 1, 3, 5, 7}, []int{1, 3, 2, 5, 7, 6, 4}},
		{"left chain", []int{5, 4, 3, 2}, []int{2, 3, 4, 5}},
		{"right chain", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"duplicate values", []int{5, 5, 5, 6}, []int{6, 5, 5, 5}},
		{"single root", []int{7}, []int{7}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tr := tree.NewBSTree(tc.input[0], postorderComparator{})
			for _, v := range tc.input[1:] {
				tr.Insert(v)
			}
			if got := tr.PostOrderTraverse(); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("postorder=%v want %v", got, tc.want)
			}
		})
	}
}
func TestBSTreePostOrderEmptyControl(t *testing.T) {
	tr := tree.NewBSTree(7, postorderComparator{})
	tr.Delete(7)
	if got := tr.PostOrderTraverse(); len(got) != 0 {
		t.Fatalf("empty traversal=%v", got)
	}
}
