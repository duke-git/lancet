# Xerror

Package xerror implements helpers for errors.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/xerror/xerror.go](https://github.com/duke-git/lancet/blob/main/xerror/xerror.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/xerror"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

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

## Documentation

### <span id="New">New</span>

<p>Creates a new XError pointer instance with message.</p>

<b>Signature:</b>

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

<b>Example:</b>

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

<p>Creates a new XError pointer instance based on error object, and add message.</p>

<b>Signature:</b>

```go
func Wrap(cause error, message ...any) *XError
```

<b>Example:</b>

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

<p>Returns unwrapped XError from err by errors.As. If no XError, returns nil.</p>

<b>Signature:</b>

```go
func Unwrap(err error) *XError
```

<b>Example:</b>

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

<p>Creates a new XError and copy message and id to new one.</p>

<b>Signature:</b>

```go
func (e *XError) Wrap(cause error) *XError
```

<b>Example:</b>

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

<p>Compatible with github.com/pkg/errors.</p>

<b>Signature:</b>

```go
func (e *XError) Unwrap() error
```

<b>Example:</b>

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

<p>Adds key and value related to the XError object.</p>

<b>Signature:</b>

```go
func (e *XError) With(key string, value any) *XError
```

<b>Example:</b>

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

<p>Sets XError object id to check equality in XError.Is.</p>

<b>Signature:</b>

```go
func (e *XError) Id(id string) *XError
```

<b>Example:</b>

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

<p>Checks if target error is XError and Error.id of two errors are matched.</p>

<b>Signature:</b>

```go
func (e *XError) Is(target error) bool
```

<b>Example:</b>

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

<p>Returns map of key and value that is set by With. All wrapped xerror.XError key and values will be merged. Key and values of wrapped error is overwritten by upper xerror.XError.</p>

<b>Signature:</b>

```go
func (e *XError) Values() map[string]any
```

<b>Example:</b>

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

<p>Returns stack trace which is compatible with pkg/errors.</p>

<b>Signature:</b>

```go
func (e *XError) StackTrace() StackTrace
```

<b>Example:</b>

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

<p>Returns information of xerror, which can be printed.</p>

<b>Signature:</b>

```go
func (e *XError) Info() *errInfo
```

<b>Example:</b>

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

<p>Error implements standard error interface.</p>

<b>Signature:</b>

```go
func (e *XError) Error() string
```

<b>Example:</b>

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

<p>TryUnwrap if err is nil then it returns a valid value. If err is not nil, Unwrap panics with err.</p>

<b>Signature:</b>

```go
func TryUnwrap[T any](val T, err error) T
```

<b>Example:</b>

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
