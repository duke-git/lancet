package function

import (
	"testing"
)

func TestWatcher(t *testing.T) {
	w := &Watcher{}
	w.Start()

	longRunningTask()

	if !w.excuting {
		t.FailNow()
	}

	w.Stop()

	eapsedTime := w.GetElapsedTime().Milliseconds()
	t.Log("Elapsed Time (milsecond)", eapsedTime)

	if w.excuting {
		t.FailNow()
	}
}

func longRunningTask() {
	var slice []int64
	for i := 0; i < 10000000; i++ {
		slice = append(slice, int64(i))
	}
}
