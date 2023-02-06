# Function

function 函数包控制函数执行流程，包含部分函数式编程。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/function/function.go](https://github.com/duke-git/lancet/blob/main/function/function.go)
-   [https://github.com/duke-git/lancet/blob/main/function/watcher.go](https://github.com/duke-git/lancet/blob/main/function/watcher.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/function"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [After](#After)
-   [Before](#Before)
-   [CurryFn](#CurryFn)
-   [Compose](#Compose)
-   [Debounced](#Debounced)
-   [Delay](#Delay)
-   [Pipeline](#Pipeline)
-   [Watcher](#Watcher)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="After">After</span>

<p>创建一个函数，当他被调用n或更多次之后将马上触发fn</p>

<b>函数签名:</b>

```go
func After(n int, fn any) func(args ...any) []reflect.Value
```

<b>示例:</b>

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

<p>创建一个函数，调用次数不超过n次，之后再调用这个函数，将返回一次最后调用fn的结果</p>

<b>函数签名:</b>

```go
func Before(n int, fn any) func(args ...any) []reflect.Value
```

<b>示例:</b>

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

<p>创建柯里化函数</p>

<b>函数签名:</b>

```go
type CurryFn[T any] func(...T) T
func (cf CurryFn[T]) New(val T) func(...T) T
```

<b>示例:</b>

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

<p>从右至左组合函数列表fnList，返回组合后的函数</p>

<b>函数签名:</b>

```go
func Compose[T any](fnList ...func(...T) T) func(...T) T
```

<b>示例:</b>

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

<p>创建一个debounced函数，该函数延迟调用fn直到自上次调用debounced函数后等待持续时间过去。</p>

<b>函数签名:</b>

```go
func Debounced(fn func(), duration time.Duration) func()
```

<b>示例:</b>

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

<p>延迟delay时间后调用函数</p>

<b>函数签名:</b>

```go
func Delay(delay time.Duration, fn any, args ...any)
```

<b>示例:</b>

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

<p>每次持续时间调用函数，直到关闭返回的 bool chan</p>

<b>函数签名:</b>

```go
func Schedule(d time.Duration, fn any, args ...any) chan bool
```

<b>示例:</b>

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

<p>执行函数pipeline.</p>

<b>函数签名:</b>

```go
func Pipeline[T any](funcs ...func(T) T) func(T) T
```

<b>示例:</b>

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

<p>Watcher用于记录代码执行时间。可以启动/停止/重置手表定时器。获取函数执行的时间。</p>

<b>函数签名:</b>

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

<b>示例:</b>

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
