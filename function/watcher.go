package function

import "time"

// Watcher is used for record code excution time
type Watcher struct {
	startTime int64
	stopTime  int64
	excuting  bool
}

// Start the watch timer.
func (w *Watcher) Start() {
	w.startTime = time.Now().UnixNano()
	w.excuting = true
}

// Stop the watch timer.
func (w *Watcher) Stop() {
	w.stopTime = time.Now().UnixNano()
	w.excuting = false
}
// GetElapsedTime get excute elapsed time.
func (w *Watcher) GetElapsedTime() time.Duration {
	if w.excuting {
		return time.Duration(time.Now().UnixNano() - w.startTime)
	} else {
		return time.Duration(w.stopTime - w.startTime)
	}
}