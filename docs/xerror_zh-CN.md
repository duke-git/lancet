# Xerror

xerror 错误处理逻辑封装

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/xerror/xerror.go](https://github.com/duke-git/lancet/blob/main/xerror/xerror.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/xerror"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [New](#New)
-   [Wrap](#Wrap)
-   [Unwrap](#Unwrap)
-   [XError_Wrap](#XError_Wrap)
-   [XError_Unwrap](#XError_Unwrap)
-   [XError_With](#XError_With)
-   [XError_Is](#XError_Is)
-   [XError_Id](#XError_Id)
-   [XError_Values](#XError_Values)
-   [XError_StackTrace](#XError_StackTrace)
-   [XError_Info](#XError_Info)
-   [XError_Error](#XError_Error)
-   [TryUnwrap](#TryUnwrap)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="New">New</span>

<p>创建XError对象实例。</p>

<b>函数签名:</b>

```go
type XError struct {
    id      string
    message string
    stack   *stack
    cause   error
    values  map[string]any
}

func New(format string, args ...any) *XError
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    err := xerror.New("error")
    fmt.Println(err.Error())

    // Output:
    // error
}
```

### <span id="Wrap">Wrap</span>

<p>根据error对象创建XError对象实例，可添加message。</p>

<b>函数签名:</b>

```go
func Wrap(cause error, message ...any) *XError
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    err := xerror.New("wrong password")
    wrapErr := xerror.Wrap(err, "error")

    fmt.Println(wrapErr.Error())

    // Output:
    // error: wrong password
}
```

### <span id="Unwrap">Unwrap</span>

<p>从error对象中解构出XError。</p>

<b>函数签名:</b>

```go
func Unwrap(err error) *XError
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/pkg/errors"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    err1 := xerror.New("error").With("level", "high")
    wrapErr := errors.Wrap(err1, "oops")

    err := xerror.Unwrap(wrapErr)

    values := err.Values()
    fmt.Println(values["level"])

    // Output:
    // high
}
```

### <span id="XError_Wrap">XError_Wrap</span>

<p>创建新的XError对象并将消息和id复制到新的对象中。</p>

<b>函数签名:</b>

```go
func (e *XError) Wrap(cause error) *XError
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    err := xerror.New("wrong password")
    wrapErr := xerror.Wrap(err, "error")

    fmt.Println(wrapErr.Error())

    // Output:
    // error: wrong password
}
```

### <span id="XError_Unwrap">XError_Unwrap</span>

<p>解构XEerror为error对象。适配github.com/pkg/errors。</p>

<b>函数签名:</b>

```go
func (e *XError) Unwrap() error
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    err1 := xerror.New("error").With("level", "high")
    err2 := err1.Wrap(errors.New("invalid username"))

    err := err2.Unwrap()

    fmt.Println(err.Error())

    // Output:
    // invalid username
}
```

### <span id="XError_With">XError_With</span>

<p>添加与XError对象的键和值。</p>

<b>函数签名:</b>

```go
func (e *XError) With(key string, value any) *XError
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    err := xerror.New("error").With("level", "high")

    errLevel := err.Values()["level"]

    fmt.Println(errLevel)

    // Output:
    // high
}
```

### <span id="XError_Id">XError_Id</span>

<p>设置XError对象的id。</p>

<b>函数签名:</b>

```go
func (e *XError) Id(id string) *XError
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    err1 := xerror.New("error").Id("e001")
    err2 := xerror.New("error").Id("e001")
    err3 := xerror.New("error").Id("e003")

    equal := err1.Is(err2)
    notEqual := err1.Is(err3)

    fmt.Println(equal)
    fmt.Println(notEqual)

    // Output:
    // true
    // false
}
```

### <span id="XError_Is">XError_Is</span>

<p>检查目标error是否为XError，两个错误中的error.id是否匹配。</p>

<b>函数签名:</b>

```go
func (e *XError) Is(target error) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    err1 := xerror.New("error").Id("e001")
    err2 := xerror.New("error").Id("e001")
    err3 := xerror.New("error").Id("e003")

    equal := err1.Is(err2)
    notEqual := err1.Is(err3)

    fmt.Println(equal)
    fmt.Println(notEqual)

    // Output:
    // true
    // false
}
```

### <span id="XError_Values">XError_Values</span>

<p>返回由With设置的键和值的映射。将合并所有XError键和值。</p>

<b>函数签名:</b>

```go
func (e *XError) Values() map[string]any
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    err := xerror.New("error").With("level", "high")

    errLevel := err.Values()["level"]

    fmt.Println(errLevel)

    // Output:
    // high
}
```


### <span id="XError_StackTrace">XError_StackTrace</span>

<p>返回与pkg/error兼容的堆栈信息。</p>

<b>函数签名:</b>

```go
func (e *XError) StackTrace() StackTrace
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    err := xerror.New("error")

    stacks := err.Stacks()

    fmt.Println(stacks[0].Func)
    fmt.Println(stacks[0].Line)

    containFile := strings.Contains(stacks[0].File, "xxx.go")
    fmt.Println(containFile)
}
```


### <span id="XError_Info">XError_Info</span>

<p>返回可打印的XError对象信息。</p>

<b>函数签名:</b>

```go
func (e *XError) Info() *errInfo
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    cause := errors.New("error")
    err := xerror.Wrap(cause, "invalid username").Id("e001").With("level", "high")

    errInfo := err.Info()

    fmt.Println(errInfo.Id)
    fmt.Println(errInfo.Cause)
    fmt.Println(errInfo.Values["level"])
    fmt.Println(errInfo.Message)

    // Output:
    // e001
    // error
    // high
    // invalid username
}
```


### <span id="XError_Error">XError_Error</span>

<p>实现标准库的error接口。</p>

<b>函数签名:</b>

```go
func (e *XError) Error() string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    err := xerror.New("error")
    fmt.Println(err.Error())

    // Output:
    // error
}
```

### <span id="TryUnwrap">TryUnwrap</span>

<p>检查error, 如果err为nil则展开，则它返回一个有效值，如果err不是nil则TryUnwrap使用err发生panic。</p>

<b>函数签名:</b>

```go
func TryUnwrap[T any](val T, err error) T
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/xerror"
)

func main() {
    result1 := xerror.TryUnwrap(strconv.Atoi("42"))
    fmt.Println(result1)

    _, err := strconv.Atoi("4o2")
    defer func() {
        v := recover()
        result2 := reflect.DeepEqual(err.Error(), v.(*strconv.NumError).Error())
        fmt.Println(result2)
    }()

    xerror.TryUnwrap(strconv.Atoi("4o2"))

    // Output:
    // 42
    // true
}
```
