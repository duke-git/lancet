# Retry

retry 重试执行函数直到函数运行成功或被 context cancel。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/retry/retry.go](https://github.com/duke-git/lancet/blob/main/retry/retry.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/retry"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [Context](#Context)
-   [Retry](#Retry)
-   [RetryFunc](#RetryFunc)
-   [RetryDuration](#RetryDuration)
-   [RetryTimes](#RetryTimes)
-   [BackoffStrategy](#BackoffStrategy)
-   [RetryWithCustomBackoff](#RetryWithCustomBackoff)
-   [RetryWithLinearBackoff](#RetryWithLinearBackoff)
-   [RetryWithExponentialWithJitterBackoff](#RetryWithExponentialWithJitterBackoff)

<div STYLE="page-break-after: always;"></div>


## 文档

### <span id="Context">Context</span>

<p>设置重试context参数</p>

<b>函数签名:</b>

```go
func Context(ctx context.Context)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/xnAOOXv9GkS)</span></b>

```go
import (
    "context"
    "errors"
    "fmt"
    "lancet-demo/retry"
    "time"
)

func main() {
    ctx, cancel := context.WithCancel(context.TODO())

    number := 0
    increaseNumber := func() error {
        number++
        if number > 3 {
            cancel()
        }
        return errors.New("error occurs")
    }

    duration := retry.RetryWithLinearBackoff(time.Microsecond*50)

    retry.Retry(increaseNumber,
        duration,
        retry.Context(ctx),
    )

    fmt.Println(number)

    // Output:
    // 4
}
```

### <span id="RetryFunc">RetryFunc</span>

<p>被重试执行的函数</p>

<b>函数签名:</b>

```go
type RetryFunc func() error
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/nk2XRmagfVF)</span></b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/duke-git/lancet/v2/retry"
)

func main() {
    number := 0
    var increaseNumber retry.RetryFunc = func() error {
        number++
        if number == 3 {
            return nil
        }
        return errors.New("error occurs")
    }

    duration := retry.RetryWithLinearBackoff(time.Microsecond*50)

    err := retry.Retry(increaseNumber, duration)
    if err != nil {
        return
    }

    fmt.Println(number)

    // Output:
    // 3
}
```

### <span id="RetryTimes">RetryTimes</span>

<p>设置重试次数，默认5</p>

<b>函数签名:</b>

```go
func RetryTimes(n uint)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ssfVeU2SwLO)</span></b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/duke-git/lancet/v2/retry"
)

func main() {
    number := 0

    increaseNumber := func() error {
        number++
        if number == 3 {
            return nil
        }
        return errors.New("error occurs")
    }

    err := retry.Retry(increaseNumber, retry.RetryTimes(2))
    if err != nil {
        fmt.Println(err)
    }

    // Output:
    // function main.main.func1 run failed after 2 times retry
}
```

### <span id="Retry">Retry</span>

<p>重试执行函数retryFunc，直到函数运行成功，或被context停止</p>

<b>函数签名:</b>

```go
func Retry(retryFunc RetryFunc, opts ...Option) error
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/nk2XRmagfVF)</span></b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/duke-git/lancet/v2/retry"
)

func main() {
    number := 0
    increaseNumber := func() error {
        number++
        if number == 3 {
            return nil
        }
        return errors.New("error occurs")
    }

    duration := retry.RetryWithLinearBackoff(time.Microsecond*50)

    err := retry.Retry(increaseNumber, duration)
    if err != nil {
        return
    }

    fmt.Println(number)

    // Output:
    // 3
}
```

### <span id="BackoffStrategy">BackoffStrategy</span>

<p>定义计算退避间隔的方法的接口。</p>

<b>函数签名:</b>

```go
// BackoffStrategy is an interface that defines a method for calculating backoff intervals.
type BackoffStrategy interface {
    // CalculateInterval returns the time.Duration after which the next retry attempt should be made.
    CalculateInterval() time.Duration
}
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/duke-git/lancet/v2/retry"
)

type ExampleCustomBackoffStrategy struct {
    interval time.Duration
}

func (c *ExampleCustomBackoffStrategy) CalculateInterval() time.Duration {
    return c.interval + 1
}

func main() {
    number := 0
    increaseNumber := func() error {
        number++
        if number == 3 {
            return nil
        }
        return errors.New("error occurs")
    }

    err := retry,Retry(increaseNumber, retry.RetryWithCustomBackoff(&示例CustomBackoffStrategy{interval: time.Microsecond * 50}))
    if err != nil {
        return
    }

    fmt.Println(number)

    // Output:
    // 3
}
```

### <span id="RetryWithCustomBackoff">RetryWithCustomBackoff</span>

<p>设置自定义退避策略。</p>

<b>函数签名:</b>

```go
func RetryWithCustomBackoff(backoffStrategy BackoffStrategy) Option 
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/jIm_o2vb5Y4)</span></b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/duke-git/lancet/v2/retry"
)

type ExampleCustomBackoffStrategy struct {
    interval time.Duration
}

func (c *ExampleCustomBackoffStrategy) CalculateInterval() time.Duration {
    return c.interval + 1
}

func main() {
    number := 0
    increaseNumber := func() error {
        number++
        if number == 3 {
            return nil
        }
        return errors.New("error occurs")
    }

    err := retry,Retry(increaseNumber, retry.RetryWithCustomBackoff(&示例CustomBackoffStrategy{interval: time.Microsecond * 50}))
    if err != nil {
        return
    }

    fmt.Println(number)

    // Output:
    // 3
}
```


### <span id="RetryWithLinearBackoff">RetryWithLinearBackoff</span>

<p>设置线性策略退避。</p>

<b>函数签名:</b>

```go
func RetryWithLinearBackoff(interval time.Duration) Option
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/PDet2ZQZwcB)</span></b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/duke-git/lancet/v2/retry"
)

func main() {
    number := 0
    increaseNumber := func() error {
        number++
        if number == 3 {
            return nil
        }
        return errors.New("error occurs")
    }

    err := retry.Retry(increaseNumber, retry.RetryWithLinearBackoff(time.Microsecond*50))
    if err != nil {
        return
    }

    fmt.Println(number)

    // Output:
    // 3
}
```


### <span id="RetryWithExponentialWithJitterBackoff">RetryWithExponentialWithJitterBackoff</span>

<p>设置指数策略退避。</p>

<b>函数签名:</b>

```go
func RetryWithExponentialWithJitterBackoff(interval time.Duration, base uint64, maxJitter time.Duration) Option
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/xp1avQmn16X)</span></b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/duke-git/lancet/v2/retry"
)

func main() {
    number := 0
    increaseNumber := func() error {
        number++
        if number == 3 {
            return nil
        }
        return errors.New("error occurs")
    }

    err := retry.Retry(increaseNumber, retry.RetryWithExponentialWithJitterBackoff(time.Microsecond*50, 2, time.Microsecond*25))
    if err != nil {
        return
    }

    fmt.Println(number)

    // Output:
    // 3
}
```
