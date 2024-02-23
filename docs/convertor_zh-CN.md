# Convertor

convertor 转换器包支持一些常见的数据类型转换

<div STYLE="page-break-after: always;"></div>

## 源码:

[https://github.com/duke-git/lancet/blob/v1/convertor/convertor.go](https://github.com/duke-git/lancet/blob/v1/convertor/convertor.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/convertor"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [ColorHexToRGB](#ColorHexToRGB)
-   [ColorRGBToHex](#ColorRGBToHex)
-   [ToBool](#ToBool)
-   [ToBytes](#ToBytes)
-   [ToChar](#ToChar)
-   [ToChannel](#ToChannel)
-   [ToFloat](#ToFloat)
-   [ToInt](#ToInt)
-   [ToJson](#ToJson)
-   [ToString](#ToString)
-   [StructToMap](#StructToMap)
-   [MapToStruct](#MapToStruct)
-   [EncodeByte](#EncodeByte)
-   [DecodeByte](#DecodeByte)
-   [DeepClone](#DeepClone)
-   [CopyProperties](#CopyProperties)
-   [ToInterface](#ToInterface)
-   [Utf8ToGbk](#Utf8ToGbk)
-   [GbkToUtf8](#GbkToUtf8)

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
    r, g, b := ColorHexToRGB(colorHex)
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
    colorHex := ColorRGBToHex(r, g, b)

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

### <span id="ToChannel">ToChannel</span>

<p>将切片转为只读channel</p>

<b>函数签名:</b>

```go
func ToChannel(array []interface{}) <-chan interface{}
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    ch := convertor.ToChannel([]int{1, 2, 3})

    val1, _ := <-ch
    fmt.Println(val1) //1

    val2, _ := <-ch
    fmt.Println(val2) //2

    val3, _ := <-ch
    fmt.Println(val3) //3

    _, ok := <-ch
    fmt.Println(ok) //false
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

### <span id="MapToStruct">MapToStruct</span>

<p>将map转成struct，struct中导出字段需要设置json tag标记</p>

<b>函数签名:</b>

```go
func MapToStruct(m map[string]interface{}, structObj interface{}) error
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    type Address struct {
        Street string `json:"street"`
        Number int    `json:"number"`
    }

    type Person struct {
        Name  string   `json:"name"`
        Age   int      `json:"age"`
        Phone string   `json:"phone"`
        Addr  *Address `json:"address"`
    }

    m := map[string]interface{}{
        "name":  "Nothin",
        "age":   28,
        "phone": "123456789",
        "address": map[string]interface{}{
            "street": "test",
            "number": 1,
        },
    }

    var p Person
    err := MapToStruct(m, &p)
    if err != nil {
        return
    }

    fmt.Printf("p.Addr.Street: %s", p.Addr.Street) //test
}
```

### <span id="EncodeByte">EncodeByte</span>

<p>将data编码成字节切片</p>

<b>函数签名:</b>

```go
func EncodeByte(data any) ([]byte, error)
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    byteData, _ := convertor.EncodeByte("abc")
    fmt.Println(byteData) //[]byte{6, 12, 0, 3, 97, 98, 99}
}
```

### <span id="DecodeByte">DecodeByte</span>

<p>解码字节切片到目标对象，目标对象需要传入一个指针实例子</p>

<b>函数签名:</b>

```go
func DecodeByte(data []byte, target any) error
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    var result string
    byteData := []byte{6, 12, 0, 3, 97, 98, 99}
    convertor.DecodeByte(byteData, &result)
    fmt.Println(result) //"abc"
}
```

### <span id="DeepClone">DeepClone</span>

<p>创建一个传入值的深拷贝, 无法克隆结构体的非导出字段。</p>

<b>函数签名:</b>

```go
func DeepClone[T any](src T) T
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    type Struct struct {
        Str        string
        Int        int
        Float      float64
        Bool       bool
        Nil        interface{}
        unexported string
    }

    cases := []interface{}{
        true,
        1,
        0.1,
        map[string]int{
            "a": 1,
            "b": 2,
        },
        &Struct{
            Str:   "test",
            Int:   1,
            Float: 0.1,
            Bool:  true,
            Nil:   nil,
            // unexported: "can't be cloned",
        },
    }

    for _, item := range cases {
        cloned := convertor.DeepClone(item)

        isPointerEqual := &cloned == &item
        fmt.Println(cloned, isPointerEqual)
    }

    // Output:
    // true false
    // 1 false
    // 0.1 false
    // map[a:1 b:2] false
    // &{test 1 0.1 true <nil> } false
}
```


### <span id="CopyProperties">CopyProperties</span>

<p>拷贝不同结构体之间的同名字段。使用json.Marshal序列化，需要设置dst和src struct字段的json tag。</p>

<b>函数签名:</b>

```go
func CopyProperties(dst, src interface{}) error
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    type Disk struct {
        Name    string  `json:"name"`
        Total   string  `json:"total"`
        Used    string  `json:"used"`
        Percent float64 `json:"percent"`
    }

    type DiskVO struct {
        Name    string  `json:"name"`
        Total   string  `json:"total"`
        Used    string  `json:"used"`
        Percent float64 `json:"percent"`
    }

    type Indicator struct {
        Id      string    `json:"id"`
        Ip      string    `json:"ip"`
        UpTime  string    `json:"upTime"`
        LoadAvg string    `json:"loadAvg"`
        Cpu     int       `json:"cpu"`
        Disk    []Disk    `json:"disk"`
        Stop    chan bool `json:"-"`
    }

    type IndicatorVO struct {
        Id      string   `json:"id"`
        Ip      string   `json:"ip"`
        UpTime  string   `json:"upTime"`
        LoadAvg string   `json:"loadAvg"`
        Cpu     int64    `json:"cpu"`
        Disk    []DiskVO `json:"disk"`
    }

    indicator := &Indicator{Id: "001", Ip: "127.0.0.1", Cpu: 1, Disk: []Disk{
        {Name: "disk-001", Total: "100", Used: "1", Percent: 10},
        {Name: "disk-002", Total: "200", Used: "1", Percent: 20},
        {Name: "disk-003", Total: "300", Used: "1", Percent: 30},
    }}

    indicatorVO := IndicatorVO{}

    err := convertor.CopyProperties(&indicatorVO, indicator)

    if err != nil {
        return
    }

    fmt.Println(indicatorVO.Id)
    fmt.Println(indicatorVO.Ip)
    fmt.Println(len(indicatorVO.Disk))

    // Output:
    // 001
    // 127.0.0.1
    // 3
}
```

### <span id="ToInterface">ToInterface</span>

<p>将反射值转换成对应的interface类型。</p>

<b>函数签名:</b>

```go
func ToInterface(v reflect.Value) (value interface{}, ok bool)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    val := reflect.ValueOf("abc")
    iVal, ok := convertor.ToInterface(val)

    fmt.Printf("%T\n", iVal)
    fmt.Printf("%v\n", iVal)
    fmt.Println(ok)

    // Output:
    // string
    // abc
    // true    
}
```

### <span id="Utf8ToGbk">Utf8ToGbk</span>

<p>utf8编码转GBK编码。</p>

<b>函数签名:</b>

```go
func Utf8ToGbk(bs []byte) ([]byte, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
    "github.com/duke-git/lancet/validator"
)

func main() {
    utf8Data := []byte("hello")
    gbkData, _ := convertor.Utf8ToGbk(utf8Data)

    fmt.Println(utf8.Valid(utf8Data))
    fmt.Println(validator.IsGBK(gbkData))

    // Output:
    // true
    // true   
}
```

### <span id="GbkToUtf8">GbkToUtf8</span>

<p>GBK编码转utf8编码。</p>

<b>函数签名:</b>

```go
func GbkToUtf8(bs []byte) ([]byte, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    gbkData, _ := convertor.Utf8ToGbk([]byte("hello"))
    utf8Data, _ := convertor.GbkToUtf8(gbkData)

    fmt.Println(utf8.Valid(utf8Data))
    fmt.Println(string(utf8Data))

    // Output:
    // true
    // hello   
}
```