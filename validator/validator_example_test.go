package validator

import (
	"fmt"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func ExampleContainChinese() {
	result1 := ContainChinese("你好")
	result2 := ContainChinese("你好hello")
	result3 := ContainChinese("hello")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// true
	// false
}

func ExampleContainLetter() {
	result1 := ContainLetter("你好")
	result2 := ContainLetter("&@#$%^&*")
	result3 := ContainLetter("ab1")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// false
	// false
	// true
}

func ExampleContainLower() {
	result1 := ContainLower("abc")
	result2 := ContainLower("aBC")
	result3 := ContainLower("ABC")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// true
	// false
}

func ExampleContainUpper() {
	result1 := ContainUpper("ABC")
	result2 := ContainUpper("abC")
	result3 := ContainUpper("abc")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// true
	// false
}

func ExampleIsAlpha() {
	result1 := IsAlpha("abc")
	result2 := IsAlpha("ab1")
	result3 := IsAlpha("")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// false
	// false
}

func ExampleIsAllUpper() {
	result1 := IsAllUpper("ABC")
	result2 := IsAllUpper("ABc")
	result3 := IsAllUpper("AB1")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// false
	// false
}

func ExampleIsAllLower() {
	result1 := IsAllLower("abc")
	result2 := IsAllLower("abC")
	result3 := IsAllLower("ab1")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// false
	// false
}

func ExampleIsBase64() {
	result1 := IsBase64("aGVsbG8=")
	result2 := IsBase64("123456")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsChineseMobile() {
	result1 := IsChineseMobile("13263527980")
	result2 := IsChineseMobile("434324324")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsChineseIdNum() {
	result1 := IsChineseIdNum("210911192105130715")
	result2 := IsChineseIdNum("123456")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsChinesePhone() {
	result1 := IsChinesePhone("010-32116675")
	result2 := IsChinesePhone("123-87562")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsCreditCard() {
	result1 := IsCreditCard("4111111111111111")
	result2 := IsCreditCard("123456")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsDns() {
	result1 := IsDns("abc.com")
	result2 := IsDns("a.b.com")
	result3 := IsDns("http://abc.com")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// false
	// false
}

func ExampleIsUrl() {
	result1 := IsUrl("abc.com")
	result2 := IsUrl("http://abc.com")
	result3 := IsUrl("abc")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// true
	// false
}

func ExampleIsEmail() {
	result1 := IsEmail("abc@xyz.com")
	result2 := IsEmail("a.b@@com")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsEmptyString() {
	result1 := IsEmptyString("")
	result2 := IsEmptyString(" ")
	result3 := IsEmptyString("\t")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// false
	// false
}

func ExampleIsFloatStr() {
	result1 := IsFloatStr("3.")
	result2 := IsFloatStr("+3.")
	result3 := IsFloatStr("12")
	result4 := IsFloatStr("abc")

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

func ExampleIsNumberStr() {
	result1 := IsNumberStr("3.")
	result2 := IsNumberStr("+3.")
	result3 := IsNumberStr("+3e2")
	result4 := IsNumberStr("abc")

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

func ExampleIsIntStr() {
	result1 := IsIntStr("+3")
	result2 := IsIntStr("-3")
	result3 := IsIntStr("3.")
	result4 := IsIntStr("abc")

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

func ExampleIsJSON() {
	result1 := IsJSON("{}")
	result2 := IsJSON("{\"name\": \"test\"}")
	result3 := IsIntStr("")
	result4 := IsIntStr("abc")

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

func ExampleIsRegexMatch() {
	result1 := IsRegexMatch("abc", `^[a-zA-Z]+$`)
	result2 := IsRegexMatch("ab1", `^[a-zA-Z]+$`)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsIp() {
	result1 := IsIp("127.0.0.1")
	result2 := IsIp("::0:0:0:0:0:0:1")
	result3 := IsIp("127.0.0")
	result4 := IsIp("::0:0:0:0:")

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

func ExampleIsIpV4() {
	result1 := IsIpV4("127.0.0.1")
	result2 := IsIpV4("::0:0:0:0:0:0:1")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsIpV6() {
	result1 := IsIpV6("127.0.0.1")
	result2 := IsIpV6("::0:0:0:0:0:0:1")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// false
	// true
}

func ExampleIsStrongPassword() {
	result1 := IsStrongPassword("abcABC", 6)
	result2 := IsStrongPassword("abcABC123@#$", 10)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// false
	// true
}

func ExampleIsWeakPassword() {
	result1 := IsWeakPassword("abcABC")
	result2 := IsWeakPassword("abc123@#$")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleIsZeroValue() {
	result1 := IsZeroValue("")
	result2 := IsZeroValue(0)
	result3 := IsZeroValue("abc")
	result4 := IsZeroValue(1)

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

func ExampleIsGBK() {
	str := "你好"
	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))

	result := IsGBK(gbkData)

	fmt.Println(result)

	// Output:
	// true
}

func ExampleIsASCII() {
	result1 := IsASCII("ABC")
	result2 := IsASCII("123")
	result3 := IsASCII("")
	result4 := IsASCII("😄")
	result5 := IsASCII("你好")

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
