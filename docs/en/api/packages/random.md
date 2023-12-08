# Random

Package random implements some basic functions to generate random int and string.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/random/random.go](https://github.com/duke-git/lancet/blob/main/random/random.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/random"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [RandBytes](#RandBytes)
-   [RandInt](#RandInt)
-   [RandString](#RandString)
-   [RandUpper](#RandUpper)
-   [RandLower](#RandLower)
-   [RandNumeral](#RandNumeral)
-   [RandNumeralOrLetter](#RandNumeralOrLetter)
-   [RandSymbolChar](#RandSymbolChar)
-   [UUIdV4](#UUIdV4)
-   [RandUniqueIntSlice](#RandUniqueIntSlice)
-   [RandFloat](#RandFloat)
-   [RandFloats](#RandFloats)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="RandBytes">RandBytes</span>

<p>Generate random byte slice.</p>

<b>Signature:</b>

```go
func RandBytes(length int) []byte
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/EkiLESeXf8d)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randBytes := random.RandBytes(4)
    fmt.Println(randBytes)
}
```

### <span id="RandInt">RandInt</span>

<p>Generate random int between min and max, may contain min, not max.</p>

<b>Signature:</b>

```go
func RandInt(min, max int) int
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/pXyyAAI5YxD)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    rInt := random.RandInt(1, 10)
    fmt.Println(rInt)
}
```

### <span id="RandString">RandString</span>

<p>Generate random given length string. only contains letter (a-zA-Z)</p>

<b>Signature:</b>

```go
func RandString(length int) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/W2xvRUXA7Mi)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandString(6)
    fmt.Println(randStr) //pGWsze
}
```

### <span id="RandUpper">RandUpper</span>

<p>Generate a random upper case string</p>

<b>Signature:</b>

```go
func RandUpper(length int) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/29QfOh0DVuh)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandString(6)
    fmt.Println(randStr) //PACWGF
}
```

### <span id="RandLower">RandLower</span>

<p>Generate a random lower case string</p>

<b>Signature:</b>

```go
func RandLower(length int) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/XJtZ471cmtI)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandLower(6)
    fmt.Println(randStr) //siqbew
}
```

### <span id="RandNumeral">RandNumeral</span>

<p>Generate a random numeral string</p>

<b>Signature:</b>

```go
func RandNumeral(length int) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/g4JWVpHsJcf)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandNumeral(6)
    fmt.Println(randStr) //035172
}
```

### <span id="RandNumeralOrLetter">RandNumeralOrLetter</span>

<p>generate a random numeral or letter string</p>

<b>Signature:</b>

```go
func RandNumeralOrLetter(length int) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/19CEQvpx2jD)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandNumeralOrLetter(6)
    fmt.Println(randStr) //0aW7cQ
}
```

### <span id="RandSymbolChar">RandSymbolChar</span>

<p>Generate a random symbol char of specified length. Symbol chars: !@#$%^&*()_+-=[]{}|;':\",./<>?.</p>

<b>Signature:</b>

```go
func RandSymbolChar(length int) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    randStr := random.RandSymbolChar(6)
    fmt.Println(randStr) //@#(_")
}
```

### <span id="UUIdV4">UUIdV4</span>

<p>Generate a random UUID of version 4 according to RFC 4122.</p>

<b>Signature:</b>

```go
func UUIdV4() (string, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/_Z9SFmr28ft)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    uuid, err := random.UUIdV4()
    if err != nil {
        return
    }
    fmt.Println(uuid)
}
```


### <span id="RandUniqueIntSlice">RandUniqueIntSlice</span>

<p>Generate a slice of random int of length n that do not repeat.</p>

<b>Signature:</b>

```go
func RandUniqueIntSlice(n, min, max int) []int
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/uBkRSOz73Ec)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    result := random.RandUniqueIntSlice(5, 0, 10)
    fmt.Println(result) //[0 4 7 1 5] (random)
}
```

### <span id="RandFloat">RandFloat</span>

<p>Generate random float64 number between [min, max) with specific precision.</p>

<b>Signature:</b>

```go
func RandFloat(min, max float64, precision int) float64
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    floatNumber := random.RandFloat(1.0, 5.0, 2)
    fmt.Println(floatNumber) //2.14 (random number)
}
```

### <span id="RandFloats">RandFloats</span>

<p>Generate a slice of random float64 numbers of length n that do not repeat.</p>

<b>Signature:</b>

```go
func RandFloats(n int, min, max float64, precision int) []float64
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/random"
)

func main() {
    floatNumbers := random.RandFloats(5, 1.0, 5.0, 2)
    fmt.Println(floatNumber) //[3.42 3.99 1.3 2.38 4.23] (random)
}
```