# Maputil

Package maputil includes some functions to manipulate map.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/maputil/map.go](https://github.com/duke-git/lancet/blob/main/maputil/map.go)

<div STYLE="page-break-after: always;"></div>

## Example:

```go
import (
    "github.com/duke-git/lancet/v2/maputil"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [ForEach](#ForEach)
-   [Filter](#Filter)
-   [Intersect](#Intersect)
-   [Keys](#Keys)
-   [Merge](#Merge)
-   [Minus](#Minus)
-   [Values](#Values)
-   [IsDisjoint](#IsDisjoint)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="ForEach">ForEach</span>

<p>Executes iteratee funcation for every key and value pair in map.</p>

<b>Signature:</b>

```go
func ForEach[K comparable, V any](m map[K]V, iteratee func(key K, value V))
```

<b>Example:</b>

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

<p>Iterates over map, return a new map contains all key and value pairs pass the predicate function.</p>

<b>Signature:</b>

```go
func Filter[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V
```

<b>Example:</b>

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
    
    result := maputil.Filter(m, isEven)

    fmt.Println(result)

    // Output:
    // map[b:2 d:4]
}
```

### <span id="Intersect">Intersect</span>

<p>Iterates over maps, return a new map of key and value pairs in all given maps.</p>

<b>Signature:</b>

```go
func Intersect[K comparable, V any](maps ...map[K]V) map[K]V
```

<b>Example:</b>

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

<p>Returns a slice of the map's keys.</p>

<b>Signature:</b>

```go
func Keys[K comparable, V any](m map[K]V) []K
```

<b>Example:</b>

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

<p>Merge maps, next key will overwrite previous key.</p>

<b>Signature:</b>

```go
func Merge[K comparable, V any](maps ...map[K]V) map[K]V
```

<b>Example:</b>

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

<p>Creates an map of whose key in mapA but not in mapB.</p>

<b>Signature:</b>

```go
func Minus[K comparable, V any](mapA, mapB map[K]V) map[K]V
```

<b>Example:</b>

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

<p>Returns a slice of the map's values.</p>

<b>Signature:</b>

```go
func Values[K comparable, V any](m map[K]V) []V
```

<b>Example:</b>

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

    fmt.Println(values)

    // Output:
    // [a a b c d]
}
```

### <span id="IsDisjoint">IsDisjoint</span>

<p>Checks two maps are disjoint if they have no keys in common</p>

<b>Signature:</b>

```go
func IsDisjoint[K comparable, V any](mapA, mapB map[K]V) bool
```

<b>Example:</b>

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
