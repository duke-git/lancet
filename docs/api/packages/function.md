# Function

function 函数包控制函数执行流程，包含部分函数式编程。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/function/function.go](https://github.com/duke-git/lancet/blob/main/function/function.go)
-   [https://github.com/duke-git/lancet/blob/main/function/predicate.go](https://github.com/duke-git/lancet/blob/main/function/predicate.go)
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
-   [Schedule](#Schedule)
-   [Pipeline](#Pipeline)
-   [Watcher](#Watcher)
-   [And](#And)
-   [Or](#Or)
-   [Negate](#Negate)
-   [Nor](#Nor)
-   [Xnor](#Xnor)
-   [Nand](#Nand)
-   [AcceptIf](#AcceptIf)


<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="After">After</span>

<p>创建一个函数，当他被调用n或更多次之后将马上触发fn</p>

<b>函数签名:</b>

```go
func After(n int, fn any) func(args ...any) []reflect.Value
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/eRD5k2vzUVX)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/0HqUDIFZ3IL)</span></b>

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

<p>创建柯里化函数</p>

<b>函数签名:</b>

```go
type CurryFn[T any] func(...T) T
func (cf CurryFn[T]) New(val T) func(...T) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/5HopfDwANKX)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/KKfugD4PKYF)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/absuEGB_GN7)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Ivtc2ZE-Tye)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/hbON-Xeyn5N)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/mPdUVvj6HD6)</span></b>

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
func (w *Watcher) Start()
func (w *Watcher) Stop()
func (w *Watcher) Reset()
func (w *Watcher) GetElapsedTime() time.Duration

```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/l2yrOpCLd1I)</span></b>

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

<p>返回一个复合谓词判断函数，该判断函数表示一组谓词的逻辑and操作。只有当所有谓词判断函数对于给定的值都返回true时，返回true, 否则返回false。</p>

<b>函数签名:</b>

```go
func And[T any](predicates ...func(T) bool) func(T) bool
```

<b>示例:</b>

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

<p>返回一个复合谓词判断函数，该判断函数表示一组谓词的逻辑or操作。只有当所有谓词判断函数对于给定的值都返回false时，返回false, 否则返回true。</p>

<b>函数签名:</b>

```go
func Or[T any](predicates ...func(T) bool) func(T) bool
```

<b>示例:</b>

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

<p>返回一个谓词函数，该谓词函数表示当前谓词的逻辑否定。</p>

<b>函数签名:</b>

```go
func Negate[T any](predicate func(T) bool) func(T) bool
```

<b>示例:</b>

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

<p>返回一个组合谓词函数，表示给定值上所有谓词逻辑非或 (nor) 的结果。只有当所有谓词函数对给定值都返回false时，该组合谓词函数才返回true。</p>

<b>函数签名:</b>

```go
func Nor[T any](predicates ...func(T) bool) func(T) bool
```

<b>示例:</b>

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

<p>返回一个复合谓词函数，表示给定谓词函数列表的逻辑非与 (NAND)。仅当列表中所有函数对给定参数返回false时，才返回true，否则返回false。</p>

<b>函数签名:</b>

```go
func Nand[T any](predicates ...func(T) bool) func(T) bool
```

<b>示例:</b>

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

<p>返回一个复合谓词函数，表示给定一组谓词函数的逻辑异或 (XNOR)。只有当所有 谓词函数对给参数都返回true或false时，该谓词函数才返回true。</p>

<b>函数签名:</b>

```go
func Xnor[T any](predicates ...func(T) bool) func(T) bool
```

<b>示例:</b>

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

### <span id="AcceptIf">AcceptIf</span>

<p>TBD</p>

<b>函数签名:</b>

```go
func AcceptIf[T any](predicate func(T) bool, apply func(T) T) func(T) (T, bool)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/function"
)

func main() {

	adder := AcceptIf(
		And(
			func(x int) bool {
				return x > 10
			}, func(x int) bool {
				return x%2 == 0
			}),
		func(x int) int {
			return x + 1
		},
	)

	result, ok := adder(20)
	fmt.Println(result)
	fmt.Println(ok)

	result, ok = adder(21)
	fmt.Println(result)
	fmt.Println(ok)

	// Output:
	// 21
	// true
	// 0
	// false
}

```