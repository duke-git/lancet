# Maputil

maputil 包包括一些操作 map 的函数。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/maputil/map.go](https://github.com/duke-git/lancet/blob/main/maputil/map.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/maputil"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录:

-   [MapTo](#MapTo)
-   [ForEach](#ForEach)
-   [Filter](#Filter)
-   [FilterByKeys](#FilterByKeys)
-   [FilterByValues](#FilterByValues)
-   [OmitBy](#OmitBy)
-   [OmitByKeys](#OmitByKeys)
-   [OmitByValues](#OmitByValues)
-   [Intersect](#Intersect)
-   [Keys](#Keys)
-   [Values](#Values)
-   [KeysBy](#KeysBy)
-   [ValuesBy](#ValuesBy)
-   [MapKeys](#MapKeys)
-   [MapValues](#MapValues)
-   [Entries](#Entries)
-   [FromEntries](#FromEntries)
-   [Transform](#Transform)
-   [Merge](#Merge)
-   [Minus](#Minus)
-   [IsDisjoint](#IsDisjoint)

<div STYLE="page-break-after: always;"></div>

## API 文档:

### <span id="MapTo">MapTo</span>

<p>快速将map或者其他类型映射到结构体或者指定类型。</p>

<b>函数签名:</b>

```go
func MapTo(src any, dst any) error
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
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

    personInfo := map[string]interface{}{
        "name":  "Nothin",
        "age":   28,
        "phone": "123456789",
        "address": map[string]interface{}{
            "street": "test",
            "number": 1,
        },
    }

    var p Person
    err := MapTo(personInfo, &p)

    fmt.Println(err)
    fmt.Println(p)

    // Output:
    // <nil>
    // {Nothin 28 123456789 {test 1}}
}
```

### <span id="ForEach">ForEach</span>

<p>对map中的每对key和value执行iteratee函数</p>

<b>函数签名:</b>

```go
func ForEach[K comparable, V any](m map[K]V, iteratee func(key K, value V))
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
        "d": 4,
    }

    var sum int

    maputil.ForEach(m, func(_ string, value int) {
        sum += value
    })

    fmt.Println(sum)

    // Output:
    // 10
}
```

### <span id="Filter">Filter</span>

<p>迭代map中的每对key和value, 返回符合predicate函数的key, value。</p>

<b>函数签名:</b>

```go
func Filter[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
        "d": 4,
        "e": 5,
    }
    isEven := func(_ string, value int) bool {
        return value%2 == 0
    }

    maputil.Filter(m, func(_ string, value int) {
        sum += value
    })

    result := Filter(m, isEven)

    fmt.Println(result)

    // Output:
    // map[b:2 d:4]
}
```

### <span id="FilterByKeys">FilterByKeys</span>

<p>迭代map, 返回一个新map，其key都是给定的key值。</p>

<b>函数签名:</b>

```go
func FilterByKeys[K comparable, V any](m map[K]V, keys []K) map[K]V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
        "d": 4,
        "e": 5,
    }

    result := maputil.FilterByKeys(m, []string{"a", "b"})

    fmt.Println(result)

    // Output:
    // map[a:1 b:2]
}
```

### <span id="FilterByValues">FilterByValues</span>

<p>迭代map, 返回一个新map，其value都是给定的value值。</p>

<b>函数签名:</b>

```go
func FilterByValues[K comparable, V comparable](m map[K]V, values []V) map[K]V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
        "d": 4,
        "e": 5,
    }

    result := maputil.FilterByValues(m, []int{3, 4})

    fmt.Println(result)

    // Output:
    // map[c:3 d:4]
}
```

### <span id="OmitBy">OmitBy</span>

<p>Filter的反向操作, 迭代map中的每对key和value, 删除符合predicate函数的key, value, 返回新map。</p>

<b>函数签名:</b>

```go
func OmitBy[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
        "d": 4,
        "e": 5,
    }
    isEven := func(_ string, value int) bool {
        return value%2 == 0
    }

    result := maputil.OmitBy(m, isEven)

    fmt.Println(result)

    // Output:
    // map[a:1 c:3 e:5]
}
```

### <span id="OmitByKeys">OmitByKeys</span>

<p>FilterByKeys的反向操作, 迭代map, 返回一个新map，其key不包括给定的key值。</p>

<b>函数签名:</b>

```go
func OmitByKeys[K comparable, V any](m map[K]V, keys []K) map[K]V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
        "d": 4,
        "e": 5,
    }

    result := maputil.OmitByKeys(m, []string{"a", "b"})

    fmt.Println(result)

    // Output:
    // map[c:3 d:4 e:5]
}
```

### <span id="OmitByValues">OmitByValues</span>

<p>FilterByValues的反向操作, 迭代map, 返回一个新map，其value不包括给定的value值。</p>

<b>函数签名:</b>

```go
func OmitByValues[K comparable, V comparable](m map[K]V, values []V) map[K]V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
        "d": 4,
        "e": 5,
    }

    result := maputil.OmitByValues(m, []int{4, 5})

    fmt.Println(result)

    // Output:
    // map[a:1 b:2 c:3]
}
```

### <span id="Intersect">Intersect</span>

<p>多个map的交集操作</p>

<b>函数签名:</b>

```go
func Intersect[K comparable, V any](maps ...map[K]V) map[K]V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m1 := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }

    m2 := map[string]int{
        "a": 1,
        "b": 2,
        "c": 6,
        "d": 7,
    }

    m3 := map[string]int{
        "a": 1,
        "b": 9,
        "e": 9,
    }

    result1 := maputil.Intersect(m1)
    result2 := maputil.Intersect(m1, m2)
    result3 := maputil.Intersect(m1, m2, m3)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // map[a:1 b:2 c:3]
    // map[a:1 b:2]
    // map[a:1]
}
```

### <span id="Keys">Keys</span>

<p>返回map中所有key的切片</p>

<b>函数签名:</b>

```go
func Keys[K comparable, V any](m map[K]V) []K
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[int]string{
        1: "a",
        2: "a",
        3: "b",
        4: "c",
        5: "d",
    }

    keys := maputil.Keys(m)
    sort.Ints(keys)

    fmt.Println(keys)

    // Output:
    // [1 2 3 4 5]
}
```

### <span id="Merge">Merge</span>

<p>合并多个maps, 相同的key会被后来的key覆盖</p>

<b>函数签名:</b>

```go
func Merge[K comparable, V any](maps ...map[K]V) map[K]V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m1 := map[int]string{
        1: "a",
        2: "b",
    }
    m2 := map[int]string{
        1: "1",
        3: "2",
    }

    result := maputil.Merge(m1, m2)

    fmt.Println(result)

    // Output:
    // map[1:c 2:b 3:d]
}
```

### <span id="Minus">Minus</span>

<p>返回一个map，其中的key存在于mapA，不存在于mapB.</p>

<b>函数签名:</b>

```go
func Minus[K comparable, V any](mapA, mapB map[K]V) map[K]V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m1 := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }

    m2 := map[string]int{
        "a": 11,
        "b": 22,
        "d": 33,
    }

    result := maputil.Minus(m1, m2)

    fmt.Println(result)

    // Output:
    // map[c:3]
}
```

### <span id="Values">Values</span>

<p>返回map中所有value的切片</p>

<b>函数签名:</b>

```go
func Values[K comparable, V any](m map[K]V) []V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[int]string{
        1: "a",
        2: "a",
        3: "b",
        4: "c",
        5: "d",
    }

    values := maputil.Values(m)
    sort.Strings(values)

    // Output:
    // [a a b c d]
}
```

### <span id="KeysBy">KeysBy</span>

<p>创建一个切片，其元素是每个map的key调用mapper函数的结果。</p>

<b>函数签名:</b>

```go
func KeysBy[K comparable, V any, T any](m map[K]V, mapper func(item K) T) []T
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "sort"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[int]string{
        1: "a",
        2: "a",
        3: "b",
    }

    keys := maputil.KeysBy(m, func(n int) int {
        return n + 1
    })

    sort.Ints(keys)

    fmt.Println(keys)

    // Output:
    // [2 3 4]
}
```

### <span id="ValuesBy">ValuesBy</span>

<p>创建一个切片，其元素是每个map的value调用mapper函数的结果。</p>

<b>函数签名:</b>

```go
func ValuesBy[K comparable, V any, T any](m map[K]V, mapper func(item V) T) []T
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "sort"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[int]string{
        1: "a",
        2: "b",
        3: "c",
    }
    values := maputil.ValuesBy(m, func(v string) string {
        switch v {
        case "a":
            return "a-1"
        case "b":
            return "b-2"
        case "c":
            return "c-3"
        default:
            return ""
        }
    })

    sort.Strings(values)

    fmt.Println(values)

    // Output:
    // [a-1 b-2 c-3]
}
```

### <span id="MapKeys">MapKeys</span>

<p>操作map的每个key，然后转为新的map。</p>

<b>函数签名:</b>

```go
func MapKeys[K comparable, V any, T comparable](m map[K]V, iteratee func(key K, value V) T) map[T]V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "strconv"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[int]string{
        1: "a",
        2: "b",
        3: "c",
    }

    result := maputil.MapKeys(m, func(k int, _ string) string {
        return strconv.Itoa(k)
    })

    fmt.Println(result)

    // Output:
    // map[1:a 2:b 3:c]
}
```

### <span id="MapValues">MapValues</span>

<p>操作map的每个value，然后转为新的map。</p>

<b>函数签名:</b>

```go
func MapValues[K comparable, V any, T any](m map[K]V, iteratee func(key K, value V) T) map[K]T
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "strconv"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[int]string{
        1: "a",
        2: "b",
        3: "c",
    }

    result := maputil.MapValues(m, func(k int, v string) string {
        return v + strconv.Itoa(k)
    })

    fmt.Println(result)

    // Output:
    // map[1:a1 2:b2 3:c3]
}
```

### <span id="Entry">Entry</span>

<p>将map转换为键/值对切片。</p>

<b>函数签名:</b>

```go
type Entry[K comparable, V any] struct {
    Key   K
    Value V
}
func Entries[K comparable, V any](m map[K]V) []Entry[K, V]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "sort"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }

    result := maputil.Entries(m)

    sort.Slice(result, func(i, j int) bool {
        return result[i].Value < result[j].Value
    })

    fmt.Println(result)

    // Output:
    // [{a 1} {b 2} {c 3}]
}
```

### <span id="FromEntries">FromEntries</span>

<p>基于键/值对的切片创建map。</p>

<b>函数签名:</b>

```go
type Entry[K comparable, V any] struct {
    Key   K
    Value V
}
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    result := maputil.FromEntries([]Entry[string, int]{
        {Key: "a", Value: 1},
        {Key: "b", Value: 2},
        {Key: "c", Value: 3},
    })

    fmt.Println(result)

    // Output:
    // map[a:1 b:2 c:3]
}
```

### <span id="Transform">Transform</span>

<p>将map转换为其他类型的map。</p>

<b>函数签名:</b>

```go
func Transform[K1 comparable, V1 any, K2 comparable, V2 any](m map[K1]V1, iteratee func(key K1, value V1) (K2, V2)) map[K2]V2
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "strconv"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }

    result := Transform(m, func(k string, v int) (string, string) {
        return k, strconv.Itoa(v)
    })

    fmt.Println(result)

    // Output:
    // map[a:1 b:2 c:3]
}
```

### <span id="IsDisjoint">IsDisjoint</span>

<p>验证两个map是否具有不同的key</p>

<b>函数签名:</b>

```go
func IsDisjoint[K comparable, V any](mapA, mapB map[K]V) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/maputil"
)

func main() {
    m1 := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }

    m2 := map[string]int{
        "d": 22,
    }

    m3 := map[string]int{
        "a": 22,
    }

    result1 := maputil.IsDisjoint(m1, m2)
    result2 := maputil.IsDisjoint(m1, m3)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```
