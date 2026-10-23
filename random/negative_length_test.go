package random

import "testing"

func TestUniqueRandomSlicesNonPositiveLength(t *testing.T) {
	for _, length := range []int{-1, 0} {
		if got := RandUniqueIntSlice(length, 1, 10); len(got) != 0 {
			t.Errorf("unexpected integers: %v", got)
		}
		if got := RandFloats(length, 1, 10, 2); len(got) != 0 {
			t.Errorf("unexpected floats: %v", got)
		}
	}
}
