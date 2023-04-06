package formatter

import (
	"bytes"
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
	assert.Equal("12,345,678.9", Comma("12345678.9", ""))

	assert.Equal("12,345", Comma(12345, ""))
	assert.Equal("$12,345", Comma(12345, "$"))
	assert.Equal("¥12,345", Comma(12345, "¥"))
	assert.Equal("12,345.6789", Comma(12345.6789, ""))
	assert.Equal("12,345.6789", Comma(+12345.6789, ""))
	assert.Equal("12,345,678.9", Comma(12345678.9, ""))
	assert.Equal("123,456,789,000", Comma(123456789000, ""))
}

func TestPretty(t *testing.T) {
	assert := internal.NewAssert(t, "TestPretty")

	cases := []any{
		"",
		"abc",
		123,
		[]string{"a", "b", "c"},
		map[string]int{"a": 1},
		struct {
			Abc int `json:"abc"`
		}{Abc: 123},
	}

	expects := []string{
		"\"\"",
		`"abc"`,
		"123",
		"[\n    \"a\",\n    \"b\",\n    \"c\"\n]",
		"{\n    \"a\": 1\n}",
		"{\n    \"abc\": 123\n}",
	}

	for i, v := range cases {
		result, err := Pretty(v)

		assert.IsNil(err)

		t.Log("result -> ", result)
		assert.Equal(expects[i], result)
	}
}

func TestPrettyToWriter(t *testing.T) {
	assert := internal.NewAssert(t, "TestPrettyToWriter")

	type User struct {
		Name string `json:"name"`
		Aage uint   `json:"age"`
	}
	user := User{Name: "King", Aage: 10000}

	expects := "{\n    \"name\": \"King\",\n    \"age\": 10000\n}\n"

	buf := &bytes.Buffer{}
	err := PrettyToWriter(user, buf)

	assert.IsNil(err)
	assert.Equal(expects, buf.String())
}
