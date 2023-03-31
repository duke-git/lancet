# Validator

Package validator contains some functions for data validation.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/validator/validator.go](https://github.com/duke-git/lancet/blob/main/validator/validator.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/validator"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [ContainChinese](#ContainChinese)
-   [ContainLetter](#ContainLetter)
-   [ContainLower](#ContainLower)
-   [ContainUpper](#ContainUpper)
-   [IsAlpha](#IsAlpha)
-   [IsAllUpper](#IsAllUpper)
-   [IsAllLower](#IsAllLower)
-   [IsASCII](#IsASCII)
-   [IsBase64](#IsBase64)
-   [IsChineseMobile](#IsChineseMobile)
-   [IsChineseIdNum](#IsChineseIdNum)
-   [IsChinesePhone](#IsChinesePhone)
-   [IsCreditCard](#IsCreditCard)
-   [IsDns](#IsDns)
-   [IsEmail](#IsEmail)
-   [IsEmptyString](#IsEmptyString)
-   [IsFloatStr](#IsFloatStr)
-   [IsNumberStr](#IsNumberStr)
-   [IsJSON](#IsJSON)
-   [IsRegexMatch](#IsRegexMatch)
-   [IsIntStr](#IsIntStr)
-   [IsIp](#IsIp)
-   [IsIpV4](#IsIpV4)
-   [IsIpV6](#IsIpV6)
-   [IsStrongPassword](#IsStrongPassword)
-   [IsUrl](#IsUrl)
-   [IsWeakPassword](#IsWeakPassword)
-   [IsZeroValue](#IsZeroValue)
-   [IsGBK](#IsGBK)
-   [IsPrintable](#IsPrintable)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="ContainChinese">ContainChinese</span>

<p>Check if the string contain mandarin chinese.</p>

<b>Signature:</b>

```go
func ContainChinese(s string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.ContainChinese("你好")
    result2 := validator.ContainChinese("你好hello")
    result3 := validator.ContainChinese("hello")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // true
    // true
    // false
}
```

### <span id="ContainLetter">ContainLetter</span>

<p>Check if the string contain at least one letter.</p>

<b>Signature:</b>

```go
func ContainLetter(str string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.ContainLetter("你好")
    result2 := validator.ContainLetter("&@#$%^&*")
    result3 := validator.ContainLetter("ab1")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // false
    // false
    // true
}
```

### <span id="ContainLower">ContainLower</span>

<p>Check if the string contain at least one lower case letter a-z.</p>

<b>Signature:</b>

```go
func ContainLower(str string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.ContainLower("abc")
    result2 := validator.ContainLower("aBC")
    result3 := validator.ContainLower("ABC")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // true
    // true
    // false
}
```

### <span id="ContainUpper">ContainUpper</span>

<p>Check if the string contain at least one upper case letter A-Z.</p>

<b>Signature:</b>

```go
func ContainUpper(str string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.ContainUpper("ABC")
    result2 := validator.ContainUpper("abC")
    result3 := validator.ContainUpper("abc")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // true
    // true
    // false
}
```

### <span id="IsAlpha">IsAlpha</span>

<p>Check if the string contains only letters (a-zA-Z).</p>

<b>Signature:</b>

```go
func IsAlpha(s string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsAlpha("abc")
    result2 := validator.IsAlpha("ab1")
    result3 := validator.IsAlpha("")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // true
    // false
    // false
}
```

### <span id="IsAllUpper">IsAllUpper</span>

<p>Check if string is all upper case letters A-Z.</p>

<b>Signature:</b>

```go
func IsAllUpper(str string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsAllUpper("ABC")
    result2 := validator.IsAllUpper("ABc")
    result3 := validator.IsAllUpper("AB1")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // true
    // false
    // false
}
```

### <span id="IsAllLower">IsAllLower</span>

<p>Check if string is all lower case letters a-z.</p>

<b>Signature:</b>

```go
func IsAllLower(str string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsAllLower("abc")
    result2 := validator.IsAllLower("abC")
    result3 := validator.IsAllLower("ab1")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // true
    // false
    // false
}
```

### <span id="IsASCII">IsASCII</span>

<p>Checks if string is all ASCII char.</p>

<b>Signature:</b>

```go
func IsASCII(str string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsASCII("ABC")
    result2 := validator.IsASCII("123")
    result3 := validator.IsASCII("")
    result4 := validator.IsASCII("😄")
    result5 := validator.IsASCII("你好")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)

    // Output:
    // true
    // true
    // true
    // false
    // false
}
```

### <span id="IsBase64">IsBase64</span>

<p>Check if the string is base64 string.</p>

<b>Signature:</b>

```go
func IsBase64(base64 string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsBase64("aGVsbG8=")
    result2 := validator.IsBase64("123456")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsChineseMobile">IsChineseMobile</span>

<p>Check if the string is valid chinese mobile number.</p>

<b>Signature:</b>

```go
func IsChineseMobile(mobileNum string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsChineseMobile("13263527980")
    result2 := validator.IsChineseMobile("434324324")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsChineseIdNum">IsChineseIdNum</span>

<p>Check if the string is chinese id number.</p>

<b>Signature:</b>

```go
func IsChineseIdNum(id string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsChineseIdNum("210911192105130715")
    result2 := validator.IsChineseIdNum("123456")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsChinesePhone">IsChinesePhone</span>

<p>Check if the string is chinese phone number.</p>

<b>Signature:</b>

```go
func IsChinesePhone(phone string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsChinesePhone("010-32116675")
    result2 := validator.IsChinesePhone("123-87562")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsCreditCard">IsCreditCard</span>

<p>Check if the string is credit card.</p>

<b>Signature:</b>

```go
func IsCreditCard(creditCart string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsCreditCard("4111111111111111")
    result2 := validator.IsCreditCard("123456")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsDns">IsDns</span>

<p>Check if the string is valid dns.</p>

<b>Signature:</b>

```go
func IsDns(dns string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsDns("abc.com")
    result2 := validator.IsDns("a.b.com")
    result3 := validator.IsDns("http://abc.com")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // true
    // false
    // false
}
```

### <span id="IsEmail">IsEmail</span>

<p>Check if the string is email address.</p>

<b>Signature:</b>

```go
func IsEmail(email string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsEmail("abc@xyz.com")
    result2 := validator.IsEmail("a.b@@com")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsEmptyString">IsEmptyString</span>

<p>Check if the string is empty or not.</p>

<b>Signature:</b>

```go
func IsEmptyString(s string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsEmptyString("")
    result2 := validator.IsEmptyString(" ")
    result3 := validator.IsEmptyString("\t")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // true
    // false
    // false
}
```

### <span id="IsFloatStr">IsFloatStr</span>

<p>Check if the string can convert to a float.</p>

<b>Signature:</b>

```go
func IsFloatStr(s string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsFloatStr("3.")
    result2 := validator.IsFloatStr("+3.")
    result3 := validator.IsFloatStr("12")
    result4 := validator.IsFloatStr("abc")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // true
    // true
    // true
    // false
}
```

### <span id="IsNumberStr">IsNumberStr</span>

<p>Check if the string can convert to a number.</p>

<b>Signature:</b>

```go
func IsNumberStr(s string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsNumberStr("3.")
    result2 := validator.IsNumberStr("+3.")
    result3 := validator.IsNumberStr("+3e2")
    result4 := validator.IsNumberStr("abc")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // true
    // true
    // true
    // false
}
```

### <span id="IsJSON">IsJSON</span>

<p>Check if the string is valid JSON.</p>

<b>Signature:</b>

```go
func IsJSON(str string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsJSON("{}")
    result2 := validator.IsJSON("{\"name\": \"test\"}")
    result3 := validator.IsIntStr("")
    result4 := validator.IsIntStr("abc")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // true
    // true
    // false
    // false
}
```

### <span id="IsRegexMatch">IsRegexMatch</span>

<p>Check if the string match the regexp.</p>

<b>Signature:</b>

```go
func IsRegexMatch(s, regex string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsRegexMatch("abc", `^[a-zA-Z]+$`)
    result2 := validator.IsRegexMatch("ab1", `^[a-zA-Z]+$`)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsIntStr">IsIntStr</span>

<p>Check if the string can convert to a integer.</p>

<b>Signature:</b>

```go
func IsIntStr(s string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsIntStr("+3")
    result2 := validator.IsIntStr("-3")
    result3 := validator.IsIntStr("3.")
    result4 := validator.IsIntStr("abc")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // true
    // true
    // false
    // false
}
```

### <span id="IsIp">IsIp</span>

<p>Check if the string is a ip address.</p>

<b>Signature:</b>

```go
func IsIp(ipstr string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsIp("127.0.0.1")
    result2 := validator.IsIp("::0:0:0:0:0:0:1")
    result3 := validator.IsIp("127.0.0")
    result4 := validator.IsIp("::0:0:0:0:")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // true
    // true
    // false
    // false
}
```

### <span id="IsIpV4">IsIpV4</span>

<p>Check if the string is a ipv4 address.</p>

<b>Signature:</b>

```go
func IsIpV4(ipstr string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsIpV4("127.0.0.1")
    result2 := validator.IsIpV4("::0:0:0:0:0:0:1")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsIpV6">IsIpV6</span>

<p>Check if the string is a ipv6 address.</p>

<b>Signature:</b>

```go
func IsIpV6(ipstr string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsIpV6("127.0.0.1")
    result2 := validator.IsIpV6("::0:0:0:0:0:0:1")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // false
    // true
}
```

### <span id="IsStrongPassword">IsStrongPassword</span>

<p>Check if the string is strong password (alpha(lower+upper) + number + special chars(!@#$%^&*()?><)).</p>

<b>Signature:</b>

```go
func IsStrongPassword(password string, length int) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsStrongPassword("abcABC", 6)
    result2 := validator.IsStrongPassword("abcABC123@#$", 10)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // false
    // true
}
```

### <span id="IsUrl">IsUrl</span>

<p>Check if the string is url.</p>

<b>Signature:</b>

```go
func IsUrl(str string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsUrl("abc.com")
    result2 := validator.IsUrl("http://abc.com")
    result3 := validator.IsUrl("abc")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // true
    // true
    // false
}
```

### <span id="IsWeakPassword">IsWeakPassword</span>

<p>Checks if the string is weak password（only letter or only number or letter + number）
.</p>

<b>Signature:</b>

```go
func IsWeakPassword(password string, length int) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsWeakPassword("abcABC")
    result2 := validator.IsWeakPassword("abc123@#$")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsZeroValue">IsZeroValue</span>

<p>Checks if passed value is a zero value.</p>

<b>Signature:</b>

```go
func IsZeroValue(value any) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsZeroValue("")
    result2 := validator.IsZeroValue(0)
    result3 := validator.IsZeroValue("abc")
    result4 := validator.IsZeroValue(1)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // true
    // true
    // false
    // false
}
```

### <span id="IsGBK">IsGBK</span>

<p>Checks if data encoding is gbk(Chinese character internal code extension specification). this function is implemented by whether double bytes fall within the encoding range of gbk,while each byte of utf-8 encoding format falls within the encoding range of gbk.Therefore, utf8.valid() should be called first to check whether it is not utf-8 encoding and then call IsGBK() to check gbk encoding. like the example.</p>

<b>Signature:</b>

```go
func IsGBK(data []byte) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "golang.org/x/text/encoding/simplifiedchinese"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    str := "你好"
    gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))

    result := validator.IsGBK(gbkData)

    fmt.Println(result)

    // Output:
    // true
}
```


### <span id="IsPrintable">IsPrintable</span>

<p>Checks if string is all printable chars.</p>

<b>Signature:</b>

```go
func IsPrintable(str string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsPrintable("ABC")
    result2 := validator.IsPrintable("{id: 123}")
    result3 := validator.IsPrintable("")
    result4 := validator.IsPrintable("😄")
    result5 := validator.IsPrintable("\u0000")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)

    // Output:
    // true
    // true
    // true
    // true
    // false
}
```