# Convertor
convertor转换器包支持一些常见的数据类型转换

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

- [Convertor](#convertor)
  - [源码:](#源码)
  - [用法:](#用法)
  - [目录](#目录)
  - [文档](#文档)
    - [<span id="ColorHexToRGB">ColorHexToRGB</span>](#colorhextorgb)
    - [<span id="ColorRGBToHex">ColorRGBToHex</span>](#colorrgbtohex)
    - [<span id="ToBool">ToBool</span>](#tobool)
    - [<span id="ToBytes">ToBytes</span>](#tobytes)
    - [<span id="ToChar">ToChar</span>](#tochar)
    - [<span id="ToChannel">ToChannel</span>](#tochannel)
    - [<span id="ToFloat">ToFloat</span>](#tofloat)
    - [<span id="ToInt">ToInt</span>](#toint)
    - [<span id="ToJson">ToJson</span>](#tojson)
    - [<span id="ToString">ToString</span>](#tostring)
    - [<span id="StructToMap">StructToMap</span>](#structtomap)
    - [<span id="EncodeByte">EncodeByte</span>](#encodebyte)
    - [<span id="DecodeByte">DecodeByte</span>](#decodebyte)
  

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
    "github.com/duke-git/lancet/v2/convertor"
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
    "github.com/duke-git/lancet/v2/convertor"
)

func main() {
    var result string
	byteData := []byte{6, 12, 0, 3, 97, 98, 99}
	convertor.DecodeByte(byteData, &result)
    fmt.Println(result) //"abc"
}
```