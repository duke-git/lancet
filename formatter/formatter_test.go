package formatter

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestComma(t *testing.T) {
	assert := internal.NewAssert(t, "TestComma")

	assert.Equal("", Comma("", ""))
	assert.Equal("", Comma("aa", ""))
	assert.Equal("", Comma("aa.a", ""))
	assert.Equal("", Comma([]int{1}, ""))
	assert.Equal("123", Comma("123", ""))
	assert.Equal("12,345", Comma("12345", ""))

	assert.Equal("12,345", Comma(12345, ""))
	assert.Equal("$12,345", Comma(12345, "$"))
	assert.Equal("¥12,345", Comma(12345, "¥"))
	assert.Equal("12,345.6789", Comma(12345.6789, ""))
}
