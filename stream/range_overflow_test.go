package stream

import "testing"

func TestFromRangeSmallIntegerSpan(t *testing.T) {
	got := FromRange[int8](-120, 120, 10).ToSlice()
	if len(got) != 25 || got[0] != -120 || got[len(got)-1] != 120 {
		t.Errorf("unexpected range: %v", got)
	}
	if got := FromRange[uint8](250, 255, 5).ToSlice(); len(got) != 2 || got[1] != 255 {
		t.Errorf("unexpected unsigned range: %v", got)
	}
}
