# HashMap

HashMap 数据结构实现

<div STYLE="page-break-after: always;"></div>

## 源码

- [https://github.com/duke-git/lancet/blob/main/datastructure/hashmap/hashmap.go](https://github.com/duke-git/lancet/blob/main/datastructure/hashmap/hashmap.go)

<div STYLE="page-break-after: always;"></div>

## 用法

```go
import (
    hashmap "github.com/duke-git/lancet/v2/datastructure/hashmap"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

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

## API 文档

### <span id="NewHashMap">NewHashMap</span>

<p>新建默认容量（1 << 10）的HashMap指针实例</p>

<b>函数签名:</b>

```go
func NewHashMap() *HashMap
```

<b>示例:</b>

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

### <span id="NewHashMapWithCapacity">NewHashMapWithCapacity</span>

<p>新建指定容量和长度的HashMap指针实例.</p>

<b>函数签名:</b>

```go
func NewHashMapWithCapacity(size, capacity uint64) *HashMap
```

<b>示例:</b>

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

<p>在hashmap中根据key获取值</p>

<b>函数签名:</b>

```go
func (hm *HashMap) Get(key any) any
```

<b>示例:</b>

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

<p>将key-value放入hashmap中</p>

<b>函数签名:</b>

```go
func (hm *HashMap) Put(key any, value any) any
```

<b>示例:</b>

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

<p>将指定的key从hashmap中删除</p>

<b>函数签名:</b>

```go
func (hm *HashMap) Delete(key any)
```

<b>示例:</b>

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

<p>判断hashmap中是否包含指定的key</p>

<b>函数签名:</b>

```go
func (hm *HashMap) Contains(key any) bool
```

<b>示例:</b>

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

<p>迭代hashmap，对每个key和value执行iteratee函数</p>

<b>函数签名:</b>

```go
func (hm *HashMap) Iterate(iteratee func(key, value any))
```

<b>示例:</b>

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

<p>返回hashmap所有key的切片 (随机顺序)</p>

<b>函数签名:</b>

```go
func (hm *HashMap) Keys() []any
```

<b>示例:</b>

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

<p>返回hashmap所有值的切片 (随机顺序).</p>

<b>函数签名:</b>

```go
func (hm *HashMap) Values() []any
```

<b>示例:</b>

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


