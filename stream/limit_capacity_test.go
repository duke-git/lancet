package stream

import (
	"math"
	"testing"
)

func TestLimitLargerThanSource(t *testing.T) {
	got := Of(1, 2).Limit(math.MaxInt).ToSlice()
	if len(got) != 2 || got[0] != 1 || got[1] != 2 {
		t.Errorf("unexpected limited stream: %v", got)
	}
}
