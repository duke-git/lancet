package pointer

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestOf(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestOf")

	result1 := Of(123)
	result2 := Of("abc")

	assert.Equal(123, *result1)
	assert.Equal("abc", *result2)
}

func TestUnwrap(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUnwrap")

	a := 123
	b := "abc"

	assert.Equal(a, Unwrap(&a))
	assert.Equal(b, Unwrap(&b))
}

func TestUnwarpOr(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUnwarpOr")

	a := 123
	b := "abc"

	var c *int
	var d *string

	assert.Equal(a, UnwarpOr(&a, 456))
	assert.Equal(b, UnwarpOr(&b, "abc"))
	assert.Equal(456, UnwarpOr(c, 456))
	assert.Equal("def", UnwarpOr(d, "def"))
}

func TestUnwrapOrDefault(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUnwrapOrDefault")

	a := 123
	b := "abc"

	var c *int
	var d *string

	assert.Equal(a, UnwarpOrDefault(&a))
	assert.Equal(b, UnwarpOrDefault(&b))
	assert.Equal(0, UnwarpOrDefault(c))
	assert.Equal("", UnwarpOrDefault(d))
}

func TestExtractPointer(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestExtractPointer")

	a := 1
	b := &a
	c := &b
	d := &c

	assert.Equal(1, ExtractPointer(d))
}
