# Random

Package random implements some basic functions to generate random int and string.

<div STYLE="page-break-after: always;"></div>

## Source:

[https://github.com/duke-git/lancet/blob/v1/random/random.go](https://github.com/duke-git/lancet/blob/v1/random/random.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/random"
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
-   [UUIdV4](#UUIdV4)
-   [RandFloat](#RandFloat)
-   [RandFloats](#RandFloats)
-   [RandUniqueIntSlice](#RandUniqueIntSlice)
-   [RandIntSlice](#RandIntSlice)
-   [RandStringSlice](#RandStringSlice)
-   [RandBool](#RandBool)
-   [RandBoolSlice](#RandBoolSlice)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="RandBytes">RandBytes</span>

<p>Generate random byte slice.</p>

<b>Signature:</b>

```go
func RandBytes(length int) []byte
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    randStr := random.RandNumeralOrLetter(6)
    fmt.Println(randStr) //0aW7cQ
}
```

### <span id="UUIdV4">UUIdV4</span>

<p>Generate a random UUID of version 4 according to RFC 4122.</p>

<b>Signature:</b>

```go
func UUIdV4() (string, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    uuid, err := random.UUIdV4()
    if err != nil {
        return
    }
    fmt.Println(uuid)
}
```

### <span id="RandFloat">RandFloat</span>

<p>Generate a random float64 number between [min, max) with specific precision.</p>

<b>Signature:</b>

```go
func RandFloat(min, max float64, precision int) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/zbD_tuobJtr)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    floatNumber := random.RandFloat(1.0, 5.0, 2)
    fmt.Println(floatNumber) //2.14 (random number)
}
```

### <span id="RandFloats">RandFloats</span>

<p>Generate a slice of random float64 numbers of length n that do not repeat. Number range in [min, max)</p>

<b>Signature:</b>

```go
func RandFloats(length int, min, max float64, precision int) []float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/I3yndUQ-rhh)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    floatNumbers := random.RandFloats(5, 1.0, 5.0, 2)
    fmt.Println(floatNumbers) //[3.42 3.99 1.3 2.38 4.23] (random)
}
```

### <span id="RandUniqueIntSlice">RandUniqueIntSlice</span>

<p>Generate a slice of random int of length n that do not repeat.</p>

<b>Signature:</b>

```go
func RandUniqueIntSlice(n, min, max int) []int
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    result := RandUniqueIntSlice(5, 0, 10)
    fmt.Println(result) //[0 4 7 1 5] (random)
}
```

### <span id="RandIntSlice">RandIntSlice</span>

<p>Generate a slice of random int. Number range in [min, max)</p>

<b>Signature:</b>

```go
func RandIntSlice(length, min, max int) []int
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    result := random.RandIntSlice(5, 0, 10)
    fmt.Println(result) //[1 4 7 1 5] (random)
}
```

### <span id="RandStringSlice">RandStringSlice</span>

<p>Generate a slice of random string of length strLen based on charset. chartset should be one of the following: random.Numeral, random.LowwerLetters, random.UpperLetters random.Letters, random.SymbolChars, random.AllChars. or a combination of them.</p>

<b>Signature:</b>

```go
func RandStringSlice(charset string, sliceLen, strLen int) []string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    strs := random.RandStringSlice(random.Letters, 4, 6)
    fmt.Println(strs)

    // output random string slice like below:
    //[CooSMq RUFjDz FAeMPf heRyGv]
}
```

### <span id="RandBool">RandBool</span>

<p>Generate a random boolean value (true or false).</p>

<b>Signature:</b>

```go
func RandBool() bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    result := random.RandBool()
    fmt.Println(result) // true or false (random)
}
```

### <span id="RandBoolSlice">RandBoolSlice</span>

<p>Generates a random boolean slice of specified length.</p>

<b>Signature:</b>

```go
func RandBoolSlice(length int) []bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/random"
)

func main() {
    result := random.RandBoolSlice(2)
    fmt.Println(result) // [true false] (random)
}
```