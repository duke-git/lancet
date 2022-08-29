package strutil

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestCamelCase(t *testing.T) {
	assert := internal.NewAssert(t, "TestCamelCase")

	assert.Equal("fooBar", CamelCase("foo_bar"))
	assert.Equal("fooBar", CamelCase("Foo-Bar"))
	assert.Equal("fooBar", CamelCase("Foo&bar"))
	assert.Equal("fooBar", CamelCase("foo bar"))

	assert.NotEqual("FooBar", CamelCase("foo_bar"))
}

func TestCapitalize(t *testing.T) {
	assert := internal.NewAssert(t, "TestCapitalize")

	assert.Equal("Foo", Capitalize("foo"))
	assert.Equal("Foo", Capitalize("Foo"))
	assert.Equal("Foo", Capitalize("Foo"))

	assert.NotEqual("foo", Capitalize("Foo"))
}

func TestKebabCase(t *testing.T) {
	assert := internal.NewAssert(t, "TestKebabCase")

	assert.Equal("foo-bar", KebabCase("Foo Bar-"))
	assert.Equal("foo-bar", KebabCase("foo_Bar"))
	assert.Equal("foo-bar", KebabCase("fooBar"))
	assert.Equal("f-o-o-b-a-r", KebabCase("__FOO_BAR__"))

	assert.NotEqual("foo_bar", KebabCase("fooBar"))
}

func TestSnakeCase(t *testing.T) {
	assert := internal.NewAssert(t, "TestSnakeCase")

	assert.Equal("foo_bar", SnakeCase("Foo Bar-"))
	assert.Equal("foo_bar", SnakeCase("foo_Bar"))
	assert.Equal("foo_bar", SnakeCase("fooBar"))
	assert.Equal("f_o_o_b_a_r", SnakeCase("__FOO_BAR__"))
	assert.Equal("a_bbc_s_a_b_b_c", SnakeCase("aBbc-s$@a&%_B.B^C"))

	assert.NotEqual("foo-bar", SnakeCase("foo_Bar"))
}

func TestUpperFirst(t *testing.T) {
	assert := internal.NewAssert(t, "TestLowerFirst")

	assert.Equal("Foo", UpperFirst("foo"))
	assert.Equal("BAR", UpperFirst("bAR"))
	assert.Equal("FOo", UpperFirst("FOo"))
	assert.Equal("FOo大", UpperFirst("fOo大"))

	assert.NotEqual("Bar", UpperFirst("BAR"))
}

func TestLowerFirst(t *testing.T) {
	assert := internal.NewAssert(t, "TestLowerFirst")

	assert.Equal("foo", LowerFirst("foo"))
	assert.Equal("bAR", LowerFirst("BAR"))
	assert.Equal("fOo", LowerFirst("FOo"))
	assert.Equal("fOo大", LowerFirst("FOo大"))

	assert.NotEqual("Bar", LowerFirst("BAR"))
}

func TestPadEnd(t *testing.T) {
	assert := internal.NewAssert(t, "TestPadEnd")

	assert.Equal("a", PadEnd("a", 1, "b"))
	assert.Equal("ab", PadEnd("a", 2, "b"))
	assert.Equal("abcdmn", PadEnd("abcd", 6, "mno"))
	assert.Equal("abcdmm", PadEnd("abcd", 6, "m"))
	assert.Equal("abcaba", PadEnd("abc", 6, "ab"))

	assert.NotEqual("ba", PadEnd("a", 2, "b"))
}

func TestPadStart(t *testing.T) {
	assert := internal.NewAssert(t, "TestPadStart")

	assert.Equal("a", PadStart("a", 1, "b"))
	assert.Equal("ba", PadStart("a", 2, "b"))
	assert.Equal("mnabcd", PadStart("abcd", 6, "mno"))
	assert.Equal("mmabcd", PadStart("abcd", 6, "m"))
	assert.Equal("abaabc", PadStart("abc", 6, "ab"))

	assert.NotEqual("ab", PadStart("a", 2, "b"))
}

func TestBefore(t *testing.T) {
	assert := internal.NewAssert(t, "TestBefore")

	assert.Equal("lancet", Before("lancet", ""))
	assert.Equal("github.com", Before("github.com/test/lancet", "/"))
	assert.Equal("github.com/", Before("github.com/test/lancet", "test"))
}

func TestBeforeLast(t *testing.T) {
	assert := internal.NewAssert(t, "TestBeforeLast")

	assert.Equal("lancet", BeforeLast("lancet", ""))
	assert.Equal("github.com/test", BeforeLast("github.com/test/lancet", "/"))
	assert.Equal("github.com/test/", BeforeLast("github.com/test/test/lancet", "test"))

	assert.NotEqual("github.com/", BeforeLast("github.com/test/test/lancet", "test"))
}

func TestAfter(t *testing.T) {
	assert := internal.NewAssert(t, "TestAfter")

	assert.Equal("lancet", After("lancet", ""))
	assert.Equal("test/lancet", After("github.com/test/lancet", "/"))
	assert.Equal("/lancet", After("github.com/test/lancet", "test"))
}

func TestAfterLast(t *testing.T) {
	assert := internal.NewAssert(t, "TestAfterLast")

	assert.Equal("lancet", AfterLast("lancet", ""))
	assert.Equal("lancet", AfterLast("github.com/test/lancet", "/"))
	assert.Equal("/lancet", AfterLast("github.com/test/lancet", "test"))
	assert.Equal("/lancet", AfterLast("github.com/test/test/lancet", "test"))

	assert.NotEqual("/test/lancet", AfterLast("github.com/test/test/lancet", "test"))
}

func TestIsString(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsString")

	assert.Equal(true, IsString("lancet"))
	assert.Equal(true, IsString(""))
	assert.Equal(false, IsString(1))
	assert.Equal(false, IsString(true))
	assert.Equal(false, IsString([]string{}))
}

func TestReverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverse")

	assert.Equal("cba", Reverse("abc"))
	assert.Equal("54321", Reverse("12345"))
}

func TestWrap(t *testing.T) {
	assert := internal.NewAssert(t, "TestWrap")

	assert.Equal("ab", Wrap("ab", ""))
	assert.Equal("", Wrap("", "*"))
	assert.Equal("*ab*", Wrap("ab", "*"))
	assert.Equal("\"ab\"", Wrap("ab", "\""))
	assert.Equal("'ab'", Wrap("ab", "'"))
}

func TestUnwrap(t *testing.T) {
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
	assert := internal.NewAssert(t, "TestSplitEx")

	assert.Equal([]string{}, SplitEx(" a b c ", "", true))

	assert.Equal([]string{"", "a", "b", "c", ""}, SplitEx(" a b c ", " ", false))
	assert.Equal([]string{"a", "b", "c"}, SplitEx(" a b c ", " ", true))

	assert.Equal([]string{" a", "b", "c", ""}, SplitEx(" a = b = c = ", " = ", false))
	assert.Equal([]string{" a", "b", "c"}, SplitEx(" a = b = c = ", " = ", true))
}
