package formatter

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestComma(t *testing.T) {
	assert := internal.NewAssert(t, "TestComma")

	assert.Equal("", Comma("", ""))
	assert.Equal("", Comma("aa", ""))
	assert.Equal("", Comma("aa.a", ""))
	assert.Equal("123", Comma("123", ""))
	assert.Equal("12,345", Comma("12345", ""))
	assert.Equal("12,345.6789", Comma("12345.6789", ""))
	assert.Equal("123,456,789,000", Comma("123456789000", ""))

	assert.Equal("12,345", Comma(12345, ""))
	assert.Equal("$12,345", Comma(12345, "$"))
	assert.Equal("¥12,345", Comma(12345, "¥"))
	assert.Equal("12,345.6789", Comma(12345.6789, ""))
	assert.Equal("12,345.6789", Comma(+12345.6789, ""))
	assert.Equal("12,345,678.9", Comma(12345678.9, ""))
	assert.Equal("123,456,789,000", Comma(123456789000, ""))
}
