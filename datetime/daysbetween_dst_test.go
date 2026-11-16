package datetime

import (
	"testing"
	"time"
	_ "time/tzdata"
)

func TestDaysBetweenDaylightSaving(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal(err)
	}
	start := time.Date(2024, time.March, 1, 0, 0, 0, 0, loc)
	end := time.Date(2024, time.March, 31, 0, 0, 0, 0, loc)
	if got := DaysBetween(start, end); got != 30 {
		t.Errorf("got %d, want 30", got)
	}
}
