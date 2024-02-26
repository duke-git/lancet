# Function

Package function can control the flow of function execution and support part of functional programming.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/function/function.go](https://github.com/duke-git/lancet/blob/main/function/function.go)
-   [https://github.com/duke-git/lancet/blob/main/function/predicate.go](https://github.com/duke-git/lancet/blob/main/function/predicate.go)
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
-   [And](#And)
-   [Or](#Or)
-   [Negate](#Negate)
-   [Nor](#Nor)
-   [Xnor](#Xnor)
-   [Nand](#Nand)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="After">After</span>

<p>Creates a function that invokes given func once it's called n or more times.</p>

<b>Signature:</b>

```go
func After(n int, fn any) func(args ...any) []reflect.Value
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/eRD5k2vzUVX)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/0HqUDIFZ3IL)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/5HopfDwANKX)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/KKfugD4PKYF)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/absuEGB_GN7)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/Ivtc2ZE-Tye)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/hbON-Xeyn5N)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/mPdUVvj6HD6)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/l2yrOpCLd1I)</span></b>

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


### <span id="And">And</span>

<p>Returns a composed predicate that represents the logical AND of a list of predicates. It evaluates to true only if all predicates evaluate to true for the given value.</p>

<b>Signature:</b>

```go
func And[T any](predicates ...func(T) bool) func(T) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    isNumericAndLength5 := function.And(
        func(s string) bool { return strings.ContainsAny(s, "0123456789") },
        func(s string) bool { return len(s) == 5 },
    )

    fmt.Println(isNumericAndLength5("12345"))
    fmt.Println(isNumericAndLength5("1234"))
    fmt.Println(isNumericAndLength5("abcde"))

    // Output:
    // true
    // false
    // false
}
```

### <span id="Or">Or</span>

<p>Returns a composed predicate that represents the logical OR of a list of predicates. It evaluates to true if at least one of the predicates evaluates to true for the given value.</p>

<b>Signature:</b>

```go
func Or[T any](predicates ...func(T) bool) func(T) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    containsDigitOrSpecialChar := function.Or(
        func(s string) bool { return strings.ContainsAny(s, "0123456789") },
        func(s string) bool { return strings.ContainsAny(s, "!@#$%") },
    )

    fmt.Println(containsDigitOrSpecialChar("hello!"))
    fmt.Println(containsDigitOrSpecialChar("hello"))

    // Output:
    // true
    // false
}
```

### <span id="Negate">Negate</span>

<p>Returns a predicate that represents the logical negation of this predicate.</p>

<b>Signature:</b>

```go
func Negate[T any](predicate func(T) bool) func(T) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    // Define some simple predicates for demonstration
    isUpperCase := func(s string) bool {
        return strings.ToUpper(s) == s
    }
    isLowerCase := func(s string) bool {
        return strings.ToLower(s) == s
    }
    isMixedCase := function.Negate(function.Or(isUpperCase, isLowerCase))

    fmt.Println(isMixedCase("ABC"))
    fmt.Println(isMixedCase("AbC"))

    // Output:
    // false
    // true
}
```


### <span id="Nor">Nor</span>

<p>Returns a composed predicate that represents the logical NOR of a list of predicates. It evaluates to true only if all predicates evaluate to false for the given value.</p>

<b>Signature:</b>

```go
func Nor[T any](predicates ...func(T) bool) func(T) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    match := function.Nor(
        func(s string) bool { return strings.ContainsAny(s, "0123456789") },
        func(s string) bool { return len(s) == 5 },
    )

    fmt.Println(match("dbcdckkeee"))


    match = function.Nor(
        func(s string) bool { return strings.ContainsAny(s, "0123456789") },
        func(s string) bool { return len(s) == 5 },
    )

    fmt.Println(match("0123456789"))

    // Output:
    // true
    // false
}
```

### <span id="Nand">Nand</span>

<p>Returns a composed predicate that represents the logical NAND of a list of predicates. It evaluates to true only if all predicates evaluate to false for the given value.</p>

<b>Signature:</b>

```go
func Nand[T any](predicates ...func(T) bool) func(T) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    isNumericAndLength5 := function.Nand(
        func(s string) bool { return strings.ContainsAny(s, "0123456789") },
        func(s string) bool { return len(s) == 5 },
    )

    fmt.Println(isNumericAndLength5("12345"))
    fmt.Println(isNumericAndLength5("1234"))
    fmt.Println(isNumericAndLength5("abcdef"))

    // Output:
    // false
    // false
    // true
}
```

### <span id="Xnor">Xnor</span>

<p>Returns a composed predicate that represents the logical XNOR of a list of predicates. It evaluates to true only if all predicates evaluate to true or false for the given value.</p>

<b>Signature:</b>

```go
func Xnor[T any](predicates ...func(T) bool) func(T) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {
    isEven := func(i int) bool { return i%2 == 0 }
    isPositive := func(i int) bool { return i > 0 }

    match := function.Xnor(isEven, isPositive)

    fmt.Println(match(2))
    fmt.Println(match(-3))
    fmt.Println(match(3))

    // Output:
    // true
    // true
    // false
}
```