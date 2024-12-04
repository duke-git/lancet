# Convertor

Package convertor contains some functions for data type convertion.

<div STYLE="page-break-after: always;"></div>

## Source:

[https://github.com/duke-git/lancet/blob/v1/convertor/convertor.go](https://github.com/duke-git/lancet/blob/v1/convertor/convertor.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/convertor"
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

<b>Example:</b>

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

<p>Convert color rgb to color hex.</p>

<b>Signature:</b>

```go
func ColorRGBToHex(red, green, blue int) string
```

<b>Example:</b>

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

<p>Convert string to a boolean value. Use strconv.ParseBool</p>

<b>Signature:</b>

```go
func ToBool(s string) (bool, error)
```

<b>Example:</b>

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

<p>Convert interface to byte slice.</p>

<b>Signature:</b>

```go
func ToBytes(data interface{}) ([]byte, error)
```

<b>Example:</b>

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

<p>Convert string to char slice.</p>

<b>Signature:</b>

```go
func ToChar(s string) []string
```

<b>Example:</b>

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

<p>Convert a collection of elements to a read-only channels.</p>

<b>Signature:</b>

```go
func ToChannel(array []interface{}) <-chan interface{}
```

<b>Example:</b>

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

<p>Convert interface to a float64 value. If param is a invalid floatable, will return 0 and error. </p>

<b>Signature:</b>

```go
func ToFloat(value interface{}) (float64, error)
```

<b>Example:</b>

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

<p>Convert interface to a int64 value. If param is a invalid intable, will return 0 and error. </p>

<b>Signature:</b>

```go
func ToInt(value interface{}) (int64, error)
```

<b>Example:</b>

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

<p>Convert interface to json string. If param can't be converted, will return "" and error. </p>

<b>Signature:</b>

```go
func ToJson(value interface{}) (string, error)
```

<b>Example:</b>

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

<p>Convert interface to string. </p>

<b>Signature:</b>

```go
func ToString(value interface{}) string
```

<b>Example:</b>

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

<p>Converts struct to map, only convert exported field, struct field tag `json` should be set.</p>

<b>Signature:</b>

```go
func StructToMap(value interface{}) (map[string]interface{}, error)
```

<b>Example:</b>

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

<p>Converts map to struct, only convert exported field, struct field tag `json` should be set.</p>

<b>Signature:</b>

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

<p>Encode data to byte slice.</p>

<b>Signature:</b>

```go
func EncodeByte(data interface{}) ([]byte, error)
```

<b>Example:</b>

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

<p>Decode byte data to target object. target should be a pointer instance.</p>

<b>Signature:</b>

```go
func DecodeByte(data []byte, target interface{}) error
```

<b>Example:</b>

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

<p>Creates a deep copy of passed item, can't clone unexported field of struct.</p>

<b>Signature:</b>

```go
func DeepClone[T interface{}](src T) T
```

<b>Example:</b>

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

<p>Copies each field from the source struct into the destination struct. Use json.Marshal/Unmarshal, so json tag should be set for fields of dst and src struct.</p>

<b>Signature:</b>

```go
func CopyProperties(dst, src interface{}) error
```

<b>Example:</b>

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

<p>Converts reflect value to its interface type.</p>

<b>Signature:</b>

```go
func ToInterface(v reflect.Value) (value interface{}, ok bool)
```

<b>Example:</b>

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

<p>Converts utf8 encoding data to GBK encoding data.</p>

<b>Signature:</b>

```go
func Utf8ToGbk(bs []byte) ([]byte, error)
```

<b>Example:</b>

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

<p>Converts GBK encoding data to utf8 encoding data.</p>

<b>Signature:</b>

```go
func GbkToUtf8(bs []byte) ([]byte, error)
```

<b>Example:</b>

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

### <span id="ToStdBase64">ToStdBase64</span>

<p>Convert a value to a string encoded in standard Base64. Error data of type "error" will also be encoded, and complex structures will be converted to a JSON-formatted string.</p>

<b>Signature:</b>

```go
func ToStdBase64(value interface{}) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
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

    mapVal := map[string]interface{}{"a": "hi", "b": 2, "c": struct {
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

<p>Convert a value to a string encoded in url Base64. Error data of type "error" will also be encoded, and complex structures will be converted to a JSON-formatted string.</p>

<b>Signature:</b>

```go
func ToUrlBase64(value interface{}) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
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

    mapVal := map[string]interface{}{"a": "hi", "b": 2, "c": struct {
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

<p>Convert a value to a string encoded in raw standard Base64. Error data of type "error" will also be encoded, and complex structures will be converted to a JSON-formatted string.</p>

<b>Signature:</b>

```go
func ToRawStdBase64(value interface{}) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
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

    mapVal := map[string]interface{}{"a": "hi", "b": 2, "c": struct {
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

<p> Convert a value to a string encoded in raw url Base64. Error data of type "error" will also be encoded, and complex structures will be converted to a JSON-formatted string.</p>

<b>Signature:</b>

```go
func ToRawUrlBase64(value interface{}) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {

    stringVal := "hello"
    afterEncode = convertor.ToRawUrlBase64(stringVal)
    fmt.Println(afterEncode)

    byteSliceVal := []byte("hello")
    afterEncode = convertor.ToRawUrlBase64(byteSliceVal)
    fmt.Println(afterEncode)

    intVal := 123
    afterEncode = convertor.ToRawUrlBase64(intVal)
    fmt.Println(afterEncode)

    mapVal := map[string]interface{}{"a": "hi", "b": 2, "c": struct {
        A string
        B int
    }{"hello", 3}}
    afterEncode = convertor.ToRawUrlBase64(mapVal)
    fmt.Println(afterEncode)

    floatVal := 123.456
    afterEncode = convertor.ToRawUrlBase64(floatVal)
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

### <span id="ToBigInt">ToBigInt</span>

<p>Convert value to bigInt.</p>

<b>Signature:</b>

```go
func ToBigInt(v interface{}) (*big.Int, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/convertor"
)

func main() {
    n := 9876543210
    bigInt, _ := convertor.ToBigInt(n)

    fmt.Println(bigInt)
    // Output:
    // 9876543210
}
```