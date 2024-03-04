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

-   [After](#After)
-   [AfterLast](#AfterLast)
-   [Before](#Before)
-   [BeforeLast](#BeforeLast)
-   [CamelCase](#CamelCase)
-   [Capitalize](#Capitalize)
-   [ContainsAll](#ContainsAll)
-   [ContainsAny](#ContainsAny)
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

<p>Coverts string to camelCase string, non letters and numbers will be ignored.</p>

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

    s4 := strutil.CamelCase("Foo-#1😄$_%^&*(1bar")
    fmt.Println(s4) //foo11Bar
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

<p>KebabCase covert string to kebab-case, non letters and numbers will be ignored.</p>

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
    fmt.Println(s4) //foo-bar
}
```

### <span id="UpperKebabCase">UpperKebabCase</span>

<p>UpperKebabCase covert string to upper KEBAB-CASE, non letters and numbers will be ignored.</p>

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

<p>Coverts string to snake_case, non letters and numbers will be ignored.</p>

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
    fmt.Println(s4) //foo_bar

    s5 := strutil.SnakeCase("Foo-#1😄$_%^&*(1bar")
    fmt.Println(s5) //foo_1_1_bar
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

<p>Splits a string into words, word only contains alphabetic characters.</p>

<b>Signature:</b>

```go
func SplitWords(s string) []string
```

<b>Example:</b>

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

<p>Return the number of meaningful word, word only contains alphabetic characters.</p>

<b>Signature:</b>

```go
func WordCount(s string) int
```

<b>Example:</b>

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

### <span id="StringToBytes">StringToBytes</span>

<p>Converts a string to byte slice without a memory allocation.</p>

<b>Signature:</b>

```go
func StringToBytes(str string) (b []byte)
```

<b>Example:</b>

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

<p>Converts a byte slice to string without a memory allocation.</p>

<b>Signature:</b>

```go
func BytesToString(bytes []byte) string
```

<b>Example:</b>

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

<p>Checks if a string is whitespace or empty.</p>

<b>Signature:</b>

```go
func IsBlank(str string) bool
```

<b>Example:</b>

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

<p>Checks if a string starts with any of an array of specified strings.</p>

<b>Signature:</b>

```go
func HasPrefixAny(str string, prefixes []string) bool
```

<b>Example:</b>

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

<p>Checks if a string ends with any of an array of specified strings.</p>

<b>Signature:</b>

```go
func HasSuffixAny(str string, suffixes []string) bool
```

<b>Example:</b>

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

<p>Returns the index of the first instance of substr in string after offsetting the string by `idxFrom`, or -1 if substr is not present in string.</p>

<b>Signature:</b>

```go
func IndexOffset(str string, substr string, idxFrom int) int
```

<b>Example:</b>

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

<p>Returns a copy of `str`, which is replaced by a map in unordered way, case-sensitively.</p>

<b>Signature:</b>

```go
func ReplaceWithMap(str string, replaces map[string]string) string
```

<b>Example:</b>

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

<p>Strips whitespace (or other characters) from the beginning and end of a string. The optional parameter `characterMask` specifies the additional stripped characters.</p>

<b>Signature:</b>

```go
func Trim(str string, characterMask ...string) string
```

<b>Example:</b>

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

<p>Splits string `str` by a string `delimiter` to a slice, and calls Trim to every element of slice. It ignores the elements which are empty after Trim.</p>

<b>Signature:</b>

```go
func SplitAndTrim(str, delimiter string, characterMask ...string) []string
```

<b>Example:</b>

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

<p>HideString hide some chars in source string with param `replaceChar`. replace range is origin[start : end]. [start, end).</p>

<b>Signature:</b>

```go
func HideString(origin string, start, end int, replaceChar string) string
```

<b>Example:</b>

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

<p>Return true if target string contains all the substrings.</p>

<b>Signature:</b>

```go
func ContainsAll(str string, substrs []string) bool
```

<b>Example:</b>

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

<p>Return true if target string contains any one of the substrings.</p>

<b>Signature:</b>

```go
func ContainsAny(str string, substrs []string) bool
```

<b>Example:</b>

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

<p>Remove whitespace characters from a string. when set repalceAll is true removes all whitespace, false only replaces consecutive whitespace characters with one space.</p>

<b>Signature:</b>

```go
func RemoveWhiteSpace(str string, repalceAll bool) string
```

<b>Example:</b>

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

<p>Return substring between the start and end position(excluded) of source string.</p>

<b>Signature:</b>

```go
func SubInBetween(str string, start string, end string) string
```

<b>Example:</b>

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

<p>HammingDistance calculates the Hamming distance between two strings. The Hamming distance is the number of positions at which the corresponding symbols are different.</p>

<b>Signature:</b>

```go
HammingDistance(a, b string) (int, error)
```

<b>Example:</b>

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