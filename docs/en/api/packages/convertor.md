# Convertor

Package convertor contains some functions for data type convertion.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/convertor/convertor.go](https://github.com/duke-git/lancet/blob/main/convertor/convertor.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/convertor"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

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
-   [ToBigInt](#ToBigInt)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="ColorHexToRGB">ColorHexToRGB</span>

<p>Convert color hex to color rgb.</p>

<b>Signature:</b>

```go
func ColorHexToRGB(colorHex string) (red, green, blue int)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/o7_ft-JCJBV)</span></b>

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

<p>Convert color rgb to color hex.</p>

<b>Signature:</b>

```go
func ColorRGBToHex(red, green, blue int) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/nzKS2Ro87J1)</span></b>

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

<p>Convert string to bool. Use strconv.ParseBool.</p>

<b>Signature:</b>

```go
func ToBool(s string) (bool, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/ARht2WnGdIN)</span></b>

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

<p>Convert value to byte slice.</p>

<b>Signature:</b>

```go
func ToBytes(data any) ([]byte, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/fAMXYFDvOvr)</span></b>

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

<p>Convert string to char slice.</p>

<b>Signature:</b>

```go
func ToChar(s string) []string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/JJ1SvbFkVdM)</span></b>

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

<p>Convert a collection of elements to a read-only channel.</p>

<b>Signature:</b>

```go
func ToChannel[T any](array []T) <-chan T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/hOx_oYZbAnL)</span></b>

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

<p>Convert value to a float64 value. If param is a invalid floatable, will return 0.0 and error. </p>

<b>Signature:</b>

```go
func ToFloat(value any) (float64, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/4YTmPCibqHJ)</span></b>

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

<p>Convert value to a int64 value. If param is a invalid intable, will return 0 and error. </p>

<b>Signature:</b>

```go
func ToInt(value any) (int64, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/9_h9vIt-QZ_b)</span></b>

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

<p>Convert interface to json string. If param can't be converted, will return "" and error. </p>

<b>Signature:</b>

```go
func ToJson(value any) (string, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/2rLIkMmXWvR)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    aMap := map[string]int{"a": 1, "b": 2, "c": 3}
    result, err := convertor.ToJson(aMap)

    if err != nil {
        fmt.Printf("%v", err)
    }

    fmt.Println(result)

    // Output:
    // {"a":1,"b":2,"c":3}
}
```

### <span id="ToMap">ToMap</span>

<p>Convert a slice of structs to a map based on iteratee function. </p>

<b>Signature:</b>

```go
func ToMap[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K]V
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/tVFy7E-t24l)</span></b>

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

<p>Returns a pointer to passed value. </p>

<b>Signature:</b>

```go
func ToPointer[T any](value T) *T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/ASf_etHNlw1)</span></b>

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

<p>ToString convert value to string, for number, string, []byte, will convert to string. For other type (slice, map, array, struct) will call json.Marshal</p>

<b>Signature:</b>

```go
func ToString(value any) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/nF1zOOslpQq)</span></b>

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

<p>Convert struct to map, only convert exported field, struct field tag `json` should be set.</p>

<b>Signature:</b>

```go
func StructToMap(value any) (map[string]any, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/KYGYJqNUBOI)</span></b>

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

<p>Convert a map to a slice based on iteratee function.</p>

<b>Signature:</b>

```go
func MapToSlice[T any, K comparable, V any](aMap map[K]V, iteratee func(K, V) T) []T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/dmX4Ix5V6Wl)</span></b>

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

<p>Encode data to byte slice.</p>

<b>Signature:</b>

```go
func EncodeByte(data any) ([]byte, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/DVmM1G5JfuP)</span></b>

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

<p>Decode byte data to target object. target should be a pointer instance.</p>

<b>Signature:</b>

```go
func DecodeByte(data []byte, target any) error
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/zI6xsmuQRbn)</span></b>

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


### <span id="CopyProperties">CopyProperties</span>

<p>Copies each field from the source struct into the destination struct. Use json.Marshal/Unmarshal, so json tag should be set for fields of dst and src struct.</p>

<b>Signature:</b>

```go
func CopyProperties[T, U any](dst T, src U) (err error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/oZujoB5Sgg5)</span></b>

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

### <span id="Utf8ToGbk">Utf8ToGbk</span>

<p>Converts utf8 encoding data to GBK encoding data.</p>

<b>Signature:</b>

```go
func Utf8ToGbk(bs []byte) ([]byte, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/9FlIaFLArIL)</span></b>

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

<p>Converts GBK encoding data to utf8 encoding data.</p>

<b>Signature:</b>

```go
func GbkToUtf8(bs []byte) ([]byte, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/OphmHCN_9u8)</span></b>

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

<p>Convert a value to a string encoded in standard Base64. Error data of type "error" will also be encoded, and complex structures will be converted to a JSON formatted string.</p>

<b>Signature:</b>

```go
func ToStdBase64(value any) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/_fLJqJD3NMo)</span></b>

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

<p>Convert a value to a string encoded in url Base64. Error data of type "error" will also be encoded, and complex structures will be converted to a JSON formatted string.</p>

<b>Signature:</b>

```go
func ToUrlBase64(value any) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/C_d0GlvEeUR)</span></b>

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

<p>Convert a value to a string encoded in raw standard Base64. Error data of type "error" will also be encoded, and complex structures will be converted to a JSON formatted string.</p>

<b>Signature:</b>

```go
func ToRawStdBase64(value any) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/wSAr3sfkDcv)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {

    stringVal := "hello"
    afterEncode := convertor.ToRawStdBase64(stringVal)
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
    afterEncode = convertor.ToRawStdBase64(floatVal)
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

<p> Convert a value to a string encoded in raw url Base64. Error data of type "error" will also be encoded, and complex structures will be converted to a JSON formatted string.</p>

<b>Signature:</b>

```go
func ToRawUrlBase64(value any) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/HwdDPFcza1O)</span></b>

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

### <span id="DeepClone">DeepClone</span>

<p>Creates a deep copy of passed item, can't clone unexported field of struct.</p>

<b>Signature:</b>

```go
func DeepClone[T any](src T) T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/j4DP5dquxnk)</span></b>

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

### <span id="ToBigInt">ToBigInt</span>

<p>Converts an integer of any supported type (int, int64, uint64, etc.) to *big.Int</p>

<b>Signature:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/X3itkCxwB_x)</span></b>

```go
func ToBigInt[T any](v T) (*big.Int, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    n := 9876543210
    bigInt, _ := convertor.ToBigInt(n)

    fmt.Println(bigInt)
    // Output:
    // 9876543210
}
```