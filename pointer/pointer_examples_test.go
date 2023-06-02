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
