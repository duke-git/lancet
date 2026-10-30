package datastructure_test

import (
	"reflect"
	"testing"
	"time"

	list "github.com/duke-git/lancet/v2/datastructure/list"
)

func TestCopyOnWriteSortReleasesLock(t *testing.T) {
	for _, tc := range []struct {
		name         string
		values       []int
		panicCompare bool
		sort         bool
	}{
		{"sorted values", []int{3, 1, 2}, false, true},
		{"empty values", []int{}, false, true},
		{"comparator panic", []int{3, 1, 2}, true, true},
		{"without sort control", []int{1, 2, 3}, false, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := list.NewCopyOnWriteList(tc.values)
			marker := new(int)
			if tc.sort {
				func() {
					defer func() {
						r := recover()
						if tc.panicCompare {
							if r != marker {
								t.Errorf("panic identity changed: %v", r)
							}
						} else if r != nil {
							t.Errorf("unexpected panic: %v", r)
						}
					}()
					c.Sort(func(a, b int) bool {
						if tc.panicCompare {
							panic(marker)
						}
						return a < b
					})
				}()
			}
			if tc.sort && !tc.panicCompare {
				got := []int{}
				c.ForEach(func(v int) { got = append(got, v) })
				want := []int{1, 2, 3}
				if len(tc.values) == 0 {
					want = []int{}
				}
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("sort got=%v want=%v", got, want)
				}
			}
			done := make(chan struct{})
			go func() { c.Add(4); close(done) }()
			select {
			case <-done:
				if c.Size() != len(tc.values)+1 {
					t.Fatal("next mutation did not persist")
				}
			case <-time.After(time.Second):
				t.Fatal("Sort left the mutation lock held")
			}
		})
	}
}
