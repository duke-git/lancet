# Validator
validator验证器包，包含常用字符串格式验证函数。

<div STYLE="page-break-after: always;"></div>

## 源码:

[https://github.com/duke-git/lancet/blob/v1/validator/validator.go](https://github.com/duke-git/lancet/blob/v1/validator/validator.go)


<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/validator"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录：
- [ContainChinese](#ContainChinese)
- [ContainLetter](#ContainLetter)
- [ContainLower](#ContainLower)
- [ContainUpper](#ContainUpper)
- [IsAlpha](#IsAlpha)
- [IsAllUpper](#IsAllUpper)
- [IsAllLower](#IsAllLower)
- [IsBase64](#IsBase64)
- [IsChineseMobile](#IsChineseMobile)
- [IsChineseIdNum](#IsChineseIdNum)
- [IsChinesePhone](#IsChinesePhone)
- [IsCreditCard](#IsCreditCard)
- [IsDns](#IsDns)
- [IsEmail](#IsEmail)
- [IsEmptyString](#IsEmptyString)
- [IsFloatStr](#IsFloatStr)
- [IsNumberStr](#IsNumberStr)
- [IsJSON](#IsJSON)
- [IsRegexMatch](#IsRegexMatch)
- [IsIntStr](#IsIntStr)
- [IsIp](#IsIp)
- [IsIpV4](#IsIpV4)
- [IsIpV6](#IsIpV6)
- [IsStrongPassword](#IsStrongPassword)
- [IsUrl](#IsUrl)
- [IsWeakPassword](#IsWeakPassword)
- [IsZeroValue](#IsZeroValue)
  

<div STYLE="page-break-after: always;"></div>

## 文档


### <span id="ContainChinese">ContainChinese</span>
<p>验证字符串是否包含中文字符</p>

<b>函数签名:</b>

```go
func ContainChinese(s string) bool
```
<b>例子:</b>

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
<p>验证字符串是否包含至少一个英文字母</p>

<b>函数签名:</b>

```go
func ContainLetter(str string) bool
```
<b>例子:</b>

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
<p>验证字符串是否包含至少一个英文小写字母</p>

<b>函数签名:</b>

```go
func ContainLower(str string) bool
```
<b>例子:</b>

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
<p>验证字符串是否包含至少一个英文大写字母.</p>

<b>函数签名:</b>

```go
func ContainUpper(str string) bool
```
<b>例子:</b>

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
<p>验证字符串是否只包含英文字母</p>

<b>函数签名:</b>

```go
func IsAlpha(s string) bool
```
<b>例子:</b>

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
<p>验证字符串是否全是大写英文字母</p>

<b>函数签名:</b>

```go
func IsAllUpper(str string) bool
```
<b>例子:</b>

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
<p>验证字符串是否全是小写英文字母</p>

<b>函数签名:</b>

```go
func IsAllLower(str string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是base64编码</p>

<b>函数签名:</b>

```go
func IsBase64(base64 string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是中国手机号码</p>

<b>函数签名:</b>

```go
func IsChineseMobile(mobileNum string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是中国身份证号码</p>

<b>函数签名:</b>

```go
func IsChineseIdNum(id string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是中国电话座机号码</p>

<b>函数签名:</b>

```go
func IsChinesePhone(phone string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是信用卡号码</p>

<b>函数签名:</b>

```go
func IsCreditCard(creditCart string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是有效dns</p>

<b>函数签名:</b>

```go
func IsDns(dns string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是有效电子邮件地址</p>

<b>函数签名:</b>

```go
func IsEmail(email string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是空字符串</p>

<b>函数签名:</b>

```go
func IsEmptyString(s string) bool
```
<b>例子:</b>

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




### <span id="IsFloatStr">IsFloatStr</span>
<p>验证字符串是否是可以转换为浮点数</p>

<b>函数签名:</b>

```go
func IsFloatStr(s string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是可以转换为数字</p>

<b>函数签名:</b>

```go
func IsNumberStr(s string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是有效json</p>

<b>函数签名:</b>

```go
func IsJSON(str string) bool
```
<b>例子:</b>

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
<p>验证字符串是否可以匹配正则表达式</p>

<b>函数签名:</b>

```go
func IsRegexMatch(s, regex string) bool
```
<b>例子:</b>

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




### <span id="IsIntStr">IsIntStr</span>
<p>验证字符串是否是可以转换为整数</p>

<b>函数签名:</b>

```go
func IsIntStr(s string) bool
```
<b>例子:</b>

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




### <span id="IsIp">IsIp</span>
<p>验证字符串是否是ip地址</p>

<b>函数签名:</b>

```go
func IsIp(ipstr string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是ipv4地址</p>

<b>函数签名:</b>

```go
func IsIpV4(ipstr string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是ipv6地址</p>

<b>函数签名:</b>

```go
func IsIpV6(ipstr string) bool
```
<b>例子:</b>

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
<p>验证字符串是否是强密码：(alpha(lower+upper) + number + special chars(!@#$%^&*()?><))</p>

<b>函数签名:</b>

```go
func IsStrongPassword(password string, length int) bool
```
<b>例子:</b>

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
<p>验证字符串是否是url</p>

<b>函数签名:</b>

```go
func IsUrl(str string) bool 
```
<b>例子:</b>

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
<p>验证字符串是否是弱密码：（only letter or only number or letter + number）
.</p>

<b>函数签名:</b>

```go
func IsWeakPassword(password string, length int) bool
```
<b>例子:</b>

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
<p>判断传入的参数值是否为零值</p>

<b>函数签名:</b>

```go
func IsZeroValue(value any) bool
```
<b>例子:</b>

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






