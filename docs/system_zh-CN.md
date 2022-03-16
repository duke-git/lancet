# System
system包含os, runtime, shell command相关函数。

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/system/os.go](https://github.com/duke-git/lancet/blob/main/system/os.go)


<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/system"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [IsWindows](#IsWindows)
- [IsLinux](#IsLinux)
- [IsMac](#IsMac)
- [GetOsEnv](#GetOsEnv)
- [SetOsEnv](#SetOsEnv)
- [RemoveOsEnv](#RemoveOsEnv)
- [CompareOsEnv](#CompareOsEnv)
- [ExecCommand](#ExecCommand)
  

<div STYLE="page-break-after: always;"></div>

## Documentation文档


### <span id="IsWindows">IsWindows</span>
<p>检查当前操作系统是否是windows</p>

<b>Signature:</b>

```go
func IsWindows() bool
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/system"
)

func main() {
	isOsWindows := system.IsWindows()
	fmt.Println(isOsWindows)
}
```




### <span id="IsLinux">IsLinux</span>
<p>检查当前操作系统是否是linux</p>

<b>Signature:</b>

```go
func IsLinux() bool
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/system"
)

func main() {
	isOsLinux := system.IsLinux()
	fmt.Println(isOsLinux)
}
```



### <span id="IsMac">IsMac</span>
<p>检查当前操作系统是否是macos</p>

<b>Signature:</b>

```go
func IsMac() bool
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/system"
)

func main() {
	isOsMac := system.IsMac
	fmt.Println(isOsMac)
}
```



### <span id="GetOsEnv">GetOsEnv</span>
<p>获取key命名的环境变量的值</p>

<b>Signature:</b>

```go
func GetOsEnv(key string) string
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/system"
)

func main() {
	fooEnv := system.GetOsEnv("foo")
	fmt.Println(fooEnv)
}
```



### <span id="SetOsEnv">SetOsEnv</span>
<p>设置由key命名的环境变量的值</p>

<b>Signature:</b>

```go
func SetOsEnv(key, value string) error
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/system"
)

func main() {
	err := system.SetOsEnv("foo", "foo_value")
	fmt.Println(err)
}
```




### <span id="RemoveOsEnv">RemoveOsEnv</span>
<p>删除单个环境变量</p>

<b>Signature:</b>

```go
func RemoveOsEnv(key string) error
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/system"
)

func main() {
	err := system.RemoveOsEnv("foo")
	if err != nil {
		fmt.Println(err)
	}
}
```



### <span id="CompareOsEnv">CompareOsEnv</span>
<p>获取key命名的环境变量值并与compareEnv进行比较</p>

<b>Signature:</b>

```go
func CompareOsEnv(key, comparedEnv string) bool
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/system"
)

func main() {
	system.SetOsEnv("foo", "foo_value")
	res := system.CompareOsEnv("foo", "foo_value")
	fmt.Println(res) //true
}
```




### <span id="ExecCommand">CompareOsEnv</span>
<p>使用shell /bin/bash -c(linux) 或 cmd (windows) 执行shell命令</p>

<b>Signature:</b>

```go
func ExecCommand(command string) (stdout, stderr string, err error)
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/system"
)

func main() {
	out, errout, err := system.ExecCommand("ls")
	fmt.Println(out)
	fmt.Println(errout)
	fmt.Println(err)
}
```








