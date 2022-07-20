# Concurrency
Package concurrency contain some functions to support concurrent programming. eg, goroutine, channel, async.

<div STYLE="page-break-after: always;"></div>

## Source:

- [https://github.com/duke-git/lancet/blob/main/concurrency/channel.go](https://github.com/duke-git/lancet/blob/main/concurrency/channel.go)

<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/duke-git/lancet/v2/concurrency"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
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

## Documentation


## Channel
### <span id="NewChannel">NewChannel</span>
<p>return a Channel pointer instance.</p>

<b>Signature:</b>

```go
type Channel struct {}
func NewChannel() *Channel
```
<b>Example:</b>

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

<p>Link multiple channels into one channel until cancel the context.</p>

<b>Signature:</b>

```go
func (c *Channel) Bridge(ctx context.Context, chanStream <-chan <-chan any) <-chan any
```
<b>Example:</b>

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

<p>merge multiple channels into one channel until cancel the context.</p>

<b>Signature:</b>

```go
func (c *Channel) FanIn(ctx context.Context, channels ...<-chan any) <-chan any
```
<b>Example:</b>

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

<p>Return a chan, put param `values` into the chan repeatly until cancel the context.</p>

<b>Signature:</b>

```go
func (c *Channel) Repeat(ctx context.Context, values ...any) <-chan any
```
<b>Example:</b>

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

<p>Return a chan, excutes fn repeatly, and put the result into retruned chan until cancel context.</p>

<b>Signature:</b>

```go
func (c *Channel) RepeatFn(ctx context.Context, fn func() any) <-chan any
```
<b>Example:</b>

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

<p>Read one or more channels into one channel, will close when any readin channel is closed.</p>

<b>Signature:</b>

```go
func (c *Channel) Or(channels ...<-chan any) <-chan any
```
<b>Example:</b>

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

<p>Read a channel into another channel, will close until cancel context.</p>

<b>Signature:</b>

```go
func (c *Channel) OrDone(ctx context.Context, channel <-chan any) <-chan any
```
<b>Example:</b>

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

<p>Return a chan whose values are tahken from another chan until cancel context.</p>

<b>Signature:</b>

```go
func (c *Channel) Take(ctx context.Context, valueStream <-chan any, number int) <-chan any
```
<b>Example:</b>

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

<p>Split one chanel into two channels until cancel context.</p>

<b>Signature:</b>

```go
func (c *Channel) Tee(ctx context.Context, in <-chan any) (<-chan any, <-chan any)
```
<b>Example:</b>

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