# Pointer

pointer包支持一些指针类型的操作。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/pointer/pointer.go](https://github.com/duke-git/lancet/blob/main/pointer/pointer.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/pointer"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [Of](#Of)
-   [Unwrap](#Unwrap)
-   [ExtractPointer](#ExtractPointer)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="Of">Of</span>

<p>返回传入参数的指针值。</p>

<b>函数签名:</b>

```go
func Of[T any](v T) *T
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/pointer"
)

func main() {
    result1 := pointer.Of(123)
    result2 := pointer.Of("abc")

    fmt.Println(*result1)
    fmt.Println(*result2)

    // Output:
    // 123
    // abc
}
```


### <span id="Unwrap">Unwrap</span>

<p>返回传入指针指向的值。</p>

<b>函数签名:</b>

```go
func Unwrap[T any](p *T) T
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/pointer"
)

func main() {
    a := 123
    b := "abc"

    result1 := pointer.Unwrap(&a)
    result2 := pointer.Unwrap(&b)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 123
    // abc
}
```

### <span id="ExtractPointer">ExtractPointer</span>

<p>返回传入interface的底层值。</p>

<b>函数签名:</b>

```go
func ExtractPointer(value any) any
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/pointer"
)

func main() {
    a := 1
    b := &a
    c := &b
    d := &c

    result := pointer.ExtractPointer(d)

    fmt.Println(result)

    // Output:
    // 1
}
```