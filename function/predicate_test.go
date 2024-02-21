package function

import (
	"strings"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestPredicatesNegatePure(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPredicatesNegatePure")

	// Define some simple predicates for demonstration
	isUpperCase := func(s string) bool {
		return strings.ToUpper(s) == s
	}
	isLowerCase := func(s string) bool {
		return strings.ToLower(s) == s
	}
	isMixedCase := Negate(Or(isUpperCase, isLowerCase))

	assert.ShouldBeFalse(isMixedCase("ABC"))
	assert.ShouldBeTrue(isMixedCase("AbC"))
}

func TestPredicatesOrPure(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPredicatesOrPure")

	containsDigitOrSpecialChar := Or(
		func(s string) bool { return strings.ContainsAny(s, "0123456789") },
		func(s string) bool { return strings.ContainsAny(s, "!@#$%") },
	)

	assert.ShouldBeTrue(containsDigitOrSpecialChar("hello!"))
	assert.ShouldBeFalse(containsDigitOrSpecialChar("hello"))
}

func TestPredicatesAndPure(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPredicatesAndPure")

	isNumericAndLength5 := And(
		func(s string) bool { return strings.ContainsAny(s, "0123456789") },
		func(s string) bool { return len(s) == 5 },
	)

	assert.ShouldBeTrue(isNumericAndLength5("12345"))
	assert.ShouldBeFalse(isNumericAndLength5("1234"))
	assert.ShouldBeFalse(isNumericAndLength5("abcde"))
}

func TestPredicatesNorPure(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPredicatesNorPure")

	match := Nor(
		func(s string) bool { return strings.ContainsAny(s, "0123456789") },
		func(s string) bool { return len(s) == 5 },
	)

	assert.ShouldBeTrue(match("dbcdckkeee"))

	match = Nor(
		func(s string) bool { return strings.ContainsAny(s, "0123456789") },
		func(s string) bool { return len(s) == 5 },
	)

	assert.ShouldBeFalse(match("0123456789"))
}

func TestPredicatesMix(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPredicatesMix")

	a := Or(
		func(s string) bool { return strings.ContainsAny(s, "0123456789") },
		func(s string) bool { return strings.ContainsAny(s, "!") },
	)

	b := And(
		func(s string) bool { return strings.ContainsAny(s, "hello") },
		func(s string) bool { return strings.ContainsAny(s, "!") },
	)

	c := Negate(And(a, b))

	assert.ShouldBeFalse(c("hello!"))

	c = Nor(a, b)
	assert.ShouldBeFalse(c("hello!"))
}
