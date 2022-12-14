// Use of this source code is governed by MIT license

// Package strutil implements some functions to manipulate string.
package strutil

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// CamelCase covert string to camelCase string.
func CamelCase(s string) string {
	var runes [][]rune
	lastCharType := 0
	charType := 0

	// split into fields based on type of unicode character
	for _, r := range s {
		switch true {
		case unicode.IsLower(r):
			charType = 1
		case unicode.IsUpper(r):
			charType = 2
		case unicode.IsDigit(r):
			charType = 3
		default:
			charType = 4
		}

		if charType == lastCharType {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastCharType = charType
	}

	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}

	// filter all non letters and none digit
	var filterRunes [][]rune
	for _, r := range runes {
		if len(r) == 0 {
			continue
		}
		if unicode.IsLetter(r[0]) || unicode.IsDigit(r[0]) {
			filterRunes = append(filterRunes, r)
		}
	}

	result := ""

	// capitalize
	for i, r := range filterRunes {
		if i == 0 {
			for j := range r {
				r[j] = unicode.ToLower(r[j])
			}
		} else {
			for j := range r {
				if j == 0 {
					r[0] = unicode.ToUpper(r[0])
				} else {
					r[j] = unicode.ToLower(r[j])
				}
			}
		}
		if len(r) > 0 {
			result = result + string(r)
		}
	}

	return result
}

// Capitalize converts the first character of a string to upper case and the remaining to lower case.
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
// non letters and numbers will be ignored
// eg. "Foo-#1😄$_%^&*(1bar" => "foo-1-1-bar"
func KebabCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "-")
}

// UpperKebabCase covert string to upper KEBAB-CASE
// non letters and numbers will be ignored
// eg. "Foo-#1😄$_%^&*(1bar" => "FOO-1-1-BAR"
func UpperKebabCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "-")
}

// SnakeCase covert string to snake_case
// non letters and numbers will be ignored
// eg. "Foo-#1😄$_%^&*(1bar" => "foo_1_1_bar"
func SnakeCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "_")
}

// UpperSnakeCase covert string to upper SNAKE_CASE
// non letters and numbers will be ignored
// eg. "Foo-#1😄$_%^&*(1bar" => "FOO_1_1_BAR"
func UpperSnakeCase(s string) string {
	result := splitIntoStrings(s, true)
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
