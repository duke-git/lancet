# Validator

validator 验证器包，包含常用字符串格式验证函数。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/validator/validator.go](https://github.com/duke-git/lancet/blob/main/validator/validator.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/validator"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录：

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
-   [IsPrintable](#IsPrintable)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="ContainChinese">ContainChinese</span>

<p>验证字符串是否包含中文字符</p>

<b>函数签名:</b>

```go
func ContainChinese(s string) bool
```

<b>示例:</b>

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

<p>验证字符串是否包含至少一个英文字母</p>

<b>函数签名:</b>

```go
func ContainLetter(str string) bool
```

<b>示例:</b>

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

<p>验证字符串是否包含至少一个英文小写字母</p>

<b>函数签名:</b>

```go
func ContainLower(str string) bool
```

<b>示例:</b>

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

<p>验证字符串是否包含至少一个英文大写字母.</p>

<b>函数签名:</b>

```go
func ContainUpper(str string) bool
```

<b>示例:</b>

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

<p>验证字符串是否只包含英文字母</p>

<b>函数签名:</b>

```go
func IsAlpha(s string) bool
```

<b>示例:</b>

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

<p>验证字符串是否全是大写英文字母</p>

<b>函数签名:</b>

```go
func IsAllUpper(str string) bool
```

<b>示例:</b>

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

<p>验证字符串是否全是小写英文字母</p>

<b>函数签名:</b>

```go
func IsAllLower(str string) bool
```

<b>示例:</b>

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

<p>验证字符串全部为ASCII字符。</p>

<b>函数签名:</b>

```go
func IsASCII(str string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是base64编码</p>

<b>函数签名:</b>

```go
func IsBase64(base64 string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是中国手机号码</p>

<b>函数签名:</b>

```go
func IsChineseMobile(mobileNum string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是中国身份证号码</p>

<b>函数签名:</b>

```go
func IsChineseIdNum(id string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是中国电话座机号码</p>

<b>函数签名:</b>

```go
func IsChinesePhone(phone string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是信用卡号码</p>

<b>函数签名:</b>

```go
func IsCreditCard(creditCart string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是有效dns</p>

<b>函数签名:</b>

```go
func IsDns(dns string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是有效电子邮件地址</p>

<b>函数签名:</b>

```go
func IsEmail(email string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是空字符串</p>

<b>函数签名:</b>

```go
func IsEmptyString(s string) bool
```

<b>示例:</b>

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

### <span id="IsInt">IsInt</span>

<p>验证参数是否是整数(int, unit)。</p>

<b>函数签名:</b>

```go
func IsInt(v any) bool
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
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

<p>验证参数是否是浮点数((float32, float34)。</p>

<b>函数签名:</b>

```go
func IsFloat(v any) bool
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
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

<p>验证参数是否是数字(integer or float)</p>

<b>函数签名:</b>

```go
func IsNumber(v any) bool
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
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

<p>验证字符串是否是可以转换为整数</p>

<b>函数签名:</b>

```go
func IsIntStr(s string) bool
```

<b>示例:</b>

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

### <span id="IsFloatStr">IsFloatStr</span>

<p>验证字符串是否是可以转换为浮点数</p>

<b>函数签名:</b>

```go
func IsFloatStr(s string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是可以转换为数字</p>

<b>函数签名:</b>

```go
func IsNumberStr(s string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是有效json</p>

<b>函数签名:</b>

```go
func IsJSON(str string) bool
```

<b>示例:</b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/validator"
)

func main() {
    result1 := validator.IsJSON("{}")
    result2 := validator.IsJSON("{\"name\": \"test\"}")
    result3 := validator.IsJSON("")
    result4 := validator.IsJSON("abc")

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

<p>验证字符串是否可以匹配正则表达式</p>

<b>函数签名:</b>

```go
func IsRegexMatch(s, regex string) bool
```

<b>示例:</b>

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



### <span id="IsIp">IsIp</span>

<p>验证字符串是否是ip地址</p>

<b>函数签名:</b>

```go
func IsIp(ipstr string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是ipv4地址</p>

<b>函数签名:</b>

```go
func IsIpV4(ipstr string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是ipv6地址</p>

<b>函数签名:</b>

```go
func IsIpV6(ipstr string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是强密码：(alpha(lower+upper) + number + special chars(!@#$%^&*()?><))</p>

<b>函数签名:</b>

```go
func IsStrongPassword(password string, length int) bool
```

<b>示例:</b>

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

<p>验证字符串是否是url</p>

<b>函数签名:</b>

```go
func IsUrl(str string) bool
```

<b>示例:</b>

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

<p>验证字符串是否是弱密码：（only letter or only number or letter + number）
.</p>

<b>函数签名:</b>

```go
func IsWeakPassword(password string, length int) bool
```

<b>示例:</b>

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

<p>判断传入的参数值是否为零值</p>

<b>函数签名:</b>

```go
func IsZeroValue(value any) bool
```

<b>示例:</b>

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

<p>检查数据编码是否为gbk（汉字内部代码扩展规范）。该函数的实现取决于双字节是否在gbk的编码范围内，而utf-8编码格式的每个字节都在gbk编码范围内。因此，应该首先调用utf8.valid检查它是否是utf-8编码，然后调用IsGBK检查gbk编码。如示例所示。</p>

<b>函数签名:</b>

```go
func IsGBK(data []byte) bool
```

<b>示例:</b>

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

<p>检查字符串是否全部为可打印字符。</p>

<b>函数签名:</b>

```go
func IsPrintable(str string) bool
```

<b>示例:</b>

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