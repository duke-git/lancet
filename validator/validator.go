// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package validator implements some validate function for string.
package validator

import (
	"encoding/json"
	"net"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var isAlphaRegexMatcher *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z]+$`)

// IsAlpha checks if the string contains only letters (a-zA-Z)
func IsAlpha(str string) bool {
	return isAlphaRegexMatcher.MatchString(str)
}

// IsAllUpper check if the string is all upper case letters A-Z
func IsAllUpper(str string) bool {
	for _, r := range str {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return str != ""
}

// IsAllLower check if the string is all lower case letters a-z
func IsAllLower(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return str != ""
}

// ContainUpper check if the string contain at least one upper case letter A-Z
func ContainUpper(str string) bool {
	for _, r := range str {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// ContainLower check if the string contain at least one lower case letter A-Z
func ContainLower(str string) bool {
	for _, r := range str {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

var containLetterRegexMatcher *regexp.Regexp = regexp.MustCompile(`[a-zA-Z]`)

// ContainLetter check if the string contain at least one letter
func ContainLetter(str string) bool {
	return containLetterRegexMatcher.MatchString(str)
}

// IsJSON checks if the string is valid JSON
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsNumberStr check if the string can convert to a number.
func IsNumberStr(s string) bool {
	return IsIntStr(s) || IsFloatStr(s)
}

// IsFloatStr check if the string can convert to a float.
func IsFloatStr(str string) bool {
	_, e := strconv.ParseFloat(str, 64)
	return e == nil
}

var isIntStrRegexMatcher *regexp.Regexp = regexp.MustCompile(`^[\+-]?\d+$`)

// IsIntStr check if the string can convert to a integer.
func IsIntStr(str string) bool {
	return isIntStrRegexMatcher.MatchString(str)
}

// IsIp check if the string is a ip address.
func IsIp(ipstr string) bool {
	ip := net.ParseIP(ipstr)
	return ip != nil
}

// IsIpV4 check if the string is a ipv4 address.
func IsIpV4(ipstr string) bool {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return false
	}
	return strings.Contains(ipstr, ".")
}

// IsIpV6 check if the string is a ipv6 address.
func IsIpV6(ipstr string) bool {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return false
	}
	return strings.Contains(ipstr, ":")
}

// IsPort check if the string is a valid net port.
func IsPort(str string) bool {
	if i, err := strconv.ParseInt(str, 10, 64); err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

var isDnsRegexMatcher *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`)

// IsDns check if the string is dns.
func IsDns(dns string) bool {
	return isDnsRegexMatcher.MatchString(dns)
}

var isEmailRegexMatcher *regexp.Regexp = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)

// IsEmail check if the string is a email address.
func IsEmail(email string) bool {
	return isEmailRegexMatcher.MatchString(email)
}

var isChineseMobileRegexMatcher *regexp.Regexp = regexp.MustCompile("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$")

// IsChineseMobile check if the string is chinese mobile number.
func IsChineseMobile(mobileNum string) bool {
	return isChineseMobileRegexMatcher.MatchString(mobileNum)
}

var isChineseIdNumRegexMatcher *regexp.Regexp = regexp.MustCompile(`^[1-9]\d{5}(18|19|20|21|22)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)

// IsChineseIdNum check if the string is chinese id number.
func IsChineseIdNum(id string) bool {
	return isChineseIdNumRegexMatcher.MatchString(id)
}

var containChineseRegexMatcher *regexp.Regexp = regexp.MustCompile("[\u4e00-\u9fa5]")

// ContainChinese check if the string contain mandarin chinese.
func ContainChinese(s string) bool {
	return containChineseRegexMatcher.MatchString(s)
}

var isChinesePhoneRegexMatcher *regexp.Regexp = regexp.MustCompile(`\d{3}-\d{8}|\d{4}-\d{7}`)

// IsChinesePhone check if the string is chinese phone number.
// Valid chinese phone is xxx-xxxxxxxx or xxxx-xxxxxxx
func IsChinesePhone(phone string) bool {
	return isChinesePhoneRegexMatcher.MatchString(phone)
}

var isCreditCardRegexMatcher *regexp.Regexp = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11}|6[27][0-9]{14})$`)

// IsCreditCard check if the string is credit card.
func IsCreditCard(creditCart string) bool {
	return isCreditCardRegexMatcher.MatchString(creditCart)
}

var isBase64RegexMatcher *regexp.Regexp = regexp.MustCompile(`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`)

// IsBase64 check if the string is base64 string.
func IsBase64(base64 string) bool {
	return isBase64RegexMatcher.MatchString(base64)
}

// IsEmptyString check if the string is empty.
func IsEmptyString(str string) bool {
	return len(str) == 0
}

// IsRegexMatch check if the string match the regexp
func IsRegexMatch(str, regex string) bool {
	reg := regexp.MustCompile(regex)
	return reg.MatchString(str)
}

// IsStrongPassword check if the string is strong password, if len(password) is less than the length param, return false
// Strong password: alpha(lower+upper) + number + special chars(!@#$%^&*()?><)
func IsStrongPassword(password string, length int) bool {
	if len(password) < length {
		return false
	}
	var num, lower, upper, special bool
	for _, r := range password {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsLower(r):
			lower = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			special = true
		}
	}

	return num && lower && upper && special

	// go doesn't support regexp (?=re)
	//pattern := `^(?=.*[0-9])(?=.*[a-zA-Z])(?=.*[@#$%^&+=])(?=\S+$).$`
	//reg := regexp.MustCompile(pattern)
	//return reg.MatchString(password)
}

// IsWeakPassword check if the string is weak password
// Weak password: only letter or only number or letter + number
func IsWeakPassword(password string) bool {
	var num, letter, special bool
	for _, r := range password {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsLetter(r):
			letter = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			special = true
		}
	}

	return (num || letter) && !special
}
