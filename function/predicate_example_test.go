package function

import (
	"fmt"
	"strings"
)

func ExampleNegate() {
	// Define some simple predicates for demonstration
	isUpperCase := func(s string) bool {
		return strings.ToUpper(s) == s
	}
	isLowerCase := func(s string) bool {
		return strings.ToLower(s) == s
	}
	isMixedCase := Negate(Or(isUpperCase, isLowerCase))

	fmt.Println(isMixedCase("ABC"))
	fmt.Println(isMixedCase("AbC"))

	// Output:
	// false
	// true
}

func ExampleOr() {
	containsDigitOrSpecialChar := Or(
		func(s string) bool { return strings.ContainsAny(s, "0123456789") },
		func(s string) bool { return strings.ContainsAny(s, "!@#$%") },
	)

	fmt.Println(containsDigitOrSpecialChar("hello!"))
	fmt.Println(containsDigitOrSpecialChar("hello"))

	// Output:
	// true
	// false
}

func ExampleAnd() {
	isNumericAndLength5 := And(
		func(s string) bool { return strings.ContainsAny(s, "0123456789") },
		func(s string) bool { return len(s) == 5 },
	)

	fmt.Println(isNumericAndLength5("12345"))
	fmt.Println(isNumericAndLength5("1234"))
	fmt.Println(isNumericAndLength5("abcde"))

	// Output:
	// true
	// false
	// false
}

func ExampleNor() {
	match := Nor(
		func(s string) bool { return strings.ContainsAny(s, "0123456789") },
		func(s string) bool { return len(s) == 5 },
	)

	fmt.Println(match("dbcdckkeee"))

	match = Nor(
		func(s string) bool { return strings.ContainsAny(s, "0123456789") },
		func(s string) bool { return len(s) == 5 },
	)

	fmt.Println(match("0123456789"))

	// Output:
	// true
	// false
}

func ExamplePredicatesMix() {
	a := Or(
		func(s string) bool { return strings.ContainsAny(s, "0123456789") },
		func(s string) bool { return strings.ContainsAny(s, "!") },
	)

	b := And(
		func(s string) bool { return strings.ContainsAny(s, "hello") },
		func(s string) bool { return strings.ContainsAny(s, "!") },
	)

	c := Negate(And(a, b))
	fmt.Println(c("hello!"))

	c = Nor(a, b)
	fmt.Println(c("hello!"))

	// Output:
	// false
	// false
}
