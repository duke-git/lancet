# Function
Package function can control the flow of function execution and support part of functional programming.

<div STYLE="page-break-after: always;"></div>

## Source:

[https://github.com/duke-git/lancet/blob/main/function/function.go](https://github.com/duke-git/lancet/blob/main/function/function.go)
[https://github.com/duke-git/lancet/blob/main/function/watcher.go](https://github.com/duke-git/lancet/blob/main/function/watcher.go)

<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/duke-git/lancet/function"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
- [After](#After)
- [Before](#Before)
- [Curry](#Curry)
- [Compose](#Compose)
- [Debounced](#Debounced)
- [Delay](#Delay)
- [Delay](#Delay)
- [Watcher](#Watcher)

<div STYLE="page-break-after: always;"></div>

## Documentation



### <span id="After">After</span>
<p>Creates a function that invokes given func once it's called n or more times.</p>

<b>Signature:</b>

```go
func After(n int, fn interface{}) func(args ...interface{}) []reflect.Value
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/function"
)

func main() {
	arr := []string{"a", "b"}
	f := function.After(len(arr), func(i int) int {
		fmt.Println("last print")
		return i
	})

	type cb func(args ...interface{}) []reflect.Value
	print := func(i int, s string, fn cb) {
		fmt.Printf("arr[%d] is %s \n", i, s)
		fn(i)
	}

	fmt.Println("arr is", arr)
	for i := 0; i < len(arr); i++ {
		print(i, arr[i], f)
	}

    //output:
    // arr is [a b]
    // arr[0] is a 
    // arr[1] is b 
    // last print
}
```



### <span id="Before">Before</span>

<p>creates a function that invokes func once it's called less than n times.</p>

<b>Signature:</b>

```go
func Before(n int, fn interface{}) func(args ...interface{}) []reflect.Value
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/function"
    "github.com/duke-git/lancet/internal"
)

func main() {
    assert := internal.NewAssert(t, "TestBefore")

	arr := []string{"a", "b", "c", "d", "e"}
	f := function.Before(3, func(i int) int {
		return i
	})

	var res []int64
	type cb func(args ...interface{}) []reflect.Value
	appendStr := func(i int, s string, fn cb) {
		v := fn(i)
		res = append(res, v[0].Int())
	}

	for i := 0; i < len(arr); i++ {
		appendStr(i, arr[i], f)
	}

	expected := []int64{0, 1, 2, 2, 2}
	assert.Equal(expected, res)
}
```



### <span id="Curry">Curry</span>

<p>Make a curry function.</p>

<b>Signature:</b>

```go
type Fn func(...interface{}) interface{}
func (f Fn) Curry(i interface{}) func(...interface{}) interface{}
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/function"
)

func main() {
    add := func(a, b int) int {
		return a + b
	}
	var addCurry function.Fn = func(values ...interface{}) interface{} {
		return add(values[0].(int), values[1].(int))
	}
	add1 := addCurry.Curry(1)
	result := add1(2)
    fmt.Println(result) //3
}
```



### <span id="Compose">Compose</span>

<p>Compose the function list from right to left, then return the composed function.</p>

<b>Signature:</b>

```go
func Compose(fnList ...func(...interface{}) interface{}) func(...interface{}) interface{}
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/function"
)

func main() {
    add1 := func(v ...interface{}) interface{} {
		return v[0].(int) + 1
	}
    add2 := func(v ...interface{}) interface{} {
		return v[0].(int) + 2
	}

    add3 := function.Compose(add1, add2)
	result := add3(1)

    fmt.Println(result) //4
}
```



### <span id="Debounced">Debounced</span>

<p>Creates a debounced function that delays invoking fn until after wait duration have elapsed since the last time the debounced function was invoked.</p>

<b>Signature:</b>

```go
func Debounced(fn func(), duration time.Duration) func()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/function"
)

func main() {
    count := 0
	add := func() {
		count++
	}

	debouncedAdd := function.Debounced(add, 50*time.Microsecond)
	function.debouncedAdd()
	function.debouncedAdd()
	function.debouncedAdd()
	function.debouncedAdd()

	time.Sleep(100 * time.Millisecond)
	fmt.Println(count) //1

	function.debouncedAdd()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(count) //2
}
```



### <span id="Delay">Delay</span>

<p>Invoke function after delayed time.</p>

<b>Signature:</b>

```go
func Delay(delay time.Duration, fn interface{}, args ...interface{})
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/function"
)

func main() {
	var print = func(s string) {
		fmt.Println(count) //test delay
	}
	function.Delay(2*time.Second, print, "test delay")
}
```



### <span id="Schedule">Schedule</span>

<p>Invoke function every duration time, until close the returned bool chan.</p>

<b>Signature:</b>

```go
func Schedule(d time.Duration, fn interface{}, args ...interface{}) chan bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/function"
)

func main() {
    var res []string
	appendStr := func(s string) {
		res = append(res, s)
	}

	stop := function.Schedule(1*time.Second, appendStr, "*")
	time.Sleep(5 * time.Second)
	close(stop)

	fmt.Println(res) //[* * * * *]
}
```



### <span id="Watcher">Watcher</span>

<p>Watcher is used for record code excution time. can start/stop/reset the watch timer. get the elapsed time of function execution.</p>

<b>Signature:</b>

```go
type Watcher struct {
	startTime int64
	stopTime  int64
	excuting  bool
}
func (w *Watcher) Start() //start the watcher
func (w *Watcher) Stop() //stop the watcher
func (w *Watcher) Reset() //reset the watcher
func (w *Watcher) GetElapsedTime() time.Duration //get the elapsed time of function execution
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/function"
)

func main() {
    w := &function.Watcher{}
	w.Start()

	longRunningTask()

	fmt.Println(w.excuting) //true

	w.Stop()

	eapsedTime := w.GetElapsedTime().Milliseconds()
	fmt.Println(eapsedTime)

	w.Reset()

	fmt.Println(w.excuting) //false

	fmt.Println(w.startTime) //0
	fmt.Println(w.stopTime) //0
}

func longRunningTask() {
	var slice []int64
	for i := 0; i < 10000000; i++ {
		slice = append(slice, int64(i))
	}
}

```



