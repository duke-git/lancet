package datastructure

import (
	"sort"
	"strconv"
	"testing"
)

func TestBSTreeDeleteReplacesRoot(t *testing.T) {
	for i, tt := range []struct {
		values []int
		remove int
		want   []int
	}{
		{nil, 4, nil},
		{[]int{2, 1, 3}, 4, []int{1, 2, 3}},
		{[]int{6, 5, 7}, 4, []int{5, 6, 7}},
		{[]int{2, 6, 1, 3, 5, 7}, 4, []int{1, 2, 3, 5, 6, 7}},
		{[]int{2, 6}, 9, []int{2, 4, 6}},
		{[]int{2, 6}, 2, []int{4, 6}},
		{[]int{2, 6, 1}, 2, []int{1, 4, 6}},
		{[]int{2, 6, 1, 3}, 2, []int{1, 3, 4, 6}},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			tree := NewBSTree(4, &intComparator{})
			for _, value := range tt.values {
				tree.Insert(value)
			}
			tree.Delete(tt.remove)
			checkBSTValues(t, tree, tt.want)
			tree.Delete(tt.remove)
			checkBSTValues(t, tree, tt.want)
			tree.Insert(8)
			want := append(append([]int(nil), tt.want...), 8)
			sort.Ints(want)
			checkBSTValues(t, tree, want)
		})
	}
}

func checkBSTValues(t *testing.T, tree *BSTree[int], want []int) {
	t.Helper()
	inOrder := tree.InOrderTraverse()
	if !sort.IntsAreSorted(inOrder) {
		t.Fatalf("in-order traversal is not sorted: %v", inOrder)
	}
	for name, got := range map[string][]int{
		"in-order":    inOrder,
		"pre-order":   tree.PreOrderTraverse(),
		"post-order":  tree.PostOrderTraverse(),
		"level-order": tree.LevelOrderTraverse(),
	} {
		sort.Ints(got)
		if len(got) != len(want) {
			t.Fatalf("%s traversal = %v, want values %v", name, got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("%s traversal = %v, want values %v", name, got, want)
			}
		}
	}
	if len(want) == 0 && tree.Depth() != 0 {
		t.Fatalf("empty tree depth = %d", tree.Depth())
	}
}
