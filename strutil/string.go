// Use of this source code is governed by MIT license

// Package strutil implements some functions to manipulate string.
package strutil

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// CamelCase covert string to camelCase string.
func CamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""
	blankSpace := " "
	regex, _ := regexp.Compile("[-_&]+")
	ss := regex.ReplaceAllString(s, blankSpace)
	for i, v := range strings.Split(ss, blankSpace) {
		vv := []rune(v)
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 96 {
				vv[0] += 32
			}
			result += string(vv)
		} else {
			result += Capitalize(v)
		}
	}

	return result
}

// Capitalize converts the first character of a string to upper case and the remaining to lower case.
func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}

	out := make([]rune, len(s))
	for i, v := range s {
		if i == 0 {
			out[i] = unicode.ToUpper(v)
		} else {
			out[i] = unicode.ToLower(v)
		}
	}

	return string(out)
}

// UpperFirst converts the first character of string to upper case.
func UpperFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToUpper(r)

	return string(r) + s[size:]
}

// LowerFirst converts the first character of string to lower case.
func LowerFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToLower(r)

	return string(r) + s[size:]
}

// PadEnd pads string on the right side if it's shorter than size.
// Padding characters are truncated if they exceed size.
func PadEnd(source string, size int, padStr string) string {
	len1 := len(source)
	len2 := len(padStr)

	if len1 >= size {
		return source
	}

	fill := ""
	if len2 >= size-len1 {
		fill = padStr[0 : size-len1]
	} else {
		fill = strings.Repeat(padStr, size-len1)
	}
	return source + fill[0:size-len1]
}

// PadStart pads string on the left side if it's shorter than size.
// Padding characters are truncated if they exceed size.
func PadStart(source string, size int, padStr string) string {
	len1 := len(source)
	len2 := len(padStr)

	if len1 >= size {
		return source
	}

	fill := ""
	if len2 >= size-len1 {
		fill = padStr[0 : size-len1]
	} else {
		fill = strings.Repeat(padStr, size-len1)
	}
	return fill[0:size-len1] + source
}

// KebabCase covert string to kebab-case
func KebabCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	regex := regexp.MustCompile(`[\W|_]+`)
	blankSpace := " "
	match := regex.ReplaceAllString(s, blankSpace)
	rs := strings.Split(match, blankSpace)

	var result []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			result = append(result, splitWords...)
		}
	}

	return strings.Join(result, "-")
}

// SnakeCase covert string to snake_case
func SnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	regex := regexp.MustCompile(`[\W|_]+`)
	blankSpace := " "
	match := regex.ReplaceAllString(s, blankSpace)
	rs := strings.Split(match, blankSpace)

	var result []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			result = append(result, splitWords...)
		}
	}

	return strings.Join(result, "_")
}

// Before create substring in source string before position when char first appear
func Before(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[0:i]
}

// BeforeLast create substring in source string before position when char last appear
func BeforeLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[0:i]
}

// After create substring in source string after position when char first appear
func After(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[i+len(char):]
}

// AfterLast create substring in source string after position when char last appear
func AfterLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[i+len(char):]
}

// IsString check if the value data type is string or not.
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

// Reverse return string whose char order is reversed to the given string
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Wrap a string with another string.
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

// Unwrap a given string from anther string. will change str value
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

// SplitEx split a given string whether the result contains empty string
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
