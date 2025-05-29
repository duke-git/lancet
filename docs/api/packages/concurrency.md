# Concurrency

并发包包含一些支持并发编程的功能。例如：goroutine, channel 等。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/concurrency/channel.go](https://github.com/duke-git/lancet/blob/main/concurrency/channel.go)
-   [https://github.com/duke-git/lancet/blob/main/concurrency/keyed_locker.go](https://github.com/duke-git/lancet/blob/main/concurrency/keyed_locker.go)

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
-   [KeyedLocker_Do](#Do)
-   [NewRWKeyedLocker](#NewRWKeyedLocker)
-   [RLock](#RLock)
-   [Lock](#Lock)
-   [NewTryKeyedLocker](#NewTryKeyedLocker)
-   [TryLock](#TryLock)
-   [Unlock](#Unlock)

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/7aB4KyMMp9A)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/qmWSy1NVF-Y)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/2VYFMexEvTm)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/7aB4KyMMp9A)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/k5N_ALVmYjE)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/4J1zAWttP85)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Wqz9rwioPww)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/lm_GoS6aDjo)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/9Utt-1pDr2J)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/3TQPKnCirrP)</span></b>

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

<p>NewKeyedLocker创建一个新的KeyedLocker，并为锁的过期设置指定的 TTL。KeyedLocker 是一个简单的键值锁实现，允许非阻塞的锁获取。</p>

<b>函数签名:</b>

```go
func NewKeyedLocker[K comparable](ttl time.Duration) *KeyedLocker[K]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/GzeyC33T5rw)</span></b>

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

<p>为指定的键获取锁并执行提供的函数。</p>

<b>函数签名:</b>

```go
func (l *KeyedLocker[K]) Do(ctx context.Context, key K, fn func()) error
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/GzeyC33T5rw)</span></b>

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

<p>NewRWKeyedLocker创建一个新的RWKeyedLocker，并为锁的过期设置指定的 TTL。RWKeyedLocker 是一个简单的键值读写锁实现，允许非阻塞的锁获取。</p>

<b>函数签名:</b>

```go
func NewRWKeyedLocker[K comparable](ttl time.Duration) *RWKeyedLocker[K]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/CkaJWWwZm9)</span></b>

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

<p>RLock为指定的键获取读锁并执行提供的函数。</p>

<b>函数签名:</b>

```go
func (l *RWKeyedLocker[K]) RLock(ctx context.Context, key K, fn func()) error
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ZrCr8sMo77T)</span></b>

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

<p>Lock为指定的键获取锁并执行提供的函数。</p>

<b>函数签名:</b>

```go
func (l *RWKeyedLocker[K]) Lock(ctx context.Context, key K, fn func()) error
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/WgAcXbOPKGk)</span></b>

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

<p>创建一个TryKeyedLocker实例，TryKeyedLocker是KeyedLocker的非阻塞版本。</p>

<b>函数签名:</b>

```go
func NewTryKeyedLocker[K comparable]() *TryKeyedLocker[K]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/VG9qLvyetE2)</span></b>

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

<p>TryLock尝试获取指定键的锁。如果锁成功获取，则返回true，否则返回false。</p>

<b>函数签名:</b>

```go
func (l *TryKeyedLocker[K]) TryLock(key K) bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/VG9qLvyetE2)</span></b>

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

<p>释放指定键的锁。</p>

<b>函数签名:</b>

```go
func (l *TryKeyedLocker[K]) Unlock(key K)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/VG9qLvyetE2)</span></b>

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
