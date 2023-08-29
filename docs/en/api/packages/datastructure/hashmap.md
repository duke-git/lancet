# HashMap

HashMap is a key value map data structure.

<div STYLE="page-break-after: always;"></div>

## Source

- [https://github.com/duke-git/lancet/blob/main/datastructure/hashmap/hashmap.go](https://github.com/duke-git/lancet/blob/main/datastructure/hashmap/hashmap.go)

<div STYLE="page-break-after: always;"></div>

## Usage

```go
import (
    hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)
```

<div STYLE="page-break-after: always;"></div>


## Index

- [NewHashMap](#NewHashMap)
- [NewHashMapWithCapacity](#NewHashMapWithCapacity)
- [Get](#Get)
- [Put](#Put)
- [Delete](#Delete)
- [Contains](#Contains)
- [Iterate](#Iterate)
- [Keys](#Keys)
- [Values](#Values)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="NewHashMap">NewHashMap</span>

<p>Make a HashMap instance with default capacity is 1 &lt&lt 10.</p>

<b>Signature:</b>

```go
func NewHashMap() *HashMap
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)

func main() {
    hm := heap.NewHashMap()
    fmt.Println(hm)
}
```

### <span id="NewHashMap">NewHashMap</span>

<p>Make a HashMap instance with given size and capacity.</p>

<b>Signature:</b>

```go
func NewHashMapWithCapacity(size, capacity uint64) *HashMap
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)

func main() {
    hm := heap.NewHashMapWithCapacity(uint64(100), uint64(1000))
    fmt.Println(hm)
}
```

### <span id="Get">Get</span>

<p>Get the value of given key in hashmap</p>

<b>Signature:</b>

```go
func (hm *HashMap) Get(key any) any
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)

func main() {
    hm := heap.NewHashMap()
    val := hm.Get("a")

    fmt.Println(val) //nil
}
```

### <span id="Put">Put</span>

<p>Put new key value in hashmap, then return value</p>

<b>Signature:</b>

```go
func (hm *HashMap) Put(key any, value any) any
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)

func main() {
    hm := heap.NewHashMap()
    hm.Put("a", 1)

    val := hm.Get("a")
    fmt.Println(val) //1
}
```



### <span id="Delete">Delete</span>

<p>Delete key-value item by given key in hashmap.</p>

<b>Signature:</b>

```go
func (hm *HashMap) Delete(key any)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)

func main() {
    hm := heap.NewHashMap()
    hm.Put("a", 1)
    val := hm.Get("a")
    fmt.Println(val) //1

    hm.Delete("a")
    val = hm.Get("a")
    fmt.Println(val) //nil
}
```



### <span id="Contains">Contains</span>

<p>Checks if given key is in hashmap or not.</p>

<b>Signature:</b>

```go
func (hm *HashMap) Contains(key any) bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)

func main() {
    hm := heap.NewHashMap()
    hm.Put("a", 1)

    fmt.Println(hm.Contains("a")) //true
    fmt.Println(hm.Contains("b")) //false
}
```



### <span id="Iterate">Iterate</span>

<p>Executes iteratee funcation for every key and value pair of hashmap.</p>

<b>Signature:</b>

```go
func (hm *HashMap) Iterate(iteratee func(key, value any))
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)

func main() {
    hm := heap.NewHashMap()
    hm.Put("a", 1)
    hm.Put("b", 2)
    hm.Put("c", 3)

    hm.Iterate(func(key, value any) {
        fmt.Println(key)
        fmt.Println(value)
    })
}
```



### <span id="Keys">Keys</span>

<p>Return a slice of the hashmap's keys (random order).</p>

<b>Signature:</b>

```go
func (hm *HashMap) Keys() []any
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)

func main() {
    hm := heap.NewHashMap()
    hm.Put("a", 1)
    hm.Put("b", 2)
    hm.Put("c", 3)

    keys := hm.Keys()
    fmt.Println(keys) //[]interface{"a", "b", "c"}
}
```


### <span id="Values">Values</span>

<p>Return a slice of the hashmap's values (random order).</p>

<b>Signature:</b>

```go
func (hm *HashMap) Values() []any
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)

func main() {
    hm := heap.NewHashMap()
    hm.Put("a", 1)
    hm.Put("b", 2)
    hm.Put("c", 3)

    values := hm.Values()
    fmt.Println(values) //[]interface{2, 1, 3}
}
```


