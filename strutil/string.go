// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package strutil implements some functions to manipulate string.
package strutil

import (
	"errors"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"math/rand"
)

// used in `Shuffle` function
var rng = rand.New(rand.NewSource(int64(time.Now().UnixNano())))

// CamelCase covert string to camelCase string.
// non letters and numbers will be ignored
// eg. "Foo-#1😄$_%^&*(1bar" => "foo11Bar"
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

// PadStart pads string on the left and right side if it's shorter than size.
// Padding characters are truncated if they exceed size.
func Pad(source string, size int, padStr string) string {
	return padAtPosition(source, size, padStr, 0)
}

// PadStart pads string on the left side if it's shorter than size.
// Padding characters are truncated if they exceed size.
func PadStart(source string, size int, padStr string) string {
	return padAtPosition(source, size, padStr, 1)
}

// PadEnd pads string on the right side if it's shorter than size.
// Padding characters are truncated if they exceed size.
func PadEnd(source string, size int, padStr string) string {
	return padAtPosition(source, size, padStr, 2)
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
	i := strings.Index(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

	return s[0:i]
}

// BeforeLast create substring in source string before position when char last appear
func BeforeLast(s, char string) string {
	i := strings.LastIndex(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

	return s[0:i]
}

// After create substring in source string after position when char first appear
func After(s, char string) string {
	i := strings.Index(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

	return s[i+len(char):]
}

// AfterLast create substring in source string after position when char last appear
func AfterLast(s, char string) string {
	i := strings.LastIndex(s, char)

	if s == "" || char == "" || i == -1 {
		return s
	}

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

// SplitWords splits a string into words, word only contains alphabetic characters.
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
func StringToBytes(str string) (b []byte) {
	return *(*[]byte)(unsafe.Pointer(&str))
}

// BytesToString converts a byte slice to string without a memory allocation.
func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// IsBlank checks if a string is whitespace, empty.
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

// HasPrefixAny check if a string starts with any of an array of specified strings.
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

// HasSuffixAny check if a string ends with any of an array of specified strings.
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
func IndexOffset(str string, substr string, idxFrom int) int {
	if idxFrom > len(str)-1 || idxFrom < 0 {
		return -1
	}

	return strings.Index(str[idxFrom:], substr) + idxFrom
}

// ReplaceWithMap returns a copy of `str`, which is replaced by a map in unordered way, case-sensitively.
func ReplaceWithMap(str string, replaces map[string]string) string {
	for k, v := range replaces {
		str = strings.ReplaceAll(str, k, v)
	}

	return str
}

// SplitAndTrim splits string `str` by a string `delimiter` to a slice,
// and calls Trim to every element of this slice. It ignores the elements
// which are empty after Trim.
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
func Trim(str string, characterMask ...string) string {
	trimChars := DefaultTrimChars

	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}

	return strings.Trim(str, trimChars)
}

// HideString hide some chars in source string with param `replaceChar`.
// replace range is origin[start : end]. [start, end)
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
func ContainsAll(str string, substrs []string) bool {
	for _, v := range substrs {
		if !strings.Contains(str, v) {
			return false
		}
	}

	return true
}

// ContainsAny return true if target string contains any one of the substrs.
func ContainsAny(str string, substrs []string) bool {
	for _, v := range substrs {
		if strings.Contains(str, v) {
			return true
		}
	}

	return false
}

var (
	whitespaceRegexMatcher     *regexp.Regexp = regexp.MustCompile(`\s`)
	mutiWhitespaceRegexMatcher *regexp.Regexp = regexp.MustCompile(`[[:space:]]{2,}|[\s\p{Zs}]{2,}`)
)

// RemoveWhiteSpace remove whitespace characters from a string.
// when set repalceAll is true removes all whitespace, false only replaces consecutive whitespace characters with one space.
func RemoveWhiteSpace(str string, repalceAll bool) string {
	if repalceAll && str != "" {
		return strings.Join(strings.Fields(str), "")
	} else if str != "" {
		str = mutiWhitespaceRegexMatcher.ReplaceAllString(str, " ")
		str = whitespaceRegexMatcher.ReplaceAllString(str, " ")
	}

	return strings.TrimSpace(str)
}

// SubInBetween return substring between the start and end position(excluded) of source string.
func SubInBetween(str string, start string, end string) string {
	if _, after, ok := strings.Cut(str, start); ok {
		if before, _, ok := strings.Cut(after, end); ok {
			return before
		}
	}

	return ""
}

// HammingDistance calculates the Hamming distance between two strings.
// The Hamming distance is the number of positions at which the corresponding symbols are different.
// This func returns an error if the input strings are of unequal lengths.
func HammingDistance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("a length and b length are unequal")
	}

	ar := []rune(a)
	br := []rune(b)

	var distance int
	for i, codepoint := range ar {
		if codepoint != br[i] {
			distance++
		}
	}

	return distance, nil
}

// Concat uses the strings.Builder to concatenate the input strings.
//   - `length` is the expected length of the concatenated string.
//   - if you are unsure about the length of the string to be concatenated, please pass 0 or a negative number.
func Concat(length int, str ...string) string {
	if len(str) == 0 {
		return ""
	}

	sb := strings.Builder{}
	if length <= 0 {
		sb.Grow(len(str[0]) * len(str))
	} else {
		sb.Grow(length)
	}

	for _, s := range str {
		sb.WriteString(s)
	}
	return sb.String()
}

// Ellipsis truncates a string to a specified length and appends an ellipsis.
func Ellipsis(str string, length int) string {
	str = strings.TrimSpace(str)

	if length <= 0 {
		return ""
	}

	runes := []rune(str)

	if len(runes) <= length {
		return str
	}

	return string(runes[:length]) + "..."
}

// Shuffle the order of characters of given string.
func Shuffle(str string) string {
	runes := []rune(str)

	for i := len(runes) - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// Rotate rotates the string by the specified number of characters.
func Rotate(str string, shift int) string {
	if shift == 0 {
		return str
	}

	runes := []rune(str)
	length := len(runes)
	if length == 0 {
		return str
	}

	shift = shift % length

	if shift < 0 {
		shift = length + shift
	}

	var sb strings.Builder
	sb.Grow(length)

	sb.WriteString(string(runes[length-shift:]))
	sb.WriteString(string(runes[:length-shift]))

	return sb.String()
}

// TemplateReplace replaces the placeholders in the template string with the corresponding values in the data map.
// The placeholders are enclosed in curly braces, e.g. {key}.
// for example, the template string is "Hello, {name}!", and the data map is {"name": "world"},
// the result will be "Hello, world!".
func TemplateReplace(template string, data map[string]string) string {
	re := regexp.MustCompile(`\{(\w+)\}`)

	result := re.ReplaceAllStringFunc(template, func(s string) string {
		key := strings.Trim(s, "{}")
		if val, ok := data[key]; ok {
			return val
		}

		return s
	})

	result = strings.ReplaceAll(result, "{{", "{")
	result = strings.ReplaceAll(result, "}}", "}")

	return result
}

// RegexMatchAllGroups Matches all subgroups in a string using a regular expression and returns the result.
func RegexMatchAllGroups(pattern, str string) [][]string {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(str, -1)
	return matches
}
