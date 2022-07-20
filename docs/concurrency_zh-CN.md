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
<p>返回一个 Channel 指针实例</p>

<b>函数签名:</b>

```go
type Channel struct {}
func NewChannel() *Channel
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    c := concurrency.NewChannel()
}
```



### <span id="Bridge">Bridge</span>

<p>将多个通道链接到一个通道，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel) Bridge(ctx context.Context, chanStream <-chan <-chan any) <-chan any
```
<b>例子:</b>

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

	c := concurrency.NewChannel()
	genVals := func() <-chan <-chan any {
		chanStream := make(chan (<-chan any))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan any, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	index := 0
	for val := range c.Bridge(ctx, genVals()) {
		fmt.Printf("%v ", val) //0 1 2 3 4 5 6 7 8 9
	}
}
```




### <span id="FanIn">FanIn</span>

<p>将多个通道合并为一个通道，直到取消上下文</p>

<b>函数签名:</b>

```go
func (c *Channel) FanIn(ctx context.Context, channels ...<-chan any) <-chan any
```
<b>例子:</b>

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

	c := concurrency.NewChannel()
	channels := make([]<-chan any, 3)

	for i := 0; i < 3; i++ {
		channels[i] = c.Take(ctx, c.Repeat(ctx, i), 3)
	}

	mergedChannel := c.FanIn(ctx, channels...)

	for val := range mergedChannel {
		fmt.Println("\t%d\n", val) //1,2,1,0,0,1,0,2,2 (order not for sure)
	}
}
```


### <span id="Repeat">Repeat</span>

<p>返回一个chan，将参数`values`重复放入chan，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel) Repeat(ctx context.Context, values ...any) <-chan any
```
<b>例子:</b>

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

	c := concurrency.NewChannel()
	intStream := c.Take(ctx, c.Repeat(ctx, 1, 2), 5)

	for v := range intStream {
		fmt.Println(v) //1, 2, 1, 2, 1
	}
}
```




### <span id="RepeatFn">RepeatFn</span>

<p>返回一个chan，重复执行函数fn，并将结果放入返回的chan，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel) RepeatFn(ctx context.Context, fn func() any) <-chan any
```
<b>例子:</b>

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

	fn := func() any {
		s := "a"
		return s
	}
	c := concurrency.NewChannel()
	dataStream := c.Take(ctx, c.RepeatFn(ctx, fn), 3)

	for v := range dataStream {
		fmt.Println(v) //a, a, a
	}
}
```



### <span id="Or">Or</span>

<p>将一个或多个通道读取到一个通道中，当任何读取通道关闭时将结束读取。</p>

<b>函数签名:</b>

```go
func (c *Channel) Or(channels ...<-chan any) <-chan any
```
<b>例子:</b>

```go
package main

import (
    "context"
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
	sig := func(after time.Duration) <-chan any {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	c := concurrency.NewChannel()
	<-c.Or(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
		sig(5*time.Second),
	)

	fmt.Println("done after %v", time.Since(start)) //1.003s
}
```




### <span id="OrDone">OrDone</span>

<p>将一个通道读入另一个通道，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel) OrDone(ctx context.Context, channel <-chan any) <-chan any
```
<b>例子:</b>

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

	c := concurrency.NewChannel()
	intStream := c.Take(ctx, c.Repeat(ctx, 1), 3)

	for val := range c.OrDone(ctx, intStream) {
		fmt.Println(val)  //1
	}
}
```




### <span id="Take">Take</span>

<p>返回一个chan，其值从另一个chan获取，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel) Take(ctx context.Context, valueStream <-chan any, number int) <-chan any
```
<b>例子:</b>

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

	numbers := make(chan any, 5)
	numbers <- 1
	numbers <- 2
	numbers <- 3
	numbers <- 4
	numbers <- 5
	defer close(numbers)

	c := concurrency.NewChannel()
	intStream := c.Take(ctx, numbers, 3)

	for val := range intStream {
		fmt.Println(val) //1, 2, 3
	}
}
```



### <span id="Tee">Tee</span>

<p>将一个通道分成两个通道，直到取消上下文。</p>

<b>函数签名:</b>

```go
func (c *Channel) Tee(ctx context.Context, in <-chan any) (<-chan any, <-chan any)
```
<b>例子:</b>

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

	c := concurrency.NewChannel()
	inStream := c.Take(ctx, c.Repeat(ctx, 1, 2), 4)

	out1, out2 := c.Tee(ctx, inStream)
	for val := range out1 {
		fmt.Println(val) //1
		fmt.Println(<-out2) //1
	}
}
```