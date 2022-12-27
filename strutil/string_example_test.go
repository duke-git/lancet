package strutil

import "fmt"

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

func ExamplePadEnd() {
	result1 := PadEnd("foo", 1, "bar")
	result2 := PadEnd("foo", 2, "bar")
	result3 := PadEnd("foo", 3, "bar")
	result4 := PadEnd("foo", 4, "bar")
	result5 := PadEnd("foo", 5, "bar")
	result6 := PadEnd("foo", 6, "bar")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	// Output:
	// foo
	// foo
	// foo
	// foob
	// fooba
	// foobar
}

func ExamplePadStart() {
	result1 := PadStart("foo", 1, "bar")
	result2 := PadStart("foo", 2, "bar")
	result3 := PadStart("foo", 3, "bar")
	result4 := PadStart("foo", 4, "bar")
	result5 := PadStart("foo", 5, "bar")
	result6 := PadStart("foo", 6, "bar")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	// Output:
	// foo
	// foo
	// foo
	// bfoo
	// bafoo
	// barfoo
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
