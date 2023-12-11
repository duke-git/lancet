package random

import (
	"fmt"
	"regexp"
	"strconv"
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

func ExampleRandSymbolChar() {
	pattern := `^[\W|_]+$`
	reg := regexp.MustCompile(pattern)

	s := RandSymbolChar(6)

	result1 := reg.MatchString(s)
	result2 := len(s)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// 6
}

func ExampleRandFloat() {
	pattern := `^[\d{1}.\d{2}]+$`
	reg := regexp.MustCompile(pattern)

	num := RandFloat(1.0, 5.0, 2)

	// check num is a random float in [1.0, 5.0)
	result1 := num >= 1.0 && num < 5.0
	result2 := reg.MatchString(strconv.FormatFloat(num, 'f', -1, 64))

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// true
}

func ExampleRandFloats() {
	isInRange := true
	numbers := RandFloats(5, 1.0, 5.0, 2)
	for _, n := range numbers {
		isInRange = (n >= 1.0 && n < 5.0)
	}

	fmt.Println(isInRange)
	fmt.Println(len(numbers))

	// Output:
	// true
	// 5
}
