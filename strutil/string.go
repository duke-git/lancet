// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package strutil implements some functions to manipulate string.
package strutil

import (
	"regexp"
	"strings"
	"unicode"
)

// CamelCase covert string to camelCase string.
func CamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	res := ""
	blankSpace := " "
	regex, _ := regexp.Compile("[-_&]+")
	ss := regex.ReplaceAllString(s, blankSpace)
	for i, v := range strings.Split(ss, blankSpace) {
		vv := []rune(v)
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 96 {
				vv[0] += 32
			}
			res += string(vv)
		} else {
			res += Capitalize(v)
		}
	}

	return res
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

// LowerFirst converts the first character of string to lower case.
func LowerFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	res := ""
	for i, v := range []rune(s) {
		if i == 0 {
			if v >= 65 && v <= 96 {
				v += 32
				res += string(v)
			} else {
				return s
			}
		} else {
			res += string(v)
		}
	}
	return res
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

	var res []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			res = append(res, splitWords...)
		}
	}

	return strings.Join(res, "-")
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

	var res []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			res = append(res, splitWords...)
		}
	}

	return strings.Join(res, "_")
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
func IsString(v interface{}) bool {
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

// ReverseStr return string whose char order is reversed to the given string
func ReverseStr(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
