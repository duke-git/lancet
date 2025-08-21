package validator

import (
	"fmt"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/duke-git/lancet/v2/internal"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func TestIsAllUpper(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsAllUpper")

	tests := []struct {
		input    string
		expected bool
	}{
		{"ABC", true},
		{"", false},
		{"abc", false},
		{"aBC", false},
		{"1BC", false},
		{"1bc", false},
		{"123", false},
		{"你好", false},
		{"A&", false},
		{"&@#$%^&*", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsAllUpper(tt.input))
	}
}

func TestIsAllLower(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsAllLower")

	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},
		{"", false},
		{"ABC", false},
		{"aBC", false},
		{"1BC", false},
		{"1bc", false},
		{"123", false},
		{"你好", false},
		{"A&", false},
		{"&@#$%^&*", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsAllLower(tt.input))
	}
}

func TestContainLower(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestContainLower")

	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},
		{"aBC", true},
		{"1bc", true},
		{"a&", true},
		{"ABC", false},
		{"", false},
		{"1BC", false},
		{"123", false},
		{"你好", false},
		{"&@#$%^&*", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, ContainLower(tt.input))
	}
}

func TestContainUpper(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestContainUpper")

	tests := []struct {
		input    string
		expected bool
	}{
		{"ABC", true},
		{"aBC", true},
		{"1BC", true},
		{"A&", true},
		{"abc", false},
		{"", false},
		{"1bc", false},
		{"123", false},
		{"你好", false},
		{"&@#$%^&*", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, ContainUpper(tt.input))
	}
}

func TestContainLetter(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestContainLetter")

	tests := []struct {
		input    string
		expected bool
	}{
		{"ABC", true},
		{"1Bc", true},
		{"1ab", true},
		{"A&", true},
		{"", false},
		{"123", false},
		{"你好", false},
		{"&@#$%^&*", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, ContainLetter(tt.input))
	}
}

func TestContainNumber(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestContainNumber")

	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"1Bc", true},
		{"a2c", true},
		{"ab3", true},
		{"a23", true},
		{"a23c", true},
		{"1%%%", true},
		{"ABC", false},
		{"", false},
		{"你好", false},
		{"&@#$%^&*", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, ContainNumber(tt.input))
	}
}

func TestIsJSON(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsJSON")

	tests := []struct {
		input    string
		expected bool
	}{
		{"{}", true},
		{"{\"name\": \"test\"}", true},
		{"[]", true},
		{"123", true},
		{"", false},
		{"abc", false},
		{"你好", false},
		{"&@#$%^&*", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsJSON(tt.input))
	}
}

func TestIsNumber(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsNumber")

	tests := []struct {
		input    interface{}
		expected bool
	}{
		{"", false},
		{"3", false},
		{0, true},
		{0.1, true},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsNumber(tt.input))
	}
}

func TestIsFloat(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsFloat")

	tests := []struct {
		input    interface{}
		expected bool
	}{
		{"", false},
		{"3", false},
		{0, false},
		{0.1, true},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsFloat(tt.input))
	}
}

func TestIsInt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsInt")

	tests := []struct {
		input    interface{}
		expected bool
	}{
		{"", false},
		{"3", false},
		{0.1, false},
		{0, true},
		{-1, true},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsInt(tt.input))
	}
}

func TestIsNumberStr(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsNumberStr")

	tests := []struct {
		input    string
		expected bool
	}{
		{"3", true},
		{"3.", true},
		{"+3.", true},
		{"-3.", true},
		{"+3e2", true},
		{"abc", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsNumberStr(tt.input))
	}
}

func TestIsFloatStr(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsFloatStr")

	tests := []struct {
		input    string
		expected bool
	}{
		{"3", true},
		{"3.", true},
		{"+3.", true},
		{"-3.", true},
		{"12", true},
		{"abc", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsFloatStr(tt.input))
	}
}

func TestIsIntStr(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsIntStr")

	tests := []struct {
		input    string
		expected bool
	}{
		{"3", true},
		{"3.", false},
		{"+3", true},
		{"-3", true},
		{"abc", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsIntStr(tt.input))
	}
}

func TestIsPort(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsPort")

	tests := []struct {
		input    string
		expected bool
	}{
		{"1", true},
		{"65535", true},
		{"abc", false},
		{"123abc", false},
		{"", false},
		{"-1", false},
		{"65536", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsPort(tt.input))
	}
}

func TestIsIp(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsIntStr")

	tests := []struct {
		input    string
		expected bool
	}{
		{"127.0.0.1", true},
		{"::0:0:0:0:0:0:1", true},
		{"127.0.0", false},
		{"127", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsIp(tt.input))
	}
}

func TestIsIpPort(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsIpPort")

	tests := []struct {
		input    string
		expected bool
	}{
		{"[::0:0:0:0:0:0:1]:8080", true},
		{"127.0.0.1:8080", true},

		{"", false},
		{":8080", false},
		{"127.0.0.1", false},
		{"0:0:0:0:0:0:0:1", false},
		{"256.256.256.256:8080", false},
		{"256.256.256.256:abc", false},
		{"127.0.0.1:70000", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsIpPort(tt.input))
	}
}

func TestIsIpV4(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsIpV4")

	tests := []struct {
		input    string
		expected bool
	}{
		{"127.0.0.1", true},
		{"::0:0:0:0:0:0:1", false},

		{"::0:0:0:0:0:0:1", false},
		{"127.0.0.1.1", false},
		{"256.0.0.1", false},
		{"127.0.0.a", false},
		{"", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsIpV4(tt.input))
	}
}

func TestIsIpV6(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsIpV6")

	tests := []struct {
		input    string
		expected bool
	}{
		{"::0:0:0:0:0:0:1", true},
		{"::1", true},
		{"::", true},
		{"127.0.0.1", false},
		{"2001:db8::8a2e:37023:7334", false},
		{"2001::25de::cade", false},
		{"", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsIpV6(tt.input))
	}

}

func TestIsUrl(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsUrl")

	tests := []struct {
		input    string
		expected bool
	}{
		{"http://abc.com", true},
		{"https://abc.com", true},
		{"ftp://abc.com", true},
		{"http://abc.com/path?query=123", true},
		{"https://abc.com/path/to/resource", true},
		{"ws://abc.com", true},
		{"wss://abc.com", true},
		{"mailto://abc.com", true},
		{"file://path/to/file", true},
		{"data://text/plain;base64,SGVsbG8sIFdvcmxkIQ==", true},
		{"http://abc.com/path/to/resource?query=123#fragment", true},

		{"abc", false},
		{"http://", false},
		{"http://abc", false},
		{"http://abc:8080", false},
		{"http://abc:99999999", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsUrl(tt.input))
	}
}

func TestIsDns(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsDns")

	tests := []struct {
		input    string
		expected bool
	}{
		{"abc.com", true},
		{"123.cn", true},
		{"a.b.com", true},
		{"a.b.c", false},
		{"a@b.com", false},
		{"http://abc.com", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsDns(tt.input))
	}
}

func TestIsEmail(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsEmail")

	assert.Equal(true, IsEmail("abc@xyz.com"))
	assert.Equal(false, IsEmail("@abc@xyz.com"))
	assert.Equal(false, IsEmail("a.b@@com"))
}

func TestContainChinese(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestContainChinese")

	tests := []struct {
		input    string
		expected bool
	}{
		{"你好", true},
		{"hello", false},
		{"你好hello", true},
		{"hello你好", true},
		{"", false},
		{"123", false},
		{"!@#$%^&*()", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, ContainChinese(tt.input))
	}
}

func TestIsChineseMobile(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsChineseMobile")

	tests := []struct {
		input    string
		expected bool
	}{
		{"13263527980", true},
		{"1326352798", false},
		{"132635279801", false},
		{"1326352798a", false},
		{"1326352798@", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsChineseMobile(tt.input))
	}
}

func TestIsChinesePhone(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsChinesePhone")

	tests := []struct {
		input    string
		expected bool
	}{
		{"010-32116675", true},
		{"0464-8756213", true},
		{"0731-82251545", true},
		{"123-87562", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsChinesePhone(tt.input))
	}
}

func TestIsChineseIdNum(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsChineseIdNum")

	tests := []struct {
		input    string
		expected bool
	}{
		{"210911192105130714", true},
		{"11010519491231002X", true},
		{"11010519491231002x", true},
		{"123456", false},
		{"990911192105130714", false},
		{"990911189905130714", false},
		{"210911222205130714", false},
		{"", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsChineseIdNum(tt.input))
	}
}

func TestIsCreditCard(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsCreditCard")

	assert.Equal(true, IsCreditCard("4111111111111111"))
	assert.Equal(false, IsCreditCard("123456"))
}

func TestIsBase64(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsBase64")

	assert.Equal(true, IsBase64("aGVsbG8="))
	assert.Equal(false, IsBase64("123456"))
}

func TestIsEmptyString(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsEmptyString")

	assert.Equal(true, IsEmptyString(""))
	assert.Equal(false, IsEmptyString("111"))
	assert.Equal(false, IsEmptyString(" "))
	assert.Equal(false, IsEmptyString("\t"))
}

func TestIsAlpha(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsAlpha")

	assert.Equal(true, IsAlpha("abc"))
	assert.Equal(false, IsAlpha("111"))
	assert.Equal(false, IsAlpha(" "))
	assert.Equal(false, IsAlpha("\t"))
	assert.Equal(false, IsAlpha(""))
}

func TestIsRegexMatch(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsRegexMatch")

	assert.Equal(true, IsRegexMatch("abc", `^[a-zA-Z]+$`))
	assert.Equal(false, IsRegexMatch("1ab", `^[a-zA-Z]+$`))
	assert.Equal(false, IsRegexMatch("", `^[a-zA-Z]+$`))
}

func TestIsStrongPassword(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsStrongPassword")

	tests := []struct {
		input    string
		length   int
		expected bool
	}{
		{"abc", 3, false},
		{"abc123", 6, false},
		{"abcABC", 6, false},
		{"abc123@#$", 9, false},
		{"abcABC123@#$", 16, false},
		{"abcABC123@#$", 12, true},
		{"abcABC123@#$", 10, true},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsStrongPassword(tt.input, tt.length))
	}
}

func TestIsWeakPassword(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsWeakPassword")

	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},
		{"123", true},
		{"abc123", true},
		{"abcABC123", true},
		{"abc123@#$", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsWeakPassword(tt.input))
	}
}

func TestIsZeroValue(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsGBK")

	str := "你好"
	gbkData, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))

	assert.Equal(true, IsGBK(gbkData))
	assert.Equal(false, utf8.Valid(gbkData))
}

func TestIsASCII(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsASCII")

	assert.Equal(true, IsASCII("ABC"))
	assert.Equal(true, IsASCII("123"))
	assert.Equal(true, IsASCII(""))
	assert.Equal(false, IsASCII("😄"))
	assert.Equal(false, IsASCII("你好"))
}

func TestIsPrintable(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsPrintable")

	tests := []struct {
		input    string
		expected bool
	}{
		{"ABC", true},
		{"123", true},
		{"你好", true},
		{"", true},
		{"😄", true},
		{"\u0000", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsPrintable(tt.input))
	}
}

func TestIsBin(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsBin")

	tests := []struct {
		input    string
		expected bool
	}{
		{"0101", true},
		{"0b1101", true},
		{"b1101", false},
		{"1201", false},
		{"", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsBin(tt.input))
	}
}

func TestIsHex(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsHex")

	tests := []struct {
		input    string
		expected bool
	}{
		{"ABCDE", true},
		{"abcde", true},
		{"0xabcde", true},
		{"0Xabcde", true},
		{"#abcde", true},
		{"cdfeg", false},
		{"0xcdfeg", false},
		{"", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsHex(tt.input))
	}
}

func TestIsBase64URL(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestIsBase64URL")

	tests := []struct {
		input    string
		expected bool
	}{
		{"SAGsbG8sIHdvcmxkIQ", true},
		{"SAGsbG8sIHdvcmxkIQ==", true},
		{"SAGsbG8sIHdvcmxkIQ=", false},
		{"SAGsbG8sIHdvcmxkIQ===", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsBase64URL(tt.input))
	}
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

func TestIsAlphaNumeric(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsAlphaNumeric")

	tests := []struct {
		input    string
		expected bool
	}{
		{"ABC", true},
		{"abc", true},
		{"aBC", true},
		{"1BC", true},
		{"1bc", true},
		{"123", true},
		{"你好", false},
		{"A&", false},
		{"&@#$%^&*", false},
		{"", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsAlphaNumeric(tt.input))
	}
}

func TestIsPassport(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsPassport")

	tests := []struct {
		passport    string
		countryCode string
		expected    bool
	}{
		{"P123456789", "CN", true},
		{"123456789", "US", true},
		{"A12345678", "GB", true},
		{"AB1234567", "FR", true},
		{"12345678", "JP", true},
		{"M12345678", "HK", true},
		{"A12345678", "MO", true},
		{"A1234567", "IN", true},
		{"12345678", "IT", true},
		{"A12345678", "AU", true},
		{"123456789", "BR", true},
		{"AB1234567", "RU", true},
		{"123456789", "CN", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, IsPassport(tt.passport, tt.countryCode))
	}
}
