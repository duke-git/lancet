# System

system 包含 os, runtime, shell command 相关函数。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/system/os.go](https://github.com/duke-git/lancet/blob/main/system/os.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/system"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [IsWindows](#IsWindows)
-   [IsLinux](#IsLinux)
-   [IsMac](#IsMac)
-   [GetOsEnv](#GetOsEnv)
-   [SetOsEnv](#SetOsEnv)
-   [RemoveOsEnv](#RemoveOsEnv)
-   [CompareOsEnv](#CompareOsEnv)
-   [ExecCommand](#ExecCommand)
-   [GetOsBits](#GetOsBits)

<div STYLE="page-break-after: always;"></div>

## Documentation 文档

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
    "github.com/duke-git/lancet/v2/system"
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
    "github.com/duke-git/lancet/v2/system"
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
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    isOsMac := system.IsMac()
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
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    err := system.SetOsEnv("foo", "abc")
    result := system.GetOsEnv("foo")

    fmt.Println(err)
    fmt.Println(result)
    // Output:
    // <nil>
    // abc
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
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    err := system.SetOsEnv("foo", "abc")
    result := system.GetOsEnv("foo")

    fmt.Println(err)
    fmt.Println(result)
    // Output:
    // <nil>
    // abc
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
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    err1 := system.SetOsEnv("foo", "abc")
    result1 := GetOsEnv("foo")

    err2 := system.RemoveOsEnv("foo")
    result2 := GetOsEnv("foo")

    fmt.Println(err1)
    fmt.Println(err2)
    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // <nil>
    // <nil>
    // abc
    //
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
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    err := system.SetOsEnv("foo", "abc")
    if err != nil {
        return
    }

    result := system.CompareOsEnv("foo", "abc")

    fmt.Println(result)

    // Output:
    // true
}
```

### <span id="ExecCommand">ExecCommand</span>

<p>执行shell命令，返回命令的stdout和stderr字符串，如果出现错误，则返回错误。参数`command`是一个完整的命令字符串，如ls-a（linux），dir（windows），ping 127.0.0.1。在linux中，使用/bin/bash-c执行命令，在windows中，使用powershell.exe执行命令。</p>

<b>Signature:</b>

```go
func ExecCommand(command string) (stdout, stderr string, err error)
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    // linux or mac
    stdout, stderr, err := system.ExecCommand("ls")
    fmt.Println("std out: ", stdout)
    fmt.Println("std err: ", stderr)
    assert.Equal("", stderr)

    // windows
    stdout, stderr, err = system.ExecCommand("dir")
    fmt.Println("std out: ", stdout)
    fmt.Println("std err: ", stderr)

    // error command
    stdout, stderr, err = system.ExecCommand("abc")
    fmt.Println("std out: ", stdout)
    fmt.Println("std err: ", stderr)
    if err != nil {
        fmt.Println(err.Error())
    }
}
```

### <span id="GetOsBits">GetOsBits</span>

<p>获取当前操作系统位数，返回32或64</p>

<b>函数签名:</b>

```go
func GetOsBits() int
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    osBit := system.GetOsBits()
    fmt.Println(osBit) // 32 or 64
}
```
