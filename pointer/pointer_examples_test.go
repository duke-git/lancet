package pointer

import "fmt"

func ExampleOf() {
	result1 := Of(123)
	result2 := Of("abc")

	fmt.Println(*result1)
	fmt.Println(*result2)

	// Output:
	// 123
	// abc
}

func ExampleUnwrap() {
	a := 123
	b := "abc"

	result1 := Unwrap(&a)
	result2 := Unwrap(&b)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 123
	// abc
}

func ExampleUnwarpOr() {
	a := 123
	b := "abc"

	var c *int
	var d *string

	result1 := UnwarpOr(&a, 456)
	result2 := UnwarpOr(&b, "abc")
	result3 := UnwarpOr(c, 456)
	result4 := UnwarpOr(d, "def")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 123
	// abc
	// 456
	// def
}

func ExampleUnwarpOrDefault() {
	a := 123
	b := "abc"

	var c *int
	var d *string

	result1 := UnwarpOrDefault(&a)
	result2 := UnwarpOrDefault(&b)
	result3 := UnwarpOrDefault(c)
	result4 := UnwarpOrDefault(d)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 123
	// abc
	// 0
	//
}

func ExampleExtractPointer() {
	a := 1
	b := &a
	c := &b
	d := &c

	result := ExtractPointer(d)

	fmt.Println(result)

	// Output:
	// 1
}

func ExampleIsNil() {
	a := 1
	b := &a
	c := &b
	d := &c
	e := &d
	var f *int

	result1 := IsNil(a)
	result2 := IsNil(b)
	result3 := IsNil(c)
	result4 := IsNil(d)
	result5 := IsNil(e)
	result6 := IsNil(f)
	result7 := IsNil(nil)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)

	// Output:
	// false
	// false
	// false
	// false
	// false
	// true
	// true
}

func ExampleUnwrapOr() {
	a := 123
	b := "abc"

	var c *int
	var d *string

	result1 := UnwrapOr(&a, 456)
	result2 := UnwrapOr(&b, "efg")
	result3 := UnwrapOr(c, 456)
	result4 := UnwrapOr(d, "def")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 123
	// abc
	// 456
	// def
}
