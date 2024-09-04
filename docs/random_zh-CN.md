# Random

random 随机数生成器包，可以生成随机[]bytes, int, string。

<div STYLE="page-break-after: always;"></div>

## 源码:

[https://github.com/duke-git/lancet/blob/v1/random/random.go](https://github.com/duke-git/lancet/blob/v1/random/random.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/random"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [RandBytes](#RandBytes)
-   [RandInt](#RandInt)
-   [RandString](#RandString)
-   [RandUpper](#RandUpper)
-   [RandLower](#RandLower)
-   [RandNumeral](#RandNumeral)
-   [RandNumeralOrLetter](#RandNumeralOrLetter)
-   [UUIdV4](#UUIdV4)
-   [RandFloat](#RandFloat)
-   [RandFloats](#RandFloats)
-   [RandUniqueIntSlice](#RandUniqueIntSlice)
-   [RandIntSlice](#RandIntSlice)
-   [RandStringSlice](#RandStringSlice)
-   [RandBool](#RandBool)
-   [RandBoolSlice](#RandBoolSlice)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="RandBytes">RandBytes</span>

<p>生成随机字节切片</p>

<b>函数签名:</b>

```go
func RandBytes(length int) []byte
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    randBytes := random.RandBytes(4)
    fmt.Println(randBytes)
}
```

### <span id="RandInt">RandInt</span>

<p>生成随机int, 范围[min, max)</p>

<b>函数签名:</b>

```go
func RandInt(min, max int) int
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    rInt := random.RandInt(1, 10)
    fmt.Println(rInt)
}
```

### <span id="RandString">RandString</span>

<p>生成给定长度的随机字符串，只包含字母(a-zA-Z)</p>

<b>函数签名:</b>

```go
func RandString(length int) string
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    randStr := random.RandString(6)
    fmt.Println(randStr) //pGWsze
}
```

### <span id="RandUpper">RandUpper</span>

<p>生成给定长度的随机大写字母字符串</p>

<b>函数签名:</b>

```go
func RandUpper(length int) string
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    randStr := random.RandString(6)
    fmt.Println(randStr) //PACWGF
}
```

### <span id="RandLower">RandLower</span>

<p>生成给定长度的随机小写字母字符串</p>

<b>函数签名:</b>

```go
func RandLower(length int) string
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    randStr := random.RandLower(6)
    fmt.Println(randStr) //siqbew
}
```

### <span id="RandNumeral">RandNumeral</span>

<p>生成给定长度的随机数字字符串</p>

<b>函数签名:</b>

```go
func RandNumeral(length int) string
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    randStr := random.RandNumeral(6)
    fmt.Println(randStr) //035172
}
```

### <span id="RandNumeralOrLetter">RandNumeralOrLetter</span>

<p>生成给定长度的随机字符串（数字+字母)</p>

<b>函数签名:</b>

```go
func RandNumeralOrLetter(length int) string
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    randStr := random.RandNumeralOrLetter(6)
    fmt.Println(randStr) //0aW7cQ
}
```

### <span id="UUIdV4">UUIdV4</span>

<p>生成UUID v4字符串</p>

<b>函数签名:</b>

```go
func UUIdV4() (string, error)
```

<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    uuid, err := random.UUIdV4()
    if err != nil {
        return
    }
    fmt.Println(uuid)
}
```

### <span id="RandFloat">RandFloat</span>

<p>Generate a random float64 number between [min, max) with specific precision.</p>

<b>Signature:</b>

```go
func RandFloat(min, max float64, precision int) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/zbD_tuobJtr)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    floatNumber := random.RandFloat(1.0, 5.0, 2)
    fmt.Println(floatNumber) //2.14 (random number)
}
```

### <span id="RandFloats">RandFloats</span>

<p>Generate a slice of random float64 numbers of length n that do not repeat. Number range in [min, max)</p>

<b>Signature:</b>

```go
func RandFloats(length int, min, max float64, precision int) []float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/I3yndUQ-rhh)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    floatNumbers := random.RandFloats(5, 1.0, 5.0, 2)
    fmt.Println(floatNumbers) //[3.42 3.99 1.3 2.38 4.23] (random)
}
```

### <span id="RandUniqueIntSlice">RandUniqueIntSlice</span>

<p>生成一个不重复的长度为n的随机int切片。</p>

<b>函数签名:</b>

```go
func RandUniqueIntSlice(n, min, max int) []int
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    result := RandUniqueIntSlice(5, 0, 10)
    fmt.Println(result) //[0 4 7 1 5] (random)
}
```

### <span id="RandIntSlice">RandIntSlice</span>

<p>生成一个特定长度的随机int切片，数值范围[min, max)。</p>

<b>函数签名:</b>

```go
func RandIntSlice(length, min, max int) []int
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    result := random.RandIntSlice(5, 0, 10)
    fmt.Println(result) //[1 2 7 1 5] (random)
}
```

### <span id="RandStringSlice">RandStringSlice</span>

<p>生成随机字符串slice. 字符串类型需要是以下几种或者它们的组合: random.Numeral, random.LowwerLetters, random.UpperLetters random.Letters, random.SymbolChars, random.AllChars。</p>

<b>函数签名:</b>

```go
func RandStringSlice(charset string, sliceLen, strLen int) []string
```

<b>实例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    strs := random.RandStringSlice(random.Letters, 4, 6)
    fmt.Println(strs)

    // output random string slice like below:
    //[CooSMq RUFjDz FAeMPf heRyGv]
}
```

### <span id="RandBool">RandBool</span>

<p>生成随机bool值(true or false)。</p>

<b>函数签名:</b>

```go
func RandBool() bool
```

<b>实例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    result := random.RandBool()
    fmt.Println(result) // true or false (random)
}
```

### <span id="RandBoolSlice">RandBoolSlice</span>

<p>生成特定长度的随机bool slice。</p>

<b>函数签名:</b>

```go
func RandBoolSlice(length int) []bool
```

<b>实例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    result := random.RandBoolSlice(2)
    fmt.Println(result) // [true false] (random)
}
```