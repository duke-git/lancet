# Formatter

formatter 格式化器包含一些数据格式化处理方法。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/formatter/formatter.go](https://github.com/duke-git/lancet/blob/main/formatter/formatter.go)
-   [https://github.com/duke-git/lancet/blob/main/formatter/byte.go](https://github.com/duke-git/lancet/blob/main/formatter/byte.go)

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
-   [DecimalBytes](#DecimalBytes)
-   [BinaryBytes](#BinaryBytes)
-   [ParseDecimalBytes](#ParseDecimalBytes)
-   [ParseBinaryBytes](#ParseBinaryBytes)

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

### <span id="DecimalBytes">DecimalBytes</span>

<p>返回十进制标准（以1000为基数）下的可读字节单位字符串。precision参数指定小数点后的位数，默认为4。</p>

<b>函数签名:</b>

```go
func DecimalBytes(size float64, precision ...int) string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    result1 := formatter.DecimalBytes(1000)
    result2 := formatter.DecimalBytes(1024)
    result3 := formatter.DecimalBytes(1234567)
    result4 := formatter.DecimalBytes(1234567, 3)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // 1KB
    // 1.024KB
    // 1.2346MB
    // 1.235MB
}
```

### <span id="BinaryBytes">BinaryBytes</span>

<p>返回binary标准（以1024为基数）下的可读字节单位字符串。precision参数指定小数点后的位数，默认为4。</p>

<b>函数签名:</b>

```go
func BinaryBytes(size float64, precision ...int) string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    result1 := formatter.BinaryBytes(1024)
    result2 := formatter.BinaryBytes(1024 * 1024)
    result3 := formatter.BinaryBytes(1234567)
    result4 := formatter.BinaryBytes(1234567, 2)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // 1KiB
    // 1MiB
    // 1.1774MiB
    // 1.18MiB
}
```

### <span id="ParseDecimalBytes">ParseDecimalBytes</span>

<p>将字节单位字符串转换成其所表示的字节数（以1000为基数）。</p>

<b>函数签名:</b>

```go
func ParseDecimalBytes(size string) (uint64, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    result1, _ := formatter.ParseDecimalBytes("12")
    result2, _ := formatter.ParseDecimalBytes("12k")
    result3, _ := formatter.ParseDecimalBytes("12 Kb")
    result4, _ := formatter.ParseDecimalBytes("12.2 kb")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // 12
    // 12000
    // 12000
    // 12200
}
```

### <span id="ParseBinaryBytes">ParseBinaryBytes</span>

<p>将字节单位字符串转换成其所表示的字节数（以1024为基数）。</p>

<b>函数签名:</b>

```go
func ParseBinaryBytes(size string) (uint64, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    result1, _ := formatter.ParseBinaryBytes("12")
    result2, _ := formatter.ParseBinaryBytes("12ki")
    result3, _ := formatter.ParseBinaryBytes("12 KiB")
    result4, _ := formatter.ParseBinaryBytes("12.2 kib")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // 12
    // 12288
    // 12288
    // 12492
}
```
