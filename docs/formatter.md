# Formatter

formatter contains some functions for data formatting.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/formatter/formatter.go](https://github.com/duke-git/lancet/blob/main/formatter/formatter.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/formatter"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [Comma](#Comma)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="Comma">Comma</span>

<p>Add comma to a number value by every 3 numbers from right to left. ahead by symbol char. if value is a invalid number string like "aa", return empty string.</p>

<b>Signature:</b>

```go
func Comma[T constraints.Float | constraints.Integer | string](value T, symbol string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/formatter"
)

func main() {
    result1 := formatter.Comma("123", "")
    result2 := formatter.Comma("12345", "$")
    result3 := formatter.Comma(1234567, "￥")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 123
    // $12,345
    // ￥1,234,567
}
```
