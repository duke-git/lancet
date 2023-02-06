# Concurrency
并发包包含一些支持并发编程的功能。例如：goroutine, channel, async等。

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/concurrency/channel.go](https://github.com/duke-git/lancet/blob/main/concurrency/channel.go)

<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/v2/concurrency"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
### Channel
- [NewChannel](#NewChannel)
- [Bridge](#Bridge)
- [FanIn](#FanIn)
- [Generate](#Generate)
- [Or](#Or)
- [OrDone](#OrDone)
- [Repeat](#Repeat)
- [RepeatFn](#RepeatFn)
- [Take](#Take)
- [Tee](#Tee)

<div STYLE="page-break-after: always;"></div>

## 文档


### Channel
### <span id="NewChannel">NewChannel</span>
<p>返回一个Channel指针实例</p>

<b>函数签名:</b>

```go
type Channel[T any] struct
func NewChannel[T any]() *Channel[T]
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    c := concurrency.NewChannel[int]()
}
```

### <span id="Bridge">Bridge</span>

<p>将多个channel链接到一个channel，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel[T]) Bridge(ctx context.Context, chanStream <-chan <-chan T) <-chan T
```
<b>示例:</b>

```go
package main

import (
    "context"
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    c := concurrency.NewChannel[int]()
    genVals := func() <-chan <-chan int {
        out := make(chan (<-chan int))
        go func() {
            defer close(out)
            for i := 1; i <= 5; i++ {
                stream := make(chan int, 1)
                stream <- i
                close(stream)
                out <- stream
            }
        }()
        return out
    }

    for v := range c.Bridge(ctx, genVals()) {
        fmt.Println(v)
    }

    // Output:
    // 1
    // 2
    // 3
    // 4
    // 5
}
```

### <span id="FanIn">FanIn</span>

<p>将多个channel合并为一个channel，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel[T]) FanIn(ctx context.Context, channels ...<-chan T) <-chan T
```
<b>示例:</b>

```go
package main

import (
    "context"
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    c := concurrency.NewChannel[int]()
    channels := make([]<-chan int, 2)

    for i := 0; i < 2; i++ {
        channels[i] = c.Take(ctx, c.Repeat(ctx, i), 2)
    }

    chs := c.FanIn(ctx, channels...)

    for v := range chs {
        fmt.Println(v) //1 1 0 0 or 0 0 1 1
    }
}
```

### <span id="Generate">Generate</span>

<p>根据传入的值，生成channel.</p>

<b>函数签名:</b>

```go
func (c *Channel[T]) Generate(ctx context.Context, values ...T) <-chan T
```
<b>示例:</b>

```go
package main

import (
    "context"
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    c := concurrency.NewChannel[int]()
    intStream := c.Generate(ctx, 1, 2, 3)

    fmt.Println(<-intStream)
    fmt.Println(<-intStream)
    fmt.Println(<-intStream)

    // Output:
    // 1
    // 2
    // 3
}
```

### <span id="Repeat">Repeat</span>

<p>返回一个channel，将参数`values`重复放入channel，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel[T]) Repeat(ctx context.Context, values ...T) <-chan T
```
<b>示例:</b>

```go
package main

import (
    "context"
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    c := concurrency.NewChannel[int]()
    intStream := c.Take(ctx, c.Repeat(ctx, 1, 2), 4)

    for v := range intStream {
        fmt.Println(v)
    }

    // Output:
    // 1
    // 2
    // 1
    // 2
}
```

### <span id="RepeatFn">RepeatFn</span>

<p>返回一个channel，重复执行函数fn，并将结果放入返回的channel，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel[T]) RepeatFn(ctx context.Context, fn func() T) <-chan T
```
<b>示例:</b>

```go
package main

import (
    "context"
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    fn := func() string {
        return "hello"
    }

    c := concurrency.NewChannel[string]()
    intStream := c.Take(ctx, c.RepeatFn(ctx, fn), 3)

    for v := range intStream {
        fmt.Println(v)
    }
    // Output:
    // hello
    // hello
    // hello
}
```

### <span id="Or">Or</span>

<p>将一个或多个channel读取到一个channel中，当任何读取channel关闭时将结束读取。</p>

<b>函数签名:</b>

```go
func (c *Channel[T]) Or(channels ...<-chan T) <-chan T
```
<b>示例:</b>

```go
package main

import (
    "context"
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    sig := func(after time.Duration) <-chan any {
        c := make(chan any)
        go func() {
            defer close(c)
            time.Sleep(after)
        }()
        return c
    }

    start := time.Now()

    c := concurrency.NewChannel[any]()
    <-c.Or(
        sig(1*time.Second),
        sig(2*time.Second),
        sig(3*time.Second),
    )

    fmt.Println("done after %v", time.Since(start)) //1.003s
}
```

### <span id="OrDone">OrDone</span>

<p>将一个channel读入另一个channel，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel[T]) OrDone(ctx context.Context, channel <-chan T) <-chan T
```
<b>示例:</b>

```go
package main

import (
    "context"
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    c := concurrency.NewChannel[int]()
    intStream := c.Take(ctx, c.Repeat(ctx, 1), 3)

    for v := range c.OrDone(ctx, intStream) {
        fmt.Println(v)
    }

    // Output:
    // 1
    // 1
    // 1
}
```

### <span id="Take">Take</span>

<p>返回一个channel，其值从另一个channel获取，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel[T]) Take(ctx context.Context, valueStream <-chan T, number int) <-chan T
```
<b>示例:</b>

```go
package main

import (
    "context"
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    numbers := make(chan int, 5)
    numbers <- 1
    numbers <- 2
    numbers <- 3
    numbers <- 4
    numbers <- 5
    defer close(numbers)

    c := concurrency.NewChannel[int]()
    intStream := c.Take(ctx, numbers, 3)

    for v := range intStream {
        fmt.Println(v)
    }

    // Output:
    // 1
    // 2
    // 3
}
```

### <span id="Tee">Tee</span>

<p>将一个channel分成两个channel，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel[T]) Tee(ctx context.Context, in <-chan T) (<-chan T, <-chan T)
```
<b>示例:</b>

```go
package main

import (
    "context"
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    c := concurrency.NewChannel[int]()
    intStream := c.Take(ctx, c.Repeat(ctx, 1), 2)

    ch1, ch2 := c.Tee(ctx, intStream)

    for v := range ch1 {
        fmt.Println(v)
        fmt.Println(<-ch2)
    }
    // Output:
    // 1
    // 1
    // 1
    // 1
}
```