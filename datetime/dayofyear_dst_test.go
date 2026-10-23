package datetime

import (
	"testing"
	"time"
	_ "time/tzdata"
)

func TestDayOfYearDaylightSaving(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatal(err)
	}
	for _, month := range []time.Month{time.January, time.July, time.December} {
		date := time.Date(2024, month, 1, 12, 0, 0, 0, loc)
		if got, want := DayOfYear(date), date.YearDay()-1; got != want {
			t.Errorf("%v: got %d, want %d", date, got, want)
		}
	}
}
