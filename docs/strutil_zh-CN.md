# Strutil

strutil 包含处理字符串的相关函数。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/strutil/string.go](https://github.com/duke-git/lancet/blob/main/strutil/string.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/strutil"
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
-   [SplitEx](#SplitEx)
-   [Substring](#Substring)
-   [Wrap](#Wrap)
-   [Unwrap](#Unwrap)
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
-   [ContainsAll](#ContainsAll)
-   [ContainsAny](#ContainsAny)


<div STYLE="page-break-after: always;"></div>

## Documentation 文档

### <span id="After">After</span>

<p>返回源字符串中特定字符串首次出现时的位置之后的子字符串。</p>

<b>函数签名:</b>

```go
func After(s, char string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.After("foo", "")
    result2 := strutil.After("foo", "foo")
    result3 := strutil.After("foo/bar", "foo")
    result4 := strutil.After("foo/bar", "/")
    result5 := strutil.After("foo/bar/baz", "/")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)

    // Output:
    // foo
    //
    // /bar
    // bar
    // bar/baz
}
```

### <span id="AfterLast">AfterLast</span>

<p>返回源字符串中指定字符串最后一次出现时的位置之后的子字符串。</p>

<b>函数签名:</b>

```go
func AfterLast(s, char string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.AfterLast("foo", "")
    result2 := strutil.AfterLast("foo", "foo")
    result3 := strutil.AfterLast("foo/bar", "/")
    result4 := strutil.AfterLast("foo/bar/baz", "/")
    result5 := strutil.AfterLast("foo/bar/foo/baz", "foo")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)

    // Output:
    // foo
    //
    // bar
    // baz
    // /baz
}
```

### <span id="Before">Before</span>

<p>返回源字符串中指定字符串第一次出现时的位置之前的子字符串。</p>

<b>函数签名:</b>

```go
func Before(s, char string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.Before("foo", "")
    result2 := strutil.Before("foo", "foo")
    result3 := strutil.Before("foo/bar", "/")
    result4 := strutil.Before("foo/bar/baz", "/")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // foo
    //
    // foo
    // foo
}
```

### <span id="BeforeLast">BeforeLast</span>

<p>返回源字符串中指定字符串最后一次出现时的位置之前的子字符串。</p>

<b>函数签名:</b>

```go
func BeforeLast(s, char string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.BeforeLast("foo", "")
    result2 := strutil.BeforeLast("foo", "foo")
    result3 := strutil.BeforeLast("foo/bar", "/")
    result4 := strutil.BeforeLast("foo/bar/baz", "/")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // foo
    //
    // foo
    // foo/bar
}
```

### <span id="CamelCase">CamelCase</span>

<p>将字符串转换为驼峰式字符串, 非字母和数字会被忽略。</p>

<b>函数签名:</b>

```go
func CamelCase(s string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    strings := []string{"", "foobar", "&FOO:BAR$BAZ", "$foo%", "Foo-#1😄$_%^&*(1bar"}

    for _, v := range strings {
        s := strutil.CamelCase(v)
        fmt.Println(s)
    }

    // Output:
    //
    // foobar
    // fooBarBaz
    // foo
    // foo11Bar
}
```

### <span id="KebabCase">KebabCase</span>

<p>将字符串转换为kebab-case, 非字母和数字会被忽略。</p>

<b>函数签名:</b>

```go
func KebabCase(s string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    strings := []string{"", "foo-bar", "Foo Bar-", "FOOBAR", "Foo-#1😄$_%^&*(1bar"}

    for _, v := range strings {
        s := strutil.KebabCase(v)
        fmt.Println(s)
    }

    // Output:
    //
    // foo-bar
    // foo-bar
    // foobar
    // foo-1-1-bar
}
```

### <span id="UpperKebabCase">UpperKebabCase</span>

<p>将字符串转换为大写KEBAB-CASE, 非字母和数字会被忽略。</p>

<b>函数签名:</b>

```go
func UpperKebabCase(s string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    strings := []string{"", "foo-bar", "Foo Bar-", "FooBAR", "Foo-#1😄$_%^&*(1bar"}

    for _, v := range strings {
        s := strutil.UpperKebabCase(v)
        fmt.Println(s)
    }

    // Output:
    //
    // FOO-BAR
    // FOO-BAR
    // FOO-BAR
    // FOO-1-1-BAR
}
```

### <span id="Capitalize">Capitalize</span>

<p>将字符串的第一个字符转换为大写。</p>

<b>函数签名:</b>

```go
func Capitalize(s string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    strings := []string{"", "Foo", "_foo", "fooBar", "foo-bar"}

    for _, v := range strings {
        s := strutil.Capitalize(v)
        fmt.Println(s)
    }

    // Output:
    //
    // Foo
    // _foo
    // Foobar
    // Foo-bar
}
```

### <span id="IsString">IsString</span>

<p>判断传入参数的数据类型是否为字符串。</p>

<b>函数签名:</b>

```go
func IsString(v any) bool
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.IsString("")
    result2 := strutil.IsString("a")
    result3 := strutil.IsString(1)
    result4 := strutil.IsString(true)
    result5 := strutil.IsString([]string{"a"})

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)

    // Output:
    // true
    // true
    // false
    // false
    // false
}
```

### <span id="LowerFirst">LowerFirst</span>

<p>将字符串的第一个字符转换为小写。</p>

<b>函数签名:</b>

```go
func LowerFirst(s string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    strings := []string{"", "bar", "BAr", "Bar大"}

    for _, v := range strings {
        s := strutil.LowerFirst(v)
        fmt.Println(s)
    }

    // Output:
    //
    // bar
    // bAr
    // bar大
}
```

### <span id="UpperFirst">UpperFirst</span>

<p>将字符串的第一个字符转换为大写形式。</p>

<b>函数签名:</b>

```go
func UpperFirst(s string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    strings := []string{"", "bar", "BAr", "bar大"}

    for _, v := range strings {
        s := strutil.UpperFirst(v)
        fmt.Println(s)
    }

    // Output:
    //
    // Bar
    // BAr
    // Bar大
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

### <span id="Reverse">Reverse</span>

<p>返回字符顺序与给定字符串相反的字符串。</p>

<b>函数签名:</b>

```go
func Reverse(s string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    s := "foo"
    rs := strutil.Reverse(s)

    fmt.Println(s)
    fmt.Println(rs)

    // Output:
    // foo
    // oof
}
```

### <span id="SnakeCase">SnakeCase</span>

<p>将字符串转换为snake_case形式, 非字母和数字会被忽略。</p>

<b>函数签名:</b>

```go
func SnakeCase(s string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    strings := []string{"", "foo-bar", "Foo Bar-", "FOOBAR", "Foo-#1😄$_%^&*(1bar"}

    for _, v := range strings {
        s := strutil.SnakeCase(v)
        fmt.Println(s)
    }

    // Output:
    //
    // foo_bar
    // foo_bar
    // foobar
    // foo_1_1_bar
}
```

### <span id="UpperSnakeCase">UpperSnakeCase</span>

<p>将字符串转换为大写SNAKE_CASE形式, 非字母和数字会被忽略。</p>

<b>函数签名:</b>

```go
func SnakeCase(s string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    strings := []string{"", "foo-bar", "Foo Bar-", "FooBAR", "Foo-#1😄$_%^&*(1bar"}

    for _, v := range strings {
        s := strutil.UpperSnakeCase(v)
        fmt.Println(s)
    }

    // Output:
    //
    // FOO_BAR
    // FOO_BAR
    // FOO_BAR
    // FOO_1_1_BAR
}
```

### <span id="SplitEx">SplitEx</span>

<p>分割字符串为切片，removeEmptyString参数指定是否去除空字符串。</p>

<b>函数签名:</b>

```go
func SplitEx(s, sep string, removeEmptyString bool) []string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.SplitEx(" a b c ", "", true)

    result2 := strutil.SplitEx(" a b c ", " ", false)
    result3 := strutil.SplitEx(" a b c ", " ", true)

    result4 := strutil.SplitEx("a = b = c = ", " = ", false)
    result5 := strutil.SplitEx("a = b = c = ", " = ", true)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)

    // Output:
    // []
    // [ a b c ]
    // [a b c]
    // [a b c ]
}
```

### <span id="Substring">Substring</span>

<p>根据指定的位置和长度截取字符串。</p>

<b>函数签名:</b>

```go
func Substring(s string, offset int, length uint) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.Substring("abcde", 1, 3)
    result2 := strutil.Substring("abcde", 1, 5)
    result3 := strutil.Substring("abcde", -1, 3)
    result4 := strutil.Substring("abcde", -2, 2)
    result5 := strutil.Substring("abcde", -2, 3)
    result6 := strutil.Substring("你好，欢迎你", 0, 2)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)

    // Output:
    // bcd
    // bcde
    // e
    // de
    // de
    // 你好
}
```

### <span id="Wrap">Wrap</span>

<p>用另一个字符串包裹一个字符串。</p>

<b>函数签名:</b>

```go
func Wrap(str string, wrapWith string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.Wrap("foo", "")
    result2 := strutil.Wrap("foo", "*")
    result3 := strutil.Wrap("'foo'", "'")
    result4 := strutil.Wrap("", "*")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // foo
    // *foo*
    // ''foo''
    //
}
```

### <span id="Wrap">Wrap</span>

<p>用另一个字符串解开包裹一个字符串。</p>

<b>函数签名:</b>

```go
func Unwrap(str string, wrapToken string) string
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    result1 := strutil.Unwrap("foo", "")
    result2 := strutil.Unwrap("*foo*", "*")
    result3 := strutil.Unwrap("*foo", "*")
    result4 := strutil.Unwrap("foo*", "*")
    result5 := strutil.Unwrap("**foo**", "*")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)

    // Output:
    // foo
    // foo
    // *foo
    // foo*
    // *foo*
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
    "github.com/duke-git/lancet/v2/strutil"
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
