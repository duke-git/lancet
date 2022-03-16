# Convertor
Package convertor contains some functions for data type convertion.

<div STYLE="page-break-after: always;"></div>

## Source:

- [https://github.com/duke-git/lancet/blob/main/convertor/convertor.go](https://github.com/duke-git/lancet/blob/main/convertor/convertor.go)

<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/duke-git/lancet/convertor"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
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

<p>Convert struct to map, only convert exported field, struct field tag `json` should be set.</p>

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