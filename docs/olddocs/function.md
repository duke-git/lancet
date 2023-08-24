# Function

Package function can control the flow of function execution and support part of functional programming.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/function/function.go](https://github.com/duke-git/lancet/blob/main/function/function.go)
-   [https://github.com/duke-git/lancet/blob/main/function/watcher.go](https://github.com/duke-git/lancet/blob/main/function/watcher.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/function"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [After](#After)
-   [Before](#Before)
-   [CurryFn](#CurryFn)
-   [Compose](#Compose)
-   [Debounced](#Debounced)
-   [Delay](#Delay)
-   [Schedule](#Schedule)
-   [Pipeline](#Pipeline)
-   [Watcher](#Watcher)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="After">After</span>

<p>Creates a function that invokes given func once it's called n or more times.</p>

<b>Signature:</b>

```go
func After(n int, fn any) func(args ...any) []reflect.Value
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    fn := function.After(2, func() {
        fmt.Println("hello")
    })

    fn()
    fn()

    // Output:
    // hello
}
```

### <span id="Before">Before</span>

<p>creates a function that invokes func once it's called less than n times.</p>

<b>Signature:</b>

```go
func Before(n int, fn any) func(args ...any) []reflect.Value
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
    "github.com/duke-git/lancet/v2/internal"
)

func main() {
    fn := function.Before(2, func() {
        fmt.Println("hello")
    })

    fn()
    fn()
    fn()
    fn()

    // Output:
    // hello
    // hello
}
```

### <span id="CurryFn">CurryFn</span>

<p>Make curry function.</p>

<b>Signature:</b>

```go
type CurryFn[T any] func(...T) T
func (cf CurryFn[T]) New(val T) func(...T) T
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    add := func(a, b int) int {
        return a + b
    }

    var addCurry function.CurryFn[int] = func(values ...int) int {
        return add(values[0], values[1])
    }
    add1 := addCurry.New(1)

    result := add1(2)

    fmt.Println(result)

    // Output:
    // 3
}
```

### <span id="Compose">Compose</span>

<p>Compose the function list from right to left, then return the composed function.</p>

<b>Signature:</b>

```go
func Compose[T any](fnList ...func(...T) T) func(...T) T
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    toUpper := func(strs ...string) string {
        return strings.ToUpper(strs[0])
    }
    toLower := func(strs ...string) string {
        return strings.ToLower(strs[0])
    }
    transform := function.Compose(toUpper, toLower)

    result := transform("aBCde")

    fmt.Println(result)

    // Output:
    // ABCDE
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
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    count := 0

    add := func() {
        count++
    }

    debouncedAdd := function.Debounced(add, 50*time.Microsecond)

    debouncedAdd()
    debouncedAdd()
    debouncedAdd()
    debouncedAdd()

    time.Sleep(100 * time.Millisecond)

    fmt.Println(count)

    debouncedAdd()

    time.Sleep(100 * time.Millisecond)

    fmt.Println(count)

    // Output:
    // 1
    // 2
}
```

### <span id="Delay">Delay</span>

<p>Invoke function after delayed time.</p>

<b>Signature:</b>

```go
func Delay(delay time.Duration, fn any, args ...any)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    var print = func(s string) {
        fmt.Println(s)
    }

    function.Delay(2*time.Second, print, "hello")

    // Output:
    // hello
}
```

### <span id="Schedule">Schedule</span>

<p>Invoke function every duration time, until close the returned bool chan.</p>

<b>Signature:</b>

```go
func Schedule(d time.Duration, fn any, args ...any) chan bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    count := 0

    increase := func() {
        count++
    }

    stop := function.Schedule(2*time.Second, increase)

    time.Sleep(2 * time.Second)
    close(stop)

    fmt.Println(count)

    // Output:
    // 2
}
```

### <span id="Pipeline">Pipeline</span>

<p>Pipeline takes a list of functions and returns a function whose param will be passed into
the functions one by one.</p>

<b>Signature:</b>

```go
func Pipeline[T any](funcs ...func(T) T) func(T) T
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    addOne := func(x int) int {
        return x + 1
    }
    double := func(x int) int {
        return 2 * x
    }
    square := func(x int) int {
        return x * x
    }

    fn := function.Pipeline(addOne, double, square)

    result := fn(2)

    fmt.Println(result)

    // Output:
    // 36
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
func NewWatcher() *Watcher
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
    "github.com/duke-git/lancet/v2/function"
)

func main() {
       w := function.NewWatcher()

    w.Start()

    longRunningTask()

    fmt.Println(w.excuting) //true

    w.Stop()

    eapsedTime := w.GetElapsedTime().Milliseconds()

    fmt.Println(eapsedTime)

    w.Reset()
}

func longRunningTask() {
    var slice []int64
    for i := 0; i < 10000000; i++ {
        slice = append(slice, int64(i))
    }
}

```
