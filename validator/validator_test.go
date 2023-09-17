package validator

import (
	"fmt"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/duke-git/lancet/internal"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func TestIsAllUpper(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsAllUpper")

	assert.Equal(true, IsAllUpper("ABC"))
	assert.Equal(false, IsAllUpper(""))
	assert.Equal(false, IsAllUpper("abc"))
	assert.Equal(false, IsAllUpper("aBC"))
	assert.Equal(false, IsAllUpper("1BC"))
	assert.Equal(false, IsAllUpper("1bc"))
	assert.Equal(false, IsAllUpper("123"))
	assert.Equal(false, IsAllUpper("你好"))
	assert.Equal(false, IsAllUpper("A&"))
	assert.Equal(false, IsAllUpper("&@#$%^&*"))
}

func TestIsAllLower(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsAllLower")

	assert.Equal(true, IsAllLower("abc"))
	assert.Equal(false, IsAllLower("ABC"))
	assert.Equal(false, IsAllLower(""))
	assert.Equal(false, IsAllLower("aBC"))
	assert.Equal(false, IsAllLower("1BC"))
	assert.Equal(false, IsAllLower("1bc"))
	assert.Equal(false, IsAllLower("123"))
	assert.Equal(false, IsAllLower("你好"))
	assert.Equal(false, IsAllLower("A&"))
	assert.Equal(false, IsAllLower("&@#$%^&*"))
}

func TestContainLower(t *testing.T) {
	assert := internal.NewAssert(t, "TestContainLower")

	assert.Equal(true, ContainLower("abc"))
	assert.Equal(true, ContainLower("aBC"))
	assert.Equal(true, ContainLower("1bc"))
	assert.Equal(true, ContainLower("a&"))

	assert.Equal(false, ContainLower("ABC"))
	assert.Equal(false, ContainLower(""))
	assert.Equal(false, ContainLower("1BC"))
	assert.Equal(false, ContainLower("123"))
	assert.Equal(false, ContainLower("你好"))
	assert.Equal(false, ContainLower("&@#$%^&*"))
}

func TestContainUpper(t *testing.T) {
	assert := internal.NewAssert(t, "TestContainUpper")

	assert.Equal(true, ContainUpper("ABC"))
	assert.Equal(true, ContainUpper("aBC"))
	assert.Equal(true, ContainUpper("1BC"))
	assert.Equal(true, ContainUpper("A&"))

	assert.Equal(false, ContainUpper("abc"))
	assert.Equal(false, ContainUpper(""))
	assert.Equal(false, ContainUpper("1bc"))
	assert.Equal(false, ContainUpper("123"))
	assert.Equal(false, ContainUpper("你好"))
	assert.Equal(false, ContainUpper("&@#$%^&*"))
}

func TestContainLetter(t *testing.T) {
	assert := internal.NewAssert(t, "TestContainLetter")

	assert.Equal(true, ContainLetter("ABC"))
	assert.Equal(true, ContainLetter("1Bc"))
	assert.Equal(true, ContainLetter("1ab"))
	assert.Equal(true, ContainLetter("A&"))

	assert.Equal(false, ContainLetter(""))
	assert.Equal(false, ContainLetter("123"))
	assert.Equal(false, ContainLetter("你好"))
	assert.Equal(false, ContainLetter("&@#$%^&*"))
}

func TestIsJSON(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsJSON")

	assert.Equal(true, IsJSON("{}"))
	assert.Equal(true, IsJSON("{\"name\": \"test\"}"))
	assert.Equal(true, IsJSON("[]"))
	assert.Equal(true, IsJSON("123"))

	assert.Equal(false, IsJSON(""))
	assert.Equal(false, IsJSON("abc"))
	assert.Equal(false, IsJSON("你好"))
	assert.Equal(false, IsJSON("&@#$%^&*"))
}

func TestIsNumberStr(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsNumberStr")

	assert.Equal(true, IsNumberStr("3."))
	assert.Equal(true, IsNumberStr("+3."))
	assert.Equal(true, IsNumberStr("-3."))
	assert.Equal(true, IsNumberStr("+3e2"))
	assert.Equal(false, IsNumberStr("abc"))
}

func TestIsFloatStr(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsFloatStr")

	assert.Equal(true, IsFloatStr("3."))
	assert.Equal(true, IsFloatStr("+3."))
	assert.Equal(true, IsFloatStr("-3."))
	assert.Equal(true, IsFloatStr("12"))
	assert.Equal(false, IsFloatStr("abc"))
}

func TestIsIntStr(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsIntStr")

	assert.Equal(true, IsIntStr("+3"))
	assert.Equal(true, IsIntStr("-3"))
	assert.Equal(false, IsIntStr("3."))
	assert.Equal(false, IsIntStr("abc"))
}

func TestIsPort(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsPort")

	assert.Equal(true, IsPort("1"))
	assert.Equal(true, IsPort("65535"))
	assert.Equal(false, IsPort("abc"))
	assert.Equal(false, IsPort("123abc"))
	assert.Equal(false, IsPort(""))
	assert.Equal(false, IsPort("-1"))
	assert.Equal(false, IsPort("65536"))
}

func TestIsIp(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsIntStr")

	assert.Equal(true, IsIp("127.0.0.1"))
	assert.Equal(true, IsIp("::0:0:0:0:0:0:1"))
	assert.Equal(false, IsIp("127.0.0"))
	assert.Equal(false, IsIp("127"))
}

func TestIsIpV4(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsIpV4")

	assert.Equal(true, IsIpV4("127.0.0.1"))
	assert.Equal(false, IsIpV4("::0:0:0:0:0:0:1"))
}

func TestIsIpV6(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsIpV6")

	assert.Equal(false, IsIpV6("127.0.0.1"))
	assert.Equal(true, IsIpV6("::0:0:0:0:0:0:1"))
}

func TestIsUrl(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsUrl")

	assert.Equal(true, IsUrl("http://abc.com"))
	assert.Equal(true, IsUrl("abc.com"))
	assert.Equal(true, IsUrl("a.b.com"))
	assert.Equal(false, IsUrl("abc"))
}

func TestIsDns(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsDns")

	assert.Equal(true, IsDns("abc.com"))
	assert.Equal(false, IsDns("a.b.com"))
	assert.Equal(false, IsDns("http://abc.com"))
}

func TestIsEmail(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsEmail")

	assert.Equal(true, IsEmail("abc@xyz.com"))
	assert.Equal(false, IsEmail("a.b@@com"))
}

func TestContainChinese(t *testing.T) {
	assert := internal.NewAssert(t, "TestContainChinese")

	assert.Equal(true, ContainChinese("你好"))
	assert.Equal(true, ContainChinese("你好hello"))
	assert.Equal(false, ContainChinese("hello"))
}

func TestIsChineseMobile(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsChineseMobile")

	assert.Equal(true, IsChineseMobile("13263527980"))
	assert.Equal(false, IsChineseMobile("434324324"))
}

func TestIsChinesePhone(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsChinesePhone")

	assert.Equal(true, IsChinesePhone("010-32116675"))
	assert.Equal(true, IsChinesePhone("0464-8756213"))
	assert.Equal(false, IsChinesePhone("123-87562"))

}

func TestIsChineseIdNum(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsChineseIdNum")

	assert.Equal(true, IsChineseIdNum("210911192105130715"))
	assert.Equal(true, IsChineseIdNum("21091119210513071X"))
	assert.Equal(true, IsChineseIdNum("21091119210513071x"))
	assert.Equal(false, IsChineseIdNum("123456"))
}

func TestIsCreditCard(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsCreditCard")

	assert.Equal(true, IsCreditCard("4111111111111111"))
	assert.Equal(false, IsCreditCard("123456"))
}

func TestIsBase64(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsBase64")

	assert.Equal(true, IsBase64("aGVsbG8="))
	assert.Equal(false, IsBase64("123456"))
}

func TestIsEmptyString(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsEmptyString")

	assert.Equal(true, IsEmptyString(""))
	assert.Equal(false, IsEmptyString("111"))
	assert.Equal(false, IsEmptyString(" "))
	assert.Equal(false, IsEmptyString("\t"))
}

func TestIsAlpha(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsAlpha")

	assert.Equal(true, IsAlpha("abc"))
	assert.Equal(false, IsAlpha("111"))
	assert.Equal(false, IsAlpha(" "))
	assert.Equal(false, IsAlpha("\t"))
	assert.Equal(false, IsAlpha(""))
}

func TestIsRegexMatch(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsRegexMatch")

	assert.Equal(true, IsRegexMatch("abc", `^[a-zA-Z]+$`))
	assert.Equal(false, IsRegexMatch("1ab", `^[a-zA-Z]+$`))
	assert.Equal(false, IsRegexMatch("", `^[a-zA-Z]+$`))
}

func TestIsStrongPassword(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsStrongPassword")

	assert.Equal(false, IsStrongPassword("abc", 3))
	assert.Equal(false, IsStrongPassword("abc123", 6))
	assert.Equal(false, IsStrongPassword("abcABC", 6))
	assert.Equal(false, IsStrongPassword("abc123@#$", 9))
	assert.Equal(false, IsStrongPassword("abcABC123@#$", 16))
	assert.Equal(true, IsStrongPassword("abcABC123@#$", 12))
	assert.Equal(true, IsStrongPassword("abcABC123@#$", 10))
}

func TestIsWeakPassword(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsWeakPassword")

	assert.Equal(true, IsWeakPassword("abc"))
	assert.Equal(true, IsWeakPassword("123"))
	assert.Equal(true, IsWeakPassword("abc123"))
	assert.Equal(true, IsWeakPassword("abcABC123"))
	assert.Equal(false, IsWeakPassword("abc123@#$"))
}

func TestIsZeroValue(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsZeroValue")

	var (
		zeroPtr   *string
		zeroSlice []int
		zeroFunc  func() string
		zeroMap   map[string]string
		nilIface  interface{}
		zeroIface fmt.Formatter
	)
	zeroValues := []interface{}{
		nil,
		false,
		0,
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),

		0.0,
		float32(0.0),
		float64(0.0),

		"",

		// func
		zeroFunc,

		// array / slice
		[0]int{},
		zeroSlice,

		// map
		zeroMap,

		// interface
		nilIface,
		zeroIface,

		// pointer
		zeroPtr,

		// struct
		time.Time{},
	}

	for _, value := range zeroValues {
		assert.Equal(true, IsZeroValue(value))
	}

	var nonZeroIface fmt.Stringer = time.Now()

	nonZeroValues := []interface{}{
		// bool
		true,

		// int
		1,
		int8(1),
		int16(1),
		int32(1),
		int64(1),
		uint8(1),
		uint16(1),
		uint32(1),
		uint64(1),

		// float
		1.0,
		float32(1.0),
		float64(1.0),

		// string
		"test",

		// func
		time.Now,

		// array / slice
		[]int{},
		[]int{42},
		[1]int{42},

		// map
		make(map[string]string, 1),

		// interface
		nonZeroIface,

		// pointer
		&nonZeroIface,

		// struct
		time.Now(),
	}

	for _, value := range nonZeroValues {
		assert.Equal(false, IsZeroValue(value))
	}
}

func TestIsGBK(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsGBK")

	str := "你好"
	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))

	assert.Equal(true, IsGBK(gbkData))
	assert.Equal(false, utf8.Valid(gbkData))
}

func TestIsNumber(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsNumber")

	assert.Equal(false, IsNumber(""))
	assert.Equal(false, IsNumber("3"))
	assert.Equal(true, IsNumber(0))
	assert.Equal(true, IsNumber(0.1))
}

func TestIsFloat(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsFloat")

	assert.Equal(false, IsFloat(""))
	assert.Equal(false, IsFloat("3"))
	assert.Equal(false, IsFloat(0))
	assert.Equal(true, IsFloat(0.1))
}

func TestIsInt(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsInt")

	assert.Equal(false, IsInt(""))
	assert.Equal(false, IsInt("3"))
	assert.Equal(false, IsInt(0.1))
	assert.Equal(true, IsInt(0))
	assert.Equal(true, IsInt(-1))
}

func TestIsASCII(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsASCII")

	assert.Equal(true, IsASCII("ABC"))
	assert.Equal(true, IsASCII("123"))
	assert.Equal(true, IsASCII(""))
	assert.Equal(false, IsASCII("😄"))
	assert.Equal(false, IsASCII("你好"))
}

func TestIsPrintable(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsPrintable")

	assert.Equal(true, IsPrintable("ABC"))
	assert.Equal(true, IsPrintable("{id: 123}"))
	assert.Equal(true, IsPrintable(""))
	assert.Equal(true, IsPrintable("😄"))
	assert.Equal(false, IsPrintable("\u0000"))
}

func TestIsBin(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsBin")

	assert.Equal(true, IsBin("0101"))
	assert.Equal(true, IsBin("0b1101"))

	assert.Equal(false, IsBin("b1101"))
	assert.Equal(false, IsBin("1201"))
	assert.Equal(false, IsBin(""))
}

func TestIsHex(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsHex")

	assert.Equal(true, IsHex("ABCDE"))
	assert.Equal(true, IsHex("abcde"))
	assert.Equal(true, IsHex("0xabcde"))
	assert.Equal(true, IsHex("0Xabcde"))
	assert.Equal(true, IsHex("#abcde"))

	assert.Equal(false, IsHex("cdfeg"))
	assert.Equal(false, IsHex("0xcdfeg"))
	assert.Equal(false, IsHex(""))
}

func TestIsBase64URL(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsBase64URL")

	assert.Equal(true, IsBase64URL("SAGsbG8sIHdvcmxkIQ"))
	assert.Equal(true, IsBase64URL("SAGsbG8sIHdvcmxkIQ=="))

	assert.Equal(false, IsBase64URL("SAGsbG8sIHdvcmxkIQ="))
	assert.Equal(false, IsBase64URL("SAGsbG8sIHdvcmxkIQ==="))
	// assert.Equal(false, IsBase64URL(""))
}

func TestIsJWT(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsJWT")

	assert.Equal(true, IsJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibWVzc2FnZSI6IlB1dGluIGlzIGFic29sdXRlIHNoaXQiLCJpYXQiOjE1MTYyMzkwMjJ9.wkLWA5GtCpWdxNOrRse8yHZgORDgf8TpJp73WUQb910"))
	assert.Equal(false, IsJWT("abc"))
}

func TestIsVisa(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsVisa")

	assert.Equal(true, IsVisa("4111111111111111"))
	assert.Equal(false, IsVisa("123"))
}

func TestIsMasterCard(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsMasterCard")

	assert.Equal(true, IsMasterCard("5425233430109903"))
	assert.Equal(false, IsMasterCard("4111111111111111"))
}

func TestIsAmericanExpress(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsAmericanExpress")

	assert.Equal(true, IsAmericanExpress("342883359122187"))
	assert.Equal(false, IsAmericanExpress("3782822463100007"))
}

func TestIsUnionPay(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsUnionPay")

	assert.Equal(true, IsUnionPay("6221263430109903"))
	assert.Equal(false, IsUnionPay("3782822463100007"))
}

func TestIsChinaUnionPay(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsChinaUnionPay")

	assert.Equal(true, IsChinaUnionPay("6250941006528599"))
	assert.Equal(false, IsChinaUnionPay("3782822463100007"))
}
