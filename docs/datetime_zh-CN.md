# Datetime
datetime日期时间处理包，格式化日期，比较日期。

<div STYLE="page-break-after: always;"></div>

## 源码:

[https://github.com/duke-git/lancet/blob/main/datetime/datetime.go](https://github.com/duke-git/lancet/blob/main/datetime/datetime.go)

<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/datetime"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [AddDay](#AddDay)
- [AddHour](#AddHour)
- [AddMinute](#AddMinute)
- [GetNowDate](#GetNowDate)
- [GetNowTime](#GetNowTime)
- [GetNowDateTime](#GetNowDateTime)
- [GetZeroHourTimestamp](#GetZeroHourTimestamp)
- [GetNightTimestamp](#GetNightTimestamp)

- [FormatTimeToStr](#FormatTimeToStr)
- [FormatStrToTime](#FormatStrToTime)

<div STYLE="page-break-after: always;"></div>

## 文档

## 注:
1. 方法FormatTimeToStr和FormatStrToTime中的format参数值需要传以下类型之一：
- yyyy-mm-dd hh:mm:ss
- yyyy-mm-dd hh:mm
- yyyy-mm-dd hh
- yyyy-mm-dd
- yyyy-mm
- mm-dd
- dd-mm-yy hh:mm:ss
- yyyy/mm/dd hh:mm:ss
- yyyy/mm/dd hh:mm
- yyyy-mm-dd hh
- yyyy/mm/dd
- yyyy/mm
- mm/dd
- dd/mm/yy hh:mm:ss
- yyyy
- mm
- hh:mm:ss
- mm:ss


### <span id="AddDay">AddDay</span>
<p>将日期加/减天数</p>

<b>函数签名:</b>

```go
func AddDay(t time.Time, day int64) time.Time
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/datetime"
)

func main() {
    now := time.Now()
	after2Days := datetime.AddDay(now, 2)
	before2Days := datetime.AddDay(now, -2)

    fmt.Println(after2Days, before2Days)
}
```


### <span id="AddHour">AddHour</span>
<p>将日期加/减小时数</p>

<b>函数签名:</b>

```go
func AddHour(t time.Time, hour int64) time.Time
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/datetime"
)

func main() {
    now := time.Now()
    after2Hours := datetime.AddHour(now, 2)
    before2Hours := datetime.AddHour(now, -2)

    fmt.Println(after2Hours, after2Hours)
}
```

### <span id="AddMinute">AddMinute</span>
<p>将日期加/减分钟数</p>

<b>函数签名:</b>

```go
func AddMinute(t time.Time, minute int64) time.Time
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/datetime"
)

func main() {
    now := time.Now()
    after2Minute := datetime.AddMinute(now, 2)
    before2Minute := datetime.AddMinute(now, -2)

    fmt.Println(after2Minute, before2Minute)
}
```

### <span id="GetNowDate">GetNowDate</span>
<p>获取当天日期，返回格式：yyyy-mm-dd</p>

<b>函数签名:</b>

```go
func GetNowDate() string
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/datetime"
)

func main() {
    now := time.Now()
	currentDate := datetime.GetNowDate()
    fmt.Println(currentDate) // 2022-01-28
}
```


### <span id="GetNowTime">GetNowTime</span>
<p>获取当时时间，返回格式：hh:mm:ss</p>

<b>函数签名:</b>

```go
func GetNowTime() string
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/datetime"
)

func main() {
    now := time.Now()
    currentTime := datetime.GetNowTime()
    fmt.Println(currentDate) // 15:57:33
}
```


### <span id="GetNowDateTime">GetNowDateTime</span>
<p>获取当时日期和时间，返回格式：yyyy-mm-dd hh:mm:ss.</p>

<b>函数签名:</b>

```go
func GetNowDateTime() string
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/datetime"
)

func main() {
    now := time.Now()
    current := datetime.GetNowDateTime()
    fmt.Println(current) // 2022-01-28 15:59:33
}
```


### <span id="GetZeroHourTimestamp">GetZeroHourTimestamp</span>
<p>获取零时时间戳(timestamp of 00:00).</p>

<b>函数签名:</b>

```go
func GetZeroHourTimestamp() int64
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/datetime"
)

func main() {
    now := time.Now()
    zeroTime := datetime.GetZeroHourTimestamp()
    fmt.Println(zeroTime) // 1643299200
}
```


### <span id="GetNightTimestamp">GetNightTimestamp</span>
<p>获取午夜时间戳(timestamp of 23:59).</p>

<b>函数签名:</b>

```go
func GetNightTimestamp() int64
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/datetime"
)

func main() {
    now := time.Now()
    nightTime := datetime.GetNightTimestamp()
    fmt.Println(nightTime) // 1643385599
}
```

### <span id="FormatTimeToStr">FormatTimeToStr</span>
<p>将日期格式化成字符串，`format` 参数格式参考注<sup>1</sup></p>

<b>函数签名:</b>

```go
func FormatTimeToStr(t time.Time, format string) string
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/datetime"
)

func main() {
    now := time.Now()
    timeStr := datetime.FormatTimeToStr(now, "yyyy/mm/dd hh:mm:ss")
    fmt.Println(timeStr) //2022/01/28 16:07:44
}
```


### <span id="FormatStrToTime">FormatStrToTime</span>
<p>将字符串格式化成时间，`format` 参数格式参考注<sup>1</sup></p>

<b>函数签名:</b>

```go
func FormatStrToTime(str, format string) (time.Time, error)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datetime"
)

func main() {
    time := datetime.FormatStrToTime("2006-01-02 15:04:05", "yyyy/mm/dd hh:mm:ss")
    fmt.Println(time)
}
```



