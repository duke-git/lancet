# Random
random随机数生成器包，可以生成随机[]bytes, int, string。

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/random/random.go](https://github.com/duke-git/lancet/blob/main/random/random.go)


<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/v2/random"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [RandBytes](#RandBytes)
- [RandInt](#RandInt)
- [RandString](#RandString)
- [UUIdV4](#UUIdV4)


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
    "github.com/duke-git/lancet/v2/random"
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
    "github.com/duke-git/lancet/v2/random"
)

func main() {
	rInt := random.RandInt(1, 10)
	fmt.Println(rInt)
}
```



### <span id="RandString">RandInt</span>
<p>生成随机给定长度的随机字符串</p>

<b>函数签名:</b>

```go
func RandString(length int) string
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
	randStr := random.RandString(6)
	fmt.Println(randStr)
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
    "github.com/duke-git/lancet/v2/random"
)

func main() {
	uuid, err := random.UUIdV4()
    if err != nil {
        return
    }
	fmt.Println(uuid)
}
```


