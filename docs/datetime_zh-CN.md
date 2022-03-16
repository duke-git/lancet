# Datetime
datetime日期时间处理包，格式化日期，比较日期。

<div STYLE="page-break-after: always;"></div>

## 源码:

[https://github.com/duke-git/lancet/blob/v1/datetime/datetime.go](https://github.com/duke-git/lancet/blob/v1/datetime/datetime.go)

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
- [BeginOfMinute](#BeginOfMinute)
- [BeginOfHour](#BeginOfHour)
- [BeginOfDay](#BeginOfDay)
- [BeginOfWeek](#BeginOfWeek)
- [BeginOfMonth](#BeginOfMonth)
- [BeginOfYear](#BeginOfYear)

- [EndOfMinute](#EndOfMinute)
- [EndOfHour](#EndOfHour)
- [EndOfDay](#EndOfDay)
- [EndOfWeek](#EndOfWeek)
- [EndOfMonth](#EndOfMonth)
- [EndOfYear](#EndOfYear)
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

### <span id="BeginOfMinute">BeginOfMinute</span>
<p>返回指定时间的分钟开始时间</p>

<b>函数签名:</b>

```go
func BeginOfMinute(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.BeginOfMinute(td)
    fmt.Println(bm) //2022-02-15 15:48:00 +0800 CST
}
```

### <span id="BeginOfHour">BeginOfHour</span>
<p>返回指定时间的小时开始时间</p>

<b>函数签名:</b>

```go
func BeginOfHour(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.BeginOfHour(td)
    fmt.Println(bm) //2022-02-15 15:00:00 +0800 CST
}
```

### <span id="BeginOfDay">BeginOfDay</span>
<p>返回指定时间的当天开始时间</p>

<b>函数签名:</b>

```go
func BeginOfDay(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.BeginOfDay(td)
    fmt.Println(bm) //2022-02-15 00:00:00 +0800 CST
}
```



### <span id="BeginOfWeek">BeginOfWeek</span>
<p>返回指定时间的星期开始时间</p>

<b>函数签名:</b>

```go
func BeginOfWeek(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.BeginOfWeek(td)
    fmt.Println(bm) //2022-02-13 00:00:00 +0800 CST
}
```



### <span id="BeginOfMonth">BeginOfMonth</span>
<p>返回指定时间的当月开始时间</p>

<b>函数签名:</b>

```go
func BeginOfMonth(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.BeginOfMonth(td)
    fmt.Println(bm) //2022-02-01 00:00:00 +0800 CST
}
```


### <span id="BeginOfYear">BeginOfYear</span>
<p>返回指定时间的当年开始时间</p>

<b>函数签名:</b>

```go
func BeginOfYear(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.BeginOfYear(td)
    fmt.Println(bm) //2022-01-01 00:00:00 +0800 CST
}
```



### <span id="EndOfMinute">EndOfMinute</span>
<p>返回指定时间的分钟结束时间</p>

<b>函数签名:</b>

```go
func EndOfMinute(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.EndOfMinute(td)
    fmt.Println(bm) //2022-02-15 15:48:59.999999999 +0800 CST
}
```

### <span id="EndOfHour">EndOfHour</span>
<p>返回指定时间的小时结束时间</p>

<b>函数签名:</b>

```go
func EndOfHour(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.EndOfHour(td)
    fmt.Println(bm) //2022-02-15 15:59:59.999999999 +0800 CST
}
```

### <span id="EndOfDay">EndOfDay</span>
<p>返回指定时间的当天结束时间.</p>

<b>函数签名:</b>

```go
func EndOfDay(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.EndOfDay(td)
    fmt.Println(bm) //2022-02-15 23:59:59.999999999 +0800 CST
}
```



### <span id="EndOfWeek">EndOfWeek</span>
<p>返回指定时间的星期结束时间</p>

<b>函数签名:</b>

```go
func EndOfWeek(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.EndOfWeek(td)
    fmt.Println(bm) //2022-02-19 23:59:59.999999999 +0800 CST
}
```



### <span id="EndOfMonth">EndOfMonth</span>
<p>返回指定时间的月份结束时间</p>

<b>函数签名:</b>

```go
func EndOfMonth(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.EndOfMonth(td)
    fmt.Println(bm) //2022-02-28 23:59:59.999999999 +0800 CST
}
```


### <span id="EndOfYear">EndOfYear</span>
<p>返回指定时间的年份结束时间</p>

<b>函数签名:</b>

```go
func EndOfYear(t time.Time) time.Time
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
    td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
    bm := datetime.EndOfYear(td)
    fmt.Println(bm) //2022-12-31 23:59:59.999999999 +0800 CST
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



