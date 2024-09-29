# Formatter

formatter contains some functions for data formatting.

<div STYLE="page-break-after: always;"></div>

## Source:

[https://github.com/duke-git/lancet/blob/v1/formatter/formatter.go](https://github.com/duke-git/lancet/blob/v1/formatter/formatter.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/formatter"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [Comma](#Comma)
-   [Pretty](#Pretty)
-   [PrettyToWriter](#PrettyToWriter)
-   [DecimalBytes](#DecimalBytes)
-   [BinaryBytes](#BinaryBytes)
-   [ParseDecimalBytes](#ParseDecimalBytes)
-   [ParseBinaryBytes](#ParseBinaryBytes)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="Comma">Comma</span>

<p>Add comma to number by every 3 numbers from right. ahead by symbol char.
Param should be number or numberic string.</p>

<b>Signature:</b>

```go
func Comma(value interface{}, prefixSymbol string) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/formatter"
)

func main() {
    fmt.Println(formatter.Comma("12345", ""))   // "12,345"
    fmt.Println(formatter.Comma(12345.67, "¥")) // "¥12,345.67"
}
```

### <span id="Pretty">Pretty</span>

<p>Pretty data to JSON string.</p>

<b>Signature:</b>

```go
func Pretty(v interface{}) (string, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/formatter"
)

func main() {
    result1, _ := formatter.Pretty([]string{"a", "b", "c"})
    result2, _ := formatter.Pretty(map[string]int{"a": 1})

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // [
    //     "a",
    //     "b",
    //     "c"
    // ]
    // {
    //     "a": 1
    // }
}
```

### <span id="PrettyToWriter">PrettyToWriter</span>

<p>Pretty encode data to writer.</p>

<b>Signature:</b>

```go
func PrettyToWriter(v interface{}, out io.Writer) error
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/formatter"
)

func main() {
    type User struct {
        Name string `json:"name"`
        Aage uint   `json:"age"`
    }
    user := User{Name: "King", Aage: 10000}

    buf := &bytes.Buffer{}
    err := formatter.PrettyToWriter(user, buf)

    fmt.Println(buf)
    fmt.Println(err)

    // Output:
    // {
    //     "name": "King",
    //     "age": 10000
    // }
    //
    // <nil>
}
```

### <span id="DecimalBytes">DecimalBytes</span>

<p>Returns a human readable byte size under decimal standard (base 1000). The precision parameter specifies the number of digits after the decimal point, which is 4 for default.</p>

<b>Signature:</b>

```go
func DecimalBytes(size float64, precision ...int) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/formatter"
)

func main() {
    result1 := formatter.DecimalBytes(1000)
    result2 := formatter.DecimalBytes(1024)
    result3 := formatter.DecimalBytes(1234567)
    result4 := formatter.DecimalBytes(1234567, 3)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // 1KB
    // 1.024KB
    // 1.2346MB
    // 1.235MB
}
```

### <span id="BinaryBytes">BinaryBytes</span>

<p>Returns a human readable byte size under binary standard (base 1024). The precision parameter specifies the number of digits after the decimal point, which is 4 for default.</p>

<b>Signature:</b>

```go
func BinaryBytes(size float64, precision ...int) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/formatter"
)

func main() {
    result1 := formatter.BinaryBytes(1024)
    result2 := formatter.BinaryBytes(1024 * 1024)
    result3 := formatter.BinaryBytes(1234567)
    result4 := formatter.BinaryBytes(1234567, 2)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // 1KiB
    // 1MiB
    // 1.1774MiB
    // 1.18MiB
}
```

### <span id="ParseDecimalBytes">ParseDecimalBytes</span>

<p>Returns the human readable bytes size string into the amount it represents(base 1000).</p>

<b>Signature:</b>

```go
func ParseDecimalBytes(size string) (uint64, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/formatter"
)

func main() {
    result1, _ := formatter.ParseDecimalBytes("12")
    result2, _ := formatter.ParseDecimalBytes("12k")
    result3, _ := formatter.ParseDecimalBytes("12 Kb")
    result4, _ := formatter.ParseDecimalBytes("12.2 kb")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // 12
    // 12000
    // 12000
    // 12200
}
```

### <span id="ParseBinaryBytes">ParseBinaryBytes</span>

<p>Returns the human readable bytes size string into the amount it represents(base 1024).</p>

<b>Signature:</b>

```go
func ParseBinaryBytes(size string) (uint64, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/formatter"
)

func main() {
    result1, _ := formatter.ParseBinaryBytes("12")
    result2, _ := formatter.ParseBinaryBytes("12ki")
    result3, _ := formatter.ParseBinaryBytes("12 KiB")
    result4, _ := formatter.ParseBinaryBytes("12.2 kib")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    fmt.Println(result4)

    // Output:
    // 12
    // 12288
    // 12288
    // 12492
}
```
