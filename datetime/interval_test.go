package datetime

import (
	"testing"
	"time"
)

func TestGenerateDatetimesNonPositiveInterval(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, interval := range []string{"0s", "-1h"} {
		t.Run(interval, func(t *testing.T) {
			if _, err := GenerateDatetimesBetween(start, start.Add(time.Hour), "2006-01-02", interval); err == nil {
				t.Error("expected an error for a non-positive interval")
			}
		})
	}
}
