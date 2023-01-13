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
<p>Create a Channel pointer instance.</p>

<b>Signature:</b>

```go
type Channel[T any] struct
func NewChannel[T any]() *Channel[T]
```
<b>Example:</b>

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

<p>Link multiple channels into one channel until cancel the context.</p>

<b>Signature:</b>

```go
func (c *Channel[T]) Bridge(ctx context.Context, chanStream <-chan <-chan T) <-chan T
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

<p>Merge multiple channels into one channel until cancel the context.</p>

<b>Signature:</b>

```go
func (c *Channel[T]) FanIn(ctx context.Context, channels ...<-chan T) <-chan T
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

### <span id="Repeat">Repeat</span>

<p>Create channel, put values into the channel repeatly until cancel the context.</p>

<b>Signature:</b>

```go
func (c *Channel[T]) Repeat(ctx context.Context, values ...T) <-chan T
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

### <span id="Generate">Generate</span>

<p>Creates a channel, then put values into the channel.</p>

<b>Signature:</b>

```go
func (c *Channel[T]) Generate(ctx context.Context, values ...T) <-chan T
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

### <span id="RepeatFn">RepeatFn</span>

<p>Create a channel, excutes fn repeatly, and put the result into the channel, until close context.</p>

<b>Signature:</b>

```go
func (c *Channel[T]) RepeatFn(ctx context.Context, fn func() T) <-chan T
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

<p>Read one or more channels into one channel, will close when any readin channel is closed.</p>

<b>Signature:</b>

```go
func (c *Channel[T]) Or(channels ...<-chan T) <-chan T
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

<p>Read a channel into another channel, will close until cancel context.</p>

<b>Signature:</b>

```go
func (c *Channel[T]) OrDone(ctx context.Context, channel <-chan T) <-chan T
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

<p>Create a channel whose values are taken from another channel with limit number.</p>

<b>Signature:</b>

```go
func (c *Channel[T]) Take(ctx context.Context, valueStream <-chan T, number int) <-chan T
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

<p>Split one chanel into two channels, until cancel the context.</p>

<b>Signature:</b>

```go
func (c *Channel[T]) Tee(ctx context.Context, in <-chan T) (<-chan T, <-chan T)
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