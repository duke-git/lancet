# Strutil
Package strutil contains some functions to manipulate string.

<div STYLE="page-break-after: always;"></div>

## Source:

[https://github.com/duke-git/lancet/blob/v1/strutil/string.go](https://github.com/duke-git/lancet/blob/v1/strutil/string.go)


<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/duke-git/lancet/strutil"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
- [After](#After)
- [AfterLast](#AfterLast)
- [Before](#Before)
- [BeforeLast](#BeforeLast)
- [CamelCase](#CamelCase)
- [Capitalize](#Capitalize)
- [IsString](#IsString)
- [KebabCase](#KebabCase)
- [LowerFirst](#LowerFirst)
- [UpperFirst](#UpperFirst)
- [PadEnd](#PadEnd)
- [PadStart](#PadStart)
- [ReverseStr](#ReverseStr)
- [SnakeCase](#SnakeCase)
- [Wrap](#Wrap)
  
- [Unwrap](#Unwrap)
  

<div STYLE="page-break-after: always;"></div>

## Documentation


### <span id="After">After</span>
<p>Creates substring in source string after position when char first appear.</p>

<b>Signature:</b>

```go
func After(s, char string) string
```
<b>Example:</b>

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
<p>Creates substring in source string after position when char last appear.</p>

<b>Signature:</b>

```go
func AfterLast(s, char string) string
```
<b>Example:</b>

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
	fmt.Println(s3) // /test/lancet
}
```




### <span id="Before">Before</span>
<p>Creates substring in source string before position when char first appear.</p>

<b>Signature:</b>

```go
func Before(s, char string) string
```
<b>Example:</b>

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
<p>Creates substring in source string before position when char first appear.</p>

<b>Signature:</b>

```go
func BeforeLast(s, char string) string
```
<b>Example:</b>

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
<p>Covert string to camelCase string.</p>

<b>Signature:</b>

```go
func CamelCase(s string) string
```
<b>Example:</b>

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
}
```




### <span id="Capitalize">Capitalize</span>
<p>Convert the first character of a string to upper case.</p>

<b>Signature:</b>

```go
func Capitalize(s string) string
```
<b>Example:</b>

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
<p>Check if the value's data type is string.</p>

<b>Signature:</b>

```go
func IsString(v interface{}) bool
```
<b>Example:</b>

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
<p>Covert string to kebab-case.</p>

<b>Signature:</b>

```go
func KebabCase(s string) string
```
<b>Example:</b>

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
	fmt.Println(s4) //f-o-o-b-a-r
}
```




### <span id="LowerFirst">LowerFirst</span>
<p>Convert the first character of string to lower case.</p>

<b>Signature:</b>

```go
func LowerFirst(s string) string
```
<b>Example:</b>

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
<p>Convert the first character of string to upper case.</p>

<b>Signature:</b>

```go
func UpperFirst(s string) string
```
<b>Example:</b>

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




### <span id="PadEnd">PadEnd</span>
<p>Pads string on the right side if it's shorter than size.</p>

<b>Signature:</b>

```go
func PadEnd(source string, size int, padStr string) string
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.PadEnd("a", 1, "b")
	fmt.Println(s1) //a

	s2 := strutil.PadEnd("a", 2, "b")
	fmt.Println(s2) //ab

	s3 := strutil.PadEnd("abcd", 6, "mno")
	fmt.Println(s3) //abcdmn

	s4 := strutil.PadEnd("abc", 6, "ab")
	fmt.Println(s4) //abcaba
}
```




### <span id="PadStart">PadStart</span>
<p>Pads string on the left side if it's shorter than size.</p>

<b>Signature:</b>

```go
func PadStart(source string, size int, padStr string) string
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/strutil"
)

func main() {
	s1 := strutil.PadStart("a", 1, "b")
	fmt.Println(s1) //a

	s2 := strutil.PadStart("a", 2, "b")
	fmt.Println(s2) //ba

	s3 := strutil.PadStart("abcd", 6, "mno")
	fmt.Println(s3) //mnabcd

	s4 := strutil.PadStart("abc", 6, "ab")
	fmt.Println(s4) //abaabc
}
```




### <span id="ReverseStr">ReverseStr</span>
<p>Return string whose char order is reversed to the given string.</p>

<b>Signature:</b>

```go
func ReverseStr(s string) string
```
<b>Example:</b>

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
<p>Covert string to snake_case.</p>

<b>Signature:</b>

```go
func SnakeCase(s string) string
```
<b>Example:</b>

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
	fmt.Println(s4) //f_o_o_b_a_r

	s5 := strutil.SnakeCase("aBbc-s$@a&%_B.B^C")
	fmt.Println(s5) //a_bbc_s_a_b_b_c
}
```




### <span id="Wrap">Wrap</span>
<p>Wrap a string with another string.</p>

<b>Signature:</b>

```go
func Wrap(str string, wrapWith string) string
```
<b>Example:</b>

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




### <span id="Wrap">Wrap</span>
<p>Unwrap a given string from anther string. will change str value.</p>

<b>Signature:</b>

```go
func Unwrap(str string, wrapToken string) string
```
<b>Example:</b>

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









