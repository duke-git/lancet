package strutil

import (
	"fmt"
	"reflect"
)

func ExampleAfter() {
	result1 := After("foo", "")
	result2 := After("foo", "foo")
	result3 := After("foo/bar", "foo")
	result4 := After("foo/bar", "/")
	result5 := After("foo/bar/baz", "/")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	// Output:
	// foo
	//
	// /bar
	// bar
	// bar/baz
}

func ExampleAfterLast() {
	result1 := AfterLast("foo", "")
	result2 := AfterLast("foo", "foo")
	result3 := AfterLast("foo/bar", "/")
	result4 := AfterLast("foo/bar/baz", "/")
	result5 := AfterLast("foo/bar/foo/baz", "foo")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	// Output:
	// foo
	//
	// bar
	// baz
	// /baz
}

func ExampleBefore() {
	result1 := Before("foo", "")
	result2 := Before("foo", "foo")
	result3 := Before("foo/bar", "/")
	result4 := Before("foo/bar/baz", "/")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	// Output:
	// foo
	//
	// foo
	// foo
}

func ExampleBeforeLast() {
	result1 := BeforeLast("foo", "")
	result2 := BeforeLast("foo", "foo")
	result3 := BeforeLast("foo/bar", "/")
	result4 := BeforeLast("foo/bar/baz", "/")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	// Output:
	// foo
	//
	// foo
	// foo/bar
}

func ExampleCamelCase() {
	strings := []string{"", "foobar", "&FOO:BAR$BAZ", "$foo%", "Foo-#1😄$_%^&*(1bar"}

	for _, v := range strings {
		s := CamelCase(v)
		fmt.Println(s)
	}
	// Output:
	//
	// foobar
	// fooBarBaz
	// foo
	// foo11Bar
}

func ExampleCapitalize() {
	strings := []string{"", "Foo", "_foo", "fooBar", "foo-bar"}

	for _, v := range strings {
		s := Capitalize(v)
		fmt.Println(s)
	}
	// Output:
	//
	// Foo
	// _foo
	// Foobar
	// Foo-bar
}

func ExampleIsString() {
	result1 := IsString("")
	result2 := IsString("a")
	result3 := IsString(1)
	result4 := IsString(true)
	result5 := IsString([]string{"a"})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	// Output:
	// true
	// true
	// false
	// false
	// false
}

func ExampleKebabCase() {
	strings := []string{"", "foo-bar", "Foo Bar-", "FOOBAR", "Foo-#1😄$_%^&*(1bar"}

	for _, v := range strings {
		s := KebabCase(v)
		fmt.Println(s)
	}
	// Output:
	//
	// foo-bar
	// foo-bar
	// foobar
	// foo-1-1-bar
}

func ExampleUpperKebabCase() {
	strings := []string{"", "foo-bar", "Foo Bar-", "FooBAR", "Foo-#1😄$_%^&*(1bar"}

	for _, v := range strings {
		s := UpperKebabCase(v)
		fmt.Println(s)
	}
	// Output:
	//
	// FOO-BAR
	// FOO-BAR
	// FOO-BAR
	// FOO-1-1-BAR
}

func ExampleLowerFirst() {
	strings := []string{"", "bar", "BAr", "Bar大"}

	for _, v := range strings {
		s := LowerFirst(v)
		fmt.Println(s)
	}
	// Output:
	//
	// bar
	// bAr
	// bar大
}

func ExampleUpperFirst() {
	strings := []string{"", "bar", "BAr", "bar大"}

	for _, v := range strings {
		s := UpperFirst(v)
		fmt.Println(s)
	}
	// Output:
	//
	// Bar
	// BAr
	// Bar大
}

func ExamplePad() {
	result1 := Pad("foo", 1, "bar")
	result2 := Pad("foo", 2, "bar")
	result3 := Pad("foo", 3, "bar")
	result4 := Pad("foo", 4, "bar")
	result5 := Pad("foo", 5, "bar")
	result6 := Pad("foo", 6, "bar")
	result7 := Pad("foo", 7, "bar")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)

	// Output:
	// foo
	// foo
	// foo
	// foob
	// bfoob
	// bfooba
	// bafooba
}

func ExamplePadEnd() {
	result1 := PadEnd("foo", 1, "bar")
	result2 := PadEnd("foo", 2, "bar")
	result3 := PadEnd("foo", 3, "bar")
	result4 := PadEnd("foo", 4, "bar")
	result5 := PadEnd("foo", 5, "bar")
	result6 := PadEnd("foo", 6, "bar")
	result7 := PadEnd("foo", 7, "bar")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)
	// Output:
	// foo
	// foo
	// foo
	// foob
	// fooba
	// foobar
	// foobarb
}

func ExamplePadStart() {
	result1 := PadStart("foo", 1, "bar")
	result2 := PadStart("foo", 2, "bar")
	result3 := PadStart("foo", 3, "bar")
	result4 := PadStart("foo", 4, "bar")
	result5 := PadStart("foo", 5, "bar")
	result6 := PadStart("foo", 6, "bar")
	result7 := PadStart("foo", 7, "bar")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)
	// Output:
	// foo
	// foo
	// foo
	// bfoo
	// bafoo
	// barfoo
	// barbfoo
}

func ExampleReverse() {
	s := "foo"
	rs := Reverse(s)

	fmt.Println(s)
	fmt.Println(rs)
	// Output:
	// foo
	// oof
}

func ExampleSnakeCase() {
	strings := []string{"", "foo-bar", "Foo Bar-", "FOOBAR", "Foo-#1😄$_%^&*(1bar"}

	for _, v := range strings {
		s := SnakeCase(v)
		fmt.Println(s)
	}
	// Output:
	//
	// foo_bar
	// foo_bar
	// foobar
	// foo_1_1_bar
}

func ExampleUpperSnakeCase() {
	strings := []string{"", "foo-bar", "Foo Bar-", "FooBAR", "Foo-#1😄$_%^&*(1bar"}

	for _, v := range strings {
		s := UpperSnakeCase(v)
		fmt.Println(s)
	}
	// Output:
	//
	// FOO_BAR
	// FOO_BAR
	// FOO_BAR
	// FOO_1_1_BAR
}

func ExampleSplitEx() {
	result1 := SplitEx(" a b c ", "", true)

	result2 := SplitEx(" a b c ", " ", false)
	result3 := SplitEx(" a b c ", " ", true)

	result4 := SplitEx("a = b = c = ", " = ", false)
	result5 := SplitEx("a = b = c = ", " = ", true)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	// Output:
	// []
	// [ a b c ]
	// [a b c]
	// [a b c ]
	// [a b c]
}

func ExampleWrap() {
	result1 := Wrap("foo", "")
	result2 := Wrap("foo", "*")
	result3 := Wrap("'foo'", "'")
	result4 := Wrap("", "*")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	// Output:
	// foo
	// *foo*
	// ''foo''
	//
}

func ExampleUnwrap() {
	result1 := Unwrap("foo", "")
	result2 := Unwrap("*foo*", "*")
	result3 := Unwrap("*foo", "*")
	result4 := Unwrap("foo*", "*")
	result5 := Unwrap("**foo**", "*")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	// Output:
	// foo
	// foo
	// *foo
	// foo*
	// *foo*
}

func ExampleSubstring() {

	result1 := Substring("abcde", 1, 3)
	result2 := Substring("abcde", 1, 5)
	result3 := Substring("abcde", -1, 3)
	result4 := Substring("abcde", -2, 2)
	result5 := Substring("abcde", -2, 3)
	result6 := Substring("你好，欢迎你", 0, 2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	// Output:
	// bcd
	// bcde
	// e
	// de
	// de
	// 你好
}

func ExampleSplitWords() {

	result1 := SplitWords("a word")
	result2 := SplitWords("I'am a programmer")
	result3 := SplitWords("a -b-c' 'd'e")
	result4 := SplitWords("你好，我是一名码农")
	result5 := SplitWords("こんにちは，私はプログラマーです")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// [a word]
	// [I'am a programmer]
	// [a b-c' d'e]
	// []
	// []
}

func ExampleWordCount() {

	result1 := WordCount("a word")
	result2 := WordCount("I'am a programmer")
	result3 := WordCount("a -b-c' 'd'e")
	result4 := WordCount("你好，我是一名码农")
	result5 := WordCount("こんにちは，私はプログラマーです")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// 2
	// 3
	// 3
	// 0
	// 0
}

func ExampleRemoveNonPrintable() {
	result1 := RemoveNonPrintable("hello\u00a0 \u200bworld\n")
	result2 := RemoveNonPrintable("你好😄")

	fmt.Println(result1)
	fmt.Println(result2)
	// Output:
	// hello world
	// 你好😄
}

func ExampleStringToBytes() {
	result1 := StringToBytes("abc")
	result2 := reflect.DeepEqual(result1, []byte{'a', 'b', 'c'})

	fmt.Println(result1)
	fmt.Println(result2)
	// Output:
	// [97 98 99]
	// true
}

func ExampleBytesToString() {
	bytes := []byte{'a', 'b', 'c'}
	result := BytesToString(bytes)

	fmt.Println(result)
	// Output:
	// abc
}

func ExampleIsBlank() {
	result1 := IsBlank("")
	result2 := IsBlank("\t\v\f\n")
	result3 := IsBlank(" 中文")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	// Output:
	// true
	// true
	// false
}

func ExampleIsNotBlank() {
	result1 := IsNotBlank("")
	result2 := IsNotBlank("			")
	result3 := IsNotBlank("\t\v\f\n")
	result4 := IsNotBlank(" 中文")
	result5 := IsNotBlank(" 	world	")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	// Output:
	// false
	// false
	// false
	// true
	// true
}

func ExampleHasPrefixAny() {
	result1 := HasPrefixAny("foo bar", []string{"fo", "xyz", "hello"})
	result2 := HasPrefixAny("foo bar", []string{"oom", "world"})

	fmt.Println(result1)
	fmt.Println(result2)
	// Output:
	// true
	// false
}

func ExampleHasSuffixAny() {
	result1 := HasSuffixAny("foo bar", []string{"bar", "xyz", "hello"})
	result2 := HasSuffixAny("foo bar", []string{"oom", "world"})

	fmt.Println(result1)
	fmt.Println(result2)
	// Output:
	// true
	// false
}

func ExampleIndexOffset() {
	str := "foo bar hello world"

	result1 := IndexOffset(str, "o", 5)
	result2 := IndexOffset(str, "o", 0)
	result3 := IndexOffset(str, "d", len(str)-1)
	result4 := IndexOffset(str, "d", len(str))
	result5 := IndexOffset(str, "f", -1)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	// Output:
	// 12
	// 1
	// 18
	// -1
	// -1
}

func ExampleReplaceWithMap() {
	str := "ac ab ab ac"
	replaces := map[string]string{
		"a": "1",
		"b": "2",
	}

	result := ReplaceWithMap(str, replaces)

	fmt.Println(result)
	// Output:
	// 1c 12 12 1c
}

func ExampleTrim() {
	result1 := Trim("\nabcd")

	str := "$ ab	cd $ "

	result2 := Trim(str)
	result3 := Trim(str, "$")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// abcd
	// $ ab	cd $
	// ab	cd
}

func ExampleSplitAndTrim() {
	str := " a,b, c,d,$1 "

	result1 := SplitAndTrim(str, ",")
	result2 := SplitAndTrim(str, ",", "$")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// [a b c d $1]
	// [a b c d 1]
}

func ExampleHideString() {
	str := "13242658976"

	result1 := HideString(str, 3, 3, "*")
	result2 := HideString(str, 3, 4, "*")
	result3 := HideString(str, 3, 7, "*")
	result4 := HideString(str, 7, 11, "*")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 13242658976
	// 132*2658976
	// 132****8976
	// 1324265****
}

func ExampleContainsAll() {
	str := "hello world"

	result1 := ContainsAll(str, []string{"hello", "world"})
	result2 := ContainsAll(str, []string{"hello", "abc"})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleContainsAny() {
	str := "hello world"

	result1 := ContainsAny(str, []string{"hello", "world"})
	result2 := ContainsAny(str, []string{"hello", "abc"})
	result3 := ContainsAny(str, []string{"123", "abc"})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// true
	// false
}

func ExampleRemoveWhiteSpace() {
	str := " hello   \r\n	\t   world"

	result1 := RemoveWhiteSpace(str, true)
	result2 := RemoveWhiteSpace(str, false)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// helloworld
	// hello world
}

func ExampleSubInBetween() {
	str := "abcde"

	result1 := SubInBetween(str, "", "de")
	result2 := SubInBetween(str, "a", "d")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// abc
	// bc
}

func ExampleHammingDistance() {

	result, _ := HammingDistance("abc", "def")
	fmt.Println(result)

	result, _ = HammingDistance("name", "namf")
	fmt.Println(result)

	// Output:
	// 3
	// 1
}

func ExampleConcat() {
	result1 := Concat(12, "Hello", " ", "World", "!")
	result2 := Concat(11, "Go", " ", "Language")
	result3 := Concat(0, "An apple a ", "day，", "keeps the", " doctor away")
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// Hello World!
	// Go Language
	// An apple a day，keeps the doctor away
}

func ExampleEllipsis() {
	result1 := Ellipsis("hello world", 5)
	result2 := Ellipsis("你好，世界!", 2)
	result3 := Ellipsis("😀😃😄😁😆", 3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// hello...
	// 你好...
	// 😀😃😄...
}

func ExampleRotate() {
	result1 := Rotate("Hello", 0)
	result2 := Rotate("Hello", 1)
	result3 := Rotate("Hello", 2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// Hello
	// oHell
	// loHel
}

func ExampleTemplateReplace() {
	template := `Hello, my name is {name}, I'm {age} years old.`
	data := map[string]string{
		"name": "Bob",
		"age":  "20",
	}

	result := TemplateReplace(template, data)

	fmt.Println(result)

	// Output:
	// Hello, my name is Bob, I'm 20 years old.
}

func ExampleRegexMatchAllGroups() {
	pattern := `(\w+\.+\w+)@(\w+)\.(\w+)`
	str := "Emails: john.doe@example.com and jane.doe@example.com"

	result := RegexMatchAllGroups(pattern, str)

	fmt.Println(result[0])
	fmt.Println(result[1])

	// Output:
	// [john.doe@example.com john.doe example com]
	// [jane.doe@example.com jane.doe example com]
}
