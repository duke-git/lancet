# Pointer

pointer 包支持一些指针类型的操作。

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
-   [UnwrapOr](#UnwrapOr)
-   [UnwrapOrDefault](#UnwrapOrDefault)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="Of">Of</span>

<p>返回传入参数的指针值。</p>

<b>函数签名:</b>

```go
func Of[T any](v T) *T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/HFd70x4DrMj)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/cgeu3g7cjWb)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/D7HFjeWU2ZP)</span></b>

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

### <span id="UnwrapOr">UnwrapOr</span>

<p>返回指针的值，如果指针为零值，则返回fallback。</p>

<b>函数签名:</b>

```go
UnwrapOr[T any](p *T, fallback T) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/mmNaLC38W8C)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/pointer"
)

func main() {
	a := 123
	b := "abc"

	var c *int
	var d *string

	result1 := pointer.UnwrapOr(&a, 456)
	result2 := pointer.UnwrapOr(&b, "abc")
	result3 := pointer.UnwrapOr(c, 456)
	result4 := pointer.UnwrapOr(d, "def")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 123
	// abc
	// 456
	// def
}
```

### <span id="UnwrapOrDefault">UnwrapOrDefault</span>

<p>返回指针的值，如果指针为零值，则返回相应零值。</p>

<b>函数签名:</b>

```go
UnwrapOrDefault[T any](p *T) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ZnGIHf8_o4E)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/pointer"
)

func main() {
	a := 123
	b := "abc"

	var c *int
	var d *string

	result1 := pointer.UnwrapOrDefault(&a)
	result2 := pointer.UnwrapOrDefault(&b)
	result3 := pointer.UnwrapOrDefault(c)
	result4 := pointer.UnwrapOrDefault(d)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 123
	// abc
	// 0
	//
}
```
