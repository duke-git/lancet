package random

import (
	"fmt"
	"regexp"
)

func ExampleRandInt() {
	result := RandInt(1, 10)

	if result >= 1 && result < 10 {
		fmt.Println("ok")
	}

	// Output:
	// ok
}

func ExampleRandBytes() {
	bytes := RandBytes(4)

	fmt.Println(len(bytes))

	// Output:
	// 4
}

func ExampleRandString() {
	pattern := `^[a-zA-Z]+$`
	reg := regexp.MustCompile(pattern)

	s := RandString(6)

	result1 := reg.MatchString(s)
	result2 := len(s)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// 6
}

func ExampleRandUpper() {
	pattern := `^[A-Z]+$`
	reg := regexp.MustCompile(pattern)

	s := RandUpper(6)

	result1 := reg.MatchString(s)
	result2 := len(s)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// 6
}

func ExampleRandLower() {
	pattern := `^[a-z]+$`
	reg := regexp.MustCompile(pattern)

	s := RandLower(6)

	result1 := reg.MatchString(s)
	result2 := len(s)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// 6
}

func ExampleRandNumeral() {
	pattern := `^[0-9]+$`
	reg := regexp.MustCompile(pattern)

	s := RandNumeral(6)

	result1 := reg.MatchString(s)
	result2 := len(s)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// 6
}

func ExampleRandNumeralOrLetter() {
	pattern := `^[0-9a-zA-Z]+$`
	reg := regexp.MustCompile(pattern)

	s := RandNumeralOrLetter(6)

	result1 := reg.MatchString(s)
	result2 := len(s)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// 6
}

func ExampleUUIdV4() {
	pattern := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`
	reg := regexp.MustCompile(pattern)

	s, _ := UUIdV4()

	result := reg.MatchString(s)

	fmt.Println(result)

	// Output:
	// true
}

func ExampleRandUniqueIntSlice() {
	result := RandUniqueIntSlice(5, 0, 10)

	if len(result) == 5 {
		fmt.Println("ok")
	}

	// Output:
	// ok
}
