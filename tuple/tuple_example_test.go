// Copyright 2023 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package tuple implements tuple data type and some operations on it.
package tuple

import "fmt"

func ExampleNewTuple2() {
	t := NewTuple2(1, 0.1)
	fmt.Printf("%v %v", t.FieldA, t.FieldB)
	// Output: 1 0.1
}

func ExampleTuple2_Unbox() {
	t := NewTuple2(1, 0.1)
	v1, v2 := t.Unbox()
	fmt.Printf("%v %v", v1, v2)
	// Output: 1 0.1
}

func ExampleNewTuple3() {
	t := NewTuple3(1, 0.1, "a")
	fmt.Printf("%v %v %v", t.FieldA, t.FieldB, t.FieldC)
	// Output: 1 0.1 a
}

func ExampleTuple3_Unbox() {
	t := NewTuple3(1, 0.1, "a")
	v1, v2, v3 := t.Unbox()
	fmt.Printf("%v %v %v", v1, v2, v3)
	// Output: 1 0.1 a
}

func ExampleNewTuple4() {
	t := NewTuple4(1, 0.1, "a", true)
	fmt.Printf("%v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD)
	// Output: 1 0.1 a true
}

func ExampleTuple4_Unbox() {
	t := NewTuple4(1, 0.1, "a", true)
	v1, v2, v3, v4 := t.Unbox()
	fmt.Printf("%v %v %v %v", v1, v2, v3, v4)
	// Output: 1 0.1 a true
}

func ExampleNewTuple5() {
	t := NewTuple5(1, 0.1, "a", true, 2)
	fmt.Printf("%v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE)
	// Output: 1 0.1 a true 2
}

func ExampleTuple5_Unbox() {
	t := NewTuple5(1, 0.1, "a", true, 2)
	v1, v2, v3, v4, v5 := t.Unbox()
	fmt.Printf("%v %v %v %v %v", v1, v2, v3, v4, v5)
	// Output: 1 0.1 a true 2
}

func ExampleNewTuple6() {
	t := NewTuple6(1, 0.1, "a", true, 2, 2.2)
	fmt.Printf("%v %v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF)
	// Output: 1 0.1 a true 2 2.2
}

func ExampleTuple6_Unbox() {
	t := NewTuple6(1, 0.1, "a", true, 2, 2.2)
	v1, v2, v3, v4, v5, v6 := t.Unbox()
	fmt.Printf("%v %v %v %v %v %v", v1, v2, v3, v4, v5, v6)
	// Output: 1 0.1 a true 2 2.2
}

func ExampleNewTuple7() {
	t := NewTuple7(1, 0.1, "a", true, 2, 2.2, "b")
	fmt.Printf("%v %v %v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG)
	// Output: 1 0.1 a true 2 2.2 b
}

func ExampleTuple7_Unbox() {
	t := NewTuple7(1, 0.1, "a", true, 2, 2.2, "b")
	v1, v2, v3, v4, v5, v6, v7 := t.Unbox()
	fmt.Printf("%v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7)
	// Output: 1 0.1 a true 2 2.2 b
}

func ExampleNewTuple8() {
	t := NewTuple8(1, 0.1, "a", true, 2, 2.2, "b", "c")
	fmt.Printf("%v %v %v %v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH)
	// Output: 1 0.1 a true 2 2.2 b c
}

func ExampleTuple8_Unbox() {
	t := NewTuple8(1, 0.1, "a", true, 2, 2.2, "b", "c")
	v1, v2, v3, v4, v5, v6, v7, v8 := t.Unbox()
	fmt.Printf("%v %v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7, v8)
	// Output: 1 0.1 a true 2 2.2 b c
}

func ExampleNewTuple9() {
	t := NewTuple9(1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1})
	fmt.Printf("%v %v %v %v %v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH, t.FieldI)
	// Output: 1 0.1 a true 2 2.2 b c map[a:1]
}

func ExampleTuple9_Unbox() {
	t := NewTuple9(1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1})
	v1, v2, v3, v4, v5, v6, v7, v8, v9 := t.Unbox()
	fmt.Printf("%v %v %v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7, v8, v9)
	// Output: 1 0.1 a true 2 2.2 b c map[a:1]
}
func ExampleNewTuple10() {
	type foo struct {
		A string
	}
	t := NewTuple10(1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1}, foo{A: "a"})
	fmt.Printf("%v %v %v %v %v %v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH, t.FieldI, t.FieldJ)
	// Output: 1 0.1 a true 2 2.2 b c map[a:1] {a}
}

func ExampleTuple10_Unbox() {
	type foo struct {
		A string
	}
	t := NewTuple10(1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1}, foo{A: "a"})
	v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 := t.Unbox()
	fmt.Printf("%v %v %v %v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7, v8, v9, v10)
	// Output: 1 0.1 a true 2 2.2 b c map[a:1] {a}
}
