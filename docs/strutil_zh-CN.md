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
-   [ContainsAll](#ContainsAll)
-   [ContainsAny](#ContainsAny)
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
-   [SplitWords](#SplitWords)
-   [WordCount](#WordCount)
-   [RemoveNonPrintable](#RemoveNonPrintable)
-   [StringToBytes](#StringToBytes)
-   [BytesToString](#BytesToString)
-   [IsBlank](#IsBlank)
-   [HasPrefixAny](#HasPrefixAny)
-   [HasSuffixAny](#HasSuffixAny)
-   [IndexOffset](#IndexOffset)
-   [ReplaceWithMap](#ReplaceWithMap)
-   [Trim](#Trim)
-   [SplitAndTrim](#SplitAndTrim)
-   [HideString](#HideString)
-   [RemoveWhiteSpace](#RemoveWhiteSpace)
-   [SubInBetween](#SubInBetween)
-   [HammingDistance](#HammingDistance)
-   [Concat](#Concat)
-   [Ellipsis](#Ellipsis)
-   [Shuffle](#Shuffle)
-   [Rotate](#Rotate)
-   [TemplateReplace](#TemplateReplace)
-   [RegexMatchAllGroups](#RegexMatchAllGroups)


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
    "github.com/duke-git/lancet/strutil"
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
    "github.com/duke-git/lancet/strutil"
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
    "github.com/duke-git/lancet/strutil"
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

### <span id="SplitWords">SplitWords</span>

<p>将字符串拆分为单词，只支持字母字符单词。</p>

<b>函数签名:</b>

```go
func SplitWords(s string) []string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    result1 := strutil.SplitWords("a word")
    result2 := strutil.SplitWords("I'am a programmer")
    result3 := strutil.SplitWords("Bonjour, je suis programmeur")
    result4 := strutil.SplitWords("a -b-c' 'd'e")
    result5 := strutil.SplitWords("你好，我是一名码农")
    result6 := strutil.SplitWords("こんにちは，私はプログラマーです")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)

    // Output:
    // [a word]
    // [I'am a programmer]
    // [Bonjour je suis programmeur]
    // [a b-c' d'e]
    // []
    // []
}
```

### <span id="WordCount">WordCount</span>

<p>返回有意义单词的数量，只支持字母字符单词。</p>

<b>函数签名:</b>

```go
func WordCount(s string) int
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    result1 := strutil.WordCount("a word")
    result2 := strutil.WordCount("I'am a programmer")
    result3 := strutil.WordCount("Bonjour, je suis programmeur")
    result4 := strutil.WordCount("a -b-c' 'd'e")
    result5 := strutil.WordCount("你好，我是一名码农")
    result6 := strutil.WordCount("こんにちは，私はプログラマーです")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)

    // Output:
    // 2
    // 3
    // 4
    // 3
    // 0
    // 0
}
```

### <span id="RemoveNonPrintable">RemoveNonPrintable</span>

<p>删除字符串中不可打印的字符。</p>

<b>函数签名:</b>

```go
func RemoveNonPrintable(str string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    result1 := strutil.RemoveNonPrintable("hello\u00a0 \u200bworld\n")
    result2 := strutil.RemoveNonPrintable("你好😄")

    fmt.Println(result1)
    fmt.Println(result2)
    // Output:
    // hello world
    // 你好😄
}
```

### <span id="StringToBytes">StringToBytes</span>

<p>在不分配内存的情况下将字符串转换为字节片。</p>

<b>函数签名:</b>

```go
func StringToBytes(str string) (b []byte)
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    result1 := strutil.StringToBytes("abc")
    result2 := reflect.DeepEqual(result1, []byte{'a', 'b', 'c'})

    fmt.Println(result1)
    fmt.Println(result2)
    // Output:
    // [97 98 99]
    // true
}
```

### <span id="BytesToString">BytesToString</span>

<p>在不分配内存的情况下将字节切片转换为字符串。</p>

<b>函数签名:</b>

```go
func BytesToString(bytes []byte) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    bytes := []byte{'a', 'b', 'c'}
    result := strutil.BytesToString(bytes)

    fmt.Println(result)
    // Output:
    // abc
}
```

### <span id="IsBlank">IsBlank</span>

<p>检查字符串是否为空格或空。</p>

<b>函数签名:</b>

```go
func IsBlank(str string) bool
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    result1 := strutil.IsBlank("")
    result2 := strutil.IsBlank("\t\v\f\n")
    result3 := strutil.IsBlank(" 中文")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    // Output:
    // true
    // true
    // false
}
```

### <span id="HasPrefixAny">HasPrefixAny</span>

<p>检查字符串是否以指定字符串数组中的任何一个开头。</p>

<b>函数签名:</b>

```go
func HasPrefixAny(str string, prefixes []string) bool
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    result1 := strutil.HasPrefixAny("foo bar", []string{"fo", "xyz", "hello"})
    result2 := strutil.HasPrefixAny("foo bar", []string{"oom", "world"})

    fmt.Println(result1)
    fmt.Println(result2)
    // Output:
    // true
    // false
}
```

### <span id="HasSuffixAny">HasSuffixAny</span>

<p>检查字符串是否以指定字符串数组中的任何一个结尾。</p>

<b>函数签名:</b>

```go
func HasSuffixAny(str string, suffixes []string) bool
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    result1 := strutil.HasSuffixAny("foo bar", []string{"bar", "xyz", "hello"})
    result2 := strutil.HasSuffixAny("foo bar", []string{"oom", "world"})

    fmt.Println(result1)
    fmt.Println(result2)
    // Output:
    // true
    // false
}
```

### <span id="IndexOffset">IndexOffset</span>

<p>将字符串偏移idxFrom后，返回字符串中第一个 substr 实例的索引，如果字符串中不存在 substr，则返回 -1。</p>

<b>函数签名:</b>

```go
func IndexOffset(str string, substr string, idxFrom int) int
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    str := "foo bar hello world"

    result1 := strutil.IndexOffset(str, "o", 5)
    result2 := strutil.IndexOffset(str, "o", 0)
    result3 := strutil.IndexOffset(str, "d", len(str)-1)
    result4 := strutil.IndexOffset(str, "d", len(str))
    result5 := strutil.IndexOffset(str, "f", -1)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    // Output:
    // 12
    // 1
    // 18
    // -1
    // -1
}
```

### <span id="ReplaceWithMap">ReplaceWithMap</span>

<p>返回`str`的副本，以无序的方式被map替换，区分大小写。</p>

<b>函数签名:</b>

```go
func ReplaceWithMap(str string, replaces map[string]string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    str := "ac ab ab ac"
    replaces := map[string]string{
        "a": "1",
        "b": "2",
    }

    result := strutil.ReplaceWithMap(str, replaces)

    fmt.Println(result)
    // Output:
    // 1c 12 12 1c
}
```

### <span id="Trim">Trim</span>

<p>从字符串的开头和结尾去除空格（或其他字符）。 可选参数 characterMask 指定额外的剥离字符。</p>

<b>函数签名:</b>

```go
func Trim(str string, characterMask ...string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    result1 := strutil.Trim("\nabcd")

    str := "$ ab    cd $ "

    result2 := strutil.Trim(str)
    result3 := strutil.Trim(str, "$")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // abcd
    // $ ab    cd $
    // ab    cd
}
```


### <span id="SplitAndTrim">SplitAndTrim</span>

<p>将字符串str按字符串delimiter拆分为一个切片，并对该数组的每个元素调用Trim。忽略Trim后为空的元素。</p>

<b>函数签名:</b>

```go
func SplitAndTrim(str, delimiter string, characterMask ...string) []string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    str := " a,b, c,d,$1 "

    result1 := strutil.SplitAndTrim(str, ",")
    result2 := strutil.SplitAndTrim(str, ",", "$")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // [a b c d $1]
    // [a b c d 1]
}
```


### <span id="HideString">HideString</span>

<p>使用参数`replaceChar`隐藏源字符串中的一些字符。替换范围是 origin[start : end]。</p>

<b>函数签名:</b>

```go
func HideString(origin string, start, end int, replaceChar string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    str := "13242658976"

    result1 := strutil.HideString(str, 3, 3, "*")
    result2 := strutil.HideString(str, 3, 4, "*")
    result3 := strutil.HideString(str, 3, 7, "*")
    result4 := strutil.HideString(str, 7, 11, "*")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // 13242658976
    // 132*2658976
    // 132****8976
    // 1324265****
}
```

### <span id="ContainsAll">ContainsAll</span>

<p>判断字符串是否包括全部给定的子字符串切片。</p>

<b>函数签名:</b>

```go
func ContainsAll(str string, substrs []string) bool
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    str := "hello world"

    result1 := strutil.ContainsAll(str, []string{"hello", "world"})
    result2 := strutil.ContainsAll(str, []string{"hello", "abc"})

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="ContainsAny">ContainsAny</span>

<p>判断字符串是否包括给定的子字符串切片中任意一个子字符串。</p>

<b>函数签名:</b>

```go
func ContainsAny(str string, substrs []string) bool
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    str := "hello world"

    result1 := strutil.ContainsAny(str, []string{"hello", "world"})
    result2 := strutil.ContainsAny(str, []string{"hello", "abc"})
    result3 := strutil.ContainsAny(str, []string{"123", "abc"})

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // true
    // true
    // false
}
```

### <span id="RemoveWhiteSpace">RemoveWhiteSpace</span>

<p>删除字符串中的空格，当设置repalceAll为true时，删除全部空格，为false时，替换多个空格为1个空格。</p>

<b>函数签名:</b>

```go
func RemoveWhiteSpace(str string, repalceAll bool) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    str := " hello   \r\n   \t   world"

    result1 := strutil.RemoveWhiteSpace(str, true)
    result2 := strutil.RemoveWhiteSpace(str, false)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // helloworld
    // hello world
}
```

### <span id="SubInBetween">SubInBetween</span>

<p>获取字符串中指定的起始字符串start和终止字符串end直接的子字符串。</p>

<b>函数签名:</b>

```go
func SubInBetween(str string, start string, end string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    str := "abcde"

    result1 := strutil.SubInBetween(str, "", "de")
    result2 := strutil.SubInBetween(str, "a", "d")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // abc
    // bc
}
```

### <span id="HammingDistance">HammingDistance</span>

<p>计算两个字符串之间的汉明距离。汉明距离是指对应符号不同的位置数。</p>

<b>函数签名:</b>

```go
HammingDistance(a, b string) (int, error)
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {

    result1, _ := strutil.HammingDistance("de", "de")
    result2, _ := strutil.HammingDistance("a", "d")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 0
    // 1
}
```

### <span id="Concat">Concat</span>

<p>拼接字符串。length是拼接后字符串的长度，如果不确定则传0或负数。</p>

<b>函数签名:</b>

```go
func Concat(length int, str ...string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {

     result1 := strutil.Concat(12, "Hello", " ", "World", "!")
    result2 := strutil.Concat(11, "Go", " ", "Language")
    result3 := strutil.Concat(0, "An apple a ", "day，", "keeps the", " doctor away")
    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // Hello World!
    // Go Language
    // An apple a day，keeps the doctor away
}
```

### <span id="Ellipsis">Ellipsis</span>

<p>将字符串截断到指定长度，并在末尾添加省略号。</p>

<b>函数签名:</b>

```go
func Ellipsis(str string, length int) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    result1 := strutil.Ellipsis("hello world", 5)
    result2 := strutil.Ellipsis("你好，世界!", 2)
    result3 := strutil.Ellipsis("😀😃😄😁😆", 3)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // hello...
    // 你好...
    // 😀😃😄...
}
```

### <span id="Shuffle">Shuffle</span>

<p>打乱给定字符串中的字符顺序。</p>

<b>函数签名:</b>

```go
func Shuffle(str string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    result := strutil.Shuffle("hello")
    fmt.Println(result)  //olelh (random order)
}
```

### <span id="Rotate">Rotate</span>

<p>按指定的字符数旋转字符串。</p>

<b>函数签名:</b>

```go
func Rotate(str string, shift int) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    result1 := Rotate("Hello", 0)
    result2 := Rotate("Hello", 1)
    result3 := Rotate("Hello", 2)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // Hello
    // oHell
    // loHel
}
```

### <span id="TemplateReplace">TemplateReplace</span>

<p>将模板字符串中的占位符替换为数据映射中的相应值。占位符括在花括号中，例如 {key}。例如，模板字符串为“Hello, {name}!”，数据映射为{"name": "world"}，结果将为“Hello, world!”。</p>

<b>函数签名:</b>

```go
func TemplateReplace(template string, data map[string]string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    template := `Hello, my name is {name}, I'm {age} years old.`
    data := map[string]string{
        "name": "Bob",
        "age":  "20",
    }

    result := strutil.TemplateReplace(template, data)

    fmt.Println(result)

    // Output:
    // Hello, my name is Bob, I'm 20 years old.
}
```

### <span id="RegexMatchAllGroups">RegexMatchAllGroups</span>

<p>使用正则表达式匹配字符串中的所有子组并返回结果。</p>

<b>函数签名:</b>

```go
func RegexMatchAllGroups(pattern, str string) [][]string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/strutil"
)

func main() {
    pattern := `(\w+\.+\w+)@(\w+)\.(\w+)`
    str := "Emails: john.doe@example.com and jane.doe@example.com"

    result := strutil.RegexMatchAllGroups(pattern, str)

    fmt.Println(result[0])
    fmt.Println(result[1])

    // Output:
    // [john.doe@example.com john.doe example com]
    // [jane.doe@example.com jane.doe example com]
}
```