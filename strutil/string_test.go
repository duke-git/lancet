package strutil

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestCamelCase(t *testing.T) {
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
	assert := internal.NewAssert(t, "TestSnakeCase")

	cases := map[string]string{
		"Foo Bar-":                 "foo_bar",
		"foo_Bar":                  "foo_bar",
		"fooBar":                   "foo_bar",
		"__FOO_BAR__":              "f_o_o_b_a_r",
		"$foo@Bar":                 "foo_bar",
		"   $#$Foo   22    bar   ": "foo_22_bar",
		"aBbc-s$@a&%_B.B^C":        "a_bbc_s_a_b_b_c",
	}

	for k, v := range cases {
		assert.Equal(v, SnakeCase(k))
	}

	assert.Equal("", SnakeCase(""))
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
