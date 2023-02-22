# Strutil

Package strutil contains some functions to manipulate string.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/strutil/string.go](https://github.com/duke-git/lancet/blob/main/strutil/string.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/strutil"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

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
-   [PadStart](#PadStart)
-   [PadEnd](#PadEnd)
-   [Reverse](#Reverse)
-   [SnakeCase](#SnakeCase)
-   [UpperSnakeCase](#UpperSnakeCase)
-   [SplitEx](#SplitEx)
-   [Substring](#Substring)
-   [Wrap](#Wrap)
-   [Unwrap](#Unwrap)
-   [SplitWords](#SplitWords)
-   [WordCount](#WordCount)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="After">After</span>

<p>Returns the substring after the first occurrence of a specified string in the source string.</p>

<b>Signature:</b>

```go
func After(s, char string) string
```

<b>Example:</b>

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

<p>Returns the substring after the last occurrence of a specified string in the source string.</p>

<b>Signature:</b>

```go
func AfterLast(s, char string) string
```

<b>Example:</b>

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

<p>Returns the substring of the source string up to the first occurrence of the specified string.</p>

<b>Signature:</b>

```go
func Before(s, char string) string
```

<b>Example:</b>

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

<p>Returns the substring of the source string up to the last occurrence of the specified string.</p>

<b>Signature:</b>

```go
func BeforeLast(s, char string) string
```

<b>Example:</b>

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

<p>Coverts string to camelCase string, non letters and numbers will be ignored.</p>

<b>Signature:</b>

```go
func CamelCase(s string) string
```

<b>Example:</b>

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

<p>KebabCase covert string to kebab-case, non letters and numbers will be ignored.</p>

<b>Signature:</b>

```go
func KebabCase(s string) string
```

<b>Example:</b>

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

<p>UpperKebabCase covert string to upper KEBAB-CASE, non letters and numbers will be ignored.</p>

<b>Signature:</b>

```go
func UpperKebabCase(s string) string
```

<b>Example:</b>

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

<p>Convert the first character of a string to upper case.</p>

<b>Signature:</b>

```go
func Capitalize(s string) string
```

<b>Example:</b>

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

<p>Check if the value's data type is string.</p>

<b>Signature:</b>

```go
func IsString(v any) bool
```

<b>Example:</b>

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

<p>Convert the first character of string to lower case.</p>

<b>Signature:</b>

```go
func LowerFirst(s string) string
```

<b>Example:</b>

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

<p>Convert the first character of string to upper case.</p>

<b>Signature:</b>

```go
func UpperFirst(s string) string
```

<b>Example:</b>

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

<p>Pads string on the left and right side if it's shorter than size.</p>

<b>Signature:</b>

```go
func Pad(source string, size int, padStr string) string
```

<b>Example:</b>

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

<p>Pads string on the right side if it's shorter than size.</p>

<b>Signature:</b>

```go
func PadEnd(source string, size int, padStr string) string
```

<b>Example:</b>

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

<p>Pads string on the left side if it's shorter than size.</p>

<b>Signature:</b>

```go
func PadStart(source string, size int, padStr string) string
```

<b>Example:</b>

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

<p>Return string whose char order is reversed to the given string.</p>

<b>Signature:</b>

```go
func Reverse(s string) string
```

<b>Example:</b>

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

<p>Coverts string to snake_case, non letters and numbers will be ignored.</p>

<b>Signature:</b>

```go
func SnakeCase(s string) string
```

<b>Example:</b>

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

<p>Coverts string to upper KEBAB-CASE, non letters and numbers will be ignored.</p>

<b>Signature:</b>

```go
func SnakeCase(s string) string
```

<b>Example:</b>

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

<p>Split a given string whether the result contains empty string.</p>

<b>Signature:</b>

```go
func SplitEx(s, sep string, removeEmptyString bool) []string
```

<b>Example:</b>

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

<p>Returns a substring of the specified length starting at the specified offset position.</p>

<b>Signature:</b>

```go
func Substring(s string, offset int, length uint) string
```

<b>Example:</b>

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

<p>Wrap a string with given string.</p>

<b>Signature:</b>

```go
func Wrap(str string, wrapWith string) string
```

<b>Example:</b>

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

<p>Unwrap a given string from anther string. will change source string.</p>

<b>Signature:</b>

```go
func Unwrap(str string, wrapToken string) string
```

<b>Example:</b>

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

<p>Splits a string into words, word only contains alphabetic characters.</p>

<b>Signature:</b>

```go
func SplitWords(s string) []string
```

<b>Example:</b>

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

<p>Return the number of meaningful word, word only contains alphabetic characters.</p>

<b>Signature:</b>

```go
func WordCount(s string) int
```

<b>Example:</b>

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