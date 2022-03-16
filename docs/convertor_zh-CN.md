# Convertor
convertor转换器包支持一些常见的数据类型转换

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/convertor/convertor.go](https://github.com/duke-git/lancet/blob/main/convertor/convertor.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/convertor"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

- [ColorHexToRGB](#ColorHexToRGB)
- [ColorRGBToHex](#ColorRGBToHex)
- [ToBool](#ToBool)
- [ToBytes](#ToBytes)
- [ToChar](#ToChar)
- [ToInt](#ToInt)
- [ToJson](#ToJson)
- [ToString](#ToString)
- [StructToMap](#StructToMap)

<div STYLE="page-break-after: always;"></div>

## 文档



### <span id="ColorHexToRGB">ColorHexToRGB</span>
<p>颜色值十六进制转rgb</p>

<b>函数签名:</b>

```go
func ColorHexToRGB(colorHex string) (red, green, blue int)
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    colorHex := "#003366"
    r, g, b := convertor.ColorHexToRGB(colorHex)
    fmt.Println(r, g, b) //0,51,102
}
```



### <span id="ColorRGBToHex">ColorRGBToHex</span>

<p>颜色值rgb转十六进制</p>

<b>函数签名:</b>

```go
func ColorRGBToHex(red, green, blue int) string
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    r := 0
    g := 51
    b := 102
    colorHex := convertor.ColorRGBToHex(r, g, b)

    fmt.Println(colorHex) //#003366
}
```



### <span id="ToBool">ToBool</span>

<p>字符串转布尔类型，使用strconv.ParseBool</p>

<b>函数签名:</b>

```go
func ToBool(s string) (bool, error)
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    v1, _ := convertor.ToBool("1")
    fmt.Println(v1) //true

    v2, _ := convertor.ToBool("true")
    fmt.Println(v2) //true

    v3, _ := convertor.ToBool("True")
    fmt.Println(v3) //true

    v4, _ := convertor.ToBool("123")
    fmt.Println(v4) //false
}
```



### <span id="ToBytes">ToBytes</span>

<p>interface转字节切片.</p>

<b>函数签名:</b>

```go
func ToBytes(data interface{}) ([]byte, error)
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    bytesData, err := convertor.ToBytes("0")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(bytesData) //[]bytes{3, 4, 0, 0}
}
```



### <span id="ToChar">ToChar</span>

<p>字符串转字符切片</p>

<b>函数签名:</b>

```go
func ToChar(s string) []string
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    chars := convertor.ToChar("")
    fmt.Println(chars) //[]string{""}

    chars = convertor.ToChar("abc")
    fmt.Println(chars) //[]string{"a", "b", "c"}

    chars = convertor.ToChar("1 2#3")
    fmt.Println(chars) //[]string{"1", " ", "2", "#", "3"}
}
```



### <span id="ToFloat">ToFloat</span>

<p>将interface转成float64类型，如果参数无法转换，会返回0和error</p>

<b>函数签名:</b>

```go
func ToFloat(value interface{}) (float64, error)
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    v, err := convertor.ToFloat("")
    if err != nil {
        fmt.Println(err) //strconv.ParseFloat: parsing "": invalid syntax
    }
    fmt.Println(v) //0

    v, _ = convertor.ToFloat("-.11")
    fmt.Println(v) //-0.11
}
```



### <span id="ToInt">ToInt</span>

<p>将interface转成intt64类型，如果参数无法转换，会返回0和error</p>

<b>函数签名:</b>

```go
func ToInt(value interface{}) (int64, error)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    v, err := convertor.ToInt("")
    if err != nil {
        fmt.Println(err) //strconv.ParseInt: parsing "": invalid syntax
    }
    fmt.Println(v) //0

    v, _ = convertor.ToFloat(1.12)
    fmt.Println(v) //1
}
```



### <span id="ToJson">ToJson</span>

<p>将interface转成json字符串，如果参数无法转换，会返回""和error</p>

<b>函数签名:</b>

```go
func ToJson(value interface{}) (string, error)
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    var aMap = map[string]int{"a": 1, "b": 2, "c": 3}
    jsonStr, _ := convertor.ToJson(aMap)
    fmt.Printf("%q", jsonStr) //"{\"a\":1,\"b\":2,\"c\":3}"
}
```



### <span id="ToString">ToString</span>

<p>将interface转成字符串</p>

<b>函数签名:</b>

```go
func ToString(value interface{}) string
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    fmt.Printf("%q", convertor.ToString(1)) //"1"
    fmt.Printf("%q", convertor.ToString(1.1)) //"1.1"
    fmt.Printf("%q", convertor.ToString([]int{1, 2, 3})) //"[1,2,3]"
}
```



### <span id="StructToMap">StructToMap</span>

<p>将struct转成map，只会转换struct中可导出的字段。struct中导出字段需要设置json tag标记</p>

<b>函数签名:</b>

```go
func StructToMap(value interface{}) (map[string]interface{}, error)
```
<b>列子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    type People struct {
        Name string `json:"name"`
        age  int
    }
    p := People{
        "test",
        100,
    }
    pm, _ := convertor.StructToMap(p)

    fmt.Printf("type: %T, value: %s", pm, pm) //type: map[string]interface {}, value: map[name:test]
}
```