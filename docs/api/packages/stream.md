# Stream

Stream 流，该包仅验证简单 stream 实现，功能有限。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/stream/stream.go](https://github.com/duke-git/lancet/blob/main/stream/stream.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/stream"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [Of](#Of)
-   [FromSlice](#FromSlice)
-   [FromChannel](#FromChannel)
-   [FromRange](#FromRange)
-   [Generate](#Generate)
-   [Concat](#Concat)
-   [Distinct](#Distinct)
-   [Filter](#Filter)
-   [Map](#Map)
-   [Peek](#Peek)
-   [Skip](#Skip)
-   [Limit](#Limit)
-   [Reverse](#Reverse)
-   [Range](#Range)
-   [Sorted](#Sorted)
-   [ForEach](#ForEach)
-   [Reduce](#Reduce)
-   [FindFirst](#FindFirst)
-   [FindLast](#FindLast)
-   [Max](#Max)
-   [Min](#Min)
-   [AllMatch](#AllMatch)
-   [AnyMatch](#AnyMatch)
-   [NoneMatch](#NoneMatch)
-   [Count](#Count)
-   [ToSlice](#ToSlice)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="Of">Of</span>

<p>创建元素为指定值的stream。</p>

<b>函数签名:</b>

```go
func Of[T any](elems ...T) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/jI6_iZZuVFE)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    s := stream.Of(1, 2, 3)

    data := s.ToSlice()

    fmt.Println(data)

    // Output:
    // [1 2 3]
}
```

### <span id="FromSlice">FromSlice</span>

<p>从切片创建stream。</p>

<b>函数签名:</b>

```go
func FromSlice[T any](source []T) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/wywTO0XZtI4)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    s := stream.FromSlice([]int{1, 2, 3})

    data := s.ToSlice()

    fmt.Println(data)

    // Output:
    // [1 2 3]
}
```

### <span id="FromChannel">FromChannel</span>

<p>从通道创建stream。</p>

<b>函数签名:</b>

```go
func FromChannel[T any](source <-chan T) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/9TZYugGMhXZ)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    ch := make(chan int)
    go func() {
        for i := 1; i < 4; i++ {
            ch <- i
        }
        close(ch)
    }()

    s := stream.FromChannel(ch)

    data := s.ToSlice()

    fmt.Println(data)

    // Output:
    // [1 2 3]
}
```

### <span id="FromRange">FromRange</span>

<p>指定一个范围创建stream, 范围两端点值都包括在内。</p>

<b>函数签名:</b>

```go
func FromRange[T constraints.Integer | constraints.Float](start, end, step T) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/9Ex1-zcg-B-)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    s := stream.FromRange(1, 5, 1)

    data := s.ToSlice()
    fmt.Println(data)

    // Output:
    // [1 2 3 4 5]
}
```

### <span id="Generate">Generate</span>

<p>创建一个stream，其中每个元素都由提供的生成器函数生成</p>

<b>函数签名:</b>

```go
func Generate[T any](generator func() func() (item T, ok bool)) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/rkOWL1yA3j9)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    n := 0
    max := 4

    generator := func() func() (int, bool) {
        return func() (int, bool) {
            n++
            return n, n < max
        }
    }

    s := stream.Generate(generator)

    data := s.ToSlice()

    fmt.Println(data)

    // Output:
    // [1 2 3]
}
```

### <span id="Concat">Concat</span>

<p>创建一个延迟连接stream，其元素是第一个stream的所有元素，后跟第二个stream的全部元素。</p>

<b>函数签名:</b>

```go
func Concat[T any](a, b stream[T]) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/HM4OlYk_OUC)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    s1 := stream.FromSlice([]int{1, 2, 3})
    s2 := stream.FromSlice([]int{4, 5, 6})

    s := Concat(s1, s2)

    data := s.ToSlice()

    fmt.Println(data)

    // Output:
    // [1 2 3 4 5 6]
}
```

### <span id="Distinct">Distinct</span>

<p>创建并返回一个stream，用于删除重复的项。 <b>支持链式操作</b></p>

<b>函数签名:</b>

```go
func (s stream[T]) Distinct() stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/eGkOSrm64cB)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 2, 3, 3, 3})
    distinct := original.Distinct()

    data1 := original.ToSlice()
    data2 := distinct.ToSlice()

    fmt.Println(data1)
    fmt.Println(data2)

    // Output:
    // [1 2 2 3 3 3]
    // [1 2 3]
}
```

### <span id="Filter">Filter</span>

<p>返回一个通过判定函数的stream <b>支持链式操作</b></p>

<b>函数签名:</b>

```go
func (s stream[T]) Filter(predicate func(item T) bool) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/MFlSANo-buc)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3, 4, 5})

    isEven := func(n int) bool {
        return n%2 == 0
    }

    even := original.Filter(isEven)

    fmt.Println(even.ToSlice())

    // Output:
    // [2 4]
}
```

### <span id="Map">Map</span>

<p>返回一个stream，该stream由将给定函数应用于源stream元素的元素组成。<b>支持链式操作</b></p>

<b>函数签名:</b>

```go
func (s stream[T]) Map(mapper func(item T) T) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/OtNQUImdYko)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3})

    addOne := func(n int) int {
        return n + 1
    }

    increament := original.Map(addOne)

    fmt.Println(increament.ToSlice())

    // Output:
    // [2 3 4]
}
```

### <span id="Peek">Peek</span>

<p>返回一个由源stream的元素组成的stream，并在从生成的stream中消耗元素时对每个元素执行所提供的操作。 <b>支持链式操作</b></p>

<b>函数签名:</b>

```go
func (s stream[T]) Peek(consumer func(item T)) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/u1VNzHs6cb2)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3})

    data := []string{}
    peekStream := original.Peek(func(n int) {
        data = append(data, fmt.Sprint("value", n))
    })

    fmt.Println(original.ToSlice())
    fmt.Println(peekStream.ToSlice())
    fmt.Println(data)

    // Output:
    // [1 2 3]
    // [1 2 3]
    // [value1 value2 value3]
}
```

### <span id="Skip">Skip</span>

<p>在丢弃stream的前n个元素后，返回由源stream的其余元素组成的stream。如果此stream包含的元素少于n个，则将返回一个空stream。<b>支持链式操作</b></p>

<b>函数签名:</b>

```go
func (s stream[T]) Skip(n int) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/fNdHbqjahum)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3, 4})

    s1 := original.Skip(-1)
    s2 := original.Skip(0)
    s3 := original.Skip(1)
    s4 := original.Skip(5)

    fmt.Println(s1.ToSlice())
    fmt.Println(s2.ToSlice())
    fmt.Println(s3.ToSlice())
    fmt.Println(s4.ToSlice())

    // Output:
    // [1 2 3 4]
    // [1 2 3 4]
    // [2 3 4]
    // []
}
```

### <span id="Limit">Limit</span>

<p>返回由源stream的元素组成的stream，该stream被截断为长度不超过maxSize。<b>支持链式操作</b></p>

<b>函数签名:</b>

```go
func (s stream[T]) Limit(maxSize int) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/qsO4aniDcGf)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3, 4})

    s1 := original.Limit(-1)
    s2 := original.Limit(0)
    s3 := original.Limit(1)
    s4 := original.Limit(5)

    fmt.Println(s1.ToSlice())
    fmt.Println(s2.ToSlice())
    fmt.Println(s3.ToSlice())
    fmt.Println(s4.ToSlice())

    // Output:
    // []
    // []
    // [1]
    // [1 2 3 4]
}
```

### <span id="Reverse">Reverse</span>

<p>返回元素与源stream的顺序相反的stream。<b>支持链式操作</b></p>

<b>函数签名:</b>

```go
func (s stream[T]) Reverse() stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/A8_zkJnLHm4)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3})

    reverse := original.Reverse()

    fmt.Println(reverse.ToSlice())

    // Output:
    // [3 2 1]
}
```

### <span id="Range">Range</span>

<p>返回一个stream，该stream的元素在从源stream的开始（包含）到结束（排除）的范围内。<b>支持链式操作</b></p>

<b>函数签名:</b>

```go
func (s stream[T]) Range(start, end int) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/indZY5V2f4j)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3})

    s1 := original.Range(0, 0)
    s2 := original.Range(0, 1)
    s3 := original.Range(0, 3)
    s4 := original.Range(1, 2)

    fmt.Println(s1.ToSlice())
    fmt.Println(s2.ToSlice())
    fmt.Println(s3.ToSlice())
    fmt.Println(s4.ToSlice())

    // Output:
    // []
    // [1]
    // [1 2 3]
    // [2]
}
```

### <span id="Sorted">Sorted</span>

<p>返回一个stream，该stream由源stream的元素组成，并根据提供的less函数进行排序。<b>支持链式操作</b></p>

<b>函数签名:</b>

```go
func (s stream[T]) Sorted(less func(a, b T) bool) stream[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/XXtng5uonFj)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{4, 2, 1, 3})

    sorted := original.Sorted(func(a, b int) bool { return a < b })

    fmt.Println(original.ToSlice())
    fmt.Println(sorted.ToSlice())

    // Output:
    // [4 2 1 3]
    // [1 2 3 4]
}
```

### <span id="ForEach">ForEach</span>

<p>对stream的每个元素执行一个操作。</p>

<b>函数签名:</b>

```go
func (s stream[T]) ForEach(action func(item T))
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Dsm0fPqcidk)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3})

    result := 0
    original.ForEach(func(item int) {
        result += item
    })

    fmt.Println(result)

    // Output:
    // 6
}
```

### <span id="Reduce">Reduce</span>

<p>使用关联累加函数对stream的元素执行reduce操作，并reduce操作结果（如果有）。</p>

<b>函数签名:</b>

```go
func (s stream[T]) Reduce(initial T, accumulator func(a, b T) T) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/6uzZjq_DJLU)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3})

    result := original.Reduce(0, func(a, b int) int {
        return a + b
    })

    fmt.Println(result)

    // Output:
    // 6
}
```

### <span id="FindFirst">FindFirst</span>

<p>返回此stream的第一个元素和true，如果stream为空，则返回零值和false。</p>

<b>函数签名:</b>

```go
func (s stream[T]) FindFirst() (T, bool)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/9xEf0-6C1e3)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3})

    result, ok := original.FindFirst()

    fmt.Println(result)
    fmt.Println(ok)

    // Output:
    // 1
    // true
}
```

### <span id="FindLast">FindLast</span>

<p>返回此stream最后一个元素和true，如果stream为空，则返回零值和false。</p>

<b>函数签名:</b>

```go
func (s stream[T]) FindLast() (T, bool)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/WZD2rDAW-2h)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{3, 2, 1})

    result, ok := original.FindLast()

    fmt.Println(result)
    fmt.Println(ok)

    // Output:
    // 1
    // true
}
```

### <span id="Max">Max</span>

<p>根据提供的less函数返回stream的最大元素。less 函数: a > b</p>

<b>函数签名:</b>

```go
func (s stream[T]) Max(less func(a, b T) bool) (T, bool)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/fm-1KOPtGzn)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{4, 2, 1, 3})

    max, ok := original.Max(func(a, b int) bool { return a > b })

    fmt.Println(max)
    fmt.Println(ok)

    // Output:
    // 4
    // true
}
```

### <span id="Min">Min</span>

<p>根据提供的less函数返回stream的最小元素。less函数: a &lt b</p>

<b>函数签名:</b>

```go
func (s stream[T]) Min(less func(a, b T) bool) (T, bool)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/vZfIDgGNRe_0)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{4, 2, 1, 3})

    min, ok := original.Min(func(a, b int) bool { return a < b })

    fmt.Println(min)
    fmt.Println(ok)

    // Output:
    // 1
    // true
}
```

### <span id="AllMatch">AllMatch</span>

<p>判断stream的所有元素是否全部匹配指定判定函数。</p>

<b>函数签名:</b>

```go
func (s stream[T]) AllMatch(predicate func(item T) bool) bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/V5TBpVRs-Cx)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3})

    result1 := original.AllMatch(func(item int) bool {
        return item > 0
    })

    result2 := original.AllMatch(func(item int) bool {
        return item > 1
    })

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="AnyMatch">AnyMatch</span>

<p>判断stream是否包含匹配指定判定函数的元素。</p>

<b>函数签名:</b>

```go
func (s stream[T]) AnyMatch(predicate func(item T) bool) bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/PTCnWn4OxSn)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3})

    result1 := original.AnyMatch(func(item int) bool {
        return item > 1
    })

    result2 := original.AnyMatch(func(item int) bool {
        return item > 3
    })

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="NoneMatch">NoneMatch</span>

<p>判断stream的元素是否全部不匹配指定的判定函数。</p>

<b>函数签名:</b>

```go
func (s stream[T]) NoneMatch(predicate func(item T) bool) bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/iWS64pL1oo3)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    original := stream.FromSlice([]int{1, 2, 3})

    result1 := original.NoneMatch(func(item int) bool {
        return item > 3
    })

    result2 := original.NoneMatch(func(item int) bool {
        return item > 1
    })

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="Count">Count</span>

<p>返回stream中元素的数量。</p>

<b>函数签名:</b>

```go
func (s stream[T]) Count() int
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/r3koY6y_Xo-)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    s1 := stream.FromSlice([]int{1, 2, 3})
    s2 := stream.FromSlice([]int{})

    fmt.Println(s1.Count())
    fmt.Println(s2.Count())

    // Output:
    // 3
    // 0
}
```

### <span id="ToSlice">ToSlice</span>

<p>返回stream中的元素切片。</p>

<b>函数签名:</b>

```go
func (s stream[T]) ToSlice() []T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/jI6_iZZuVFE)</span></b>

```go
import (
    "fmt"
    "github.com/duke-git/lancet/v2/stream"
)

func main() {
    s := stream.Of(1, 2, 3)

    data := s.ToSlice()

    fmt.Println(data)

    // Output:
    // [1 2 3]
}
```
