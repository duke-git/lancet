package function

import "time"

// Watcher is used for record code excution time
// Play: Todo
type Watcher struct {
	startTime int64
	stopTime  int64
	excuting  bool
}

// Start the watch timer.
func NewWatcher() *Watcher {
	return &Watcher{}
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
	}
	return time.Duration(w.stopTime - w.startTime)
}

// Reset the watch timer.
func (w *Watcher) Reset() {
	w.startTime = 0
	w.stopTime = 0
	w.excuting = false
}
