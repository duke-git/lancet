# System

Package system contains some functions about os, runtime, shell command.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/system/os.go](https://github.com/duke-git/lancet/blob/main/system/os.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/system"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [IsWindows](#IsWindows)
-   [IsLinux](#IsLinux)
-   [IsMac](#IsMac)
-   [GetOsEnv](#GetOsEnv)
-   [SetOsEnv](#SetOsEnv)
-   [RemoveOsEnv](#RemoveOsEnv)
-   [CompareOsEnv](#CompareOsEnv)
-   [ExecCommand](#ExecCommand)
-   [GetOsBits](#GetOsBits)
-   [StartProcess](#StartProcess)
-   [StopProcess](#StopProcess)
-   [KillProcess](#KillProcess)
-   [GetProcessInfo](#GetProcessInfo)


<div STYLE="page-break-after: always;"></div>


## Documentation

### <span id="IsWindows">IsWindows</span>

<p>Check if current os is windows.</p>

<b>Signature:</b>

```go
func IsWindows() bool
```

<b>Example:<span style="float:right;display:inline-block">[Run](https://go.dev/play/p/zIflQgZNuxD)</span></b>

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

<p>Check if current os is linux.</p>

<b>Signature:</b>

```go
func IsLinux() bool
```

<b>Example:<span style="float:right;display:inline-block">[Run](https://go.dev/play/p/zIflQgZNuxD)</span></b>

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

<p>Check if current os is macos.</p>

<b>Signature:</b>

```go
func IsMac() bool
```

<b>Example:<span style="float:right;display:inline-block">[Run](https://go.dev/play/p/Mg4Hjtyq7Zc)</span></b>

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

<p>Gets the value of the environment variable named by the key.</p>

<b>Signature:</b>

```go
func GetOsEnv(key string) string
```

<b>Example:<span style="float:right;display:inline-block">[Run](https://go.dev/play/p/D88OYVCyjO-)</span></b>

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

<p>Sets the value of the environment variable named by the key.</p>

<b>Signature:</b>

```go
func SetOsEnv(key, value string) error
```

<b>Example:<span style="float:right;display:inline-block">[Run](https://go.dev/play/p/D88OYVCyjO-)</span></b>

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

<p>Remove a single environment variable.</p>

<b>Signature:</b>

```go
func RemoveOsEnv(key string) error
```

<b>Example:<span style="float:right;display:inline-block">[Run](https://go.dev/play/p/fqyq4b3xUFQ)</span></b>

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

<p>Get env named by the key and compare it with comparedEnv.</p>

<b>Signature:</b>

```go
func CompareOsEnv(key, comparedEnv string) bool
```

<b>Example:<span style="float:right;display:inline-block">[Run](https://go.dev/play/p/BciHrKYOHbp)</span></b>

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

<p>Execute shell command, return the stdout and stderr string of command, and error if error occur. param `command` is a complete command string, like, ls -a (linux), dir(windows), ping 127.0.0.1. In linux, use /bin/bash -c to execute command, In windows, use powershell.exe to execute command.
The second parameter of the function is the cmd option control parameter. The type is func(*exec.Cmd). You can set the cmd attribute through this parameter.</p>

<b>Signature:</b>

```go
type (
    Option func(*exec.Cmd)
)
func ExecCommand(command string, opts ...Option) (stdout, stderr string, err error)
```

<b>Example:<span style="float:right;display:inline-block">[Run](https://go.dev/play/p/n-2fLyZef-4)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    // linux or mac
    stdout, stderr, err := system.ExecCommand("ls", func(cmd *exec.Cmd) {
        cmd.Dir = "/tmp"
    })
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

<p>Get current os bits, 32bit or 64bit. return 32 or 64.</p>

<b>Signature:</b>

```go
func GetOsBits() int
```

<b>Example:<span style="float:right;display:inline-block">[Run](https://go.dev/play/p/ml-_XH3gJbW)</span></b>

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

### <span id="StartProcess">StartProcess</span>

<p>Start a new process with the specified name and arguments.</p>

<b>Signature:</b>

```go
func StartProcess(command string, args ...string) (int, error)
```

<b>Example:<span style="float:right;display:inline-block">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    pid, err := system.StartProcess("sleep", "2")
    if err != nil {
        return
    }

    fmt.Println(pid)
}
```

### <span id="StopProcess">StopProcess</span>

<p>Stop a process by pid.</p>

<b>Signature:</b>

```go
func StopProcess(pid int) error
```

<b>Example:<span style="float:right;display:inline-block">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    pid, err := system.StartProcess("sleep", "10")
    if err != nil {
        return
    }
    time.Sleep(1 * time.Second)

    err = system.StopProcess(pid)

    fmt.Println(err)

    // Output:
    // <nil>
}
```

### <span id="KillProcess">KillProcess</span>

<p>Kill a process by pid.</p>

<b>Signature:</b>

```go
func KillProcess(pid int) error
```

<b>Example:<span style="float:right;display:inline-block">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    pid, err := system.StartProcess("sleep", "10")
    if err != nil {
        return
    }
    time.Sleep(1 * time.Second)

    err = system.KillProcess(pid)

    fmt.Println(err)

    // Output:
    // <nil>
}
```

### <span id="GetProcessInfo">GetProcessInfo</span>

<p>Retrieves detailed process information by pid.</p>

<b>Signature:</b>

```go
func GetProcessInfo(pid int) (*ProcessInfo, error)
```

<b>Example:<span style="float:right;display:inline-block">[Run](todo)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/system"
)

func main() {
    pid, err := system.StartProcess("ls", "-a")
    if err != nil {
        return
    }

    processInfo, err := system.GetProcessInfo(pid)
    if err != nil {
        return
    }

    fmt.Println(processInfo)
}
```