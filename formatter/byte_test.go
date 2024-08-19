package formatter

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestDecimalBytes(t *testing.T) {
	t.Parallel()

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
	assert.Equal("62MB", DecimalBytes(61812496, 0))
	assert.Equal("61.8MB", DecimalBytes(61812496, 1))
	assert.Equal("401MB", DecimalBytes(401000000))
	assert.Equal("401MB", DecimalBytes(401000000, 0))
	assert.Equal("4MB", DecimalBytes(4010000, 0))

	assert.Equal("4MB", DecimalBytes(4000000, 0))
	assert.Equal("4MB", DecimalBytes(4000000, 1))
}

func TestBinaryBytes(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBinaryBytes")

	assert.Equal("1KiB", BinaryBytes(1024))
	assert.Equal("1MiB", BinaryBytes(1024*1024))
	assert.Equal("1.1774MiB", BinaryBytes(1234567))
	assert.Equal("1.18MiB", BinaryBytes(1234567, 2))

	assert.Equal("10KiB", BinaryBytes(10240, 0))
	assert.Equal("10KiB", BinaryBytes(10240, 1))
	assert.Equal("10KiB", BinaryBytes(10240, 2))
}

func TestParseDecimalBytes(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
