package datastructure_test

import (
	list "github.com/duke-git/lancet/v2/datastructure/list"
	"reflect"
	"testing"
)

func TestCopyOnWriteSubListPreservesReceiver(t *testing.T) {
	for _, tc := range []struct {
		name       string
		start, end int
		want       []int
	}{
		{"middle", 1, 3, []int{2, 3}}, {"empty range", 2, 2, []int{}}, {"prefix", 0, 2, []int{1, 2}}, {"suffix", 2, 4, []int{3, 4}}, {"full control", 0, 4, []int{1, 2, 3, 4}}, {"invalid control", 3, 1, []int{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := list.NewCopyOnWriteList([]int{1, 2, 3, 4})
			got := c.SubList(tc.start, tc.end)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("sublist=%v want%v", got, tc.want)
			}
			values := []int{}
			c.ForEach(func(v int) { values = append(values, v) })
			if c.Size() != 4 || !reflect.DeepEqual(values, []int{1, 2, 3, 4}) {
				t.Fatalf("SubList discarded source values: %v", values)
			}
			if len(got) > 0 {
				got[0] = 99
			}
			if *c.Get(0) != 1 || *c.Get(1) != 2 {
				t.Fatal("returned slice aliases receiver")
			}
			c.Add(5)
			if c.Size() != 5 || *c.Get(4) != 5 {
				t.Fatal("later append changed")
			}
		})
	}
}
