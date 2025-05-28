# Concurrency

Package concurrency contain some functions to support concurrent programming. eg, goroutine, channel.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/concurrency/channel.go](https://github.com/duke-git/lancet/blob/main/concurrency/channel.go)
-   [https://github.com/duke-git/lancet/blob/main/concurrency/keyed_locker.go](https://github.com/duke-git/lancet/blob/main/concurrency/keyed_locker.go)

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

-   [NewChannel](#NewChannel)
-   [Bridge](#Bridge)
-   [FanIn](#FanIn)
-   [Generate](#Generate)
-   [Or](#Or)
-   [OrDone](#OrDone)
-   [Repeat](#Repeat)
-   [RepeatFn](#RepeatFn)
-   [Take](#Take)
-   [Tee](#Tee)

### KeyedLocker

-   [NewKeyedLocker](#NewKeyedLocker)
-   [Do](#Do)
-   [NewRWKeyedLocker](#NewRWKeyedLocker)
-   [RLock](#RLock)
-   [Lock](#Lock)
-   [NewTryKeyedLocker](#NewTryKeyedLocker)
-   [TryLock](#TryLock)
-   [Unlock](#Unlock)

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

<b>Example: <span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/7aB4KyMMp9A)</span></b>

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

<b>Example: <span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/qmWSy1NVF-Y)</span></b>

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

<b>Example: <span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/2VYFMexEvTm)</span></b>

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

<b>Signature: <span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/k5N_ALVmYjE)</span></b>

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

<b>Example: <span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/7aB4KyMMp9A)</span></b>

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

<b>Signature: <span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/4J1zAWttP85)</span></b>

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

<b>Signature: <span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/Wqz9rwioPww)</span></b>

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

<b>Signature: <span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/lm_GoS6aDjo)</span></b>

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

<b>Signature: <span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/9Utt-1pDr2J)</span></b>

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

<b>Signature: <span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/3TQPKnCirrP)</span></b>

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

### KeyedLocker

### <span id="NewKeyedLocker">NewKeyedLocker</span>

<p>KeyedLocker is a simple implementation of a keyed locker that allows for non-blocking lock acquisition.</p>

<b>Signature:</b>

```go
func NewKeyedLocker[K comparable](ttl time.Duration) *KeyedLocker[K]
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/todo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    locker := concurrency.NewKeyedLocker[string](2 * time.Second)

    task := func() {
        fmt.Println("Executing task...")
        time.Sleep(1 * time.Second)
        fmt.Println("Task completed.")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    if err := locker.Do(ctx, "mykey", task); err != nil {
        log.Fatalf("Error executing task: %v\n", err)
    } else {
        fmt.Println("Task successfully executed.")
    }

    ctx2, cancel2 := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel2()

    if err := locker.Do(ctx2, "mykey", task); err != nil {
        log.Fatalf("Error executing task: %v\n", err)
    } else {
        fmt.Println("Task successfully executed.")
    }

    // Output:
    // Executing task...
    // Task completed.
    // Task successfully executed.
    // Executing task...
    // Task completed.
    // Task successfully executed.
}
```

### <span id="Do">Do</span>

<p>Acquires a lock for the specified key and executes the provided function.</p>

<b>Signature:</b>

```go
func (l *KeyedLocker[K]) Do(ctx context.Context, key K, fn func()) error
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/todo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    locker := concurrency.NewKeyedLocker[string](2 * time.Second)

    task := func() {
        fmt.Println("Executing task...")
        time.Sleep(1 * time.Second)
        fmt.Println("Task completed.")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    if err := locker.Do(ctx, "mykey", task); err != nil {
        log.Fatalf("Error executing task: %v\n", err)
    } else {
        fmt.Println("Task successfully executed.")
    }

    ctx2, cancel2 := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel2()

    if err := locker.Do(ctx2, "mykey", task); err != nil {
        log.Fatalf("Error executing task: %v\n", err)
    } else {
        fmt.Println("Task successfully executed.")
    }

    // Output:
    // Executing task...
    // Task completed.
    // Task successfully executed.
    // Executing task...
    // Task completed.
    // Task successfully executed.
}
```

### <span id="NewRWKeyedLocker">NewRWKeyedLocker</span>

<p>RWKeyedLocker is a read-write version of KeyedLocker.</p>

<b>Signature:</b>

```go
func NewRWKeyedLocker[K comparable](ttl time.Duration) *RWKeyedLocker[K]
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/todo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    locker := concurrency.NewRWKeyedLocker[string](2 * time.Second)

    // Simulate a key
    key := "resource_key"

    fn := func() {
        fmt.Println("Starting write operation...")
        // Simulate write operation, assuming it takes 2 seconds
        time.Sleep(200 * time.Millisecond)
        fmt.Println("Write operation completed!")
    }

    // Acquire the write lock and execute the operation
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    // Execute the lock operation with a 3-second timeout
    err := locker.Lock(ctx, key, fn)
    if err != nil {
        return
    }

    //output:
    //Starting write operation...
    //Write operation completed!
}
```

### <span id="RLock">RLock</span>

<p>Acquires a read lock for the specified key and executes the provided function.</p>

<b>Signature:</b>

```go
func (l *RWKeyedLocker[K]) RLock(ctx context.Context, key K, fn func()) error
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/todo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    locker := concurrency.NewRWKeyedLocker[string](2 * time.Second)

    // Simulate a key
    key := "resource_key"

    fn := func() {
        fmt.Println("Starting write operation...")
        // Simulate write operation, assuming it takes 2 seconds
        time.Sleep(200 * time.Millisecond)
        fmt.Println("Write operation completed!")
    }

    // Acquire the write lock and execute the operation
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    // Execute the lock operation with a 3-second timeout
    err := locker.RLock(ctx, key, fn)
    if err != nil {
        return
    }

    //output:
    //Starting write operation...
    //Write operation completed!
}
```

### <span id="Lock">Lock</span>

<p>Acquires a write lock for the specified key and executes the provided function.</p>

<b>Signature:</b>

```go
func (l *RWKeyedLocker[K]) Lock(ctx context.Context, key K, fn func()) error
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/todo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    locker := NewRWKeyedLocker[string](2 * time.Second)

    // Simulate a key
    key := "resource_key"

    fn := func() {
        fmt.Println("Starting write operation...")
        // Simulate write operation, assuming it takes 2 seconds
        time.Sleep(200 * time.Millisecond)
        fmt.Println("Write operation completed!")
    }

    // Acquire the write lock and execute the operation
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    // Execute the lock operation with a 3-second timeout
    err := locker.Lock(ctx, key, fn)
    if err != nil {
        return
    }

    //output:
    //Starting write operation...
    //Write operation completed!
}
```

### <span id="NewTryKeyedLocker">NewTryKeyedLocker</span>

<p>TryKeyedLocker is a non-blocking version of KeyedLocker.</p>

<b>Signature:</b>

```go
func NewTryKeyedLocker[K comparable]() *TryKeyedLocker[K]
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/todo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    locker := concurrency.NewTryKeyedLocker[string]()

    key := "resource_key"

    if locker.TryLock(key) {
        fmt.Println("Lock acquired")
        time.Sleep(1 * time.Second)
        // Unlock after work is done
        locker.Unlock(key)
        fmt.Println("Lock released")
    } else {
        fmt.Println("Lock failed")
    }

    //output:
    //Lock acquired
    //Lock released
}
```

### <span id="TryLock">TryLock</span>

<p>TryLock tries to acquire a lock for the specified key.</p>

<b>Signature:</b>

```go
func (l *TryKeyedLocker[K]) TryLock(key K) bool
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/todo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    locker := concurrency.NewTryKeyedLocker[string]()

    key := "resource_key"

    if locker.TryLock(key) {
        fmt.Println("Lock acquired")
        time.Sleep(1 * time.Second)
        // Unlock after work is done
        locker.Unlock(key)
        fmt.Println("Lock released")
    } else {
        fmt.Println("Lock failed")
    }

    //output:
    //Lock acquired
    //Lock released
}
```

### <span id="Unlock">Unlock</span>

<p>Unlock releases the lock for the specified key.</p>

<b>Signature:</b>

```go
func (l *TryKeyedLocker[K]) Unlock(key K)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/todo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/concurrency"
)

func main() {
    locker := concurrency.NewTryKeyedLocker[string]()

    key := "resource_key"

    if locker.TryLock(key) {
        fmt.Println("Lock acquired")
        time.Sleep(1 * time.Second)
        // Unlock after work is done
        locker.Unlock(key)
        fmt.Println("Lock released")
    } else {
        fmt.Println("Lock failed")
    }

    //output:
    //Lock acquired
    //Lock released
}
```
