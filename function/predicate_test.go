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

func TestPredicatesNandPure(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPredicatesNandPure")

	isNumericAndLength5 := Nand(
		func(s string) bool { return strings.ContainsAny(s, "0123456789") },
		func(s string) bool { return len(s) == 5 },
	)

	assert.ShouldBeFalse(isNumericAndLength5("12345"))
	assert.ShouldBeTrue(isNumericAndLength5("1234"))
	assert.ShouldBeTrue(isNumericAndLength5("abcdef"))
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

func TestPredicatesXnorPure(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPredicatesXnorPure")

	isEven := func(i int) bool { return i%2 == 0 }
	isPositive := func(i int) bool { return i > 0 }

	match := Xnor(isEven, isPositive)

	assert.ShouldBeTrue(match(2))
	assert.ShouldBeTrue(match(-3))
	assert.ShouldBeFalse(match(3))
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

	k := Nor(a, b)
	assert.ShouldBeFalse(k("hello!"))

	o := Xnor(a, b)
	assert.ShouldBeTrue(o("hello!"))

	p := Nand(c, k)
	assert.ShouldBeTrue(p("hello!"))
}
