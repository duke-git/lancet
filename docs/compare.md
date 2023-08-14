# Compare

Package compare provides a lightweight comparison function on any type.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/compare/compare.go](https://github.com/duke-git/lancet/blob/main/compare/compare.go)

-   [https://github.com/duke-git/lancet/blob/main/compare/compare_internal.go](https://github.com/duke-git/lancet/blob/main/compare/compare_internal.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/condition"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [Equal](#Equal)
-   [EqualValue](#EqualValue)
-   [LessThan](#LessThan)
-   [GreaterThan](#GreaterThan)
-   [LessOrEqual](#LessOrEqual)
-   [GreaterOrEqual](#GreaterOrEqual)
-   [InDelta](#InDelta)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="Equal">Equal</span>

<p>Checks if two values are equal or not. (check both type and value)</p>

<b>Signature:</b>

```go
func Equal(left, right any) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/compare"
)

func main() {
    result1 := compare.Equal(1, 1)
    result2 := compare.Equal("1", "1")
    result3 := compare.Equal([]int{1, 2, 3}, []int{1, 2, 3})
    result4 := compare.Equal(map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 2: "b"})

    result5 := compare.Equal(1, "1")
    result6 := compare.Equal(1, int64(1))
    result7 := compare.Equal([]int{1, 2}, []int{1, 2, 3})

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)
    fmt.Println(result7)

    // Output:
    // true
    // true
    // true
    // true
    // false
    // false
    // false
}
```

### <span id="EqualValue">EqualValue</span>

<p>Checks if two values are equal or not. (check value only)</p>

<b>Signature:</b>

```go
func EqualValue(left, right any) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/compare"
)

func main() {
    result1 := compare.EqualValue(1, 1)
    result2 := compare.EqualValue(int(1), int64(1))
    result3 := compare.EqualValue(1, "1")
    result4 := compare.EqualValue(1, "2")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // true
    // true
    // true
    // false
}
```

### <span id="LessThan">LessThan</span>

<p>Checks if value `left` less than value `right`.</p>

<b>Signature:</b>

```go
func LessThan(left, right any) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/compare"
)

func main() {
    result1 := compare.LessThan(1, 2)
    result2 := compare.LessThan(1.1, 2.2)
    result3 := compare.LessThan("a", "b")

    time1 := time.Now()
    time2 := time1.Add(time.Second)
    result4 := compare.LessThan(time1, time2)

    result5 := compare.LessThan(2, 1)
    result6 := compare.LessThan(1, int64(2))

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)

    // Output:
    // true
    // true
    // true
    // true
    // false
    // false
}
```

### <span id="GreaterThan">GreaterThan</span>

<p>Checks if value `left` greater than value `right`.</p>

<b>Signature:</b>

```go
func GreaterThan(left, right any) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/compare"
)

func main() {
    result1 := compare.GreaterThan(2, 1)
    result2 := compare.GreaterThan(2.2, 1.1)
    result3 := compare.GreaterThan("b", "a")

    time1 := time.Now()
    time2 := time1.Add(time.Second)
    result4 := compare.GreaterThan(time2, time1)

    result5 := compare.GreaterThan(1, 2)
    result6 := compare.GreaterThan(int64(2), 1)
    result7 := compare.GreaterThan("b", "c")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)
    fmt.Println(result7)

    // Output:
    // true
    // true
    // true
    // true
    // false
    // false
    // false
}
```

### <span id="LessOrEqual">LessOrEqual</span>

<p>Checks if value `left` less than or equal than value `right`.</p>

<b>Signature:</b>

```go
func LessOrEqual(left, right any) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/compare"
)

func main() {
    result1 := compare.LessOrEqual(1, 1)
    result2 := compare.LessOrEqual(1.1, 2.2)
    result3 := compare.LessOrEqual("a", "b")

    time1 := time.Now()
    time2 := time1.Add(time.Second)
    result4 := compare.LessOrEqual(time1, time2)

    result5 := compare.LessOrEqual(2, 1)
    result6 := compare.LessOrEqual(1, int64(2))

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)

    // Output:
    // true
    // true
    // true
    // true
    // false
    // false
}
```

### <span id="GreaterOrEqual">GreaterOrEqual</span>

<p>Checks if value `left` less greater or equal than value `right`.</p>

<b>Signature:</b>

```go
func GreaterOrEqual(left, right any) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/compare"
)

func main() {
    result1 := compare.GreaterOrEqual(1, 1)
    result2 := compare.GreaterOrEqual(2.2, 1.1)
    result3 := compare.GreaterOrEqual("b", "b")

    time1 := time.Now()
    time2 := time1.Add(time.Second)
    result4 := compare.GreaterOrEqual(time2, time1)

    result5 := compare.GreaterOrEqual(1, 2)
    result6 := compare.GreaterOrEqual(int64(2), 1)
    result7 := compare.GreaterOrEqual("b", "c")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)
    fmt.Println(result7)

    // Output:
    // true
    // true
    // true
    // true
    // false
    // false
    // false
}
```

### <span id="InDelta">InDelta</span>

<p>Checks if two values are equal or not within a delta.</p>

<b>Signature:</b>

```go
func InDelta[T constraints.Integer | constraints.Float](left, right T, delta float64) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/compare"
)

func main() {
    result1 := InDelta(1, 1, 0)
    result2 := InDelta(1, 2, 0)

    result3 := InDelta(2.0/3.0, 0.66667, 0.001)
    result4 := InDelta(2.0/3.0, 0.0, 0.001)

    result5 := InDelta(float64(74.96)-float64(20.48), 54.48, 0)
    result6 := InDelta(float64(74.96)-float64(20.48), 54.48, 1e-14)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)
    fmt.Println(result5)
    fmt.Println(result6)

    // Output:
    // true
    // false
    // true
    // false
    // false
    // true
}
```
