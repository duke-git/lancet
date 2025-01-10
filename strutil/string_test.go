package strutil

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestCamelCase(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCamelCase")

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"foobar", "foobar"},
		{"&FOO:BAR$BAZ", "fooBarBaz"},
		{"fooBar", "fooBar"},
		{"FOObar", "foObar"},
		{"$foo%", "foo"},
		{"   $#$Foo   22    bar   ", "foo22Bar"},
		{"Foo-#1😄$_%^&*(1bar", "foo11Bar"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, CamelCase(tt.input))
	}
}

func TestCapitalize(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCapitalize")

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"Foo", "Foo"},
		{"_foo", "_foo"},
		{"foobar", "Foobar"},
		{"fooBar", "Foobar"},
		{"foo Bar", "Foo bar"},
		{"foo-bar", "Foo-bar"},
		{"$foo%", "$foo%"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, Capitalize(tt.input))
	}
}

func TestKebabCase(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestKebabCase")

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"foo-bar", "foo-bar"},
		{"--Foo---Bar-", "foo-bar"},
		{"Foo Bar-", "foo-bar"},
		{"foo_Bar", "foo-bar"},
		{"fooBar", "foo-bar"},
		{"FOOBAR", "foobar"},
		{"FOO_BAR", "foo-bar"},
		{"__FOO_BAR__", "foo-bar"},
		{"$foo@Bar", "foo-bar"},
		{"   $#$Foo   22    bar   ", "foo-22-bar"},
		{"Foo-#1😄$_%^&*(1bar", "foo-1-1-bar"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, KebabCase(tt.input))
	}
}

func TestUpperKebabCase(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUpperKebabCase")

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"foo-bar", "FOO-BAR"},
		{"--Foo---Bar-", "FOO-BAR"},
		{"Foo Bar-", "FOO-BAR"},
		{"foo_Bar", "FOO-BAR"},
		{"fooBar", "FOO-BAR"},
		{"FOOBAR", "FOOBAR"},
		{"FOO_BAR", "FOO-BAR"},
		{"__FOO_BAR__", "FOO-BAR"},
		{"$foo@Bar", "FOO-BAR"},
		{"   $#$Foo   22    bar   ", "FOO-22-BAR"},
		{"Foo-#1😄$_%^&*(1bar", "FOO-1-1-BAR"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, UpperKebabCase(tt.input))
	}
}

func TestSnakeCase(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSnakeCase")

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"foo-bar", "foo_bar"},
		{"--Foo---Bar-", "foo_bar"},
		{"Foo Bar-", "foo_bar"},
		{"foo_Bar", "foo_bar"},
		{"fooBar", "foo_bar"},
		{"FOOBAR", "foobar"},
		{"FOO_BAR", "foo_bar"},
		{"__FOO_BAR__", "foo_bar"},
		{"$foo@Bar", "foo_bar"},
		{"   $#$Foo   22    bar   ", "foo_22_bar"},
		{"Foo-#1😄$_%^&*(1bar", "foo_1_1_bar"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, SnakeCase(tt.input))
	}
}

func TestUpperSnakeCase(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUpperSnakeCase")

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"foo-bar", "FOO_BAR"},
		{"--Foo---Bar-", "FOO_BAR"},
		{"Foo Bar-", "FOO_BAR"},
		{"foo_Bar", "FOO_BAR"},
		{"fooBar", "FOO_BAR"},
		{"FOOBAR", "FOOBAR"},
		{"FOO_BAR", "FOO_BAR"},
		{"__FOO_BAR__", "FOO_BAR"},
		{"$foo@Bar", "FOO_BAR"},
		{"   $#$Foo   22    bar   ", "FOO_22_BAR"},
		{"Foo-#1😄$_%^&*(1bar", "FOO_1_1_BAR"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, UpperSnakeCase(tt.input))
	}
}

func TestUpperFirst(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLowerFirst")

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"foo", "Foo"},
		{"bAR", "BAR"},
		{"FOo", "FOo"},
		{"fOo大", "FOo大"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, UpperFirst(tt.input))
	}
}

func TestLowerFirst(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLowerFirst")

	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"foo", "foo"},
		{"bAR", "bAR"},
		{"FOo", "fOo"},
		{"fOo大", "fOo大"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, LowerFirst(tt.input))
	}
}

func TestPad(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPad")

	tests := []struct {
		input    string
		padSize  int
		padChar  string
		expected string
	}{
		{"", 0, "", ""},
		{"a", 0, "", "a"},
		{"a", 1, "", "a"},
		{"a", 2, "", "a "},
		{"a", 1, "b", "a"},
		{"a", 2, "b", "ab"},
		{"abcd", 6, "m", "mabcdm"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, Pad(tt.input, tt.padSize, tt.padChar))
	}
}

func TestPadEnd(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPadEnd")

	tests := []struct {
		input    string
		padSize  int
		padChar  string
		expected string
	}{
		{"a", 2, " ", "a "},
		{"a", 1, "b", "a"},
		{"a", 2, "b", "ab"},
		{"abcd", 6, "mno", "abcdmn"},
		{"abcd", 6, "m", "abcdmm"},
		{"abcd", 6, "ab", "abcdab"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, PadEnd(tt.input, tt.padSize, tt.padChar))
	}
}

func TestPadStart(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPadStart")

	tests := []struct {
		input    string
		padSize  int
		padChar  string
		expected string
	}{
		{"a", 1, "b", "a"},
		{"a", 2, "b", "ba"},
		{"abcd", 6, "mno", "mnabcd"},
		{"abcd", 6, "m", "mmabcd"},
		{"abc", 6, "ab", "abaabc"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, PadStart(tt.input, tt.padSize, tt.padChar))
	}
}

func TestBefore(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBefore")

	tests := []struct {
		input    string
		char     string
		expected string
	}{
		{"lancet", "", "lancet"},
		{"lancet", "lancet", ""},
		{"lancet", "abcdef", "lancet"},
		{"github.com/test/lancet", "/", "github.com"},
		{"github.com/test/lancet", "test", "github.com/"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, Before(tt.input, tt.char))
	}
}

func TestBeforeLast(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBeforeLast")

	assert.Equal("lancet", BeforeLast("lancet", ""))
	assert.Equal("lancet", BeforeLast("lancet", "abcdef"))
	assert.Equal("github.com/test", BeforeLast("github.com/test/lancet", "/"))
	assert.Equal("github.com/test/", BeforeLast("github.com/test/test/lancet", "test"))
}

func TestAfter(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAfter")

	assert.Equal("lancet", After("lancet", ""))
	assert.Equal("", After("lancet", "lancet"))
	assert.Equal("test/lancet", After("github.com/test/lancet", "/"))
	assert.Equal("/lancet", After("github.com/test/lancet", "test"))
	assert.Equal("lancet", After("lancet", "abcdef"))
}

func TestAfterLast(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAfterLast")

	assert.Equal("lancet", AfterLast("lancet", ""))
	assert.Equal("lancet", AfterLast("github.com/test/lancet", "/"))
	assert.Equal("/lancet", AfterLast("github.com/test/lancet", "test"))
	assert.Equal("/lancet", AfterLast("github.com/test/test/lancet", "test"))
	assert.Equal("lancet", AfterLast("lancet", "abcdef"))
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

	bytes := StringToBytes("abc")
	assert.Equal(bytes, []byte{'a', 'b', 'c'})
}

func TestBytesToString(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBytesToString")

	str := BytesToString([]byte{'a', 'b', 'c'})
	assert.Equal("abc", str)
}

func TestIsBlank(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsBlank")

	assert.Equal(IsBlank(""), true)
	assert.Equal(IsBlank("\t\v\f\n"), true)

	assert.Equal(IsBlank(" 中文"), false)
}

func TestIsNotBlank(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsBlank")

	assert.Equal(IsNotBlank(""), false)
	assert.Equal(IsNotBlank("			"), false)
	assert.Equal(IsNotBlank("\t\v\f\n"), false)

	assert.Equal(IsNotBlank(" 中文"), true)
	assert.Equal(IsNotBlank(" 	world	"), true)
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

	tests := []struct {
		input        string
		start        int
		end          int
		replacedChar string
		expected     string
	}{
		{"13242658976", 0, -1, "*", "13242658976"},
		{"13242658976", 0, 0, "*", "13242658976"},
		{"13242658976", 0, 4, "*", "****2658976"},
		{"13242658976", 3, 3, "*", "13242658976"},
		{"13242658976", 3, 4, "*", "132*2658976"},
		{"13242658976", 3, 7, "*", "132****8976"},
		{"13242658976", 3, 11, "*", "132********"},
		{"13242658976", 7, 100, "*", "1324265****"},
		{"13242658976", 100, 100, "*", "13242658976"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, HideString(tt.input, tt.start, tt.end, tt.replacedChar))
	}
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
	t.Parallel()
	assert := internal.NewAssert(t, "TestRemoveWhiteSpace")

	str := " hello   \r\n	\t   world"

	assert.Equal("", RemoveWhiteSpace("", true))
	assert.Equal("helloworld", RemoveWhiteSpace(str, true))
	assert.Equal("hello world", RemoveWhiteSpace(str, false))
}

func TestSubInBetween(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestSubInBetween")

	str := "abcde"

	assert.Equal("", SubInBetween(str, "", ""))
	assert.Equal("ab", SubInBetween(str, "", "c"))
	assert.Equal("bc", SubInBetween(str, "a", "d"))
	assert.Equal("", SubInBetween(str, "a", ""))
	assert.Equal("", SubInBetween(str, "a", "f"))
}

func TestHammingDistance(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "HammingDistance")

	hd := func(a, b string) int {
		c, _ := HammingDistance(a, b)
		return c
	}

	tests := []struct {
		strA            string
		strB            string
		hammingDistance int
	}{
		{" ", " ", 0},
		{" ", "c", 1},
		{"a", "d", 1},
		{"a", " ", 1},
		{"a", "f", 1},

		{"", "", 0},
		{"abc", "ab", -1},
		{"abc", "def", 3},
		{"kitten", "sitting", -1},
		{"ö", "ü", 1},
		{"日本語", "日本語", 0},
		{"日本語", "語日本", 3},
	}

	for _, tt := range tests {
		assert.Equal(tt.hammingDistance, hd(tt.strA, tt.strB))
	}
}

func TestConcat(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestConcat")

	assert.Equal("", Concat(0))
	assert.Equal("a", Concat(1, "a"))
	assert.Equal("ab", Concat(2, "a", "b"))
	assert.Equal("abc", Concat(3, "a", "b", "c"))
	assert.Equal("abc", Concat(3, "a", "", "b", "c", ""))
	assert.Equal("你好，世界！", Concat(0, "你好", "，", "", "世界！", ""))
	assert.Equal("Hello World!", Concat(0, "Hello", " Wo", "r", "ld!", ""))
}

func TestEllipsis(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEllipsis")

	tests := []struct {
		input    string
		length   int
		expected string
	}{
		{"", 0, ""},
		{"hello world", 0, ""},
		{"hello world", -1, ""},
		{"hello world", 5, "hello..."},
		{"hello world", 11, "hello world"},
		{"你好，世界!", 2, "你好..."},
		{"😀😃😄😁😆", 3, "😀😃😄..."},
		{"This is a test.", 10, "This is a ..."},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, Ellipsis(tt.input, tt.length))
	}
}

func TestShuffle(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestShuffle")

	assert.Equal("", Shuffle(""))
	assert.Equal("a", Shuffle("a"))

	str := "hello"
	shuffledStr := Shuffle(str)
	assert.Equal(5, len(shuffledStr))
}

func TestRotate(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRotate")

	tests := []struct {
		input    string
		shift    int
		expected string
	}{
		{"", 1, ""},
		{"a", 0, "a"},
		{"a", 1, "a"},
		{"a", -1, "a"},

		{"Hello", -2, "lloHe"},
		{"Hello", 1, "oHell"},
		{"Hello, world!", 3, "ld!Hello, wor"},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, Rotate(tt.input, tt.shift))
	}
}

func TestTemplateReplace(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTemplateReplace")

	t.Run("basic", func(t *testing.T) {
		template := `Hello, my name is {name}, I'm {age} years old.`
		data := map[string]string{
			"name": "Bob",
			"age":  "20",
		}

		expected := `Hello, my name is Bob, I'm 20 years old.`
		result := TemplateReplace(template, data)

		assert.Equal(expected, result)
	})

	t.Run("not found", func(t *testing.T) {
		template := `Hello, my name is {name}, I'm {age} years old.`
		data := map[string]string{
			"name": "Bob",
		}

		expected := `Hello, my name is Bob, I'm {age} years old.`
		result := TemplateReplace(template, data)

		assert.Equal(expected, result)
	})

	t.Run("brackets", func(t *testing.T) {
		template := `Hello, my name is {name}, I'm {{age}} years old.`
		data := map[string]string{
			"name": "Bob",
			"age":  "20",
		}

		expected := `Hello, my name is Bob, I'm {20} years old.`
		result := TemplateReplace(template, data)

		assert.Equal(expected, result)
	})
}

func TestRegexMatchAllGroups(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRegexMatchAllGroups")

	tests := []struct {
		pattern  string
		str      string
		expected [][]string
	}{
		{
			pattern:  `(\w+\.+\w+)@(\w+)\.(\w+)`,
			str:      "Emails: john.doe@example.com and jane.doe@example.com",
			expected: [][]string{{"john.doe@example.com", "john.doe", "example", "com"}, {"jane.doe@example.com", "jane.doe", "example", "com"}},
		},
		{
			pattern:  `(\d+)`,
			str:      "No numbers here!",
			expected: nil,
		},
		{
			pattern:  `(\d{3})-(\d{2})-(\d{4})`,
			str:      "My number is 123-45-6789",
			expected: [][]string{{"123-45-6789", "123", "45", "6789"}},
		},
		{
			pattern:  `(\w+)\s(\d+)`,
			str:      "Item A 123, Item B 456",
			expected: [][]string{{"A 123", "A", "123"}, {"B 456", "B", "456"}},
		},
		{
			pattern:  `(\d{2})-(\d{2})-(\d{4})`,
			str:      "Dates: 01-01-2020, 12-31-1999",
			expected: [][]string{{"01-01-2020", "01", "01", "2020"}, {"12-31-1999", "12", "31", "1999"}},
		},
	}

	for _, tt := range tests {
		result := RegexMatchAllGroups(tt.pattern, tt.str)
		assert.Equal(tt.expected, result)
	}
}

func TestExtractContent(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestExtractContent")

	tests := []struct {
		name     string
		input    string
		start    string
		end      string
		expected []string
	}{
		{
			name:     "Extract content between <tag> and </tag>",
			input:    "This is <tag>content1</tag> and <tag>content2</tag> and <tag>content3</tag>",
			start:    "<tag>",
			end:      "</tag>",
			expected: []string{"content1", "content2", "content3"},
		},
		{
			name:     "No tags in the string",
			input:    "This string has no tags",
			start:    "<tag>",
			end:      "</tag>",
			expected: []string{},
		},
		{
			name:     "Single tag pair",
			input:    "<tag>onlyContent</tag>",
			start:    "<tag>",
			end:      "</tag>",
			expected: []string{"onlyContent"},
		},
		{
			name:     "Tags without end tag",
			input:    "This <tag>content without end tag",
			start:    "<tag>",
			end:      "</tag>",
			expected: []string{},
		},
		{
			name:     "Tags with nested content",
			input:    "<tag>content <nested>inner</nested> end</tag>",
			start:    "<tag>",
			end:      "</tag>",
			expected: []string{"content <nested>inner</nested> end"},
		},
		{
			name:     "Edge case with empty string",
			input:    "",
			start:    "<tag>",
			end:      "</tag>",
			expected: []string{},
		},
		{
			name:     "Edge case with no start tag",
			input:    "content without start tag",
			start:    "<tag>",
			end:      "</tag>",
			expected: []string{},
		},
		{
			name:     "Edge case with no end tag",
			input:    "<tag>content without end tag",
			start:    "<tag>",
			end:      "</tag>",
			expected: []string{},
		},
		{
			name:     "Multiple consecutive tags",
			input:    "<tag>content1</tag><tag>content2</tag><tag>content3</tag>",
			start:    "<tag>",
			end:      "</tag>",
			expected: []string{"content1", "content2", "content3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractContent(tt.input, tt.start, tt.end)
			assert.Equal(tt.expected, result)
		})
	}
}

func TestFindAllOccurrences(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFindAllOccurrences")

	var empty []int
	tests := []struct {
		input    string
		substr   string
		expected []int
	}{
		{"", "", empty},
		{"", "a", empty},
		{"a", "", []int{0}},
		{"a", "a", []int{0}},
		{"aa", "a", []int{0, 1}},
		{"ababab", "ab", []int{0, 2, 4}},
		{"ababab", "ba", []int{1, 3}},
		{"ababab", "c", empty},
		{"ababab", "ababab", []int{0}},
		{"ababab", "abababc", empty},
		{"ababab", "abababa", empty},
	}

	for _, tt := range tests {
		assert.Equal(tt.expected, FindAllOccurrences(tt.input, tt.substr))
	}
}
