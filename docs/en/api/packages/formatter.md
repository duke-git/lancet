# Formatter

formatter contains some functions for data formatting.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/formatter/formatter.go](https://github.com/duke-git/lancet/blob/main/formatter/formatter.go)
-   [https://github.com/duke-git/lancet/blob/main/formatter/byte.go](https://github.com/duke-git/lancet/blob/main/formatter/byte.go)
-   [https://github.com/duke-git/lancet/blob/main/formatter/address.go](https://github.com/duke-git/lancet/blob/main/formatter/address.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/formatter"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

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

## Documentation

### <span id="Comma">Comma</span>

<p>Add comma to a number value by every 3 numbers from right to left. ahead by a prefix symbol char. if value is a invalid number string like "aa", return empty string.</p>

<b>Signature:</b>

```go
func Comma[T constraints.Float | constraints.Integer | string](value T, prefixSymbol string) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/eRD5k2vzUVX)</span></b>

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

<p>Pretty data to JSON string.</p>

<b>Signature:</b>

```go
func Pretty(v any) (string, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/YsciGj3FH2x)</span></b>

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

<p>Pretty encode data to writer.</p>

<b>Signature:</b>

```go
func PrettyToWriter(v any, out io.Writer) error
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/LPLZ3lDi5ma)</span></b>

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

<p>Returns a human readable byte size under decimal standard (base 1000). The precision parameter specifies the number of digits after the decimal point, which is 4 for default.</p>

<b>Signature:</b>

```go
func DecimalBytes(size float64, precision ...int) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/FPXs1suwRcs)</span></b>

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

<p>Returns a human readable byte size under binary standard (base 1024). The precision parameter specifies the number of digits after the decimal point, which is 4 for default.</p>

<b>Signature:</b>

```go
func BinaryBytes(size float64, precision ...int) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/G9oHHMCAZxP)</span></b>

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

<p>Returns the human readable bytes size string into the amount it represents(base 1000).</p>

<b>Signature:</b>

```go
func ParseDecimalBytes(size string) (uint64, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/Am98ybWjvjj)</span></b>

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

<p>Returns the human readable bytes size string into the amount it represents(base 1024).</p>

<b>Signature:</b>

```go
func ParseBinaryBytes(size string) (uint64, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/69v1tTT62x8)</span></b>

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

<p>Parses a Chinese address string intelligently and extracts structured information. It can parse addresses with or without user information (name, phone, ID card, etc.). When withUser is true, it extracts user information from the address string. When withUser is false, it only parses the location information. Supports various address formats: standard format, compact format, labeled format, county-level cities format, etc.</p>

<b>Signature:</b>

```go
func ParseCNAddress(str string, withUser bool) *AddressInfo
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/o5l09hQopEV)</span></b>

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    // Parse complete address with user information
    result1 := formatter.ParseCNAddress("张三 13800138000 北京市朝阳区建国路1号", true)
    jsonData1, _ := json.MarshalIndent(result1, "", "  ")
    fmt.Println("Example 1 - With user info:")
    fmt.Println(string(jsonData1))

    // Parse address only, without extracting user information
    result2 := formatter.ParseCNAddress("北京市海淀区中关村大街1号", false)
    fmt.Printf("\nExample 2 - Address only:\n")
    fmt.Printf("Province: %s, City: %s, Region: %s, Street: %s\n",
        result2.Province, result2.City, result2.Region, result2.Street)

    // Parse county-level city address
    result3 := formatter.ParseCNAddress("河北省石家庄市新乐市经济开发区兴工街10号", false)
    fmt.Printf("\nExample 3 - County-level city:\n")
    fmt.Printf("Province: %s, City: %s, Region: %s, Street: %s\n",
        result3.Province, result3.City, result3.Region, result3.Street)

    // Compact format
    result4 := formatter.ParseCNAddress("马云13593464918陕西省西安市雁塔区丈八沟街道", true)
    fmt.Printf("\nExample 4 - Compact format:\n")
    fmt.Printf("Name: %s, Phone: %s, Address: %s%s%s%s\n",
        result4.Name, result4.Mobile, result4.Province, result4.City, result4.Region, result4.Street)

    // Output:
    // Example 1 - With user info:
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
    // Example 2 - Address only:
    // Province: 北京, City: 北京市, Region: 海淀区, Street: 中关村大街1号
    //
    // Example 3 - County-level city:
    // Province: 河北省, City: 石家庄市, Region: 新乐市, Street: 经济开发区兴工街10号
    //
    // Example 4 - Compact format:
    // Name: 马云, Phone: 13593464918, Address: 陕西省西安市雁塔区丈八沟街道
}
```

### <span id="ParsePersonInfo">ParsePersonInfo</span>

<p>Extracts user information (name, phone, ID card, postal code) from an address string. It separates personal information from the address, supporting labeled format, compact format, and formats with separators. Returns an AddressInfo with extracted user information and cleaned address string.</p>

<b>Signature:</b>

```go
func ParsePersonInfo(str string) *AddressInfo
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/JO-uTlJlTy7)</span></b>

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    // Extract name and phone
    result1 := formatter.ParsePersonInfo("张三 13800138000 北京市朝阳区")
    fmt.Println("Example 1 - Name and phone:")
    fmt.Printf("Name: %s, Phone: %s, Address: %s\n", result1.Name, result1.Mobile, result1.Addr)

    // Extract ID card number
    result2 := formatter.ParsePersonInfo("李四 110101199001011234 上海市")
    fmt.Println("\nExample 2 - ID card number:")
    fmt.Printf("Name: %s, ID Card: %s, Address: %s\n", result2.Name, result2.IDN, result2.Addr)

    // Labeled format
    result3 := formatter.ParsePersonInfo("收货人：王五 电话：13900139000 收货地址：天津市河西区友谊路20号")
    jsonData3, _ := json.MarshalIndent(result3, "", "  ")
    fmt.Println("\nExample 3 - Labeled format:")
    fmt.Println(string(jsonData3))

    // Output:
    // Example 1 - Name and phone:
    // Name: 张三, Phone: 13800138000, Address: 北京市朝阳区
    //
    // Example 2 - ID card number:
    // Name: 李四, ID Card: 110101199001011234, Address: 上海市
    //
    // Example 3 - Labeled format:
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
