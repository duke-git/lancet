# Convertor
convertor转换器包支持一些常见的数据类型转换

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/convertor/convertor.go](https://github.com/duke-git/lancet/blob/main/convertor/convertor.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/convertor"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

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
    "github.com/duke-git/lancet/v2/convertor"
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

<p>interface转字节切片.</p>

<b>函数签名:</b>

```go
func ToBytes(data any) ([]byte, error)
```
<b>列子:</b>

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

<p>将切片转为只读channel</p>

<b>函数签名:</b>

```go
func ToChannel[T any](array []T) <-chan T
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
func ToFloat(value any) (float64, error)
```
<b>列子:</b>

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

<p>将interface转成intt64类型，如果参数无法转换，会返回0和error</p>

<b>函数签名:</b>

```go
func ToInt(value any) (int64, error)
```
<b>例子:</b>

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

<p>将interface转成json字符串，如果参数无法转换，会返回""和error</p>

<b>函数签名:</b>

```go
func ToJson(value any) (string, error)
```
<b>列子:</b>

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

<p>将切片转为map</p>

<b>函数签名:</b>

```go
func ToMap[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K]V
```
<b>例子:</b>

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

<p>返回传入值的指针</p>

<b>函数签名:</b>

```go
func ToPointer[T any](value T) *T
```
<b>例子:</b>

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

<p>将interface转成字符串</p>

<b>函数签名:</b>

```go
func ToString(value any) string
```
<b>例子:</b>

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

<p>将struct转成map，只会转换struct中可导出的字段。struct中导出字段需要设置json tag标记</p>

<b>函数签名:</b>

```go
func StructToMap(value any) (map[string]any, error)
```
<b>列子:</b>

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

<p>map中key和value执行函数iteratee后，转为切片</p>

<b>函数签名:</b>

```go
func MapToSlice[T any, K comparable, V any](aMap map[K]V, iteratee func(K, V) T) []T
```
<b>例子:</b>

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