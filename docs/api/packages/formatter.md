# Formatter

formatter 格式化器包含一些数据格式化处理方法。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/formatter/formatter.go](https://github.com/duke-git/lancet/blob/main/formatter/formatter.go)
-   [https://github.com/duke-git/lancet/blob/main/formatter/byte.go](https://github.com/duke-git/lancet/blob/main/formatter/byte.go)
-   [https://github.com/duke-git/lancet/blob/main/formatter/address.go](https://github.com/duke-git/lancet/blob/main/formatter/address.go)

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
-   [ParseCNAddress](#ParseCNAddress)
-   [ParsePersonInfo](#ParsePersonInfo)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="Comma">Comma</span>

<p>用逗号每隔3位分割数字/字符串，支持添加前缀符号。参数value必须是数字或者可以转为数字的字符串, 否则返回空字符串</p>

<b>函数签名:</b>

```go
func Comma[T constraints.Float | constraints.Integer | string](value T, prefixSymbol string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/eRD5k2vzUVX)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/YsciGj3FH2x)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/LPLZ3lDi5ma)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/FPXs1suwRcs)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/G9oHHMCAZxP)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Am98ybWjvjj)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/69v1tTT62x8)</span></b>

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

### <span id="ParseCNAddress">ParseCNAddress</span>

<p>智能解析中国地址字符串并提取结构化信息。可以解析带或不带用户信息（姓名、电话、身份证等）的地址。当 withUser 为 true 时，从地址字符串中提取用户信息。当 withUser 为 false 时，仅解析位置信息。支持多种地址格式：标准格式、紧凑格式、带关键词格式、县级市格式等。</p>

<b>函数签名:</b>

```go
func ParseCNAddress(str string, withUser bool) *AddressInfo
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/o5l09hQopEV)</span></b>

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    // 解析包含用户信息的完整地址
    result1 := formatter.ParseCNAddress("张三 13800138000 北京市朝阳区建国路1号", true)
    jsonData1, _ := json.MarshalIndent(result1, "", "  ")
    fmt.Println("示例 1 - 带用户信息:")
    fmt.Println(string(jsonData1))

    // 仅解析地址，不提取用户信息
    result2 := formatter.ParseCNAddress("北京市海淀区中关村大街1号", false)
    fmt.Printf("\n示例 2 - 仅地址:\n")
    fmt.Printf("省: %s, 市: %s, 区: %s, 街道: %s\n",
        result2.Province, result2.City, result2.Region, result2.Street)

    // 解析县级市地址
    result3 := formatter.ParseCNAddress("河北省石家庄市新乐市经济开发区兴工街10号", false)
    fmt.Printf("\n示例 3 - 县级市:\n")
    fmt.Printf("省: %s, 市: %s, 区/县: %s, 街道: %s\n",
        result3.Province, result3.City, result3.Region, result3.Street)

    // 紧凑格式
    result4 := formatter.ParseCNAddress("马云13593464918陕西省西安市雁塔区丈八沟街道", true)
    fmt.Printf("\n示例 4 - 紧凑格式:\n")
    fmt.Printf("姓名: %s, 电话: %s, 地址: %s%s%s%s\n",
        result4.Name, result4.Mobile, result4.Province, result4.City, result4.Region, result4.Street)

    // Output:
    // 示例 1 - 带用户信息:
    // {
    //   "name": "张三",
    //   "mobile": "13800138000",
    //   "idn": "",
    //   "postcode": "",
    //   "province": "北京",
    //   "city": "北京市",
    //   "region": "朝阳区",
    //   "street": "建国路1号",
    //   "addr": "北京市朝阳区建国路1号"
    // }
    //
    // 示例 2 - 仅地址:
    // 省: 北京, 市: 北京市, 区: 海淀区, 街道: 中关村大街1号
    //
    // 示例 3 - 县级市:
    // 省: 河北省, 市: 石家庄市, 区/县: 新乐市, 街道: 经济开发区兴工街10号
    //
    // 示例 4 - 紧凑格式:
    // 姓名: 马云, 电话: 13593464918, 地址: 陕西省西安市雁塔区丈八沟街道
}
```

### <span id="ParsePersonInfo">ParsePersonInfo</span>

<p>从地址字符串中提取用户信息（姓名、电话、身份证、邮编）。将个人信息与地址分离，支持带标签格式、紧凑格式、带分隔符格式。返回包含提取的用户信息和清理后地址字符串的 AddressInfo。</p>

<b>函数签名:</b>

```go
func ParsePersonInfo(str string) *AddressInfo
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/JO-uTlJlTy7)</span></b>

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    // 提取姓名和手机号
    result1 := formatter.ParsePersonInfo("张三 13800138000 北京市朝阳区")
    fmt.Println("示例 1 - 姓名和手机号:")
    fmt.Printf("姓名: %s, 手机: %s, 地址: %s\n", result1.Name, result1.Mobile, result1.Addr)

    // 提取身份证号
    result2 := formatter.ParsePersonInfo("李四 110101199001011234 上海市")
    fmt.Println("\n示例 2 - 身份证号:")
    fmt.Printf("姓名: %s, 身份证: %s, 地址: %s\n", result2.Name, result2.IDN, result2.Addr)

    // 带标签格式
    result3 := formatter.ParsePersonInfo("收货人：王五 电话：13900139000 收货地址：天津市河西区友谊路20号")
    jsonData3, _ := json.MarshalIndent(result3, "", "  ")
    fmt.Println("\n示例 3 - 带标签格式:")
    fmt.Println(string(jsonData3))

    // Output:
    // 示例 1 - 姓名和手机号:
    // 姓名: 张三, 手机: 13800138000, 地址: 北京市朝阳区
    //
    // 示例 2 - 身份证号:
    // 姓名: 李四, 身份证: 110101199001011234, 地址: 上海市
    //
    // 示例 3 - 带标签格式:
    // {
    //   "name": "王五",
    //   "mobile": "13900139000",
    //   "idn": "",
    //   "postcode": "",
    //   "province": "",
    //   "city": "",
    //   "region": "",
    //   "street": "",
    //   "addr": "天津市河西区友谊路20号"
    // }
}
```
