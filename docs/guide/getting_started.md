---
outline: deep
---

# 安装

1. <b>使用 go1.18 及以上版本的用户，建议安装 v2.x.x。 因为 v2.x.x 应用 go1.18 的泛型重写了大部分函数。</b>

```go
go get github.com/duke-git/lancet/v2 // will install latest version of v2.x.x
```

2. <b>使用 go1.18 以下版本的用户，必须安装 v1.x.x。目前最新的 v1 版本是 v1.4.5。</b>

```go
go get github.com/duke-git/lancet // below go1.18, install latest version of v1.x.x
```

## 用法

lancet 是以包的结构组织代码的，使用时需要导入相应的包名。例如：如果使用字符串相关函数，需要导入 strutil 包:

```go
import "github.com/duke-git/lancet/v2/strutil"
```

## 示例

此处以字符串工具函数 Reverse（逆序字符串）为例，需要导入 strutil 包:

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    s := "hello"
    rs := strutil.Reverse(s)
    fmt.Println(rs) //olleh
}
```

## 更多

更多特性请参考[API](https://www.golancet.cn/api/overview.html).
