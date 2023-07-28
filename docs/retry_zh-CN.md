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

<div STYLE="page-break-after: always;"></div>

## Document 文档

### <span id="Context">Context</span>

<p>设置重试context参数</p>

<b>函数签名:</b>

```go
func Context(ctx context.Context)
```

<b>示例:</b>

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

    duration := retry.RetryDuration(time.Microsecond*50)

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

<b>示例:</b>

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

    duration := retry.RetryDuration(time.Microsecond*50)

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

<b>示例:</b>

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

### <span id="RetryDuration">RetryDuration</span>

<p>设置重试间隔时间，默认3秒</p>

<b>函数签名:</b>

```go
func RetryDuration(d time.Duration)
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

func main() {
    number := 0
    increaseNumber := func() error {
        number++
        if number == 3 {
            return nil
        }
        return errors.New("error occurs")
    }

    duration := retry.RetryDuration(time.Microsecond*50)

    err := retry.Retry(increaseNumber, duration)
    if err != nil {
        return
    }

    fmt.Println(number)

    // Output:
    // 3
}
```

### <span id="Retry">Retry</span>

<p>重试执行函数retryFunc，直到函数运行成功，或被context停止</p>

<b>函数签名:</b>

```go
func Retry(retryFunc RetryFunc, opts ...Option) error
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

func main() {
    number := 0
    increaseNumber := func() error {
        number++
        if number == 3 {
            return nil
        }
        return errors.New("error occurs")
    }

    duration := retry.RetryDuration(time.Microsecond*50)

    err := retry.Retry(increaseNumber, duration)
    if err != nil {
        return
    }

    fmt.Println(number)

    // Output:
    // 3
}
```
