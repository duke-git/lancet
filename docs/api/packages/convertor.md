# Convertor

convertor 转换器包支持一些常见的数据类型转换

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/convertor/convertor.go](https://github.com/duke-git/lancet/blob/main/convertor/convertor.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/convertor"
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
-   [ToMap](#ToMap)
-   [ToPointer](#ToPointer)
-   [ToString](#ToString)
-   [StructToMap](#StructToMap)
-   [MapToSlice](#MapToSlice)
-   [EncodeByte](#EncodeByte)
-   [DecodeByte](#DecodeByte)
-   [DeepClone](#DeepClone)
-   [CopyProperties](#CopyProperties)
-   [ToInterface](#ToInterface)
-   [Utf8ToGbk](#Utf8ToGbk)
-   [GbkToUtf8](#GbkToUtf8)
-   [ToStdBase64](#ToStdBase64)
-   [ToUrlBase64](#ToUrlBase64)
-   [ToRawStdBase64](#ToRawStdBase64)
-   [ToRawUrlBase64](#ToRawUrlBase64)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="ColorHexToRGB">ColorHexToRGB</span>

<p>颜色值十六进制转rgb。</p>

<b>函数签名:</b>

```go
func ColorHexToRGB(colorHex string) (red, green, blue int)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/o7_ft-JCJBV)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    colorHex := "#003366"
    r, g, b := convertor.ColorHexToRGB(colorHex)

    fmt.Println(r, g, b)

    // Output:
    // 0 51 102
}
```

### <span id="ColorRGBToHex">ColorRGBToHex</span>

<p>颜色值rgb转十六进制。</p>

<b>函数签名:</b>

```go
func ColorRGBToHex(red, green, blue int) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/nzKS2Ro87J1)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    r := 0
    g := 51
    b := 102
    colorHex := ColorRGBToHex(r, g, b)

    fmt.Println(colorHex)

    // Output:
    // #003366
}
```

### <span id="ToBool">ToBool</span>

<p>字符串转布尔类型，使用strconv.ParseBool。</p>

<b>函数签名:</b>

```go
func ToBool(s string) (bool, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ARht2WnGdIN)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    cases := []string{"1", "true", "True", "false", "False", "0", "123", "0.0", "abc"}

    for i := 0; i < len(cases); i++ {
        result, _ := convertor.ToBool(cases[i])
        fmt.Println(result)
    }

    // Output:
    // true
    // true
    // true
    // false
    // false
    // false
    // false
    // false
    // false
}
```

### <span id="ToBytes">ToBytes</span>

<p>Interface转字节切片。</p>

<b>函数签名:</b>

```go
func ToBytes(data any) ([]byte, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/fAMXYFDvOvr)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    bytesData, err := convertor.ToBytes("abc")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(bytesData)

    // Output:
    // [97 98 99]
}
```

### <span id="ToChar">ToChar</span>

<p>字符串转字符切片。</p>

<b>函数签名:</b>

```go
func ToChar(s string) []string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/JJ1SvbFkVdM)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    result1 := convertor.ToChar("")
    result2 := convertor.ToChar("abc")
    result3 := convertor.ToChar("1 2#3")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // []
    // [a b c]
    // [1   2 # 3]
}
```

### <span id="ToChannel">ToChannel</span>

<p>将切片转为只读channel。</p>

<b>函数签名:</b>

```go
func ToChannel[T any](array []T) <-chan T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/hOx_oYZbAnL)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    ch := convertor.ToChannel([]int{1, 2, 3})
    result1 := <-ch
    result2 := <-ch
    result3 := <-ch

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 1
    // 2
    // 3
}
```

### <span id="ToFloat">ToFloat</span>

<p>将interface转成float64类型，如果参数无法转换，会返回0和error。</p>

<b>函数签名:</b>

```go
func ToFloat(value any) (float64, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/4YTmPCibqHJ)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    result1, _ := convertor.ToFloat("")
    result2, err := convertor.ToFloat("abc")
    result3, _ := convertor.ToFloat("-1")
    result4, _ := convertor.ToFloat("-.11")
    result5, _ := convertor.ToFloat("1.23e3")
    result6, _ := convertor.ToFloat(true)

    fmt.Println(result1)
    fmt.Println(result2, err)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)

    // Output:
    // 0
    // 0 strconv.ParseFloat: parsing "": invalid syntax
    // -1
    // -0.11
    // 1230
    // 0
}
```

### <span id="ToInt">ToInt</span>

<p>将interface转成int64类型，如果参数无法转换，会返回0和error。</p>

<b>函数签名:</b>

```go
func ToInt(value any) (int64, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/9_h9vIt-QZ_b)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    result1, _ := convertor.ToInt("123")
    result2, _ := convertor.ToInt("-123")
    result3, _ := convertor.ToInt(float64(12.3))
    result4, err := convertor.ToInt("abc")
    result5, _ := convertor.ToInt(true)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4, err)
    fmt.Println(result5)

    // Output:
    // 123
    // -123
    // 12
    // 0 strconv.ParseInt: parsing "": invalid syntax
    // 0
}
```

### <span id="ToJson">ToJson</span>

<p>将interface转成json字符串，如果参数无法转换，会返回""和error。</p>

<b>函数签名:</b>

```go
func ToJson(value any) (string, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/2rLIkMmXWvR)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    aMap := map[string]int{"a": 1, "b": 2, "c": 3}
    result, err := ToJson(aMap)

    if err != nil {
        fmt.Printf("%v", err)
    }

    fmt.Println(result)

    // Output:
    // {"a":1,"b":2,"c":3}
}
```

### <span id="ToMap">ToMap</span>

<p>将切片转为map。</p>

<b>函数签名:</b>

```go
func ToMap[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K]V
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/tVFy7E-t24l)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    type Message struct {
        name string
        code int
    }
    messages := []Message{
        {name: "Hello", code: 100},
        {name: "Hi", code: 101},
    }

    result := convertor.ToMap(messages, func(msg Message) (int, string) {
        return msg.code, msg.name
    })

    fmt.Println(result)

    // Output:
    // map[100:Hello 101:Hi]
}
```

### <span id="ToPointer">ToPointer</span>

<p>返回传入值的指针。</p>

<b>函数签名:</b>

```go
func ToPointer[T any](value T) *T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ASf_etHNlw1)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    result := convertor.ToPointer(123)
    fmt.Println(*result)

    // Output:
    // 123
}
```

### <span id="ToString">ToString</span>

<p>将值转换为字符串，对于数字、字符串、[]byte，将转换为字符串。 对于其他类型（切片、映射、数组、结构体）将调用 json.Marshal</p>

<b>函数签名:</b>

```go
func ToString(value any) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/nF1zOOslpQq)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    result1 := convertor.ToString("")
    result2 := convertor.ToString(nil)
    result3 := convertor.ToString(0)
    result4 := convertor.ToString(1.23)
    result5 := convertor.ToString(true)
    result6 := convertor.ToString(false)
    result7 := convertor.ToString([]int{1, 2, 3})

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)
    fmt.Println(result7)

    // Output:
    //
    //
    // 0
    // 1.23
    // true
    // false
    // [1,2,3]
}
```

### <span id="StructToMap">StructToMap</span>

<p>将struct转成map，只会转换struct中可导出的字段。struct中导出字段需要设置json tag标记。</p>

<b>函数签名:</b>

```go
func StructToMap(value any) (map[string]any, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/KYGYJqNUBOI)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
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

    fmt.Println(pm)

    // Output:
    // map[name:test]
}
```

### <span id="MapToSlice">MapToSlice</span>

<p>map中key和value执行函数iteratee后，转为切片。</p>

<b>函数签名:</b>

```go
func MapToSlice[T any, K comparable, V any](aMap map[K]V, iteratee func(K, V) T) []T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/dmX4Ix5V6Wl)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    aMap := map[string]int{"a": 1, "b": 2, "c": 3}
    result := MapToSlice(aMap, func(key string, value int) string {
        return key + ":" + strconv.Itoa(value)
    })

    fmt.Println(result) //[]string{"a:1", "b:2", "c:3"}
}
```

### <span id="EncodeByte">EncodeByte</span>

<p>将data编码成字节切片。</p>

<b>函数签名:</b>

```go
func EncodeByte(data any) ([]byte, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/DVmM1G5JfuP)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    byteData, _ := convertor.EncodeByte("abc")
    fmt.Println(byteData)

    // Output:
    // [6 12 0 3 97 98 99]
}
```

### <span id="DecodeByte">DecodeByte</span>

<p>解码字节切片到目标对象，目标对象需要传入一个指针实例。</p>

<b>函数签名:</b>

```go
func DecodeByte(data []byte, target any) error
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/zI6xsmuQRbn)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    var result string
    byteData := []byte{6, 12, 0, 3, 97, 98, 99}

    err := convertor.DecodeByte(byteData, &result)
    if err != nil {
        return
    }

    fmt.Println(result)

    // Output:
    // abc
}
```

### <span id="DeepClone">DeepClone</span>

<p>创建一个传入值的深拷贝, 无法克隆结构体的非导出字段。</p>

<b>函数签名:</b>

```go
func DeepClone[T any](src T) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/j4DP5dquxnk)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
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
func CopyProperties[T, U any](dst T, src U) (err error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/oZujoB5Sgg5)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/syqw0-WG7Xd)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/9FlIaFLArIL)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
    "github.com/duke-git/lancet/v2/validator"
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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/OphmHCN_9u8)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
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

### <span id="ToStdBase64">ToStdBase64</span>

<p>将值转换为StdBase64编码的字符串。error类型的数据也会把error的原因进行编码，复杂的结构会转为JSON格式的字符串</p>

<b>函数签名:</b>

```go
func ToStdBase64(value any) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/_fLJqJD3NMo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    afterEncode := convertor.ToStdBase64(nil)
    fmt.Println(afterEncode)

    afterEncode = convertor.ToStdBase64("")
    fmt.Println(afterEncode)

    stringVal := "hello"
    afterEncode = convertor.ToStdBase64(stringVal)
    fmt.Println(afterEncode)

    byteSliceVal := []byte("hello")
    afterEncode = convertor.ToStdBase64(byteSliceVal)
    fmt.Println(afterEncode)

    intVal := 123
    afterEncode = convertor.ToStdBase64(intVal)
    fmt.Println(afterEncode)

    mapVal := map[string]any{"a": "hi", "b": 2, "c": struct {
        A string
        B int
    }{"hello", 3}}
    afterEncode = convertor.ToStdBase64(mapVal)
    fmt.Println(afterEncode)

    floatVal := 123.456
    afterEncode = convertor.ToStdBase64(floatVal)
    fmt.Println(afterEncode)

    boolVal := true
    afterEncode = convertor.ToStdBase64(boolVal)
    fmt.Println(afterEncode)

    errVal := errors.New("err")
    afterEncode = convertor.ToStdBase64(errVal)
    fmt.Println(afterEncode)

    // Output:
    //
    //
    // aGVsbG8=
    // aGVsbG8=
    // MTIz
    // eyJhIjoiaGkiLCJiIjoyLCJjIjp7IkEiOiJoZWxsbyIsIkIiOjN9fQ==
    // MTIzLjQ1Ng==
    // dHJ1ZQ==
    // ZXJy
}

```

### <span id="ToUrlBase64">ToUrlBase64</span>

<p>值转换为 ToUrlBase64 编码的字符串。error 类型的数据也会把 error 的原因进行编码，复杂的结构会转为 JSON 格式的字符串</p>

<b>函数签名:</b>

```go
func ToUrlBase64(value any) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/C_d0GlvEeUR)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    afterEncode := convertor.ToUrlBase64(nil)
    fmt.Println(afterEncode)


    stringVal := "hello"
    afterEncode = convertor.ToUrlBase64(stringVal)
    fmt.Println(afterEncode)

    byteSliceVal := []byte("hello")
    afterEncode = convertor.ToUrlBase64(byteSliceVal)
    fmt.Println(afterEncode)

    intVal := 123
    afterEncode = convertor.ToUrlBase64(intVal)
    fmt.Println(afterEncode)

    mapVal := map[string]any{"a": "hi", "b": 2, "c": struct {
        A string
        B int
    }{"hello", 3}}
    afterEncode = convertor.ToUrlBase64(mapVal)
    fmt.Println(afterEncode)

    floatVal := 123.456
    afterEncode = convertor.ToUrlBase64(floatVal)
    fmt.Println(afterEncode)

    boolVal := true
    afterEncode = convertor.ToUrlBase64(boolVal)
    fmt.Println(afterEncode)

    errVal := errors.New("err")
    afterEncode = convertor.ToUrlBase64(errVal)
    fmt.Println(afterEncode)

    // Output:
    //
    // aGVsbG8=
    // aGVsbG8=
    // MTIz
    // eyJhIjoiaGkiLCJiIjoyLCJjIjp7IkEiOiJoZWxsbyIsIkIiOjN9fQ==
    // MTIzLjQ1Ng==
    // dHJ1ZQ==
    // ZXJy
}

```

### <span id="ToRawStdBase64">ToRawStdBase64</span>

<p>值转换为 ToRawStdBase64 编码的字符串。error 类型的数据也会把 error 的原因进行编码，复杂的结构会转为 JSON 格式的字符串</p>

<b>函数签名:</b>

```go
func ToRawStdBase64(value any) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/wSAr3sfkDcv)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {

    stringVal := "hello"
    afterEncode = convertor.ToRawStdBase64(stringVal)
    fmt.Println(afterEncode)

    byteSliceVal := []byte("hello")
    afterEncode = convertor.ToRawStdBase64(byteSliceVal)
    fmt.Println(afterEncode)

    intVal := 123
    afterEncode = convertor.ToRawStdBase64(intVal)
    fmt.Println(afterEncode)

    mapVal := map[string]any{"a": "hi", "b": 2, "c": struct {
        A string
        B int
    }{"hello", 3}}
    afterEncode = convertor.ToRawStdBase64(mapVal)
    fmt.Println(afterEncode)

    floatVal := 123.456
    afterEncode := convertor.ToRawStdBase64(floatVal)
    fmt.Println(afterEncode)

    boolVal := true
    afterEncode = convertor.ToRawStdBase64(boolVal)
    fmt.Println(afterEncode)

    errVal := errors.New("err")
    afterEncode = convertor.ToRawStdBase64(errVal)
    fmt.Println(afterEncode)

    // Output:
    // aGVsbG8
    // aGVsbG8
    // MTIz
    // eyJhIjoiaGkiLCJiIjoyLCJjIjp7IkEiOiJoZWxsbyIsIkIiOjN9fQ
    // MTIzLjQ1Ng
    // dHJ1ZQ
    // ZXJy
}
```

### <span id="ToRawUrlBase64">ToRawUrlBase64</span>

<p>值转换为 ToRawUrlBase64 编码的字符串。error 类型的数据也会把 error 的原因进行编码，复杂的结构会转为 JSON 格式的字符串</p>

<b>函数签名:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/HwdDPFcza1O)</span></b>

```go
func ToRawUrlBase64(value any) string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {

    stringVal := "hello"
    afterEncode := convertor.ToRawUrlBase64(stringVal)
    fmt.Println(afterEncode)

    byteSliceVal := []byte("hello")
    afterEncode = convertor.ToRawUrlBase64(byteSliceVal)
    fmt.Println(afterEncode)

    intVal := 123
    afterEncode = convertor.ToRawUrlBase64(intVal)
    fmt.Println(afterEncode)

    mapVal := map[string]any{"a": "hi", "b": 2, "c": struct {
        A string
        B int
    }{"hello", 3}}
    afterEncode = convertor.ToRawUrlBase64(mapVal)
    fmt.Println(afterEncode)

    floatVal := 123.456
    afterEncode = convertor.ToRawUrlBase64(floatVal)
    fmt.Println(afterEncode)

    boolVal := true
    afterEncode = convertor.ToRawUrlBase64(boolVal)
    fmt.Println(afterEncode)

    errVal := errors.New("err")
    afterEncode = convertor.ToRawUrlBase64(errVal)
    fmt.Println(afterEncode)

    // Output:
    // aGVsbG8
    // aGVsbG8
    // MTIz
    // eyJhIjoiaGkiLCJiIjoyLCJjIjp7IkEiOiJoZWxsbyIsIkIiOjN9fQ
    // MTIzLjQ1Ng
    // dHJ1ZQ
    // ZXJy
}
```