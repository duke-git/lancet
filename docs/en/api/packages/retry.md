# Retry

Package retry is for executing a function repeatedly until it was successful or canceled by the context.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/retry/retry.go](https://github.com/duke-git/lancet/blob/main/retry/retry.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/retry"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [Context](#Context)
-   [Retry](#Retry)
-   [RetryFunc](#RetryFunc)
-   [RetryDuration](#RetryDuration)
-   [RetryTimes](#RetryTimes)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="Context">Context</span>

<p>Set retry context config, can cancel the retry with context.</p>

<b>Signature:</b>

```go
func Context(ctx context.Context)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/xnAOOXv9GkS)</span></b>

```go
import (
    "context"
    "errors"
    "fmt"
    "github.com/duke-git/lancet/v2/retry"
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

<p>Function that retry executes.</p>

<b>Signature:</b>

```go
type RetryFunc func() error
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/nk2XRmagfVF)</span></b>

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

<p>Set times of retry. Default times is 5.</p>

<b>Signature:</b>

```go
func RetryTimes(n uint)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/ssfVeU2SwLO)</span></b>

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

<p>Set duration of retries. Default duration is 3 second.</p>

<b>Signature:</b>

```go
func RetryDuration(d time.Duration)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/nk2XRmagfVF)</span></b>

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

<p>Executes the retryFunc repeatedly until it was successful or canceled by the context.</p>

<b>Signature:</b>

```go
func Retry(retryFunc RetryFunc, opts ...Option) error
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/nk2XRmagfVF)</span></b>

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
