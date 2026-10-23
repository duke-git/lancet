package datetime

import (
	"testing"
	"time"
)

func TestEmptyTimezoneOptions(t *testing.T) {
	empty := []string{}
	now := time.Now()
	if got := FormatTimeToStr(now, "yyyy-mm-dd", empty...); got != FormatTimeToStr(now, "yyyy-mm-dd") {
		t.Errorf("unexpected date: %s", got)
	}
	if _, err := FormatStrToTime("2026-01-02", "yyyy-mm-dd", empty...); err != nil {
		t.Fatal(err)
	}
	if NowDateOrTime("yyyy", empty...) == "" {
		t.Error("empty date")
	}
	for _, fn := range []func(...string) int64{Timestamp, TimestampMilli, TimestampMicro, TimestampNano} {
		if fn(empty...) == 0 {
			t.Error("empty timestamp")
		}
	}
}
