package formatter

import (
	"bytes"
	"testing"

	"github.com/duke-git/lancet/internal"
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

	assert.Equal("-999", Comma(-999, ""))
	assert.Equal("-1,000", Comma(-1000, ""))
	assert.Equal("-1,234,567", Comma(-1234567, ""))
}

func TestPretty(t *testing.T) {
	assert := internal.NewAssert(t, "TestPretty")

	cases := []interface{}{
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

func TestDecimalBytes(t *testing.T) {
	assert := internal.NewAssert(t, "TestDecimalBytes")

	assert.Equal("1KB", DecimalBytes(1000))
	assert.Equal("1.024KB", DecimalBytes(1024))
	assert.Equal("1.2346MB", DecimalBytes(1234567))
	assert.Equal("1.235MB", DecimalBytes(1234567, 3))
	assert.Equal("1.123GB", DecimalBytes(float64(1.123*UnitGB)))
	assert.Equal("2.123TB", DecimalBytes(float64(2.123*UnitTB)))
	assert.Equal("3.123PB", DecimalBytes(float64(3.123*UnitPB)))
	assert.Equal("4.123EB", DecimalBytes(float64(4.123*UnitEB)))
	assert.Equal("1EB", DecimalBytes(float64(1000*UnitPB)))
}

func TestBinaryBytes(t *testing.T) {
	assert := internal.NewAssert(t, "TestBinaryBytes")

	assert.Equal("1KiB", BinaryBytes(1024))
	assert.Equal("1MiB", BinaryBytes(1024*1024))
	assert.Equal("1.1774MiB", BinaryBytes(1234567))
	assert.Equal("1.18MiB", BinaryBytes(1234567, 2))
}

func TestParseDecimalBytes(t *testing.T) {
	assert := internal.NewAssert(t, "TestParseDecimalBytes")

	cases := map[string]uint64{
		"12":      uint64(12),
		"12 k":    uint64(12000),
		"12 kb":   uint64(12000),
		"12kb":    uint64(12000),
		"12k":     uint64(12000),
		"12K":     uint64(12000),
		"12KB":    uint64(12000),
		"12 K":    uint64(12000),
		"12 KB":   uint64(12000),
		"12 Kb":   uint64(12000),
		"12 kB":   uint64(12000),
		"12.2 KB": uint64(12200),
	}

	for k, v := range cases {
		result, err := ParseDecimalBytes(k)
		assert.Equal(v, result)
		assert.IsNil(err)
	}

	_, err := ParseDecimalBytes("12 AB")
	assert.IsNotNil(err)
}

func TestParseBinaryBytes(t *testing.T) {
	assert := internal.NewAssert(t, "TestParseBinaryBytes")

	cases := map[string]uint64{
		"12":       uint64(12),
		"12 ki":    uint64(12288),
		"12 kib":   uint64(12288),
		"12kib":    uint64(12288),
		"12ki":     uint64(12288),
		"12KI":     uint64(12288),
		"12KIB":    uint64(12288),
		"12KiB":    uint64(12288),
		"12 Ki":    uint64(12288),
		"12 KiB":   uint64(12288),
		"12 Kib":   uint64(12288),
		"12 kiB":   uint64(12288),
		"12.2 KiB": uint64(12492),
	}

	for k, v := range cases {
		result, err := ParseBinaryBytes(k)
		assert.Equal(v, result)
		assert.IsNil(err)
	}

	_, err := ParseDecimalBytes("12 AB")
	assert.IsNotNil(err)
}
