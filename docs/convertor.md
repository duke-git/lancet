# Convertor
Package convertor contains some functions for data type convertion.

<div STYLE="page-break-after: always;"></div>

## Source:

- [https://github.com/duke-git/lancet/blob/main/convertor/convertor.go](https://github.com/duke-git/lancet/blob/main/convertor/convertor.go)

<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/duke-git/lancet/v2/convertor"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
- [ColorHexToRGB](#ColorHexToRGB)
- [ColorRGBToHex](#ColorRGBToHex)
- [ToBool](#ToBool)
- [ToBytes](#ToBytes)
- [ToChar](#ToChar)
- [ToChannel](#ToChannel)
- [ToFloat](#ToFloat)
- [ToInt](#ToInt)
- [ToJson](#ToJson)
- [ToMap](#ToMap)
- [ToPointer](#ToPointer)
- [ToString](#ToString)
- [StructToMap](#StructToMap)
- [MapToSlice](#MapToSlice)
- [EncodeByte](#EncodeByte)
- [DecodeByte](#DecodeByte)

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
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    colorHex := "#003366"
    r, g, b := convertor.ColorHexToRGB(colorHex)
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
    "github.com/duke-git/lancet/v2/convertor"
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

<p>Convert string to bool. Use strconv.ParseBool.</p>

<b>Signature:</b>

```go
func ToBool(s string) (bool, error)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
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

<p>Convert value to byte slice.</p>

<b>Signature:</b>

```go
func ToBytes(data any) ([]byte, error)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
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
    "github.com/duke-git/lancet/v2/convertor"
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

<p>Convert a collection of elements to a read-only channel.</p>

<b>Signature:</b>

```go
func ToChannel[T any](array []T) <-chan T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
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

<p>Convert value to a float64 value. If param is a invalid floatable, will return 0.0 and error. </p>

<b>Signature:</b>

```go
func ToFloat(value any) (float64, error)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
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

<p>Convert value to a int64 value. If param is a invalid intable, will return 0 and error. </p>

<b>Signature:</b>

```go
func ToInt(value any) (int64, error)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
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
func ToJson(value any) (string, error)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    var aMap = map[string]int{"a": 1, "b": 2, "c": 3}
    jsonStr, _ := convertor.ToJson(aMap)
    fmt.Printf("%q", jsonStr) //"{\"a\":1,\"b\":2,\"c\":3}"
}
```



### <span id="ToMap">ToMap</span>

<p>Convert a slice of structs to a map based on iteratee function. </p>

<b>Signature:</b>

```go
func ToMap[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K]V
```
<b>Example:</b>

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

    fmt.Println(result) //{100: "Hello", 101: "Hi"}
}
```



### <span id="ToPointer">ToPointer</span>

<p>Returns a pointer to passed value. </p>

<b>Signature:</b>

```go
func ToPointer[T any](value T) *T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    result := convertor.ToPointer(123)
    fmt.Println(*result) //123
}
```


### <span id="ToString">ToString</span>

<p>ToString convert value to string, for number, string, []byte, will convert to string. For other type (slice, map, array, struct) will call json.Marshal</p>

<b>Signature:</b>

```go
func ToString(value any) string
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    fmt.Printf("%q", convertor.ToString(1)) //"1"
    fmt.Printf("%q", convertor.ToString(1.1)) //"1.1"
    fmt.Printf("%q", convertor.ToString([]int{1, 2, 3})) //"[1,2,3]"
}
```


### <span id="StructToMap">StructToMap</span>

<p>Convert struct to map, only convert exported field, struct field tag `json` should be set.</p>

<b>Signature:</b>

```go
func StructToMap(value any) (map[string]any, error)
```
<b>Example:</b>

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

    fmt.Printf("type: %T, value: %s", pm, pm) //type: map[string]interface {}, value: map[name:test]
}
```



### <span id="MapToSlice">MapToSlice</span>

<p>Convert a map to a slice based on iteratee function.</p>

<b>Signature:</b>

```go
func MapToSlice[T any, K comparable, V any](aMap map[K]V, iteratee func(K, V) T) []T
```
<b>Example:</b>

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
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
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
func DecodeByte(data []byte, target any) error
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    var result string
	byteData := []byte{6, 12, 0, 3, 97, 98, 99}
	convertor.DecodeByte(byteData, &result)
    fmt.Println(result) //"abc"
}
```