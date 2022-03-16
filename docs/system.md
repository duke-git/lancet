# System
Package system contains some functions about os, runtime, shell command.

<div STYLE="page-break-after: always;"></div>

## Source:

[https://github.com/duke-git/lancet/blob/v1/system/os.go](https://github.com/duke-git/lancet/blob/v1/system/os.go)


<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/duke-git/lancet/system"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
- [IsWindows](#IsWindows)
- [IsLinux](#IsLinux)
- [IsMac](#IsMac)
- [GetOsEnv](#GetOsEnv)
- [SetOsEnv](#SetOsEnv)
- [RemoveOsEnv](#RemoveOsEnv)
- [CompareOsEnv](#CompareOsEnv)
- [ExecCommand](#ExecCommand)
  

<div STYLE="page-break-after: always;"></div>

## Documentation


### <span id="IsWindows">IsWindows</span>
<p>Check if current os is windows.</p>

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
<p>Check if current os is linux.</p>

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
<p>Check if current os is macos.</p>

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
<p>Gets the value of the environment variable named by the key.</p>

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
<p>Sets the value of the environment variable named by the key.</p>

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
<p>Remove a single environment variable.</p>

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
<p>Get env named by the key and compare it with comparedEnv.</p>

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
<p>use shell /bin/bash -c(linux) or cmd (windows) to execute command.</p>

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








