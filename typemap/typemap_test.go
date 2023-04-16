package typemap

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestTransform(t *testing.T) {
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
	assert.NoError(t, err)
	assert.Equal(t, src["name"], p.Name)
	assert.Equal(t, src["age"], p.Age)
	assert.Equal(t, src["phone"], p.Phone)
	assert.Equal(t, "test", p.Addr.Street)
	assert.Equal(t, 1, p.Addr.Number)
}

func TestBaseType(t *testing.T) {
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
	assert.NoError(t, err)
	assert.EqualValues(t, -1, dist.I)
	assert.EqualValues(t, -8, dist.I8)
	assert.EqualValues(t, -16, dist.I16)
	assert.EqualValues(t, -32, dist.I32)
	assert.EqualValues(t, -64, dist.I64)
	assert.EqualValues(t, 1, dist.U)
	assert.EqualValues(t, 8, dist.U8)
	assert.EqualValues(t, 16, dist.U16)
	assert.EqualValues(t, 32, dist.U32)
	assert.EqualValues(t, 64, dist.U64)
	assert.EqualValues(t, tc["f32"], dist.F32)
	assert.EqualValues(t, tc["f64"], dist.F64)
	assert.EqualValues(t, tc["str"], dist.Str)
	assert.EqualValues(t, tc["tf"], dist.Tf)
	assert.EqualValues(t, tc["complex"], dist.Comp)
}
