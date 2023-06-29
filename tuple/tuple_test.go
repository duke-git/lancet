// Copyright 2023 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package tuple implements tuple data type and some operations on it.
package tuple

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestTuples(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestTuples")

	type foo struct {
		A string
	}

	t2 := NewTuple2[int, float64](1, 0.1)
	t3 := NewTuple3[int, float64, string](1, 0.1, "a")
	t4 := NewTuple4[int, float64, string, bool](1, 0.1, "a", true)
	t5 := NewTuple5[int, float64, string, bool, int64](1, 0.1, "a", true, 2)
	t6 := NewTuple6[int, float64, string, bool, int64, float32](1, 0.1, "a", true, 2, 2.2)
	t7 := NewTuple7[int, float64, string, bool, int64, float32, string](1, 0.1, "a", true, 2, 2.2, "b")
	t8 := NewTuple8[int, float64, string, bool, int64, float32, string, string](1, 0.1, "a", true, 2, 2.2, "b", "c")
	t9 := NewTuple9[int, float64, string, bool, int64, float32, string, string, map[string]int](1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1})
	t10 := NewTuple10[int, float64, string, bool, int64, float32, string, string, map[string]int, foo](1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1}, foo{A: "a"})

	assert.Equal(t2, Tuple2[int, float64]{FieldA: 1, FieldB: 0.1})
	assert.Equal(t3, Tuple3[int, float64, string]{FieldA: 1, FieldB: 0.1, FieldC: "a"})
	assert.Equal(t4, Tuple4[int, float64, string, bool]{FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true})
	assert.Equal(t5, Tuple5[int, float64, string, bool, int64]{FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2})
	assert.Equal(t6, Tuple6[int, float64, string, bool, int64, float32]{FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2, FieldF: 2.2})
	assert.Equal(t7, Tuple7[int, float64, string, bool, int64, float32, string]{FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2, FieldF: 2.2, FieldG: "b"})
	assert.Equal(t8, Tuple8[int, float64, string, bool, int64, float32, string, string]{FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2, FieldF: 2.2, FieldG: "b", FieldH: "c"})
	assert.Equal(t9, Tuple9[int, float64, string, bool, int64, float32, string, string, map[string]int]{FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2, FieldF: 2.2, FieldG: "b", FieldH: "c", FieldI: map[string]int{"a": 1}})
	assert.Equal(t10, Tuple10[int, float64, string, bool, int64, float32, string, string, map[string]int, foo]{FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2, FieldF: 2.2, FieldG: "b", FieldH: "c", FieldI: map[string]int{"a": 1}, FieldJ: foo{A: "a"}})
}

func TestTuple_Unbox(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestTuple_Unbox")

	type foo struct {
		A string
	}

	t2 := NewTuple2[int, float64](1, 0.1)
	v1, v2 := t2.Unbox()
	assert.Equal(1, v1)
	assert.Equal(0.1, v2)

	t3 := NewTuple3[int, float64, string](1, 0.1, "a")
	v1, v2, v3 := t3.Unbox()
	assert.Equal(1, v1)
	assert.Equal(0.1, v2)
	assert.Equal("a", v3)

	t4 := NewTuple4[int, float64, string, bool](1, 0.1, "a", true)
	v1, v2, v3, v4 := t4.Unbox()
	assert.Equal(1, v1)
	assert.Equal(0.1, v2)
	assert.Equal("a", v3)
	assert.Equal(true, v4)

	t5 := NewTuple5[int, float64, string, bool, int64](1, 0.1, "a", true, 2)
	v1, v2, v3, v4, v5 := t5.Unbox()
	assert.Equal(1, v1)
	assert.Equal(0.1, v2)
	assert.Equal("a", v3)
	assert.Equal(true, v4)
	assert.Equal(int64(2), v5)

	t6 := NewTuple6[int, float64, string, bool, int64, float32](1, 0.1, "a", true, 2, 2.2)
	v1, v2, v3, v4, v5, v6 := t6.Unbox()
	assert.Equal(1, v1)
	assert.Equal(0.1, v2)
	assert.Equal("a", v3)
	assert.Equal(true, v4)
	assert.Equal(int64(2), v5)
	assert.Equal(float32(2.2), v6)

	t7 := NewTuple7[int, float64, string, bool, int64, float32, string](1, 0.1, "a", true, 2, 2.2, "b")
	v1, v2, v3, v4, v5, v6, v7 := t7.Unbox()
	assert.Equal(1, v1)
	assert.Equal(0.1, v2)
	assert.Equal("a", v3)
	assert.Equal(true, v4)
	assert.Equal(int64(2), v5)
	assert.Equal(float32(2.2), v6)
	assert.Equal("b", v7)

	t8 := NewTuple8[int, float64, string, bool, int64, float32, string, string](1, 0.1, "a", true, 2, 2.2, "b", "c")
	v1, v2, v3, v4, v5, v6, v7, v8 := t8.Unbox()
	assert.Equal(1, v1)
	assert.Equal(0.1, v2)
	assert.Equal("a", v3)
	assert.Equal(true, v4)
	assert.Equal(int64(2), v5)
	assert.Equal(float32(2.2), v6)
	assert.Equal("b", v7)
	assert.Equal("c", v8)

	t9 := NewTuple9[int, float64, string, bool, int64, float32, string, string, map[string]int](1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1})
	v1, v2, v3, v4, v5, v6, v7, v8, v9 := t9.Unbox()
	assert.Equal(1, v1)
	assert.Equal(0.1, v2)
	assert.Equal("a", v3)
	assert.Equal(true, v4)
	assert.Equal(int64(2), v5)
	assert.Equal(float32(2.2), v6)
	assert.Equal("b", v7)
	assert.Equal("c", v8)
	assert.Equal(map[string]int{"a": 1}, v9)

	t10 := NewTuple10[int, float64, string, bool, int64, float32, string, string, map[string]int, foo](1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1}, foo{A: "a"})
	v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 := t10.Unbox()
	assert.Equal(1, v1)
	assert.Equal(0.1, v2)
	assert.Equal("a", v3)
	assert.Equal(true, v4)
	assert.Equal(int64(2), v5)
	assert.Equal(float32(2.2), v6)
	assert.Equal("b", v7)
	assert.Equal("c", v8)
	assert.Equal(map[string]int{"a": 1}, v9)
	assert.Equal(foo{A: "a"}, v10)
}

func TestZip(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestZip")

	r2 := Zip2(
		[]int{1, 2},
		[]string{"a", "b"},
	)
	assert.Equal(r2, []Tuple2[int, string]{
		{FieldA: 1, FieldB: "a"},
		{FieldA: 2, FieldB: "b"},
	})

	r3 := Zip3(
		[]int{1, 2, 3},
		[]string{"a", "b", "c"},
		[]float64{0.1, 0.2, 0.3},
	)
	assert.Equal(r3, []Tuple3[int, string, float64]{
		{FieldA: 1, FieldB: "a", FieldC: 0.1},
		{FieldA: 2, FieldB: "b", FieldC: 0.2},
		{FieldA: 3, FieldB: "c", FieldC: 0.3},
	})

	r4 := Zip4(
		[]int{1, 2, 3, 4},
		[]string{"a", "b", "c", "d"},
		[]float64{0.1, 0.2, 0.3, 0.4},
		[]bool{true, true, true, true},
	)
	assert.Equal(r4, []Tuple4[int, string, float64, bool]{
		{FieldA: 1, FieldB: "a", FieldC: 0.1, FieldD: true},
		{FieldA: 2, FieldB: "b", FieldC: 0.2, FieldD: true},
		{FieldA: 3, FieldB: "c", FieldC: 0.3, FieldD: true},
		{FieldA: 4, FieldB: "d", FieldC: 0.4, FieldD: true},
	})

	r5 := Zip5(
		[]int{1, 2, 3, 4, 5},
		[]string{"a", "b", "c", "d", "e"},
		[]float64{0.1, 0.2, 0.3, 0.4, 0.5},
		[]bool{true, true, true, true, true},
		[]int{6, 7, 8, 9, 10},
	)
	assert.Equal(r5, []Tuple5[int, string, float64, bool, int]{
		{FieldA: 1, FieldB: "a", FieldC: 0.1, FieldD: true, FieldE: 6},
		{FieldA: 2, FieldB: "b", FieldC: 0.2, FieldD: true, FieldE: 7},
		{FieldA: 3, FieldB: "c", FieldC: 0.3, FieldD: true, FieldE: 8},
		{FieldA: 4, FieldB: "d", FieldC: 0.4, FieldD: true, FieldE: 9},
		{FieldA: 5, FieldB: "e", FieldC: 0.5, FieldD: true, FieldE: 10},
	})

}

func TestUnzip(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestUnzip")

	r1, r2, r3 := Unzip3([]Tuple3[string, int, float64]{{FieldA: "a", FieldB: 1, FieldC: 0.1}, {FieldA: "b", FieldB: 2, FieldC: 0.2}})

	assert.Equal(r1, []string{"a", "b"})
	assert.Equal(r2, []int{1, 2})
	assert.Equal(r3, []float64{0.1, 0.2})
}
