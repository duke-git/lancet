---
outline: deep
---

# Installation

1. <b>For users who use go1.18 and above, it is recommended to install lancet v2.x.x. Cause in v2.x.x all functions was rewriten with generics of go1.18.</b>

```go
go get github.com/duke-git/lancet/v2 // will install latest version of v2.x.x
```

2. <b>For users who use version below go1.18, you should install v1.x.x. The latest of v1.x.x is v1.4.5. </b>

```go
go get github.com/duke-git/lancet // below go1.18, install latest version of v1.x.x
```


## Usage

Lancet organizes the code into package structure, and you need to import the corresponding package name when use it. For example, if you use string-related functions, just import the strutil package like below:

```go
import "github.com/duke-git/lancet/v2/strutil"
```

## Example

Here takes the string function `Reverse` (reverse order string) as an example, and the strutil package needs to be imported.

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/strutil"
)

func main() {
    s := "hello"
    rs := strutil.Reverse(s)
    fmt.Println(rs) //olleh
}
```


## More

Check out the [API](https://www.golancet.cn/en/api/overview.html) for details.
