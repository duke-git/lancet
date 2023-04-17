package maputil

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

type (
	Person struct {
		Name  string  `json:"name"`
		Age   int     `json:"age"`
		Phone string  `json:"phone"`
		Addr  Address `json:"address"`
	}

	Address struct {
		Street string `json:"street"`
		Number int    `json:"number"`
	}
)

func TestStructType(t *testing.T) {
	assert := internal.NewAssert(t, "TestStructType")

	src := map[string]interface{}{
		"name":  "Nothin",
		"age":   28,
		"phone": "123456789",
		"address": map[string]interface{}{
			"street": "test",
			"number": 1,
		},
	}

	var p Person
	err := MapTo(src, &p)
	assert.IsNil(err)
	assert.Equal(src["name"], p.Name)
	assert.Equal(src["age"], p.Age)
	assert.Equal(src["phone"], p.Phone)
	assert.Equal("test", p.Addr.Street)
	assert.Equal(1, p.Addr.Number)
}

func TestBaseType(t *testing.T) {
	assert := internal.NewAssert(t, "TestBaseType")

	tc := map[string]interface{}{
		"i32":     -32,
		"i8":      -8,
		"i16":     -16,
		"i64":     -64,
		"i":       -1,
		"u32":     32,
		"u8":      8,
		"u16":     16,
		"u64":     64,
		"u":       1,
		"tf":      true,
		"f32":     1.32,
		"f64":     1.64,
		"str":     "hello mapto",
		"complex": 1 + 3i,
	}

	type BaseType struct {
		I    int        `json:"i"`
		I8   int8       `json:"i8"`
		I16  int16      `json:"i16"`
		I32  int32      `json:"i32"`
		I64  int64      `json:"i64"`
		U    uint       `json:"u"`
		U8   uint8      `json:"u8"`
		U16  uint16     `json:"u16"`
		U32  uint32     `json:"u32"`
		U64  uint64     `json:"u64"`
		F32  float32    `json:"f32"`
		F64  float64    `json:"f64"`
		Tf   bool       `json:"tf"`
		Str  string     `json:"str"`
		Comp complex128 `json:"complex"`
	}

	var dist BaseType
	err := MapTo(tc, &dist)
	assert.IsNil(err)

	var number float64

	MapTo(tc["i"], &number)
	assert.EqualValues(-1, number)
	MapTo(tc["i8"], &number)
	assert.EqualValues(-8, number)
	MapTo(tc["i16"], &number)
	assert.EqualValues(-16, number)
	MapTo(tc["i32"], &number)
	assert.EqualValues(-32, number)
	MapTo(tc["i64"], &number)
	assert.EqualValues(-64, number)
	MapTo(tc["u"], &number)
	assert.EqualValues(1, number)
	MapTo(tc["u8"], &number)
	assert.EqualValues(8, number)
	MapTo(tc["u16"], &number)
	assert.EqualValues(16, number)
	MapTo(tc["u32"], &number)
	assert.EqualValues(32, number)
	MapTo(tc["u64"], &number)
	assert.EqualValues(64, number)
}
