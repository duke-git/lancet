// Copyright 2023 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package tuple implements tuple data type and some operations on it.
package tuple

// Tuple2 represents a 2 elemnets tuple
type Tuple2[A any, B any] struct {
	FieldA A
	FieldB B
}

// Unbox returns values in tuple.
// Play: todo
func (t Tuple2[A, B]) Unbox() (A, B) {
	return t.FieldA, t.FieldB
}

// NewTuple2 creates a 2 elemnets tuple from a list of values.
// Play: todo
func NewTuple2[A any, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{FieldA: a, FieldB: b}
}

// Zip2 create a slice of Tuple2.
// Play: todo
func Zip2[A any, B any](a []A, b []B) []Tuple2[A, B] {
	if len(a) != len(b) {
		panic("slices must have same length")
	}

	size := len(a)
	tuples := make([]Tuple2[A, B], size)

	for i := 0; i < size; i++ {
		tuples[i] = Tuple2[A, B]{FieldA: a[i], FieldB: b[i]}
	}

	return tuples
}

// Tuple3 represents a 3 elemnets tuple
type Tuple3[A any, B any, C any] struct {
	FieldA A
	FieldB B
	FieldC C
}

// Unbox returns values in tuple.
// Play: todo
func (t Tuple3[A, B, C]) Unbox() (A, B, C) {
	return t.FieldA, t.FieldB, t.FieldC
}

// NewTuple3 creates a 3 elemnets tuple from a list of values.
// Play: todo
func NewTuple3[A any, B any, C any](a A, b B, c C) Tuple3[A, B, C] {
	return Tuple3[A, B, C]{FieldA: a, FieldB: b, FieldC: c}
}

// Zip3 create a slice of Tuple3. s
// Play: todo
func Zip3[A any, B any, C any](a []A, b []B) []Tuple2[A, B] {
	if len(a) != len(b) {
		panic("slices must have same length")
	}

	size := len(a)
	tuples := make([]Tuple2[A, B], size)

	for i := 0; i < size; i++ {
		tuples[i] = Tuple2[A, B]{FieldA: a[i], FieldB: b[i]}
	}

	return tuples
}

// Tuple4 represents a 4 elemnets tuple
type Tuple4[A any, B any, C any, D any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
}

// Unbox returns values in tuple.
// Play: todo
func (t Tuple4[A, B, C, D]) Unbox() (A, B, C, D) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD
}

// NewTuple4 creates a 4 elemnets tuple from a list of values.
// Play: todo
func NewTuple4[A any, B any, C any, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return Tuple4[A, B, C, D]{FieldA: a, FieldB: b, FieldC: c, FieldD: d}
}

// Tuple5 represents a 5 elemnets tuple
type Tuple5[A any, B any, C any, D any, E any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
}

// Unbox returns values in tuple.
// Play: todo
func (t Tuple5[A, B, C, D, E]) Unbox() (A, B, C, D, E) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE
}

// NewTuple5 creates a 5 elemnets tuple from a list of values.
// Play: todo
func NewTuple5[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return Tuple5[A, B, C, D, E]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e}
}

// Tuple6 represents a 6 elemnets tuple
type Tuple6[A any, B any, C any, D any, E any, F any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
	FieldF F
}

// Unbox returns values in tuple.
// Play: todo
func (t Tuple6[A, B, C, D, E, F]) Unbox() (A, B, C, D, E, F) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF
}

// NewTuple6 creates a 6 elemnets tuple from a list of values.
// Play: todo
func NewTuple6[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F] {
	return Tuple6[A, B, C, D, E, F]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e, FieldF: f}
}

// Tuple7 represents a 7 elemnets tuple
type Tuple7[A any, B any, C any, D any, E any, F any, G any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
	FieldF F
	FieldG G
}

// Unbox returns values in tuple.
// Play: todo
func (t Tuple7[A, B, C, D, E, F, G]) Unbox() (A, B, C, D, E, F, G) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG
}

// NewTuple7 creates a 7 elemnets tuple from a list of values.
// Play: todo
func NewTuple7[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G] {
	return Tuple7[A, B, C, D, E, F, G]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e, FieldF: f, FieldG: g}
}

// Tuple8 represents a 8 elemnets tuple
type Tuple8[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
	FieldF F
	FieldG G
	FieldH H
}

// Unbox returns values in tuple.
// Play: todo
func (t Tuple8[A, B, C, D, E, F, G, H]) Unbox() (A, B, C, D, E, F, G, H) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH
}

// NewTuple8 creates a 8 elemnets tuple from a list of values.
// Play: todo
func NewTuple8[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H] {
	return Tuple8[A, B, C, D, E, F, G, H]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e, FieldF: f, FieldG: g, FieldH: h}
}

// Tuple9 represents a 9 elemnets tuple
type Tuple9[A any, B any, C any, D any, E any, F any, G any, H any, I any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
	FieldF F
	FieldG G
	FieldH H
	FieldI I
}

// Unbox returns values in tuple.
// Play: todo
func (t Tuple9[A, B, C, D, E, F, G, H, I]) Unbox() (A, B, C, D, E, F, G, H, I) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH, t.FieldI
}

// NewTuple9 creates a 9 elemnets tuple from a list of values.
// Play: todo
func NewTuple9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I] {
	return Tuple9[A, B, C, D, E, F, G, H, I]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e, FieldF: f, FieldG: g, FieldH: h, FieldI: i}
}

// Tuple10 represents a 10 elemnets tuple
type Tuple10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
	FieldE E
	FieldF F
	FieldG G
	FieldH H
	FieldI I
	FieldJ J
}

// Unbox returns values in tuple.
// Play: todo
func (t Tuple10[A, B, C, D, E, F, G, H, I, J]) Unbox() (A, B, C, D, E, F, G, H, I, J) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH, t.FieldI, t.FieldJ
}

// NewTuple10 creates a 10 elemnets tuple from a list of values.
// Play: todo
func NewTuple10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) Tuple10[A, B, C, D, E, F, G, H, I, J] {
	return Tuple10[A, B, C, D, E, F, G, H, I, J]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e, FieldF: f, FieldG: g, FieldH: h, FieldI: i, FieldJ: j}
}

// func mustBeSameLenght()  {

// }
