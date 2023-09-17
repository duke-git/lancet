# Validator

Package validator contains some functions for data validation.

<div STYLE="page-break-after: always;"></div>

## Source:

[https://github.com/duke-git/lancet/blob/v1/validator/validator.go](https://github.com/duke-git/lancet/blob/v1/validator/validator.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/validator"
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
-   [IsBase64](#IsBase64)
-   [IsChineseMobile](#IsChineseMobile)
-   [IsChineseIdNum](#IsChineseIdNum)
-   [IsChinesePhone](#IsChinesePhone)
-   [IsCreditCard](#IsCreditCard)
-   [IsDns](#IsDns)
-   [IsEmail](#IsEmail)
-   [IsEmptyString](#IsEmptyString)
-   [IsInt](#IsInt)
-   [IsFloat](#IsFloat)
-   [IsNumber](#IsNumber)
-   [IsIntStr](#IsIntStr)
-   [IsFloatStr](#IsFloatStr)
-   [IsNumberStr](#IsNumberStr)
-   [IsJSON](#IsJSON)
-   [IsRegexMatch](#IsRegexMatch)
-   [IsIp](#IsIp)
-   [IsIpV4](#IsIpV4)
-   [IsIpV6](#IsIpV6)
-   [IsStrongPassword](#IsStrongPassword)
-   [IsUrl](#IsUrl)
-   [IsWeakPassword](#IsWeakPassword)
-   [IsZeroValue](#IsZeroValue)
-   [IsGBK](#IsGBK)
-   [IsASCII](#IsASCII)
-   [IsAIsPrintableSCII](#IsPrintable)
-   [IsBin](#IsBin)
-   [IsHex](#IsHex)
-   [IsBase64URL](#IsBase64URL)
-   [IsJWT](#IsJWT)
-   [IsVisa](#IsVisa)
-   [IsMasterCard](#IsMasterCard)
-   [IsAmericanExpress](#IsAmericanExpress)
-   [IsUnionPay](#IsUnionPay)
-   [IsChinaUnionPay](#IsChinaUnionPay)

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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.ContainChinese("你好")
    fmt.Println(res1) //true

    res2 := validator.ContainChinese("你好hello")
    fmt.Println(res2) //true

    res3 := validator.ContainChinese("hello")
    fmt.Println(res3) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.ContainLetter("1bc")
    fmt.Println(res1) //true

    res2 := validator.ContainLetter("123")
    fmt.Println(res2) //false

    res3 := validator.ContainLetter("&@#$%^&*")
    fmt.Println(res3) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.ContainLower("1bc")
    fmt.Println(res1) //true

    res2 := validator.ContainLower("123")
    fmt.Println(res2) //false

    res3 := validator.ContainLower("1BC")
    fmt.Println(res3) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.ContainUpper("1bc")
    fmt.Println(res1) //false

    res2 := validator.ContainUpper("123")
    fmt.Println(res2) //false

    res3 := validator.ContainUpper("1BC")
    fmt.Println(res3) //true
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.IsAlpha("abc")
    fmt.Println(res1) //true

    res2 := validator.IsAlpha("1bc")
    fmt.Println(res2) //false

    res3 := validator.IsAlpha("")
    fmt.Println(res3) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.IsAllUpper("ABC")
    fmt.Println(res1) //true

    res2 := validator.IsAllUpper("aBC")
    fmt.Println(res2) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.IsAllLower("abc")
    fmt.Println(res1) //true

    res2 := validator.IsAllLower("abC")
    fmt.Println(res2) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.IsBase64("aGVsbG8=")
    fmt.Println(res1) //true

    res2 := validator.IsBase64("123456")
    fmt.Println(res2) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.IsChineseMobile("13263527980")
    fmt.Println(res1) //true

    res2 := validator.IsChineseMobile("434324324")
    fmt.Println(res2) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.IsChineseIdNum("210911192105130715")
    fmt.Println(res1) //true

    res2 := validator.IsChineseIdNum("123456")
    fmt.Println(res2) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.IsChinesePhone("010-32116675")
    fmt.Println(res1) //true

    res2 := validator.IsChinesePhone("123-87562")
    fmt.Println(res2) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.IsCreditCard("4111111111111111")
    fmt.Println(res1) //true

    res2 := validator.IsCreditCard("123456")
    fmt.Println(res2) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.IsDns("abc.com")
    fmt.Println(res1) //true

    res2 := validator.IsDns("a.b.com")
    fmt.Println(res2) //false

    res3 := validator.IsDns("http://abc.com")
    fmt.Println(res3) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.IsEmail("abc@xyz.com")
    fmt.Println(res1) //true

    res2 := validator.IsEmail("a.b@@com")
    fmt.Println(res2) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    res1 := validator.IsEmptyString("")
    fmt.Println(res1) //true

    res2 := validator.IsEmptyString("abc")
    fmt.Println(res2) //false
}
```

### <span id="IsInt">IsInt</span>

<p>Check if the value is integer(int, unit) or not.</p>

<b>Signature:</b>

```go
func IsInt(v interface{}) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    result1 := validator.IsInt("")
    result2 := validator.IsInt("3")
    result3 := validator.IsInt(0.1)
    result4 := validator.IsInt(0)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // false
    // false
    // false
    // true
}
```

### <span id="IsFloat">IsFloat</span>

<p>Check if the value is float(float32, float34) or not.</p>

<b>Signature:</b>

```go
func IsFloat(v interface{}) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    result1 := validator.IsFloat("")
    result2 := validator.IsFloat("3")
    result3 := validator.IsFloat(0)
    result4 := validator.IsFloat(0.1)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // false
    // false
    // false
    // true
}
```

### <span id="IsNumber">IsNumber</span>

<p>Check if the value is number(integer, float) or not.</p>

<b>Signature:</b>

```go
func IsNumber(v interface{}) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    result1 := validator.IsNumber("")
    result2 := validator.IsNumber("3")
    result3 := validator.IsNumber(0.1)
    result4 := validator.IsNumber(0)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // false
    // false
    // true
    // true
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsIntStr("+3")) //true
    fmt.Println(validator.IsIntStr("-3")) //true
    fmt.Println(validator.IsIntStr("3.")) //false
    fmt.Println(validator.IsIntStr("abc")) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsFloatStr("")) //false
    fmt.Println(validator.IsFloatStr("12a")) //false
    fmt.Println(validator.IsFloatStr("3.")) //true
    fmt.Println(validator.IsFloatStr("+3.")) //true
    fmt.Println(validator.IsFloatStr("-3.")) //true
    fmt.Println(validator.IsFloatStr("12")) //true
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsNumberStr("")) //false
    fmt.Println(validator.IsNumberStr("12a")) //false
    fmt.Println(validator.IsNumberStr("3.")) //true
    fmt.Println(validator.IsNumberStr("+3.")) //true
    fmt.Println(validator.IsNumberStr("-3.")) //true
    fmt.Println(validator.IsNumberStr("+3e2")) //true
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsJSON("")) //false
    fmt.Println(validator.IsJSON("abc")) //false
    fmt.Println(validator.IsJSON("{}")) //true
    fmt.Println(validator.IsJSON("[]")) //true
    fmt.Println(validator.IsJSON("123")) //true
    fmt.Println(validator.IsJSON("{\"name\": \"test\"}")) //true
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsRegexMatch("abc", `^[a-zA-Z]+$`)) //true
    fmt.Println(validator.IsRegexMatch("1ab", `^[a-zA-Z]+$`)) //false
    fmt.Println(validator.IsRegexMatch("", `^[a-zA-Z]+$`)) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsIp("127.0.0.1")) //true
    fmt.Println(validator.IsIp("::0:0:0:0:0:0:1")) //true
    fmt.Println(validator.IsIp("127.0.0")) //false
    fmt.Println(validator.IsIp("127")) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsIpV4("127.0.0.1")) //true
    fmt.Println(validator.IsIpV4("::0:0:0:0:0:0:1")) //false
    fmt.Println(validator.IsIpV4("127.0.0")) //false
    fmt.Println(validator.IsIpV4("127")) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsIpV6("127.0.0.1")) //false
    fmt.Println(validator.IsIpV6("::0:0:0:0:0:0:1")) //true
    fmt.Println(validator.IsIpV6("127.0.0")) //false
    fmt.Println(validator.IsIpV6("127")) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsStrongPassword("abc", 3)) //false
    fmt.Println(validator.IsStrongPassword("abc123", 6)) //false
    fmt.Println(validator.IsStrongPassword("abcABC", 6)) //false
    fmt.Println(validator.IsStrongPassword("abcABC123@#$", 16)) //false
    fmt.Println(validator.IsStrongPassword("abcABC123@#$", 12)) //true
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsUrl("http://abc.com")) //true
    fmt.Println(validator.IsUrl("abc.com")) //true
    fmt.Println(validator.IsUrl("a.b.com")) //true
    fmt.Println(validator.IsUrl("abc")) //false
}
```

### <span id="IsWeakPassword">IsWeakPassword</span>

<p>Check if the string is weak password（only letter or only number or letter + number）
.</p>

<b>Signature:</b>

```go
func IsWeakPassword(password string, length int) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsWeakPassword("abc")) //true
    fmt.Println(validator.IsWeakPassword("123")) //true
    fmt.Println(validator.IsWeakPassword("abc123")) //true
    fmt.Println(validator.IsWeakPassword("abc123@#$")) //false
}
```

### <span id="IsZeroValue">IsZeroValue</span>

<p>Checks if passed value is a zero value.</p>

<b>Signature:</b>

```go
func IsZeroValue(value interface{}) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    fmt.Println(validator.IsZeroValue(nil)) //true
    fmt.Println(validator.IsZeroValue(0)) //true
    fmt.Println(validator.IsZeroValue("")) //true
    fmt.Println(validator.IsZeroValue([]int)) //true
    fmt.Println(validator.IsZeroValue(interface{})) //true

    fmt.Println(validator.IsZeroValue("0")) //false
    fmt.Println(validator.IsZeroValue("nil")) //false
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
    "github.com/duke-git/lancet/validator"
)

func main() {
    data := []byte("你好")

    // check utf8 first
    if utf8.Valid(data) {
        fmt.Println("data encoding is utf-8")
    }else if(validator.IsGBK(data)) {
        fmt.Println("data encoding is GBK")
    }
    fmt.Println("data encoding is unknown")
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
    "github.com/duke-git/lancet/validator"
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
    "github.com/duke-git/lancet/validator"
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

### <span id="IsBin">IsBin</span>

<p>Checks if a give string is a valid binary value or not.</p>

<b>Signature:</b>

```go
func IsBin(v string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    result1 := validator.IsBin("0101")
    result2 := validator.IsBin("0b1101")
    result3 := validator.IsBin("b1101")
    result4 := validator.IsBin("1201")

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

### <span id="IsHex">IsHex</span>

<p>Checks if a give string is a valid hexadecimal value or not.</p>

<b>Signature:</b>

```go
func IsHex(v string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    result1 := validator.IsHex("0xabcde")
    result2 := validator.IsHex("0XABCDE")
    result3 := validator.IsHex("cdfeg")
    result4 := validator.IsHex("0xcdfeg")

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

### <span id="IsBase64URL">IsBase64URL</span>

<p>Checks if a give string is a valid URL-safe Base64 encoded string.</p>

<b>Signature:</b>

```go
func IsBase64URL(v string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    result1 := validator.IsBase64URL("SAGsbG8sIHdvcmxkIQ")
    result2 := validator.IsBase64URL("SAGsbG8sIHdvcmxkIQ==")
    result3 := validator.IsBase64URL("SAGsbG8sIHdvcmxkIQ=")
    result4 := validator.IsBase64URL("SAGsbG8sIHdvcmxkIQ===")

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

### <span id="IsJWT">IsJWT</span>

<p>Checks if a give string is is a valid JSON Web Token (JWT).</p>

<b>Signature:</b>

```go
func IsJWT(v string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    result1 := validator.IsJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibWVzc2FnZSI6IlB1dGluIGlzIGFic29sdXRlIHNoaXQiLCJpYXQiOjE1MTYyMzkwMjJ9.wkLWA5GtCpWdxNOrRse8yHZgORDgf8TpJp73WUQb910")
    result2 := validator.IsJWT("abc")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsVisa">IsVisa</span>

<p>Checks if a give string is a valid visa card nubmer or not.</p>

<b>Signature:</b>

```go
func IsVisa(v string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    result1 := validator.IsVisa("4111111111111111")
    result2 := validator.IsVisa("123")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsMasterCard">IsMasterCard</span>

<p>Checks if a give string is a valid mastercard nubmer or not.</p>

<b>Signature:</b>

```go
func IsMasterCard(v string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsMasterCard("5425233430109903")
    result2 := validator.IsMasterCard("4111111111111111")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsAmericanExpress">IsAmericanExpress</span>

<p>Checks if a give string is a valid american express nubmer or not.</p>

<b>Signature:</b>

```go
func IsAmericanExpress(v string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    result1 := validator.IsAmericanExpress("342883359122187")
    result2 := validator.IsAmericanExpress("3782822463100007")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsUnionPay">IsVisa</span>

<p>Checks if a give string is a valid union pay nubmer or not.</p>

<b>Signature:</b>

```go
func IsUnionPay(v string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    result1 := validator.IsUnionPay("6221263430109903")
    result2 := validator.IsUnionPay("3782822463100007")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="IsChinaUnionPay">IsChinaUnionPay</span>

<p>Checks if a give string is a valid china union pay nubmer or not.</p>

<b>Signature:</b>

```go
func IsChinaUnionPay(v string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/validator"
)

func main() {
    result1 := validator.IsChinaUnionPay("6250941006528599")
    result2 := validator.IsChinaUnionPay("3782822463100007")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```
