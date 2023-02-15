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

-   [TryUnwrap](#TryUnwrap)

<div STYLE="page-break-after: always;"></div>

## Documentation

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
