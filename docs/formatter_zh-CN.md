# Formatter

formatter 格式化器包含一些数据格式化处理方法。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/formatter/formatter.go](https://github.com/duke-git/lancet/blob/main/formatter/formatter.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/formatter"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [Comma](#Comma)
-   [Pretty](#Pretty)
-   [PrettyToWriter](#PrettyToWriter)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="Comma">Comma</span>

<p>用逗号每隔3位分割数字/字符串，支持前缀添加符号。参数value必须是数字或者可以转为数字的字符串, 否则返回空字符串</p>

<b>函数签名:</b>

```go
func Comma[T constraints.Float | constraints.Integer | string](value T, symbol string) string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    result1 := formatter.Comma("123", "")
    result2 := formatter.Comma("12345", "$")
    result3 := formatter.Comma(1234567, "￥")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 123
    // $12,345
    // ￥1,234,567
}
```

### <span id="Pretty">Pretty</span>

<p>返回pretty JSON字符串.</p>

<b>函数签名:</b>

```go
func Pretty(v any) (string, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    result1, _ := formatter.Pretty([]string{"a", "b", "c"})
    result2, _ := formatter.Pretty(map[string]int{"a": 1})

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // [
    //     "a",
    //     "b",
    //     "c"
    // ]
    // {
    //     "a": 1
    // }
}
```

### <span id="PrettyToWriter">PrettyToWriter</span>

<p>Pretty encode数据到writer。</p>

<b>函数签名:</b>

```go
func PrettyToWriter(v any, out io.Writer) error
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    type User struct {
        Name string `json:"name"`
        Aage uint   `json:"age"`
    }
    user := User{Name: "King", Aage: 10000}

    buf := &bytes.Buffer{}
    err := formatter.PrettyToWriter(user, buf)

    fmt.Println(buf)
    fmt.Println(err)

    // Output:
    // {
    //     "name": "King",
    //     "age": 10000
    // }
    //
    // <nil>
}
```
