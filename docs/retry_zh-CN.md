# Retry
retry重试执行函数直到函数运行成功或被context cancel。

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/retry/retry.go](https://github.com/duke-git/lancet/blob/main/retry/retry.go)


<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/retry"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [Context](#Context)
- [Retry](#Retry)
- [RetryFunc](#RetryFunc)
- [RetryDuration](#RetryDuration)
- [RetryTimes](#RetryTimes)


<div STYLE="page-break-after: always;"></div>


## Document文档


### <span id="Context">Context</span>
<p>设置重试context参数</p>

<b>函数签名:</b>

```go
func Context(ctx context.Context)
```
<b>例子:</b>

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
	var number int
	increaseNumber := func() error {
		number++
		if number > 3 {
			cancel()
		}
		return errors.New("error occurs")
	}

	err := retry.Retry(increaseNumber,
		retry.RetryDuration(time.Microsecond*50),
		retry.Context(ctx),
	)

	if err != nil {
		fmt.Println(err) //retry is cancelled
	}
}
```




### <span id="RetryFunc">RetryFunc</span>
<p>被重试执行的函数</p>

<b>函数签名:</b>

```go
type RetryFunc func() error
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/duke-git/lancet/retry"
)

func main() {
    var number int
    var increaseNumber retry.RetryFunc
	increaseNumber = func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := retry.Retry(increaseNumber, retry.RetryDuration(time.Microsecond*50))
    if err != nil {
		log.Fatal(err)
	}

    fmt.Println(number) //3
}
```



### <span id="RetryTimes">RetryTimes</span>
<p>设置重试次数，默认5</p>

<b>函数签名:</b>

```go
func RetryTimes(n uint)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/duke-git/lancet/retry"
)

func main() {
    var number int
	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := retry.Retry(increaseNumber, retry.RetryTimes(2))
    if err != nil {
		log.Fatal(err) //2022/02/01 18:42:25 function main.main.func1 run failed after 2 times retry exit status 1
	}
}
```



### <span id="RetryDuration">RetryDuration</span>
<p>设置重试间隔时间，默认3秒</p>

<b>函数签名:</b>

```go
func RetryDuration(d time.Duration)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/duke-git/lancet/retry"
)

func main() {
    var number int
	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := retry.Retry(increaseNumber, retry.RetryDuration(time.Microsecond*50))
    if err != nil {
		log.Fatal(err)
	}

    fmt.Println(number) //3
}
```


### <span id="Retry">Retry</span>
<p>重试执行函数retryFunc，直到函数运行成功，或被context停止</p>

<b>函数签名:</b>

```go
func Retry(retryFunc RetryFunc, opts ...Option) error
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "errors"
    "log"
    "github.com/duke-git/lancet/retry"
)

func main() {
    var number int
	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := retry.Retry(increaseNumber, retry.RetryDuration(time.Microsecond*50))
    if err != nil {
		log.Fatal(err)
	}

    fmt.Println(number) //3
}
```
