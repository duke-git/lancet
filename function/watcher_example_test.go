package function

import "fmt"

func ExampleWatcher() {
	w := NewWatcher()

	w.Start()

	longRunningTask()

	w.Stop()

	// eapsedTime := w.GetElapsedTime().Milliseconds()

	fmt.Println("foo")

	w.Reset()

	// Output:
	// foo
}
