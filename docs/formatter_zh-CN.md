# Formatter
formatter格式化器包含一些数据格式化处理方法。

<div STYLE="page-break-after: always;"></div>

## 源码:

[https://github.com/duke-git/lancet/blob/main/formatter/formatter.go](https://github.com/duke-git/lancet/blob/main/formatter/formatter.go)

<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/formatter"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [Comma](#Comma)

<div STYLE="page-break-after: always;"></div>

## 文档



### <span id="Comma">Comma</span>
<p>用逗号每隔3位分割数字/字符串，签名添加符号。参数必须是数字或者可以转为数字的字符串</p>

<b>函数签名:</b>

```go
func Comma(v interface{}, symbol string) string
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/formatter"
)

func main() {
    fmt.Println(formatter.Comma("12345", ""))   // "12,345"
    fmt.Println(formatter.Comma(12345.67, "¥")) // "¥12,345.67"
}
```
