// Use of this source code is governed by MIT license

// Package strutil implements some functions to manipulate string.
package strutil

import (
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// CamelCase coverts string to camelCase string. Non letters and numbers will be ignored.
// Play: https://go.dev/play/p/9eXP3tn2tUy
func CamelCase(s string) string {
	var builder strings.Builder

	strs := splitIntoStrings(s, false)
	for i, str := range strs {
		if i == 0 {
			builder.WriteString(strings.ToLower(str))
		} else {
			builder.WriteString(Capitalize(str))
		}
	}

	return builder.String()
}

// Capitalize converts the first character of a string to upper case and the remaining to lower case.
// Play: https://go.dev/play/p/2OAjgbmAqHZ
func Capitalize(s string) string {
	result := make([]rune, len(s))
	for i, v := range s {
		if i == 0 {
			result[i] = unicode.ToUpper(v)
		} else {
			result[i] = unicode.ToLower(v)
		}
	}

	return string(result)
}

// UpperFirst converts the first character of string to upper case.
// Play: https://go.dev/play/p/sBbBxRbs8MM
func UpperFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToUpper(r)

	return string(r) + s[size:]
}

// LowerFirst converts the first character of string to lower case.
// Play: https://go.dev/play/p/CbzAyZmtJwL
func LowerFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToLower(r)

	return string(r) + s[size:]
}

// PadStart pads string on the left and right side if it's shorter than size.
// Padding characters are truncated if they exceed size.
// Play: https://go.dev/play/p/NzImQq-VF8q
func Pad(source string, size int, padStr string) string {
	return padAtPosition(source, size, padStr, 0)
}

// PadStart pads string on the left side if it's shorter than size.
// Padding characters are truncated if they exceed size.
// Play: https://go.dev/play/p/xpTfzArDfvT
func PadStart(source string, size int, padStr string) string {
	return padAtPosition(source, size, padStr, 1)
}

// PadEnd pads string on the right side if it's shorter than size.
// Padding characters are truncated if they exceed size.
// Play: https://go.dev/play/p/9xP8rN0vz--
func PadEnd(source string, size int, padStr string) string {
	return padAtPosition(source, size, padStr, 2)
}

// KebabCase coverts string to kebab-case, non letters and numbers will be ignored.
// Play: https://go.dev/play/p/dcZM9Oahw-Y
func KebabCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "-")
}

// UpperKebabCase coverts string to upper KEBAB-CASE, non letters and numbers will be ignored
// Play: https://go.dev/play/p/zDyKNneyQXk
func UpperKebabCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "-")
}

// SnakeCase coverts string to snake_case, non letters and numbers will be ignored
// Play: https://go.dev/play/p/tgzQG11qBuN
func SnakeCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "_")
}

// UpperSnakeCase coverts string to upper SNAKE_CASE, non letters and numbers will be ignored
// Play: https://go.dev/play/p/4COPHpnLx38
func UpperSnakeCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "_")
}

// Before returns the substring of the source string up to the first occurrence of the specified string.
// Play: https://go.dev/play/p/JAWTZDS4F5w
func Before(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[0:i]
}

// BeforeLast returns the substring of the source string up to the last occurrence of the specified string.
// Play: https://go.dev/play/p/pJfXXAoG_Te
func BeforeLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[0:i]
}

// After returns the substring after the first occurrence of a specified string in the source string.
// Play: https://go.dev/play/p/RbCOQqCDA7m
func After(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[i+len(char):]
}

// AfterLast returns the substring after the last occurrence of a specified string in the source string.
// Play: https://go.dev/play/p/1TegARrb8Yn
func AfterLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[i+len(char):]
}

// IsString check if the value data type is string or not.
// Play: https://go.dev/play/p/IOgq7oF9ERm
func IsString(v any) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}

// Reverse returns string whose char order is reversed to the given string.
// Play: https://go.dev/play/p/adfwalJiecD
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Wrap a string with given string.
// Play: https://go.dev/play/p/KoZOlZDDt9y
func Wrap(str string, wrapWith string) string {
	if str == "" || wrapWith == "" {
		return str
	}
	var sb strings.Builder
	sb.WriteString(wrapWith)
	sb.WriteString(str)
	sb.WriteString(wrapWith)

	return sb.String()
}

// Unwrap a given string from anther string. will change source string.
// Play: https://go.dev/play/p/Ec2q4BzCpG-
func Unwrap(str string, wrapToken string) string {
	if str == "" || wrapToken == "" {
		return str
	}

	firstIndex := strings.Index(str, wrapToken)
	lastIndex := strings.LastIndex(str, wrapToken)

	if firstIndex == 0 && lastIndex > 0 && lastIndex <= len(str)-1 {
		if len(wrapToken) <= lastIndex {
			str = str[len(wrapToken):lastIndex]
		}
	}

	return str
}

// SplitEx split a given string which can control the result slice contains empty string or not.
// Play: https://go.dev/play/p/Us-ySSbWh-3
func SplitEx(s, sep string, removeEmptyString bool) []string {
	if sep == "" {
		return []string{}
	}

	n := strings.Count(s, sep) + 1
	a := make([]string, n)
	n--
	i := 0
	sepSave := 0
	ignore := false

	for i < n {
		m := strings.Index(s, sep)
		if m < 0 {
			break
		}
		ignore = false
		if removeEmptyString {
			if s[:m+sepSave] == "" {
				ignore = true
			}
		}
		if !ignore {
			a[i] = s[:m+sepSave]
			s = s[m+len(sep):]
			i++
		} else {
			s = s[m+len(sep):]
		}
	}

	var ret []string
	if removeEmptyString {
		if s != "" {
			a[i] = s
			ret = a[:i+1]
		} else {
			ret = a[:i]
		}
	} else {
		a[i] = s
		ret = a[:i+1]
	}

	return ret
}

// Substring returns a substring of the specified length starting at the specified offset position.
// Play: https://go.dev/play/p/q3sM6ehnPDp
func Substring(s string, offset int, length uint) string {
	rs := []rune(s)
	size := len(rs)

	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}
	if offset > size {
		return ""
	}

	if length > uint(size)-uint(offset) {
		length = uint(size - offset)
	}

	str := string(rs[offset : offset+int(length)])

	return strings.Replace(str, "\x00", "", -1)
}

// SplitWords splits a string into words, word only contains alphabetic characters.
// Play: https://go.dev/play/p/KLiX4WiysMM
func SplitWords(s string) []string {
	var word string
	var words []string
	var r rune
	var size, pos int

	isWord := false

	for len(s) > 0 {
		r, size = utf8.DecodeRuneInString(s)

		switch {
		case isLetter(r):
			if !isWord {
				isWord = true
				word = s
				pos = 0
			}

		case isWord && (r == '\'' || r == '-'):
			// is word

		default:
			if isWord {
				isWord = false
				words = append(words, word[:pos])
			}
		}

		pos += size
		s = s[size:]
	}

	if isWord {
		words = append(words, word[:pos])
	}

	return words
}

// WordCount return the number of meaningful word, word only contains alphabetic characters.
// Play: https://go.dev/play/p/bj7_odx3vRf
func WordCount(s string) int {
	var r rune
	var size, count int

	isWord := false

	for len(s) > 0 {
		r, size = utf8.DecodeRuneInString(s)

		switch {
		case isLetter(r):
			if !isWord {
				isWord = true
				count++
			}

		case isWord && (r == '\'' || r == '-'):
			// is word

		default:
			isWord = false
		}

		s = s[size:]
	}

	return count
}

// RemoveNonPrintable remove non-printable characters from a string.
// Play: https://go.dev/play/p/og47F5x_jTZ
func RemoveNonPrintable(str string) string {
	result := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, str)

	return result
}

// StringToBytes converts a string to byte slice without a memory allocation.
// Play: https://go.dev/play/p/7OyFBrf9AxA
func StringToBytes(str string) (b []byte) {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data, bh.Len, bh.Cap = sh.Data, sh.Len, sh.Len
	return b
}

// BytesToString converts a byte slice to string without a memory allocation.
// Play: https://go.dev/play/p/6c68HRvJecH
func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// IsBlank checks if a string is whitespace, empty.
// Play: https://go.dev/play/p/6zXRH_c0Qd3
func IsBlank(str string) bool {
	if len(str) == 0 {
		return true
	}
	// memory copies will occur here, but UTF8 will be compatible
	runes := []rune(str)
	for _, r := range runes {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// HasPrefixAny check if a string starts with any of a slice of specified strings.
// Play: https://go.dev/play/p/8UUTl2C5slo
func HasPrefixAny(str string, prefixes []string) bool {
	if len(str) == 0 || len(prefixes) == 0 {
		return false
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(str, prefix) {
			return true
		}
	}
	return false
}

// HasSuffixAny check if a string ends with any of a slice of specified strings.
// Play: https://go.dev/play/p/sKWpCQdOVkx
func HasSuffixAny(str string, suffixes []string) bool {
	if len(str) == 0 || len(suffixes) == 0 {
		return false
	}
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}
	return false
}

// IndexOffset returns the index of the first instance of substr in string after offsetting the string by `idxFrom`,
// or -1 if substr is not present in string.
// Play: https://go.dev/play/p/qZo4lV2fomB
func IndexOffset(str string, substr string, idxFrom int) int {
	if idxFrom > len(str)-1 || idxFrom < 0 {
		return -1
	}

	return strings.Index(str[idxFrom:], substr) + idxFrom
}

// ReplaceWithMap returns a copy of `str`,
// which is replaced by a map in unordered way, case-sensitively.
// Play: https://go.dev/play/p/h3t7CNj2Vvu
func ReplaceWithMap(str string, replaces map[string]string) string {
	for k, v := range replaces {
		str = strings.ReplaceAll(str, k, v)
	}

	return str
}

// SplitAndTrim splits string `str` by a string `delimiter` to a slice,
// and calls Trim to every element of this slice. It ignores the elements
// which are empty after Trim.
// Play: https://go.dev/play/p/ZNL6o4SkYQ7
func SplitAndTrim(str, delimiter string, characterMask ...string) []string {
	result := make([]string, 0)

	for _, v := range strings.Split(str, delimiter) {
		v = Trim(v, characterMask...)
		if v != "" {
			result = append(result, v)
		}
	}

	return result
}

var (
	// DefaultTrimChars are the characters which are stripped by Trim* functions in default.
	DefaultTrimChars = string([]byte{
		'\t', // Tab.
		'\v', // Vertical tab.
		'\n', // New line (line feed).
		'\r', // Carriage return.
		'\f', // New page.
		' ',  // Ordinary space.
		0x00, // NUL-byte.
		0x85, // Delete.
		0xA0, // Non-breaking space.
	})
)

// Trim strips whitespace (or other characters) from the beginning and end of a string.
// The optional parameter `characterMask` specifies the additional stripped characters.
// Play: https://go.dev/play/p/Y0ilP0NRV3j
func Trim(str string, characterMask ...string) string {
	trimChars := DefaultTrimChars

	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}

	return strings.Trim(str, trimChars)
}

// HideString hide some chars in source string with param `replaceChar`.
// replace range is origin[start : end]. [start, end)
// Play: https://go.dev/play/p/pzbaIVCTreZ)
func HideString(origin string, start, end int, replaceChar string) string {
	size := len(origin)

	if start > size-1 || start < 0 || end < 0 || start > end {
		return origin
	}

	if end > size {
		end = size
	}

	if replaceChar == "" {
		return origin
	}

	startStr := origin[0:start]
	endStr := origin[end:size]

	replaceSize := end - start
	replaceStr := strings.Repeat(replaceChar, replaceSize)

	return startStr + replaceStr + endStr
}

// ContainsAll return true if target string contains all the substrs.
// Play: https://go.dev/play/p/KECtK2Os4zq
func ContainsAll(str string, substrs []string) bool {
	for _, v := range substrs {
		if !strings.Contains(str, v) {
			return false
		}
	}

	return true
}

// ContainsAny return true if target string contains any one of the substrs.
// Play: https://go.dev/play/p/dZGSSMB3LXE
func ContainsAny(str string, substrs []string) bool {
	for _, v := range substrs {
		if strings.Contains(str, v) {
			return true
		}
	}

	return false
}
