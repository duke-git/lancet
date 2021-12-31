package strutil

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestCamelCase(t *testing.T) {
	camelCase(t, "foo_bar", "fooBar")
	camelCase(t, "Foo-Bar", "fooBar")
	camelCase(t, "Foo&bar", "fooBar")
	camelCase(t, "foo bar", "fooBar")
}

func camelCase(t *testing.T, test string, expected string) {
	res := CamelCase(test)
	if res != expected {
		internal.LogFailedTestInfo(t, "CamelCase", test, expected, res)
		t.FailNow()
	}
}

func TestCapitalize(t *testing.T) {
	capitalize(t, "foo", "Foo")
	capitalize(t, "fOO", "Foo")
	capitalize(t, "FOo", "Foo")
}

func capitalize(t *testing.T, test string, expected string) {
	res := Capitalize(test)
	if res != expected {
		internal.LogFailedTestInfo(t, "Capitalize", test, expected, res)
		t.FailNow()
	}
}

func TestKebabCase(t *testing.T) {
	kebabCase(t, "Foo Bar-", "foo-bar")
	kebabCase(t, "foo_Bar", "foo-bar")
	kebabCase(t, "fooBar", "foo-bar")
	kebabCase(t, "__FOO_BAR__", "f-o-o-b-a-r")
}

func kebabCase(t *testing.T, test string, expected string) {
	res := KebabCase(test)
	if res != expected {
		internal.LogFailedTestInfo(t, "KebabCase", test, expected, res)
		t.FailNow()
	}
}

func TestSnakeCase(t *testing.T) {
	snakeCase(t, "Foo Bar-", "foo_bar")
	snakeCase(t, "foo_Bar", "foo_bar")
	snakeCase(t, "fooBar", "foo_bar")
	snakeCase(t, "__FOO_BAR__", "f_o_o_b_a_r")
	snakeCase(t, "aBbc-s$@a&%_B.B^C", "a_bbc_s_a_b_b_c")
}

func snakeCase(t *testing.T, test string, expected string) {
	res := SnakeCase(test)
	if res != expected {
		internal.LogFailedTestInfo(t, "SnakeCase", test, expected, res)
		t.FailNow()
	}
}

func TestLowerFirst(t *testing.T) {
	lowerFirst(t, "foo", "foo")
	lowerFirst(t, "BAR", "bAR")
	lowerFirst(t, "FOo", "fOo")
	lowerFirst(t, "FOo大", "fOo大")
}

func lowerFirst(t *testing.T, test string, expected string) {
	res := LowerFirst(test)
	if res != expected {
		internal.LogFailedTestInfo(t, "LowerFirst", test, expected, res)
		t.FailNow()
	}
}

func TestPadEnd(t *testing.T) {
	padEnd(t, "a", 1, "b", "a")
	padEnd(t, "a", 2, "b", "ab")
	padEnd(t, "abcd", 6, "mno", "abcdmn")
	padEnd(t, "abcd", 6, "m", "abcdmm")
	padEnd(t, "abc", 6, "ab", "abcaba")
}

func padEnd(t *testing.T, source string, size int, fillString string, expected string) {
	res := PadEnd(source, size, fillString)
	if res != expected {
		internal.LogFailedTestInfo(t, "PadEnd", source, expected, res)
		t.FailNow()
	}
}

func TestPadStart(t *testing.T) {
	padStart(t, "a", 1, "b", "a")
	padStart(t, "a", 2, "b", "ba")
	padStart(t, "abcd", 6, "mno", "mnabcd")
	padStart(t, "abcd", 6, "m", "mmabcd")
	padStart(t, "abc", 6, "ab", "abaabc")
}

func padStart(t *testing.T, source string, size int, fillString string, expected string) {
	res := PadStart(source, size, fillString)
	if res != expected {
		internal.LogFailedTestInfo(t, "PadEnd", source, expected, res)
		t.FailNow()
	}
}

func TestBefore(t *testing.T) {
	before(t, "lancet", "", "lancet")
	before(t, "github.com/test/lancet", "/", "github.com")
	before(t, "github.com/test/lancet", "test", "github.com/")
}

func before(t *testing.T, source, char, expected string) {
	res := Before(source, char)
	if res != expected {
		internal.LogFailedTestInfo(t, "Before", source, expected, res)
		t.FailNow()
	}
}

func TestBeforeLast(t *testing.T) {
	beforeLast(t, "lancet", "", "lancet")
	beforeLast(t, "github.com/test/lancet", "/", "github.com/test")
	beforeLast(t, "github.com/test/test/lancet", "test", "github.com/test/")
}

func beforeLast(t *testing.T, source, char, expected string) {
	res := BeforeLast(source, char)
	if res != expected {
		internal.LogFailedTestInfo(t, "BeforeLast", source, expected, res)
		t.FailNow()
	}
}

func TestAfter(t *testing.T) {
	after(t, "lancet", "", "lancet")
	after(t, "github.com/test/lancet", "/", "test/lancet")
	after(t, "github.com/test/lancet", "test", "/lancet")
}

func after(t *testing.T, source, char, expected string) {
	res := After(source, char)
	if res != expected {
		internal.LogFailedTestInfo(t, "After", source, expected, res)
		t.FailNow()
	}
}

func TestAfterLast(t *testing.T) {
	afterLast(t, "lancet", "", "lancet")
	afterLast(t, "github.com/test/lancet", "/", "lancet")
	afterLast(t, "github.com/test/test/lancet", "test", "/lancet")
}

func afterLast(t *testing.T, source, char, expected string) {
	res := AfterLast(source, char)
	if res != expected {
		internal.LogFailedTestInfo(t, "AfterLast", source, expected, res)
		t.FailNow()
	}
}

func TestIsString(t *testing.T) {
	isString(t, "lancet", true)
	isString(t, 1, false)
	isString(t, true, false)
	isString(t, []string{}, false)
}

func isString(t *testing.T, test interface{}, expected bool) {
	res := IsString(test)
	if res != expected {
		internal.LogFailedTestInfo(t, "IsString", test, expected, res)
		t.FailNow()
	}
}

func TestReverseStr(t *testing.T) {
	reverseStr(t, "abc", "cba")
	reverseStr(t, "12345", "54321")

	//failed
	//reverseStr(t, "abc", "abc")
}

func reverseStr(t *testing.T, test string, expected string) {
	res := ReverseStr(test)
	if res != expected {
		internal.LogFailedTestInfo(t, "ReverseStr", test, expected, res)
		t.FailNow()
	}
}
