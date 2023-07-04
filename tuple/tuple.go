// Copyright 2023 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package tuple implements tuple data type and some operations on it.
package tuple

import "github.com/duke-git/lancet/v2/mathutil"

// Tuple2 represents a 2 elemnets tuple
type Tuple2[A any, B any] struct {
	FieldA A
	FieldB B
}

// Unbox returns values in tuple.
// Play: https://go.dev/play/p/0fD1qfCVwjm
func (t Tuple2[A, B]) Unbox() (A, B) {
	return t.FieldA, t.FieldB
}

// NewTuple2 creates a 2 elemnets tuple from a list of values.
// Play: https://go.dev/play/p/3sHVqBQpLYN
func NewTuple2[A any, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{FieldA: a, FieldB: b}
}

// Zip2 create a slice of Tuple2, whose elements are correspond to the given slice elements.
// Play: https://go.dev/play/p/4ncWJJ77Xio
func Zip2[A any, B any](a []A, b []B) []Tuple2[A, B] {
	size := mathutil.Max(len(a), len(b))

	tuples := make([]Tuple2[A, B], size)

	for i := 0; i < size; i++ {
		v1, _ := getByIndex(a, i)
		v2, _ := getByIndex(b, i)

		tuples[i] = Tuple2[A, B]{FieldA: v1, FieldB: v2}
	}

	return tuples
}

// Unzip2 creates a group of slice from a slice of Tuple2.
// Play: https://go.dev/play/p/KBecr60feXb
func Unzip2[A any, B any](tuples []Tuple2[A, B]) ([]A, []B) {
	size := len(tuples)

	r1 := make([]A, size)
	r2 := make([]B, size)

	for i, t := range tuples {
		r1[i] = t.FieldA
		r2[i] = t.FieldB
	}

	return r1, r2
}

// Tuple3 represents a 3 elemnets tuple
type Tuple3[A any, B any, C any] struct {
	FieldA A
	FieldB B
	FieldC C
}

// Unbox returns values in tuple.
// Play: https://go.dev/play/p/YojLy-id1BS
func (t Tuple3[A, B, C]) Unbox() (A, B, C) {
	return t.FieldA, t.FieldB, t.FieldC
}

// NewTuple3 creates a 3 elemnets tuple from a list of values.
// Play: https://go.dev/play/p/FtH2sdCLlCf
func NewTuple3[A any, B any, C any](a A, b B, c C) Tuple3[A, B, C] {
	return Tuple3[A, B, C]{FieldA: a, FieldB: b, FieldC: c}
}

// Zip3 create a slice of Tuple3, whose elements are correspond to the given slice elements.
// Play: https://go.dev/play/p/97NgmsTILfu
func Zip3[A any, B any, C any](a []A, b []B, c []C) []Tuple3[A, B, C] {
	size := mathutil.Max(len(a), len(b), len(c))

	tuples := make([]Tuple3[A, B, C], size)

	for i := 0; i < size; i++ {
		v1, _ := getByIndex(a, i)
		v2, _ := getByIndex(b, i)
		v3, _ := getByIndex(c, i)

		tuples[i] = Tuple3[A, B, C]{FieldA: v1, FieldB: v2, FieldC: v3}
	}

	return tuples
}

// Unzip3 creates a group of slice from a slice of Tuple3.
// Play: https://go.dev/play/p/bba4cpAa7KO
func Unzip3[A any, B any, C any](tuples []Tuple3[A, B, C]) ([]A, []B, []C) {
	size := len(tuples)

	r1 := make([]A, size)
	r2 := make([]B, size)
	r3 := make([]C, size)

	for i, t := range tuples {
		r1[i] = t.FieldA
		r2[i] = t.FieldB
		r3[i] = t.FieldC
	}

	return r1, r2, r3
}

// Tuple4 represents a 4 elemnets tuple
type Tuple4[A any, B any, C any, D any] struct {
	FieldA A
	FieldB B
	FieldC C
	FieldD D
}

// Unbox returns values in tuple.
// Play: https://go.dev/play/p/ACj9YuACGgW
func (t Tuple4[A, B, C, D]) Unbox() (A, B, C, D) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD
}

// NewTuple4 creates a 4 elemnets tuple from a list of values.
// Play: https://go.dev/play/p/D2EqDz096tk
func NewTuple4[A any, B any, C any, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return Tuple4[A, B, C, D]{FieldA: a, FieldB: b, FieldC: c, FieldD: d}
}

// Zip4 create a slice of Tuple4, whose elements are correspond to the given slice elements.
// Play: https://go.dev/play/p/PEmTYVK5hL4
func Zip4[A any, B any, C any, D any](a []A, b []B, c []C, d []D) []Tuple4[A, B, C, D] {
	size := mathutil.Max(len(a), len(b), len(c), len(d))

	tuples := make([]Tuple4[A, B, C, D], size)

	for i := 0; i < size; i++ {
		v1, _ := getByIndex(a, i)
		v2, _ := getByIndex(b, i)
		v3, _ := getByIndex(c, i)
		v4, _ := getByIndex(d, i)

		tuples[i] = Tuple4[A, B, C, D]{FieldA: v1, FieldB: v2, FieldC: v3, FieldD: v4}
	}

	return tuples
}

// Unzip4 creates a group of slice from a slice of Tuple4.
// Play: https://go.dev/play/p/rb8z4gyYSRN
func Unzip4[A any, B any, C any, D any](tuples []Tuple4[A, B, C, D]) ([]A, []B, []C, []D) {
	size := len(tuples)

	r1 := make([]A, size)
	r2 := make([]B, size)
	r3 := make([]C, size)
	r4 := make([]D, size)

	for i, t := range tuples {
		r1[i] = t.FieldA
		r2[i] = t.FieldB
		r3[i] = t.FieldC
		r4[i] = t.FieldD
	}

	return r1, r2, r3, r4
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
// Play: https://go.dev/play/p/GyIyZHjCvoS
func (t Tuple5[A, B, C, D, E]) Unbox() (A, B, C, D, E) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE
}

// NewTuple5 creates a 5 elemnets tuple from a list of values.
// Play: https://go.dev/play/p/2WndmVxPg-r
func NewTuple5[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return Tuple5[A, B, C, D, E]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e}
}

// Zip5 create a slice of Tuple5, whose elements are correspond to the given slice elements.
// Play: https://go.dev/play/p/fCAAJLMfBIP
func Zip5[A any, B any, C any, D any, E any](a []A, b []B, c []C, d []D, e []E) []Tuple5[A, B, C, D, E] {
	size := mathutil.Max(len(a), len(b), len(c), len(d), len(e))

	tuples := make([]Tuple5[A, B, C, D, E], size)

	for i := 0; i < size; i++ {
		v1, _ := getByIndex(a, i)
		v2, _ := getByIndex(b, i)
		v3, _ := getByIndex(c, i)
		v4, _ := getByIndex(d, i)
		v5, _ := getByIndex(e, i)

		tuples[i] = Tuple5[A, B, C, D, E]{
			FieldA: v1, FieldB: v2, FieldC: v3,
			FieldD: v4, FieldE: v5}
	}

	return tuples
}

// Unzip5 creates a group of slice from a slice of Tuple5.
// Play: https://go.dev/play/p/gyl6vKfhqPb
func Unzip5[A any, B any, C any, D any, E any](tuples []Tuple5[A, B, C, D, E]) ([]A, []B, []C, []D, []E) {
	size := len(tuples)

	r1 := make([]A, size)
	r2 := make([]B, size)
	r3 := make([]C, size)
	r4 := make([]D, size)
	r5 := make([]E, size)

	for i, t := range tuples {
		r1[i] = t.FieldA
		r2[i] = t.FieldB
		r3[i] = t.FieldC
		r4[i] = t.FieldD
		r5[i] = t.FieldE
	}

	return r1, r2, r3, r4, r5
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
// Play: https://go.dev/play/p/FjIHV7lpxmW
func (t Tuple6[A, B, C, D, E, F]) Unbox() (A, B, C, D, E, F) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF
}

// NewTuple6 creates a 6 elemnets tuple from a list of values.
// Play: https://go.dev/play/p/VjqcCwEJZbs
func NewTuple6[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F] {
	return Tuple6[A, B, C, D, E, F]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e, FieldF: f}
}

// Zip6 create a slice of Tuple6, whose elements are correspond to the given slice elements.
// Play: https://go.dev/play/p/oWPrnUYuFHo
func Zip6[A any, B any, C any, D any, E any, F any](a []A, b []B, c []C, d []D, e []E, f []F) []Tuple6[A, B, C, D, E, F] {
	size := mathutil.Max(len(a), len(b), len(c), len(d), len(e), len(f))

	tuples := make([]Tuple6[A, B, C, D, E, F], size)

	for i := 0; i < size; i++ {
		v1, _ := getByIndex(a, i)
		v2, _ := getByIndex(b, i)
		v3, _ := getByIndex(c, i)
		v4, _ := getByIndex(d, i)
		v5, _ := getByIndex(e, i)
		v6, _ := getByIndex(f, i)

		tuples[i] = Tuple6[A, B, C, D, E, F]{
			FieldA: v1, FieldB: v2, FieldC: v3,
			FieldD: v4, FieldE: v5, FieldF: v6}
	}

	return tuples
}

// Unzip6 creates a group of slice from a slice of Tuple6.
// Play: https://go.dev/play/p/l41XFqCyh5E
func Unzip6[A any, B any, C any, D any, E any, F any](tuples []Tuple6[A, B, C, D, E, F]) ([]A, []B, []C, []D, []E, []F) {
	size := len(tuples)

	r1 := make([]A, size)
	r2 := make([]B, size)
	r3 := make([]C, size)
	r4 := make([]D, size)
	r5 := make([]E, size)
	r6 := make([]F, size)

	for i, t := range tuples {
		r1[i] = t.FieldA
		r2[i] = t.FieldB
		r3[i] = t.FieldC
		r4[i] = t.FieldD
		r5[i] = t.FieldE
		r6[i] = t.FieldF
	}

	return r1, r2, r3, r4, r5, r6
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
// Play: https://go.dev/play/p/R9I8qeDk0zs
func (t Tuple7[A, B, C, D, E, F, G]) Unbox() (A, B, C, D, E, F, G) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG
}

// NewTuple7 creates a 7 elemnets tuple from a list of values.
// Play: https://go.dev/play/p/dzAgv_Ezub9
func NewTuple7[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G] {
	return Tuple7[A, B, C, D, E, F, G]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e, FieldF: f, FieldG: g}
}

// Zip7 create a slice of Tuple7, whose elements are correspond to the given slice elements.
// Play: https://go.dev/play/p/WUJuo897Egf
func Zip7[A any, B any, C any, D any, E any, F any, G any](a []A, b []B, c []C, d []D, e []E, f []F, g []G) []Tuple7[A, B, C, D, E, F, G] {
	size := mathutil.Max(len(a), len(b), len(c), len(d), len(e), len(f), len(g))

	tuples := make([]Tuple7[A, B, C, D, E, F, G], size)

	for i := 0; i < size; i++ {
		v1, _ := getByIndex(a, i)
		v2, _ := getByIndex(b, i)
		v3, _ := getByIndex(c, i)
		v4, _ := getByIndex(d, i)
		v5, _ := getByIndex(e, i)
		v6, _ := getByIndex(f, i)
		v7, _ := getByIndex(g, i)

		tuples[i] = Tuple7[A, B, C, D, E, F, G]{
			FieldA: v1, FieldB: v2, FieldC: v3,
			FieldD: v4, FieldE: v5, FieldF: v6,
			FieldG: v7}
	}

	return tuples
}

// Unzip7 creates a group of slice from a slice of Tuple7.
// Play: https://go.dev/play/p/hws_P1Fr2j3
func Unzip7[A any, B any, C any, D any, E any, F any, G any](tuples []Tuple7[A, B, C, D, E, F, G]) ([]A, []B, []C, []D, []E, []F, []G) {
	size := len(tuples)

	r1 := make([]A, size)
	r2 := make([]B, size)
	r3 := make([]C, size)
	r4 := make([]D, size)
	r5 := make([]E, size)
	r6 := make([]F, size)
	r7 := make([]G, size)

	for i, t := range tuples {
		r1[i] = t.FieldA
		r2[i] = t.FieldB
		r3[i] = t.FieldC
		r4[i] = t.FieldD
		r5[i] = t.FieldE
		r6[i] = t.FieldF
		r7[i] = t.FieldG
	}

	return r1, r2, r3, r4, r5, r6, r7
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
// Play: https://go.dev/play/p/PRxLBBb4SMl
func (t Tuple8[A, B, C, D, E, F, G, H]) Unbox() (A, B, C, D, E, F, G, H) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH
}

// NewTuple8 creates a 8 elemnets tuple from a list of values.
// Play: https://go.dev/play/p/YA9S0rz3dRz
func NewTuple8[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H] {
	return Tuple8[A, B, C, D, E, F, G, H]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e, FieldF: f, FieldG: g, FieldH: h}
}

// Zip8 create a slice of Tuple8, whose elements are correspond to the given slice elements.
// Play: https://go.dev/play/p/8V9jWkuJfaQ
func Zip8[A any, B any, C any, D any, E any, F any, G any, H any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H) []Tuple8[A, B, C, D, E, F, G, H] {
	size := mathutil.Max(len(a), len(b), len(c), len(d), len(e), len(f), len(g), len(h))

	tuples := make([]Tuple8[A, B, C, D, E, F, G, H], size)

	for i := 0; i < size; i++ {
		v1, _ := getByIndex(a, i)
		v2, _ := getByIndex(b, i)
		v3, _ := getByIndex(c, i)
		v4, _ := getByIndex(d, i)
		v5, _ := getByIndex(e, i)
		v6, _ := getByIndex(f, i)
		v7, _ := getByIndex(g, i)
		v8, _ := getByIndex(h, i)

		tuples[i] = Tuple8[A, B, C, D, E, F, G, H]{
			FieldA: v1, FieldB: v2, FieldC: v3,
			FieldD: v4, FieldE: v5, FieldF: v6,
			FieldG: v7, FieldH: v8}
	}

	return tuples
}

// Unzip8 creates a group of slice from a slice of Tuple8.
// Play: https://go.dev/play/p/1SndOwGsZB4
func Unzip8[A any, B any, C any, D any, E any, F any, G any, H any](tuples []Tuple8[A, B, C, D, E, F, G, H]) ([]A, []B, []C, []D, []E, []F, []G, []H) {
	size := len(tuples)

	r1 := make([]A, size)
	r2 := make([]B, size)
	r3 := make([]C, size)
	r4 := make([]D, size)
	r5 := make([]E, size)
	r6 := make([]F, size)
	r7 := make([]G, size)
	r8 := make([]H, size)

	for i, t := range tuples {
		r1[i] = t.FieldA
		r2[i] = t.FieldB
		r3[i] = t.FieldC
		r4[i] = t.FieldD
		r5[i] = t.FieldE
		r6[i] = t.FieldF
		r7[i] = t.FieldG
		r8[i] = t.FieldH
	}

	return r1, r2, r3, r4, r5, r6, r7, r8
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
// Play: https://go.dev/play/p/oFJFGTAuOa8
func (t Tuple9[A, B, C, D, E, F, G, H, I]) Unbox() (A, B, C, D, E, F, G, H, I) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH, t.FieldI
}

// NewTuple9 creates a 9 elemnets tuple from a list of values.
// Play: https://go.dev/play/p/yS2NGGtZpQr
func NewTuple9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I] {
	return Tuple9[A, B, C, D, E, F, G, H, I]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e, FieldF: f, FieldG: g, FieldH: h, FieldI: i}
}

// Zip9 create a slice of Tuple9, whose elements are correspond to the given slice elements.
// Play: https://go.dev/play/p/cgsL15QYnfz
func Zip9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I) []Tuple9[A, B, C, D, E, F, G, H, I] {
	size := mathutil.Max(len(a), len(b), len(c), len(d), len(e), len(f), len(g), len(h), len(i))

	tuples := make([]Tuple9[A, B, C, D, E, F, G, H, I], size)

	for idx := 0; idx < size; idx++ {
		v1, _ := getByIndex(a, idx)
		v2, _ := getByIndex(b, idx)
		v3, _ := getByIndex(c, idx)
		v4, _ := getByIndex(d, idx)
		v5, _ := getByIndex(e, idx)
		v6, _ := getByIndex(f, idx)
		v7, _ := getByIndex(g, idx)
		v8, _ := getByIndex(h, idx)
		v9, _ := getByIndex(i, idx)

		tuples[idx] = Tuple9[A, B, C, D, E, F, G, H, I]{
			FieldA: v1, FieldB: v2, FieldC: v3,
			FieldD: v4, FieldE: v5, FieldF: v6,
			FieldG: v7, FieldH: v8, FieldI: v9}
	}

	return tuples
}

// Unzip9 creates a group of slice from a slice of Tuple9.
// Play: https://go.dev/play/p/91-BU_KURSA
func Unzip9[A any, B any, C any, D any, E any, F any, G any, H any, I any](tuples []Tuple9[A, B, C, D, E, F, G, H, I]) ([]A, []B, []C, []D, []E, []F, []G, []H, []I) {
	size := len(tuples)

	r1 := make([]A, size)
	r2 := make([]B, size)
	r3 := make([]C, size)
	r4 := make([]D, size)
	r5 := make([]E, size)
	r6 := make([]F, size)
	r7 := make([]G, size)
	r8 := make([]H, size)
	r9 := make([]I, size)

	for i, t := range tuples {
		r1[i] = t.FieldA
		r2[i] = t.FieldB
		r3[i] = t.FieldC
		r4[i] = t.FieldD
		r5[i] = t.FieldE
		r6[i] = t.FieldF
		r7[i] = t.FieldG
		r8[i] = t.FieldH
		r9[i] = t.FieldI
	}

	return r1, r2, r3, r4, r5, r6, r7, r8, r9
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
// Play: https://go.dev/play/p/qfyx3x_X0Cu
func (t Tuple10[A, B, C, D, E, F, G, H, I, J]) Unbox() (A, B, C, D, E, F, G, H, I, J) {
	return t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH, t.FieldI, t.FieldJ
}

// NewTuple10 creates a 10 elemnets tuple from a list of values.
// Play: https://go.dev/play/p/799qqZg0hUv
func NewTuple10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) Tuple10[A, B, C, D, E, F, G, H, I, J] {
	return Tuple10[A, B, C, D, E, F, G, H, I, J]{FieldA: a, FieldB: b, FieldC: c, FieldD: d, FieldE: e, FieldF: f, FieldG: g, FieldH: h, FieldI: i, FieldJ: j}
}

// Zip10 create a slice of Tuple10, whose elements are correspond to the given slice elements.
// Play: https://go.dev/play/p/YSR-2cXnrY4
func Zip10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I, j []J) []Tuple10[A, B, C, D, E, F, G, H, I, J] {
	size := mathutil.Max(len(a), len(b), len(c), len(d), len(e), len(f), len(g), len(h), len(i), len(j))

	tuples := make([]Tuple10[A, B, C, D, E, F, G, H, I, J], size)

	for idx := 0; idx < size; idx++ {
		v1, _ := getByIndex(a, idx)
		v2, _ := getByIndex(b, idx)
		v3, _ := getByIndex(c, idx)
		v4, _ := getByIndex(d, idx)
		v5, _ := getByIndex(e, idx)
		v6, _ := getByIndex(f, idx)
		v7, _ := getByIndex(g, idx)
		v8, _ := getByIndex(h, idx)
		v9, _ := getByIndex(i, idx)
		v10, _ := getByIndex(j, idx)

		tuples[idx] = Tuple10[A, B, C, D, E, F, G, H, I, J]{
			FieldA: v1, FieldB: v2, FieldC: v3,
			FieldD: v4, FieldE: v5, FieldF: v6,
			FieldG: v7, FieldH: v8, FieldI: v9,
			FieldJ: v10}
	}

	return tuples
}

// Unzip10 creates a group of slice from a slice of Tuple10.
// Play: https://go.dev/play/p/-taQB6Wfre_z
func Unzip10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](tuples []Tuple10[A, B, C, D, E, F, G, H, I, J]) ([]A, []B, []C, []D, []E, []F, []G, []H, []I, []J) {
	size := len(tuples)

	r1 := make([]A, size)
	r2 := make([]B, size)
	r3 := make([]C, size)
	r4 := make([]D, size)
	r5 := make([]E, size)
	r6 := make([]F, size)
	r7 := make([]G, size)
	r8 := make([]H, size)
	r9 := make([]I, size)
	r10 := make([]J, size)

	for i, t := range tuples {
		r1[i] = t.FieldA
		r2[i] = t.FieldB
		r3[i] = t.FieldC
		r4[i] = t.FieldD
		r5[i] = t.FieldE
		r6[i] = t.FieldF
		r7[i] = t.FieldG
		r8[i] = t.FieldH
		r9[i] = t.FieldI
		r10[i] = t.FieldJ
	}

	return r1, r2, r3, r4, r5, r6, r7, r8, r9, r10
}

func getByIndex[T any](slice []T, index int) (T, bool) {
	l := len(slice)
	if index >= l || -index > l {
		var zeroVal T
		return zeroVal, false
	}

	if index >= 0 {
		return slice[index], true
	}

	return slice[l+index], true
}
