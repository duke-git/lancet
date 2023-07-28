package strutil

import (
	"reflect"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestCamelCase(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCamelCase")

	cases := map[string]string{
		"":                         "",
		"foobar":                   "foobar",
		"&FOO:BAR$BAZ":             "fooBarBaz",
		"fooBar":                   "fooBar",
		"FOObar":                   "foObar",
		"$foo%":                    "foo",
		"   $#$Foo   22    bar   ": "foo22Bar",
		"Foo-#1😄$_%^&*(1bar":       "foo11Bar",
	}

	for k, v := range cases {
		assert.Equal(v, CamelCase(k))
	}
}

func TestCapitalize(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCapitalize")

	cases := map[string]string{
		"":        "",
		"Foo":     "Foo",
		"_foo":    "_foo",
		"foobar":  "Foobar",
		"fooBar":  "Foobar",
		"foo Bar": "Foo bar",
		"foo-bar": "Foo-bar",
		"$foo%":   "$foo%",
	}

	for k, v := range cases {
		assert.Equal(v, Capitalize(k))
	}
}

func TestKebabCase(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestKebabCase")

	cases := map[string]string{
		"":                         "",
		"foo-bar":                  "foo-bar",
		"--Foo---Bar-":             "foo-bar",
		"Foo Bar-":                 "foo-bar",
		"foo_Bar":                  "foo-bar",
		"fooBar":                   "foo-bar",
		"FOOBAR":                   "foobar",
		"FOO_BAR":                  "foo-bar",
		"__FOO_BAR__":              "foo-bar",
		"$foo@Bar":                 "foo-bar",
		"   $#$Foo   22    bar   ": "foo-22-bar",
		"Foo-#1😄$_%^&*(1bar":       "foo-1-1-bar",
	}

	for k, v := range cases {
		assert.Equal(v, KebabCase(k))
	}
}

func TestUpperKebabCase(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUpperKebabCase")

	cases := map[string]string{
		"":                         "",
		"foo-bar":                  "FOO-BAR",
		"--Foo---Bar-":             "FOO-BAR",
		"Foo Bar-":                 "FOO-BAR",
		"foo_Bar":                  "FOO-BAR",
		"fooBar":                   "FOO-BAR",
		"FOOBAR":                   "FOOBAR",
		"FOO_BAR":                  "FOO-BAR",
		"__FOO_BAR__":              "FOO-BAR",
		"$foo@Bar":                 "FOO-BAR",
		"   $#$Foo   22    bar   ": "FOO-22-BAR",
		"Foo-#1😄$_%^&*(1bar":       "FOO-1-1-BAR",
	}

	for k, v := range cases {
		assert.Equal(v, UpperKebabCase(k))
	}
}

func TestSnakeCase(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSnakeCase")

	cases := map[string]string{
		"":                         "",
		"foo-bar":                  "foo_bar",
		"--Foo---Bar-":             "foo_bar",
		"Foo Bar-":                 "foo_bar",
		"foo_Bar":                  "foo_bar",
		"fooBar":                   "foo_bar",
		"FOOBAR":                   "foobar",
		"FOO_BAR":                  "foo_bar",
		"__FOO_BAR__":              "foo_bar",
		"$foo@Bar":                 "foo_bar",
		"   $#$Foo   22    bar   ": "foo_22_bar",
		"Foo-#1😄$_%^&*(1bar":       "foo_1_1_bar",
	}

	for k, v := range cases {
		assert.Equal(v, SnakeCase(k))
	}
}

func TestUpperSnakeCase(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUpperSnakeCase")

	cases := map[string]string{
		"":                         "",
		"foo-bar":                  "FOO_BAR",
		"--Foo---Bar-":             "FOO_BAR",
		"Foo Bar-":                 "FOO_BAR",
		"foo_Bar":                  "FOO_BAR",
		"fooBar":                   "FOO_BAR",
		"FOOBAR":                   "FOOBAR",
		"FOO_BAR":                  "FOO_BAR",
		"__FOO_BAR__":              "FOO_BAR",
		"$foo@Bar":                 "FOO_BAR",
		"   $#$Foo   22    bar   ": "FOO_22_BAR",
		"Foo-#1😄$_%^&*(1bar":       "FOO_1_1_BAR",
	}

	for k, v := range cases {
		assert.Equal(v, UpperSnakeCase(k))
	}
}

func TestUpperFirst(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLowerFirst")

	cases := map[string]string{
		"":     "",
		"foo":  "Foo",
		"bAR":  "BAR",
		"FOo":  "FOo",
		"fOo大": "FOo大",
	}

	for k, v := range cases {
		assert.Equal(v, UpperFirst(k))
	}
}

func TestLowerFirst(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLowerFirst")

	cases := map[string]string{
		"":     "",
		"foo":  "foo",
		"bAR":  "bAR",
		"FOo":  "fOo",
		"fOo大": "fOo大",
	}

	for k, v := range cases {
		assert.Equal(v, LowerFirst(k))
	}
}

func TestPad(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPad")

	assert.Equal("a ", Pad("a", 2, ""))
	assert.Equal("a", Pad("a", 1, "b"))
	assert.Equal("ab", Pad("a", 2, "b"))
	assert.Equal("mabcdm", Pad("abcd", 6, "m"))
}

func TestPadEnd(t *testing.T) {
	assert := internal.NewAssert(t, "TestPadEnd")

	assert.Equal("a ", PadEnd("a", 2, " "))
	assert.Equal("a", PadEnd("a", 1, "b"))
	assert.Equal("ab", PadEnd("a", 2, "b"))
	assert.Equal("abcdmn", PadEnd("abcd", 6, "mno"))
	assert.Equal("abcdmm", PadEnd("abcd", 6, "m"))
	assert.Equal("abcaba", PadEnd("abc", 6, "ab"))

	assert.NotEqual("ba", PadEnd("a", 2, "b"))
}

func TestPadStart(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPadStart")

	assert.Equal("a", PadStart("a", 1, "b"))
	assert.Equal("ba", PadStart("a", 2, "b"))
	assert.Equal("mnabcd", PadStart("abcd", 6, "mno"))
	assert.Equal("mmabcd", PadStart("abcd", 6, "m"))
	assert.Equal("abaabc", PadStart("abc", 6, "ab"))

	assert.NotEqual("ab", PadStart("a", 2, "b"))
}

func TestBefore(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBefore")

	assert.Equal("lancet", Before("lancet", ""))
	assert.Equal("", Before("lancet", "lancet"))
	assert.Equal("github.com", Before("github.com/test/lancet", "/"))
	assert.Equal("github.com/", Before("github.com/test/lancet", "test"))
}

func TestBeforeLast(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBeforeLast")

	assert.Equal("lancet", BeforeLast("lancet", ""))
	assert.Equal("github.com/test", BeforeLast("github.com/test/lancet", "/"))
	assert.Equal("github.com/test/", BeforeLast("github.com/test/test/lancet", "test"))

	assert.NotEqual("github.com/", BeforeLast("github.com/test/test/lancet", "test"))
}

func TestAfter(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAfter")

	assert.Equal("lancet", After("lancet", ""))
	assert.Equal("", After("lancet", "lancet"))
	assert.Equal("test/lancet", After("github.com/test/lancet", "/"))
	assert.Equal("/lancet", After("github.com/test/lancet", "test"))
}

func TestAfterLast(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAfterLast")

	assert.Equal("lancet", AfterLast("lancet", ""))
	assert.Equal("lancet", AfterLast("github.com/test/lancet", "/"))
	assert.Equal("/lancet", AfterLast("github.com/test/lancet", "test"))
	assert.Equal("/lancet", AfterLast("github.com/test/test/lancet", "test"))

	assert.NotEqual("/test/lancet", AfterLast("github.com/test/test/lancet", "test"))
}

func TestIsString(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsString")

	assert.Equal(true, IsString("lancet"))
	assert.Equal(true, IsString(""))
	assert.Equal(false, IsString(1))
	assert.Equal(false, IsString(true))
	assert.Equal(false, IsString([]string{}))
}

func TestReverse(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestReverse")

	assert.Equal("cba", Reverse("abc"))
	assert.Equal("54321", Reverse("12345"))
}

func TestWrap(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestWrap")

	assert.Equal("ab", Wrap("ab", ""))
	assert.Equal("", Wrap("", "*"))
	assert.Equal("*ab*", Wrap("ab", "*"))
	assert.Equal("\"ab\"", Wrap("ab", "\""))
	assert.Equal("'ab'", Wrap("ab", "'"))
}

func TestUnwrap(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUnwrap")

	assert.Equal("", Unwrap("", "*"))
	assert.Equal("ab", Unwrap("ab", ""))
	assert.Equal("ab", Unwrap("ab", "*"))
	assert.Equal("*ab*", Unwrap("**ab**", "*"))
	assert.Equal("ab", Unwrap("**ab**", "**"))
	assert.Equal("ab", Unwrap("\"ab\"", "\""))
	assert.Equal("*ab", Unwrap("*ab", "*"))
	assert.Equal("ab*", Unwrap("ab*", "*"))
	assert.Equal("*", Unwrap("***", "*"))

	assert.Equal("", Unwrap("**", "*"))
	assert.Equal("***", Unwrap("***", "**"))
	assert.Equal("**", Unwrap("**", "**"))
}

func TestSplitEx(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSplitEx")

	assert.Equal([]string{}, SplitEx(" a b c ", "", true))

	assert.Equal([]string{"", "a", "b", "c", ""}, SplitEx(" a b c ", " ", false))
	assert.Equal([]string{"a", "b", "c"}, SplitEx(" a b c ", " ", true))

	assert.Equal([]string{"a", "b", "c", ""}, SplitEx("a = b = c = ", " = ", false))
	assert.Equal([]string{"a", "b", "c"}, SplitEx("a = b = c = ", " = ", true))
}

func TestSubstring(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSubstring")

	assert.Equal("bcd", Substring("abcde", 1, 3))
	assert.Equal("bcde", Substring("abcde", 1, 5))
	assert.Equal("e", Substring("abcde", -1, 3))
	assert.Equal("de", Substring("abcde", -2, 2))
	assert.Equal("de", Substring("abcde", -2, 3))
	assert.Equal("你好", Substring("你好，欢迎你", 0, 2))
}

func TestSplitWords(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSplitWords")

	cases := map[string][]string{
		"a word":            {"a", "word"},
		"I'am a programmer": {"I'am", "a", "programmer"},
		"a -b-c' 'd'e":      {"a", "b-c'", "d'e"},
		"你好，我是一名码农":         nil,
		"こんにちは，私はプログラマーです": nil,
	}

	for k, v := range cases {
		assert.Equal(v, SplitWords(k))
	}
}

func TestWordCount(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSplitWords")

	cases := map[string]int{
		"a word":            2, //   {"a", "word"},
		"I'am a programmer": 3, //   {"I'am", "a", "programmer"},
		"a -b-c' 'd'e":      3, // {"a", "b-c'", "d'e"},
		"你好，我是一名码农":         0, // nil,
		"こんにちは，私はプログラマーです": 0, // nil,
	}

	for k, v := range cases {
		assert.Equal(v, WordCount(k))
	}
}

func TestRemoveNonPrintable(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRemoveNonPrintable")

	assert.Equal("hello world", RemoveNonPrintable("hello\u00a0 \u200bworld\n"))
	assert.Equal("你好😄", RemoveNonPrintable("你好😄"))
}

func TestStringToBytes(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestStringToBytes")

	str := "abc"
	bytes := StringToBytes(str)

	assert.Equal(reflect.DeepEqual(bytes, []byte{'a', 'b', 'c'}), true)
}

func TestBytesToString(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBytesToString")

	bytes := []byte{'a', 'b', 'c'}
	str := BytesToString(bytes)

	assert.Equal(str == "abc", true)
}

func TestIsBlank(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsBlank")
	assert.Equal(IsBlank(""), true)
	assert.Equal(IsBlank("\t\v\f\n"), true)

	assert.Equal(IsBlank(" 中文"), false)
}

func TestHasPrefixAny(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestHasPrefixAny")
	str := "foo bar"
	prefixes := []string{"fo", "xyz", "hello"}
	notMatches := []string{"oom", "world"}

	assert.Equal(HasPrefixAny(str, prefixes), true)
	assert.Equal(HasPrefixAny(str, notMatches), false)
}

func TestHasSuffixAny(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestHasSuffixAny")
	str := "foo bar"
	suffixes := []string{"bar", "xyz", "hello"}
	notMatches := []string{"oom", "world"}

	assert.Equal(HasSuffixAny(str, suffixes), true)
	assert.Equal(HasSuffixAny(str, notMatches), false)
}

func TestIndexOffset(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIndexOffset")
	str := "foo bar hello world"

	assert.Equal(IndexOffset(str, "o", 5), 12)
	assert.Equal(IndexOffset(str, "o", 0), 1)
	assert.Equal(IndexOffset(str, "d", len(str)-1), len(str)-1)
	assert.Equal(IndexOffset(str, "d", len(str)), -1)
	assert.Equal(IndexOffset(str, "f", -1), -1)
}

func TestReplaceWithMap(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestReplaceWithMap")

	str := "ac ab ab ac"
	replaces := map[string]string{
		"a": "1",
		"b": "2",
	}

	assert.Equal(str, "ac ab ab ac")
	assert.Equal(ReplaceWithMap(str, replaces), "1c 12 12 1c")
}

func TestTrim(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTrim")

	str1 := "$ ab	cd $ "

	assert.Equal("$ ab	cd $", Trim(str1))
	assert.Equal("ab	cd", Trim(str1, "$"))
	assert.Equal("abcd", Trim("\nabcd"))
}

func TestSplitAndTrim(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTrim")

	str := " a,b, c,d,$1 "

	result1 := SplitAndTrim(str, ",")
	result2 := SplitAndTrim(str, ",", "$")

	assert.Equal([]string{"a", "b", "c", "d", "$1"}, result1)
	assert.Equal([]string{"a", "b", "c", "d", "1"}, result2)
}

func TestHideString(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTrim")

	str := "13242658976"

	assert.Equal("13242658976", HideString(str, 0, -1, "*"))
	assert.Equal("13242658976", HideString(str, 0, 0, "*"))
	assert.Equal("****2658976", HideString(str, 0, 4, "*"))

	assert.Equal("13242658976", HideString(str, 3, 3, "*"))
	assert.Equal("132*2658976", HideString(str, 3, 4, "*"))
	assert.Equal("132****8976", HideString(str, 3, 7, "*"))
	assert.Equal("1324265****", HideString(str, 7, 11, "*"))

	assert.Equal("1324265****", HideString(str, 7, 100, "*"))
	assert.Equal("13242658976", HideString(str, 100, 100, "*"))
}

func TestContainsAll(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestContainsAll")

	assert.Equal(true, ContainsAll("hello world", []string{"hello", "world"}))
	assert.Equal(true, ContainsAll("hello world", []string{""}))
	assert.Equal(false, ContainsAll("hello world", []string{"hello", "abc"}))
}

func TestContainsAny(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestContainsAny")

	assert.Equal(true, ContainsAny("hello world", []string{"hello", "world"}))
	assert.Equal(true, ContainsAny("hello world", []string{"hello", "abc"}))
	assert.Equal(false, ContainsAny("hello world", []string{"123", "abc"}))
}

func TestRemoveWhiteSpace(t *testing.T) {
	assert := internal.NewAssert(t, "TestRemoveWhiteSpace")

	str := " hello   \r\n	\t   world"

	assert.Equal("", RemoveWhiteSpace("", true))
	assert.Equal("helloworld", RemoveWhiteSpace(str, true))
	assert.Equal("hello world", RemoveWhiteSpace(str, false))
}
