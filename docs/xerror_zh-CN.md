# Xerror
xerror错误处理逻辑封装

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/xerror/xerror.go](https://github.com/duke-git/lancet/blob/main/xerror/xerror.go)

<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/xerror"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [Unwrap](#Unwrap)

<div STYLE="page-break-after: always;"></div>

## 文档



### <span id="Unwrap">Unwrap</span>
<p>如果err为nil则展开，则它返回一个有效值。 如果err不是nil则Unwrap使用err发生恐慌。</p>

<b>函数签名:</b>

```go
func Unwrap[T any](val T, err error) T
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/xerror"
)

func main() {
    _, err := strconv.Atoi("4o2")
	defer func() {
		v := recover()
		fmt.Println(err.Error()) // err.Error() == v.(*strconv.NumError).Error()
	}()

	xerror.Unwrap(strconv.Atoi("4o2"))
}
```
