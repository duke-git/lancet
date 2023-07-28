# Tuple

tuple包实现一个元组数据类型。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/tuple/tuple.go](https://github.com/duke-git/lancet/blob/main/tuple/tuple.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/pointer"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [Tuple2](#Tuple2)
-   [Tuple2_Unbox](#Tuple2_Unbox)
-   [Zip2](#Zip2)
-   [Unzip2](#Unzip2)
-   [Tuple3](#Tuple3)
-   [Tuple3_Unbox](#Tuple3_Unbox)
-   [Zip3](#Zip3)
-   [Unzip3](#Unzip3)
-   [Tuple4](#Tuple4)
-   [Tuple4_Unbox](#Tuple4_Unbox)
-   [Zip4](#Zip4)
-   [Unzip4](#Unzip4)
-   [Tuple5](#Tuple5)
-   [Tuple5_Unbox](#Tuple5_Unbox)
-   [Zip5](#Zip5)
-   [Unzip5](#Unzip5)
-   [Tuple6](#Tuple6)
-   [Tuple6_Unbox](#Tuple6_Unbox)
-   [Zip6](#Zip6)
-   [Unzip6](#Unzip6)
-   [Tuple7](#Tuple7)
-   [Tuple7_Unbox](#Tuple7_Unbox)
-   [Zip7](#Zip7)
-   [Unzip7](#Unzip7)
-   [Tuple8](#TTuple8uple6)
-   [Tuple8_Unbox](#Tuple8_Unbox)
-   [Zip8](#Zip8)
-   [Unzip8](#Unzip8)
-   [Tuple9](#Tuple9)
-   [Tuple9_Unbox](#Tuple9_Unbox)
-   [Zip9](#Zip9)
-   [Unzip9](#Unzip9)
-   [Tuple10](#Tuple10)
-   [Tuple10_Unbox](#Tuple10_Unbox)
-   [Zip10](#Zip10)
-   [Unzip10](#Unzip10)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="Tuple2">Tuple2</span>

<p>2元元组</p>

<b>函数签名:</b>

```go
type Tuple2[A any, B any] struct {
    FieldA A
    FieldB B
}

func NewTuple2[A any, B any](a A, b B) Tuple2[A, B]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple2(1, 0.1)
    fmt.Printf("%v %v", t.FieldA, t.FieldB)

    // Output: 1 0.1
}
```

### <span id="Tuple2_Unbox">Tuple2_Unbox</span>

<p>返回元组的字段值。</p>

<b>函数签名:</b>

```go
func (t Tuple2[A, B]) Unbox() (A, B)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple2(1, 0.1)
    v1, v2 := t.Unbox()
    fmt.Printf("%v %v", v1, v2)

    // Output: 1 0.1
}
```

### <span id="Zip2">Zip2</span>

<p>创建一个Tuple2元组切片, 其中元组的元素和传入切片元素相对应。</p>

<b>函数签名:</b>

```go
func Zip2[A any, B any](a []A, b []B) []Tuple2[A, B]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    result := tuple.Zip2([]int{1}, []float64{0.1})
    fmt.Println(result)

    // Output: [{1 0.1}]
}
```

### <span id="Unzip2">Unzip2</span>

<p>根据传入的Tuple2切片，创建一组和Tuple2元素相对应的切片。</p>

<b>函数签名:</b>

```go
func Unzip2[A any, B any](tuples []Tuple2[A, B]) ([]A, []B)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    v1, v2 := tuple.Unzip2([]tuple.Tuple2[int, float64]{{FieldA: 1, FieldB: 0.1}})

    fmt.Printf("%v %v", v1, v2)

    // Output: [1] [0.1]
}
```

### <span id="Tuple3">Tuple3</span>

<p>3元元组。</p>

<b>函数签名:</b>

```go
type Tuple3[A any, B any, C any] struct {
    FieldA A
    FieldB B
    FieldC C
}

func NewTuple3[A any, B any, C any](a A, b B c C) Tuple3[A, B, C]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple3(1, 0.1, "a")
    fmt.Printf("%v %v %v", t.FieldA, t.FieldB, t.FieldC)

    // Output: 1 0.1 a
}
```

### <span id="Tuple3_Unbox">Tuple3_Unbox</span>

<p>返回元组的字段值。</p>

<b>函数签名:</b>

```go
func (t Tuple3[A, B, C]) Unbox() (A, B, C)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple2(1, 0.1, "a")
    v1, v2, v3 := t.Unbox()
    fmt.Printf("%v %v %v", v1, v2, v3)

    // Output: 1 0.1 a
}
```

### <span id="Zip3">Zip3</span>

<p>创建一个Tuple3元组切片, 其中元组的元素和传入切片元素相对应。</p>

<b>函数签名:</b>

```go
func Zip3[A any, B any, C any](a []A, b []B, c []C) []Tuple3[A, B, C]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    result := tuple.Zip3([]int{1}, []float64{0.1}, []string{"a"})
    fmt.Println(result)

    // Output: [{1 0.1 a}]
}
```

### <span id="Unzip3">Unzip3</span>

<p>根据传入的Tuple3切片，创建一组和Tuple3元素相对应的切片。</p>

<b>函数签名:</b>

```go
func Unzip3[A any, B any, C any](tuples []Tuple3[A, B, C]) ([]A, []B, []C)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    v1, v2, v3 := tuple.Unzip3([]tuple.Tuple3[int, float64, string]{
        {FieldA: 1, FieldB: 0.1, FieldC: "a"},
    })

    fmt.Printf("%v %v %v", v1, v2, v3)

    // Output: [1] [0.1] [a]
}
```

### <span id="Tuple4">Tuple4</span>

<p>4元元组。</p>

<b>函数签名:</b>

```go
type Tuple4[A any, B any, C any, D any] struct {
    FieldA A
    FieldB B
    FieldC C
    FieldD D
}

func NewTuple4[A any, B any, C any, D any](a A, b B, c C, d D) Tuple4[A, B, C, D]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple4(1, 0.1, "a", true)
    fmt.Printf("%v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD)

    // Output: 1 0.1 a true
}
```

### <span id="Tuple4_Unbox">Tuple4_Unbox</span>

<p>返回元组的字段值。</p>

<b>函数签名:</b>

```go
func (t Tuple4[A, B, C, D]) Unbox() (A, B, C, D)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple4(1, 0.1, "a", true)
    v1, v2, v3, v4 := t.Unbox()
    fmt.Printf("%v %v %v %v", v1, v2, v3, v4)

    // Output: 1 0.1 a true
}
```

### <span id="Zip4">Zip4</span>

<p>创建一个Tuple4元组切片, 其中元组的元素和传入切片元素相对应。</p>

<b>函数签名:</b>

```go
func Zip4[A any, B any, C any, D any](a []A, b []B, c []C, d []D) []Tuple4[A, B, C, D]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    result := tuple.Zip4([]int{1}, []float64{0.1}, []string{"a"}, []bool{true})
    fmt.Println(result)

    // Output: [{1 0.1 a true}]
}
```

### <span id="Unzip4">Unzip4</span>

<p>根据传入的Tuple4切片，创建一组和Tuple4元素相对应的切片。</p>

<b>函数签名:</b>

```go
func Unzip4[A any, B any, C any, D any](tuples []Tuple4[A, B, C, D]) ([]A, []B, []C, []D)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    v1, v2, v3, v4 := tuple.Unzip4([]tuple.Tuple4[int, float64, string, bool]{
        {FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true},
    })

    fmt.Printf("%v %v %v %v", v1, v2, v3, v4)

    // Output: [1] [0.1] [a] [true]
}
```

### <span id="Tuple5">Tuple5</span>

<p>5元元组。</p>

<b>函数签名:</b>

```go
type Tuple5[A any, B any, C any, D any, E any] struct {
    FieldA A
    FieldB B
    FieldC C
    FieldD D
    FieldE E
}

func NewTuple5[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple5(1, 0.1, "a", true, 2)
    fmt.Printf("%v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE)

    // Output: 1 0.1 a true 2
}
```

### <span id="Tuple5_Unbox">Tuple5_Unbox</span>

<p>返回元组的字段值。</p>

<b>函数签名:</b>

```go
func (t Tuple5[A, B, C, D, E]) Unbox() (A, B, C, D, E)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple5(1, 0.1, "a", true, 2)
    v1, v2, v3, v4, v5 := t.Unbox()
    fmt.Printf("%v %v %v %v %v", v1, v2, v3, v4, v5)

    // Output: 1 0.1 a true 2
}
```

### <span id="Zip5">Zip5</span>

<p>创建一个Tuple5元组切片, 其中元组的元素和传入切片元素相对应。</p>

<b>函数签名:</b>

```go
func Zip5[A any, B any, C any, D any, E any](a []A, b []B, c []C, d []D, e []E) []Tuple5[A, B, C, D, E]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    result := tuple.Zip5([]int{1}, []float64{0.1}, []string{"a"}, []bool{true}, []int{2})
    fmt.Println(result)

    // Output: [{1 0.1 a true 2}]
}
```

### <span id="Unzip5">Unzip5</span>

<p>根据传入的Tuple5切片，创建一组和Tuple5元素相对应的切片。</p>

<b>函数签名:</b>

```go
func Unzip5[A any, B any, C any, D any, E any](tuples []Tuple5[A, B, C, D, E]) ([]A, []B, []C, []D, []E)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    v1, v2, v3, v4, v5 := tuple.Unzip5([]tuple.Tuple5[int, float64, string, bool, int]{
        {FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2},
    })

    fmt.Printf("%v %v %v %v %v", v1, v2, v3, v4, v5)

    // Output: [1] [0.1] [a] [true] [2]
}
```

### <span id="Tuple6">Tuple6</span>

<p>6元元组。</p>

<b>函数签名:</b>

```go
type Tuple6[A any, B any, C any, D any, E any, F any] struct {
    FieldA A
    FieldB B
    FieldC C
    FieldD D
    FieldE E
    FieldF F
}

func NewTuple6[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F) Tuple6[A, B, C, D, E, F]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple6(1, 0.1, "a", true, 2, 2.2)
    fmt.Printf("%v %v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF)

    // Output: 1 0.1 a true 2 2.2
}
```

### <span id="Tuple6_Unbox">Tuple6_Unbox</span>

<p>返回元组的字段值。</p>

<b>函数签名:</b>

```go
func (t Tuple6[A, B, C, D, E, F]) Unbox() (A, B, C, D, E, F)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple6(1, 0.1, "a", true, 2, 2.2)
    v1, v2, v3, v4, v5, v6 := t.Unbox()
    fmt.Printf("%v %v %v %v %v %v", v1, v2, v3, v4, v5, v6)

    // Output: 1 0.1 a true 2 2.2
}
```

### <span id="Zip6">Zip6</span>

<p>创建一个Tuple6元组切片, 其中元组的元素和传入切片元素相对应。</p>

<b>函数签名:</b>

```go
func Zip6[A any, B any, C any, D any, E any, F any](a []A, b []B, c []C, d []D, e []E, f []F) []Tuple6[A, B, C, D, E, F]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    result := tuple.Zip6([]int{1}, []float64{0.1}, []string{"a"}, []bool{true}, []int{2}, []float32{2.2})
    fmt.Println(result)

    // Output: [{1 0.1 a true 2 2.2}]
}
```

### <span id="Unzip6">Unzip6</span>

<p>根据传入的Tuple6切片，创建一组和Tuple6元素相对应的切片。</p>

<b>函数签名:</b>

```go
func Unzip6[A any, B any, C any, D any, E any, F any](tuples []Tuple6[A, B, C, D, E, F]) ([]A, []B, []C, []D, []E, []F)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    v1, v2, v3, v4, v5, v6 := tuple.Unzip6([]tuple.Tuple6[int, float64, string, bool, int, float32]{
        {FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2, FieldF: 2.2},
    })

    fmt.Printf("%v %v %v %v %v %v", v1, v2, v3, v4, v5, v6)

    // Output: [1] [0.1] [a] [true] [2] [2.2]
}
```

### <span id="Tuple7">Tuple7</span>

<p>7元元组。</p>

<b>函数签名:</b>

```go
type Tuple7[A any, B any, C any, D any, E any, F any, G any] struct {
    FieldA A
    FieldB B
    FieldC C
    FieldD D
    FieldE E
    FieldF F
    FieldG G
}

func NewTuple7[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G) Tuple7[A, B, C, D, E, F, G]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple7(1, 0.1, "a", true, 2, 2.2, "b")
    fmt.Printf("%v %v %v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG)

    // Output: 1 0.1 a true 2 2.2 b
}
```

### <span id="Tuple7_Unbox">Tuple7_Unbox</span>

<p>返回元组的字段值。</p>

<b>函数签名:</b>

```go
func (t Tuple7[A, B, C, D, E, F, G]) Unbox() (A, B, C, D, E, F, G)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple7(1, 0.1, "a", true, 2, 2.2, "b")
    v1, v2, v3, v4, v5, v6, v7 := t.Unbox()
    fmt.Printf("%v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7)

    // Output: 1 0.1 a true 2 2.2 b
}
```

### <span id="Zip7">Zip7</span>

<p>创建一个Tuple7元组切片, 其中元组的元素和传入切片元素相对应。</p>

<b>函数签名:</b>

```go
func Zip7[A any, B any, C any, D any, E any, F any, G any](a []A, b []B, c []C, d []D, e []E, f []F, g []G) []Tuple7[A, B, C, D, E, F, G]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    result := tuple.Zip7([]int{1}, []float64{0.1}, []string{"a"}, []bool{true}, []int{2}, []float32{2.2}, []string{"b"})
    fmt.Println(result)

    // Output: [{1 0.1 a true 2 2.2 b}]
}
```

### <span id="Unzip7">Unzip7</span>

<p>根据传入的Tuple7切片，创建一组和Tuple7元素相对应的切片。</p>

<b>函数签名:</b>

```go
func Unzip7[A any, B any, C any, D any, E any, F any, G any](tuples []Tuple7[A, B, C, D, E, F, G]) ([]A, []B, []C, []D, []E, []F, []G)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    v1, v2, v3, v4, v5, v6, v7 := tuple.Unzip7([]tuple.Tuple7[int, float64, string, bool, int, float32, string]{
        {FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2, FieldF: 2.2, FieldG: "b"},
    })

    fmt.Printf("%v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7)

    // Output: [1] [0.1] [a] [true] [2] [2.2] [b]
}
```

### <span id="Tuple8">Tuple8</span>

<p>8元元组。</p>

<b>函数签名:</b>

```go
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

func NewTuple8[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H) Tuple8[A, B, C, D, E, F, G, H]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple8(1, 0.1, "a", true, 2, 2.2, "b", "c")
    fmt.Printf("%v %v %v %v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH)

    // Output: 1 0.1 a true 2 2.2 b c
}
```

### <span id="Tuple8_Unbox">Tuple8_Unbox</span>

<p>返回元组的字段值。</p>

<b>函数签名:</b>

```go
func (t Tuple8[A, B, C, D, E, F, G, H]) Unbox() (A, B, C, D, E, F, G, H)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple8(1, 0.1, "a", true, 2, 2.2, "b", "c")
    v1, v2, v3, v4, v5, v6, v7, v8 := t.Unbox()
    fmt.Printf("%v %v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7, v8)

    // Output: 1 0.1 a true 2 2.2 b c
}
```

### <span id="Zip8">Zip8</span>

<p>创建一个Tuple8元组切片, 其中元组的元素和传入切片元素相对应。</p>

<b>函数签名:</b>

```go
func Zip8[A any, B any, C any, D any, E any, F any, G any, H any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H) []Tuple8[A, B, C, D, E, F, G, H]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    result := tuple.Zip8([]int{1}, []float64{0.1}, []string{"a"}, []bool{true}, []int{2}, []float32{2.2}, []string{"b"}, []string{"c"})
    fmt.Println(result)

    // Output: [{1 0.1 a true 2 2.2 b c}]
}
```

### <span id="Unzip8">Unzip8</span>

<p>根据传入的Tuple8切片，创建一组和Tuple8元素相对应的切片。</p>

<b>函数签名:</b>

```go
func Unzip8[A any, B any, C any, D any, E any, F any, G any, H any](tuples []Tuple8[A, B, C, D, E, F, G, H]) ([]A, []B, []C, []D, []E, []F, []G, []H)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    v1, v2, v3, v4, v5, v6, v7, v8 := tuple.Unzip8([]tuple.Tuple8[int, float64, string, bool, int, float32, string, string]{
        {FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2, FieldF: 2.2, FieldG: "b", FieldH: "c"},
    })

    fmt.Printf("%v %v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7, v8)

    // Output: [1] [0.1] [a] [true] [2] [2.2] [b] [c]
}
```

### <span id="Tuple9">Tuple9</span>

<p>9元元组。</p>

<b>函数签名:</b>

```go

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

func NewTuple9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) Tuple9[A, B, C, D, E, F, G, H, I]

```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple9(1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1})
    fmt.Printf("%v %v %v %v %v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH, t.FieldI)

    // Output: 1 0.1 a true 2 2.2 b c map[a:1]
}
```

### <span id="Tuple9_Unbox">Tuple9_Unbox</span>

<p>返回元组的字段值。</p>

<b>函数签名:</b>

```go
func (t Tuple9[A, B, C, D, E, F, G, H, I]) Unbox() (A, B, C, D, E, F, G, H, I)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    t := tuple.NewTuple9(1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1})
    v1, v2, v3, v4, v5, v6, v7, v8, v9 := t.Unbox()
    fmt.Printf("%v %v %v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7, v8, v9)

    // Output: 1 0.1 a true 2 2.2 b c map[a:1]
}
```

### <span id="Zip9">Zip9</span>

<p>创建一个Tuple9元组切片, 其中元组的元素和传入切片元素相对应。</p>

<b>函数签名:</b>

```go
func Zip9[A any, B any, C any, D any, E any, F any, G any, H any, I any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I) []Tuple9[A, B, C, D, E, F, G, H, I]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    result := tuple.Zip9([]int{1}, []float64{0.1}, []string{"a"}, []bool{true}, []int{2}, []float32{2.2}, []string{"b"}, []string{"c"}, []int64{3})
    fmt.Println(result)

    // Output: [{1 0.1 a true 2 2.2 b c 3}]
}
```

### <span id="Unzip9">Unzip9</span>

<p>根据传入的Tuple9切片，创建一组和Tuple9元素相对应的切片。</p>

<b>函数签名:</b>

```go
func Unzip9[A any, B any, C any, D any, E any, F any, G any, H any, I any](tuples []Tuple9[A, B, C, D, E, F, G, H, I]) ([]A, []B, []C, []D, []E, []F, []G, []H, []I)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    v1, v2, v3, v4, v5, v6, v7, v8, v9 := tuple.Unzip9([]tuple.Tuple9[int, float64, string, bool, int, float32, string, string, int64]{
        {FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2, FieldF: 2.2, FieldG: "b", FieldH: "c", FieldI: 3},
    })

    fmt.Printf("%v %v %v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7, v8, v9)

    // Output: [1] [0.1] [a] [true] [2] [2.2] [b] [c] [3]
}
```

### <span id="Tuple10">Tuple10</span>

<p>10元元组。</p>

<b>函数签名:</b>

```go

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

func NewTuple10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) Tuple10[A, B, C, D, E, F, G, H, I, J]

```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    type foo struct {
        A string
    }
    t := tuple.NewTuple10(1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1}, foo{A: "a"})
    fmt.Printf("%v %v %v %v %v %v %v %v %v %v", t.FieldA, t.FieldB, t.FieldC, t.FieldD, t.FieldE, t.FieldF, t.FieldG, t.FieldH, t.FieldI, t.FieldJ)

    // Output: 1 0.1 a true 2 2.2 b c map[a:1] {a}
}
```

### <span id="Tuple10_Unbox">Tuple10_Unbox</span>

<p>返回元组的字段值。</p>

<b>函数签名:</b>

```go
func (t Tuple10[A, B, C, D, E, F, G, H, I, J]) Unbox() (A, B, C, D, E, F, G, H, I, J)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    type foo struct {
        A string
    }
    t := tuple.NewTuple10(1, 0.1, "a", true, 2, 2.2, "b", "c", map[string]int{"a": 1}, foo{A: "a"})
    v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 := t.Unbox()
    fmt.Printf("%v %v %v %v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7, v8, v9, v10)

    // Output: 1 0.1 a true 2 2.2 b c map[a:1] {a}
}
```

### <span id="Zip10">Zip10</span>

<p>创建一个Tuple10元组切片, 其中元组的元素和传入切片元素相对应。</p>

<b>函数签名:</b>

```go
func Zip10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a []A, b []B, c []C, d []D, e []E, f []F, g []G, h []H, i []I, j []J) []Tuple10[A, B, C, D, E, F, G, H, I, J]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    result := tuple.Zip10([]int{1}, []float64{0.1}, []string{"a"}, []bool{true}, []int{2}, []float32{2.2}, []string{"b"}, []string{"c"}, []int64{3}, []bool{false})
    fmt.Println(result)

    // Output: [{1 0.1 a true 2 2.2 b c 3 false}]
}
```

### <span id="Unzip10">Unzip10</span>

<p>根据传入的Tuple10切片，创建一组和Tuple10元素相对应的切片。</p>

<b>函数签名:</b>

```go
func Unzip10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](tuples []Tuple10[A, B, C, D, E, F, G, H, I, J]) ([]A, []B, []C, []D, []E, []F, []G, []H, []I, []J)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/tuple"
)

func main() {
    v1, v2, v3, v4, v5, v6, v7, v8, v9, v10 := tuple.Unzip10([]tuple.Tuple10[int, float64, string, bool, int, float32, string, string, int64, bool]{
        {FieldA: 1, FieldB: 0.1, FieldC: "a", FieldD: true, FieldE: 2, FieldF: 2.2, FieldG: "b", FieldH: "c", FieldI: 3, FieldJ: false},
    })

    fmt.Printf("%v %v %v %v %v %v %v %v %v %v", v1, v2, v3, v4, v5, v6, v7, v8, v9, v10)

    // Output: [1] [0.1] [a] [true] [2] [2.2] [b] [c] [3] [false]
}
```
