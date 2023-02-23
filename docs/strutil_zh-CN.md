# Strutil

strutil 包含处理字符串的相关函数。

<div STYLE="page-break-after: always;"></div>

## 源码:

[https://github.com/duke-git/lancet/blob/v1/strutil/string.go](https://github.com/duke-git/lancet/blob/v1/strutil/string.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/strutil"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [After](#After)
-   [AfterLast](#AfterLast)
-   [Before](#Before)
-   [BeforeLast](#BeforeLast)
-   [CamelCase](#CamelCase)
-   [Capitalize](#Capitalize)
-   [IsString](#IsString)
-   [KebabCase](#KebabCase)
-   [UpperKebabCase](#UpperKebabCase)
-   [LowerFirst](#LowerFirst)
-   [UpperFirst](#UpperFirst)
-   [Pad](#Pad)
-   [PadEnd](#PadEnd)
-   [PadStart](#PadStart)
-   [Reverse](#Reverse)
-   [SnakeCase](#SnakeCase)
-   [UpperSnakeCase](#UpperSnakeCase)
-   [Wrap](#Wrap)
-   [Unwrap](#Unwrap)
-   [SplitEx](#SplitEx)

<div STYLE="page-break-after: always;"></div>

## Documentation 文档

### <span id="After">After</span>

<p>截取源字符串中char首次出现时的位置之后的子字符串</p>

<b>函数签名:</b>

```go
func After(s, char string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.After("lancet", "")
	fmt.Println(s1) //lancet

	s2 := strutil.After("github.com/test/lancet", "/")
	fmt.Println(s2) //test/lancet

	s3 := strutil.After("github.com/test/lancet", "test")
	fmt.Println(s3) // /lancet
}
```

### <span id="AfterLast">AfterLast</span>

<p>截取源字符串中char最后一次出现时的位置之后的子字符串</p>

<b>函数签名:</b>

```go
func AfterLast(s, char string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.AfterLast("lancet", "")
	fmt.Println(s1) //lancet

	s2 := strutil.AfterLast("github.com/test/lancet", "/")
	fmt.Println(s2) //lancet

	s3 := strutil.AfterLast("github.com/test/test/lancet", "test")
	fmt.Println(s3) // /lancet
}
```

### <span id="Before">Before</span>

<p>截取源字符串中char首次出现时的位置之前的子字符串</p>

<b>函数签名:</b>

```go
func Before(s, char string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.Before("lancet", "")
	fmt.Println(s1) //lancet

	s2 := strutil.Before("github.com/test/lancet", "/")
	fmt.Println(s2) //github.com

	s3 := strutil.Before("github.com/test/lancet", "test")
	fmt.Println(s3) // github.com/
}
```

### <span id="BeforeLast">BeforeLast</span>

<p>截取源字符串中char最后一次出现时的位置之前的子字符串</p>

<b>函数签名:</b>

```go
func BeforeLast(s, char string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.BeforeLast("lancet", "")
	fmt.Println(s1) //lancet

	s2 := strutil.BeforeLast("github.com/test/lancet", "/")
	fmt.Println(s2) //github.com/test

	s3 := strutil.BeforeLast("github.com/test/test/lancet", "test")
	fmt.Println(s3) //github.com/test/
}
```

### <span id="CamelCase">CamelCase</span>

<p>将字符串转换为驼峰式字符串, 非字母和数字会被忽略</p>

<b>函数签名:</b>

```go
func CamelCase(s string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.CamelCase("foo_bar")
	fmt.Println(s1) //fooBar

	s2 := strutil.CamelCase("Foo-Bar")
	fmt.Println(s2) //fooBar

	s3 := strutil.CamelCase("Foo&bar")
	fmt.Println(s3) //fooBar

	s4 := strutil.CamelCase("foo bar")
	fmt.Println(s4) //fooBar

	s4 := strutil.CamelCase("Foo-#1😄$_%^&*(1bar")
	fmt.Println(s4) //foo11Bar
}
```

### <span id="Capitalize">Capitalize</span>

<p>将字符串的第一个字符转换为大写</p>

<b>函数签名:</b>

```go
func Capitalize(s string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.Capitalize("foo")
	fmt.Println(s1) //foo

	s2 := strutil.Capitalize("Foo")
	fmt.Println(s2) //foo

	s3 := strutil.Capitalize("FOo"
	fmt.Println(s3) //fOo
}
```

### <span id="IsString">IsString</span>

<p>检查值的数据类型是否为字符串</p>

<b>函数签名:</b>

```go
func IsString(v interface{}) bool
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	fmt.Println(strutil.IsString("lancet")) //true
	fmt.Println(strutil.IsString("")) //true

	fmt.Println(strutil.IsString(1)) //false
	fmt.Println(strutil.IsString("")) //false
	fmt.Println(strutil.IsString([]string{})) //false
}
```

### <span id="KebabCase">KebabCase</span>

<p>将字符串转换为kebab-case, 非字母和数字会被忽略</p>

<b>函数签名:</b>

```go
func KebabCase(s string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.KebabCase("Foo Bar-")
	fmt.Println(s1) //foo-bar

	s2 := strutil.KebabCase("foo_Bar")
	fmt.Println(s2) //foo-bar

	s3 := strutil.KebabCase("fooBar")
	fmt.Println(s3) //foo-bar

	s4 := strutil.KebabCase("__FOO_BAR__")
	fmt.Println(s4) //foo-bar
}
```

### <span id="UpperKebabCase">UpperKebabCase</span>

<p>将字符串转换为大写KEBAB-CASE, 非字母和数字会被忽略</p>

<b>函数签名:</b>

```go
func KebabCase(s string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.UpperKebabCase("Foo Bar-")
	fmt.Println(s1) //FOO-BAR

	s2 := strutil.UpperKebabCase("foo_Bar")
	fmt.Println(s2) //FOO-BAR

	s3 := strutil.UpperKebabCase("fooBar")
	fmt.Println(s3) //FOO-BAR

	s4 := strutil.UpperKebabCase("__FOO_BAR__")
	fmt.Println(s4) //FOO-BAR
}
```

### <span id="LowerFirst">LowerFirst</span>

<p>将字符串的第一个字符转换为小写</p>

<b>函数签名:</b>

```go
func LowerFirst(s string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.LowerFirst("foo")
	fmt.Println(s1) //foo

	s2 := strutil.LowerFirst("BAR")
	fmt.Println(s2) //bAR

	s3 := strutil.LowerFirst("FOo")
	fmt.Println(s3) //fOo

	s4 := strutil.LowerFirst("fOo大")
	fmt.Println(s4) //fOo大
}
```

### <span id="UpperFirst">UpperFirst</span>

<p>将字符串的第一个字符转换为大写</p>

<b>函数签名:</b>

```go
func UpperFirst(s string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.UpperFirst("foo")
	fmt.Println(s1) //Foo

	s2 := strutil.UpperFirst("bAR")
	fmt.Println(s2) //BAR

	s3 := strutil.UpperFirst("FOo")
	fmt.Println(s3) //FOo

	s4 := strutil.UpperFirst("fOo大")
	fmt.Println(s4) //FOo大
}
```

### <span id="Pad">Pad</span>

<p>如果字符串长度短于size，则在左右两侧填充字符串。</p>

<b>函数签名:</b>

```go
func Pad(source string, size int, padStr string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.Pad("foo", 1, "bar")
    result2 := strutil.Pad("foo", 2, "bar")
    result3 := strutil.Pad("foo", 3, "bar")
    result4 := strutil.Pad("foo", 4, "bar")
    result5 := strutil.Pad("foo", 5, "bar")
    result6 := strutil.Pad("foo", 6, "bar")
    result7 := strutil.Pad("foo", 7, "bar")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)
    fmt.Println(result7)
    // Output:
    // foo
    // foo
    // foo
    // foob
    // bfoob
    // bfooba
    // bafooba
}
```

### <span id="PadEnd">PadEnd</span>

<p>如果字符串长度短于size，则在右侧填充字符串。</p>

<b>函数签名:</b>

```go
func PadEnd(source string, size int, padStr string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.PadEnd("foo", 1, "bar")
    result2 := strutil.PadEnd("foo", 2, "bar")
    result3 := strutil.PadEnd("foo", 3, "bar")
    result4 := strutil.PadEnd("foo", 4, "bar")
    result5 := strutil.PadEnd("foo", 5, "bar")
    result6 := strutil.PadEnd("foo", 6, "bar")
    result7 := strutil.PadEnd("foo", 7, "bar")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)
    fmt.Println(result7)

    // Output:
    // foo
    // foo
    // foo
    // foob
    // fooba
    // foobar
    // foobarb
}
```

### <span id="PadStart">PadStart</span>

<p>如果字符串长度短于size，则在左侧填充字符串。</p>

<b>函数签名:</b>

```go
func PadStart(source string, size int, padStr string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.PadStart("foo", 1, "bar")
    result2 := strutil.PadStart("foo", 2, "bar")
    result3 := strutil.PadStart("foo", 3, "bar")
    result4 := strutil.PadStart("foo", 4, "bar")
    result5 := strutil.PadStart("foo", 5, "bar")
    result6 := strutil.PadStart("foo", 6, "bar")
    result7 := strutil.PadStart("foo", 7, "bar")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)
    fmt.Println(result7)

    // Output:
    // foo
    // foo
    // foo
    // bfoo
    // bafoo
    // barfoo
    // barbfoo
}
```

### <span id="ReverseStr">ReverseStr</span>

<p>返回字符顺序与给定字符串相反的字符串</p>

<b>函数签名:</b>

```go
func ReverseStr(s string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.ReverseStr("abc")
	fmt.Println(s1) //cba

	s2 := strutil.ReverseStr("12345")
	fmt.Println(s2) //54321
}
```

### <span id="SnakeCase">SnakeCase</span>

<p>将字符串转换为snake_case形式, 非字母和数字会被忽略</p>

<b>函数签名:</b>

```go
func SnakeCase(s string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.SnakeCase("Foo Bar-")
	fmt.Println(s1) //foo_bar

	s2 := strutil.SnakeCase("foo_Bar")
	fmt.Println(s2) //foo_bar

	s3 := strutil.SnakeCase("fooBar")
	fmt.Println(s3) //foo_bar

	s4 := strutil.SnakeCase("__FOO_BAR__")
	fmt.Println(s4) //foo_bar

	s5 := strutil.SnakeCase("Foo-#1😄$_%^&*(1bar")
	fmt.Println(s5) //foo_1_1_bar
}
```

### <span id="UpperSnakeCase">UpperSnakeCase</span>

<p>将字符串转换为大写SNAKE_CASE形式, 非字母和数字会被忽略</p>

<b>函数签名:</b>

```go
func SnakeCase(s string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.UpperSnakeCase("Foo Bar-")
	fmt.Println(s1) //FOO_BAR

	s2 := strutil.UpperSnakeCase("foo_Bar")
	fmt.Println(s2) //FOO_BAR

	s3 := strutil.UpperSnakeCase("fooBar")
	fmt.Println(s3) //FOO_BAR

	s4 := strutil.UpperSnakeCase("__FOO_BAR__")
	fmt.Println(s4) //FOO_BAR

	s5 := strutil.UpperSnakeCase("Foo-#1😄$_%^&*(1bar")
	fmt.Println(s5) //FOO_1_1_BAR
}
```

### <span id="Wrap">Wrap</span>

<p>用另一个字符串包裹一个字符串</p>

<b>函数签名:</b>

```go
func Wrap(str string, wrapWith string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.Wrap("ab", "")
	fmt.Println(s1) //ab

	s2 := strutil.Wrap("", "*")
	fmt.Println(s2) //""

	s3 := strutil.Wrap("ab", "*")
	fmt.Println(s3) //*ab*

	s4 := strutil.Wrap("ab", "\"")
	fmt.Println(s4) //\"ab\"

	s5 := strutil.Wrap("ab", "'")
	fmt.Println(s5) //'ab'
}
```

### <span id="Unwrap">Unwrap</span>

<p>用另一个字符串解开包裹一个字符串</p>

<b>函数签名:</b>

```go
func Unwrap(str string, wrapToken string) string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.Unwrap("ab", "")
	fmt.Println(s1) //ab

	s2 := strutil.Unwrap("ab", "*")
	fmt.Println(s2) //ab

	s3 := strutil.Unwrap("**ab**", "*")
	fmt.Println(s3) //*ab*

	s4 := strutil.Unwrap("*ab", "*")
	fmt.Println(s4) //*ab

	s5 := strutil.Unwrap("***", "**")
	fmt.Println(s5) //***
}
```

### <span id="SplitEx">SplitEx</span>

<p>分割字符串为切片，removeEmptyString参数指定是否去除空字符串</p>

<b>函数签名:</b>

```go
func SplitEx(s, sep string, removeEmptyString bool) []string
```

<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	arr1 := strutil.SplitEx(" a b c ", "", true)
	fmt.Println(arr1) //[]string{}

	arr2 := strutil.SplitEx(" a b c ", " ", false)
	fmt.Println(arr2) //[]string{"", "a", "b", "c", ""}

	arr3 := strutil.SplitEx(" a b c ", " ", true)
	fmt.Println(arr3) //[]string{"a", "b", "c"}

	arr4 := strutil.SplitEx(" a = b = c = ", " = ", false)
	fmt.Println(arr4) //[]string{" a", "b", "c", ""}

	arr5 := strutil.SplitEx(" a = b = c = ", " = ", true)
	fmt.Println(arr5) //[]string{" a", "b", "c"}
}
```
