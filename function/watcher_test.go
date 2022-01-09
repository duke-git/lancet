package function

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestWatcher(t *testing.T) {
	assert := internal.NewAssert(t, "TestWatcher")

	w := &Watcher{}
	w.Start()

	longRunningTask()

	assert.Equal(true, w.excuting)

	w.Stop()

	eapsedTime := w.GetElapsedTime().Milliseconds()
	t.Log("Elapsed Time (milsecond)", eapsedTime)

	assert.Equal(false, w.excuting)

	w.Reset()
}

func longRunningTask() {
	var slice []int64
	for i := 0; i < 10000000; i++ {
		slice = append(slice, int64(i))
	}
}
