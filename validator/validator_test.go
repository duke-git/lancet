package validator

import (
	"testing"

	"github.com/duke-git/lancet/utils"
)

func TestIsNumberStr(t *testing.T) {
	isNumberStr(t, "3.", true)
	isNumberStr(t, "+3.", true)
	isNumberStr(t, "-3.", true)
	isNumberStr(t, "+3e2", true)
	isNumberStr(t, "abc", false)
}

func isNumberStr(t *testing.T, source string, expected bool) {
	res := IsNumberStr(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsNumberStr", source, expected, res)
		t.FailNow()
	}
}

func TestIsFloatStr(t *testing.T) {
	isFloatStr(t, "3.", true)
	isFloatStr(t, "+3.", true)
	isFloatStr(t, "-3.", true)
	isFloatStr(t, "12", true)
	isFloatStr(t, "abc", false)
}

func isFloatStr(t *testing.T, source string, expected bool) {
	res := IsFloatStr(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsFloatStr", source, expected, res)
		t.FailNow()
	}
}

func TestIsIntStr(t *testing.T) {
	isIntStr(t, "+3", true)
	isIntStr(t, "-3", true)
	isIntStr(t, "3.", false)
	isIntStr(t, "abc", false)
}

func isIntStr(t *testing.T, source string, expected bool) {
	res := IsIntStr(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsIntStr", source, expected, res)
		t.FailNow()
	}
}

func TestIsIp(t *testing.T) {
	isIp(t, "127.0.0.1", true)
	isIp(t, "::0:0:0:0:0:0:1", true)
	isIp(t, "120.0.0", false)
	isIp(t, "abc", false)
}

func isIp(t *testing.T, source string, expected bool) {
	res := IsIp(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsIp", source, expected, res)
		t.FailNow()
	}
}

func TestIsIpV4(t *testing.T) {
	isIpV4(t, "127.0.0.1", true)
	isIpV4(t, "::0:0:0:0:0:0:1", false)
}

func isIpV4(t *testing.T, source string, expected bool) {
	res := IsIpV4(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsIpV4", source, expected, res)
		t.FailNow()
	}
}

func TestIsIpV6(t *testing.T) {
	isIpV6(t, "127.0.0.1", false)
	isIpV6(t, "::0:0:0:0:0:0:1", true)
}

func isIpV6(t *testing.T, source string, expected bool) {
	res := IsIpV6(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsIpV6", source, expected, res)
		t.FailNow()
	}
}

func TestIsDns(t *testing.T) {
	isDns(t, "abc.com", true)
	isDns(t, "a.b.com", false)
}

func isDns(t *testing.T, source string, expected bool) {
	res := IsDns(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsDns", source, expected, res)
		t.FailNow()
	}
}

func TestIsEmail(t *testing.T) {
	isEmail(t, "abc@xyz.com", true)
	isEmail(t, "a.b@@com", false)
}

func isEmail(t *testing.T, source string, expected bool) {
	res := IsEmail(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsEmail", source, expected, res)
		t.FailNow()
	}
}

func TestContainChinese(t *testing.T) {
	containChinese(t, "你好", true)
	containChinese(t, "hello", false)
	containChinese(t, "hello你好", true)
}

func containChinese(t *testing.T, source string, expected bool) {
	res := ContainChinese(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsContainChineseChar", source, expected, res)
		t.FailNow()
	}
}

func TestIsChineseMobile(t *testing.T) {
	isChineseMobile(t, "13263527980", true)
	isChineseMobile(t, "434324324", false)
}

func isChineseMobile(t *testing.T, source string, expected bool) {
	res := IsChineseMobile(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsChineseMobile", source, expected, res)
		t.FailNow()
	}
}

func TestIsChinesePhone(t *testing.T) {
	isChinesePhone(t, "010-32116675", true)
	isChinesePhone(t, "0464-8756213", true)
	isChinesePhone(t, "123-87562", false)
}

func isChinesePhone(t *testing.T, source string, expected bool) {
	res := IsChinesePhone(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsChinesePhone", source, expected, res)
		t.FailNow()
	}
}

func TestIsChineseIdNum(t *testing.T) {
	isChineseIdNum(t, "210911192105130715", true)
	isChineseIdNum(t, "21091119210513071X", true)
	isChineseIdNum(t, "21091119210513071x", true)
	isChineseIdNum(t, "123456", false)
}

func isChineseIdNum(t *testing.T, source string, expected bool) {
	res := IsChineseIdNum(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsChineseIdNum", source, expected, res)
		t.FailNow()
	}
}

func TestIsCreditCard(t *testing.T) {
	isCreditCard(t, "4111111111111111", true)
	isCreditCard(t, "123456", false)
}

func isCreditCard(t *testing.T, source string, expected bool) {
	res := IsCreditCard(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsCreditCard", source, expected, res)
		t.FailNow()
	}
}

func TestIsBase64(t *testing.T) {
	isBase64(t, "aGVsbG8", true)
	isBase64(t, "123456", false)
}

func isBase64(t *testing.T, source string, expected bool) {
	res := IsBase64(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsBase64", source, expected, res)
		t.FailNow()
	}
}

func TestIsEmptyString(t *testing.T) {
	isEmptyString(t, "111", false)
	isEmptyString(t, " ", false)
	isEmptyString(t, "\t", false)
	isEmptyString(t, "", true)
}

func isEmptyString(t *testing.T, source string, expected bool) {
	res := IsEmptyString(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsEmptyString", source, expected, res)
		t.FailNow()
	}
}

func TestIsAlpha(t *testing.T) {
	isAlpha(t, "abc", true)
	isAlpha(t, "111", false)
	isAlpha(t, " ", false)
	isAlpha(t, "\t", false)
	isAlpha(t, "", false)
}

func isAlpha(t *testing.T, source string, expected bool) {
	res := IsAlpha(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsAlpha", source, expected, res)
		t.FailNow()
	}
}

func TestIsRegexMatch(t *testing.T) {
	isRegexMatch(t, "abc", `^[a-zA-Z]+$`, true)
	isRegexMatch(t, "1ab", `^[a-zA-Z]+$`, false)
	isRegexMatch(t, "111", `^[a-zA-Z]+$`, false)
	isRegexMatch(t, "", `^[a-zA-Z]+$`, false)
}

func isRegexMatch(t *testing.T, source, regex string, expected bool) {
	res := IsRegexMatch(source, regex)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsRegexMatch", source, expected, res)
		t.FailNow()
	}
}

func TestIsStrongPassword(t *testing.T) {
	isStrongPassword(t, "abc", 3, false)
	isStrongPassword(t, "abc123", 6, false)
	isStrongPassword(t, "abcABC", 6, false)
	isStrongPassword(t, "abc123@#$", 9, false)
	isStrongPassword(t, "abcABC123@#$", 16, false)
	isStrongPassword(t, "abcABC123@#$", 12, true)
	isStrongPassword(t, "abcABC123@#$", 10, true)
}

func isStrongPassword(t *testing.T, source string, length int, expected bool) {
	res := IsStrongPassword(source, length)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsStrongPassword", source, expected, res)
		t.FailNow()
	}
}

func TestIsWeakPassword(t *testing.T) {
	isWeakPassword(t, "abc", true)
	isWeakPassword(t, "123", true)
	isWeakPassword(t, "abc123", true)
	isWeakPassword(t, "abcABC123", true)
	isWeakPassword(t, "abc123@#$", false)
}

func isWeakPassword(t *testing.T, source string, expected bool) {
	res := IsWeakPassword(source)
	if res != expected {
		utils.LogFailedTestInfo(t, "IsWeakPassword", source, expected, res)
		t.FailNow()
	}
}
