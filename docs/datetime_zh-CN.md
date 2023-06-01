# Datetime

datetime 日期时间处理包，格式化日期，比较日期。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/datetime/datetime.go](https://github.com/duke-git/lancet/blob/main/datetime/datetime.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/datetime"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [AddDay](#AddDay)
-   [AddHour](#AddHour)
-   [AddMinute](#AddMinute)
-   [AddYear](#AddYear)
-   [BeginOfMinute](#BeginOfMinute)
-   [BeginOfHour](#BeginOfHour)
-   [BeginOfDay](#BeginOfDay)
-   [BeginOfWeek](#BeginOfWeek)
-   [BeginOfMonth](#BeginOfMonth)
-   [BeginOfYear](#BeginOfYear)
-   [EndOfMinute](#EndOfMinute)
-   [EndOfHour](#EndOfHour)
-   [EndOfDay](#EndOfDay)
-   [EndOfWeek](#EndOfWeek)
-   [EndOfMonth](#EndOfMonth)
-   [EndOfYear](#EndOfYear)
-   [GetNowDate](#GetNowDate)
-   [GetNowTime](#GetNowTime)
-   [GetNowDateTime](#GetNowDateTime)
-   [GetZeroHourTimestamp](#GetZeroHourTimestamp)
-   [GetNightTimestamp](#GetNightTimestamp)
-   [FormatTimeToStr](#FormatTimeToStr)
-   [FormatStrToTime](#FormatStrToTime)
-   [NewUnixNow](#NewUnixNow)
-   [NewUnix](#NewUnix)
-   [NewFormat](#NewFormat)
-   [NewISO8601](#NewISO8601)
-   [ToUnix](#ToUnix)
-   [ToFormat](#ToFormat)
-   [ToFormatForTpl](#ToFormatForTpl)
-   [ToIso8601](#ToIso8601)
-   [IsLeapYear](#IsLeapYear)
-   [BetweenSeconds](#BetweenSeconds)
-   [DayOfYear](#DayOfYear)
-   [IsWeekend](#IsWeekend)


<div STYLE="page-break-after: always;"></div>

## 文档

## 注:

1. 方法 FormatTimeToStr 和 FormatStrToTime 中的 format 参数值需要传以下类型之一：

-   yyyy-mm-dd hh:mm:ss
-   yyyy-mm-dd hh:mm
-   yyyy-mm-dd hh
-   yyyy-mm-dd
-   yyyy-mm
-   mm-dd
-   dd-mm-yy hh:mm:ss
-   yyyy/mm/dd hh:mm:ss
-   yyyy/mm/dd hh:mm
-   yyyy-mm-dd hh
-   yyyy/mm/dd
-   yyyy/mm
-   mm/dd
-   dd/mm/yy hh:mm:ss
-   yyyy
-   mm
-   hh:mm:ss
-   mm:ss

### <span id="AddDay">AddDay</span>

<p>将日期加/减天数。</p>

<b>函数签名:</b>

```go
func AddDay(t time.Time, day int64) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    now := time.Now()

    tomorrow := datetime.AddDay(now, 1)
    diff1 := tomorrow.Sub(now)

    yesterday := datetime.AddDay(now, -1)
    diff2 := yesterday.Sub(now)

    fmt.Println(diff1)
    fmt.Println(diff2)

    // Output:
    // 24h0m0s
    // -24h0m0s
}
```

### <span id="AddHour">AddHour</span>

<p>将日期加/减小时数。</p>

<b>函数签名:</b>

```go
func AddHour(t time.Time, hour int64) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    now := time.Now()

    after2Hours := datetime.AddHour(now, 2)
    diff1 := after2Hours.Sub(now)

    before2Hours := datetime.AddHour(now, -2)
    diff2 := before2Hours.Sub(now)

    fmt.Println(diff1)
    fmt.Println(diff2)

    // Output:
    // 2h0m0s
    // -2h0m0s
}
```

### <span id="AddMinute">AddMinute</span>

<p>将日期加/减分钟数。</p>

<b>函数签名:</b>

```go
func AddMinute(t time.Time, minute int64) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    now := time.Now()

    after2Minutes := datetime.AddMinute(now, 2)
    diff1 := after2Minutes.Sub(now)

    before2Minutes := datetime.AddMinute(now, -2)
    diff2 := before2Minutes.Sub(now)

    fmt.Println(diff1)
    fmt.Println(diff2)

    // Output:
    // 2m0s
    // -2m0s
}
```

### <span id="AddYear">AddYear</span>

<p>将日期加/减年数。</p>

<b>函数签名:</b>

```go
func AddYear(t time.Time, year int64) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    now := time.Now()

    after1Year := datetime.AddYear(now, 1)
    diff1 := after1Year.Sub(now)

    before1Year := datetime.AddYear(now, -1)
    diff2 := before1Year.Sub(now)

    fmt.Println(diff1)
    fmt.Println(diff2)

    // Output:
    // 8760h0m0s
    // -8760h0m0s
}
```

### <span id="BeginOfMinute">BeginOfMinute</span>

<p>返回指定时间的分钟开始时间。</p>

<b>函数签名:</b>

```go
func BeginOfMinute(t time.Time) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfMinute(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 18:50:00 +0000 UTC
}
```

### <span id="BeginOfHour">BeginOfHour</span>

<p>返回指定时间的小时开始时间。</p>

<b>函数签名:</b>

```go
func BeginOfHour(t time.Time) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfHour(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 18:00:00 +0000 UTC
}
```

### <span id="BeginOfDay">BeginOfDay</span>

<p>返回指定时间的当天开始时间。</p>

<b>函数签名:</b>

```go
func BeginOfDay(t time.Time) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfDay(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 00:00:00 +0000 UTC
}
```

### <span id="BeginOfWeek">BeginOfWeek</span>

<p>返回指定时间的每周开始时间,默认开始时间星期日。</p>

<b>函数签名:</b>

```go
func BeginOfWeek(t time.Time, beginFrom ...time.Weekday) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfWeek(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 00:00:00 +0000 UTC
}
```

### <span id="BeginOfMonth">BeginOfMonth</span>

<p>返回指定时间的当月开始时间。</p>

<b>函数签名:</b>

```go
func BeginOfMonth(t time.Time) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfMonth(input)

    fmt.Println(result)

    // Output:
    // 2023-01-01 00:00:00 +0000 UTC
}
```

### <span id="BeginOfYear">BeginOfYear</span>

<p>返回指定时间的当年开始时间</p>

<b>函数签名:</b>

```go
func BeginOfYear(t time.Time) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.BeginOfYear(input)

    fmt.Println(result)

    // Output:
    // 2023-01-01 00:00:00 +0000 UTC
}
```

### <span id="EndOfMinute">EndOfMinute</span>

<p>返回指定时间的分钟结束时间。</p>

<b>函数签名:</b>

```go
func EndOfMinute(t time.Time) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfMinute(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 18:50:59.999999999 +0000 UTC
}
```

### <span id="EndOfHour">EndOfHour</span>

<p>返回指定时间的小时结束时间。</p>

<b>函数签名:</b>

```go
func EndOfHour(t time.Time) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfHour(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 18:59:59.999999999 +0000 UTC
}
```

### <span id="EndOfDay">EndOfDay</span>

<p>返回指定时间的当天结束时间。</p>

<b>函数签名:</b>

```go
func EndOfDay(t time.Time) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfDay(input)

    fmt.Println(result)

    // Output:
    // 2023-01-08 23:59:59.999999999 +0000 UTC
}
```

### <span id="EndOfWeek">EndOfWeek</span>

<p>返回指定时间的星期结束时间,默认结束时间星期六。</p>

<b>函数签名:</b>

```go
func EndOfWeek(t time.Time, endWith ...time.Weekday) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfWeek(input)

    fmt.Println(result)

    // Output:
    // 2023-01-14 23:59:59.999999999 +0000 UTC
}
```

### <span id="EndOfMonth">EndOfMonth</span>

<p>返回指定时间的当月结束时间。</p>

<b>函数签名:</b>

```go
func EndOfMonth(t time.Time) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfMonth(input)

    fmt.Println(result)

    // Output:
    // 2023-01-31 23:59:59.999999999 +0000 UTC
}
```

### <span id="EndOfYear">EndOfYear</span>

<p>返回指定时间的当年结束时间。</p>

<b>函数签名:</b>

```go
func EndOfYear(t time.Time) time.Time
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)
    result := datetime.EndOfYear(input)

    fmt.Println(result)

    // Output:
    // 2023-12-31 23:59:59.999999999 +0000 UTC
}
```

### <span id="GetNowDate">GetNowDate</span>

<p>获取当天日期，返回格式：yyyy-mm-dd。</p>

<b>函数签名:</b>

```go
func GetNowDate() string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    now := time.Now()
    currentDate := datetime.GetNowDate()

    fmt.Println(currentDate)

    // Output:
    // 2022-01-28
}
```

### <span id="GetNowTime">GetNowTime</span>

<p>获取当时时间，返回格式：hh:mm:ss</p>

<b>函数签名:</b>

```go
func GetNowTime() string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    now := time.Now()
    currentTime := datetime.GetNowTime()

    fmt.Println(currentTime) // 15:57:33

    // Output:
    // 15:57:33
}
```

### <span id="GetNowDateTime">GetNowDateTime</span>

<p>获取当时日期和时间，返回格式：yyyy-mm-dd hh:mm:ss。</p>

<b>函数签名:</b>

```go
func GetNowDateTime() string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    now := time.Now()
    current := datetime.GetNowDateTime()

    fmt.Println(current)

    // Output:
    // 2022-01-28 15:59:33
}
```

### <span id="GetZeroHourTimestamp">GetZeroHourTimestamp</span>

<p>获取零点时间戳(timestamp of 00:00)</p>

<b>函数签名:</b>

```go
func GetZeroHourTimestamp() int64
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    now := time.Now()
    zeroTime := datetime.GetZeroHourTimestamp()

    fmt.Println(zeroTime)

    // Output:
    // 1643299200
}
```

### <span id="GetNightTimestamp">GetNightTimestamp</span>

<p>获取午夜时间戳(timestamp of 23:59)。</p>

<b>函数签名:</b>

```go
func GetNightTimestamp() int64
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    now := time.Now()
    nightTime := datetime.GetNightTimestamp()

    fmt.Println(nightTime)

    // Output:
    // 1643385599
}
```

### <span id="FormatTimeToStr">FormatTimeToStr</span>

<p>将日期格式化成字符串，`format` 参数格式参考注1。</p>

<b>函数签名:</b>

```go
func FormatTimeToStr(t time.Time, format string) string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    t, _ := time.Parse("2006-01-02 15:04:05", "2021-01-02 16:04:08")

    result1 := datetime.FormatTimeToStr(t, "yyyy-mm-dd hh:mm:ss")
    result2 := datetime.FormatTimeToStr(t, "yyyy-mm-dd")
    result3 := datetime.FormatTimeToStr(t, "dd-mm-yy hh:mm:ss")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 2021-01-02 16:04:08
    // 2021-01-02
    // 02-01-21 16:04:08
}
```

### <span id="FormatStrToTime">FormatStrToTime</span>

<p>将字符串格式化成时间，`format` 参数格式参考注1。</p>

<b>函数签名:</b>

```go
func FormatStrToTime(str, format string) (time.Time, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    result1, _ := datetime.FormatStrToTime("2021-01-02 16:04:08", "yyyy-mm-dd hh:mm:ss")
    result2, _ := datetime.FormatStrToTime("2021-01-02", "yyyy-mm-dd")
    result3, _ := datetime.FormatStrToTime("02-01-21 16:04:08", "dd-mm-yy hh:mm:ss")

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 2021-01-02 16:04:08 +0000 UTC
    // 2021-01-02 00:00:00 +0000 UTC
    // 2021-01-02 16:04:08 +0000 UTC
}
```

### <span id="NewUnixNow">NewUnixNow</span>

<p>创建一个当前时间的unix时间戳。</p>

<b>函数签名:</b>

```go
type theTime struct {
    unix int64
}
func NewUnixNow() *theTime
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    tm := datetime.NewUnixNow()
    fmt.Println(tm)

    // Output:
    // &{1647597438}
}
```

### <span id="NewUnix">NewUnix</span>

<p>创建一个unix时间戳。</p>

<b>函数签名:</b>

```go
type theTime struct {
    unix int64
}
func NewUnix(unix int64) *theTime
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    tm := datetime.NewUnix(1647597438)
    fmt.Println(tm)

    // Output:
    // &{1647597438}
}
```

### <span id="NewFormat">NewFormat</span>

<p>创建一个yyyy-mm-dd hh:mm:ss格式时间字符串的unix时间戳。</p>

<b>函数签名:</b>

```go
type theTime struct {
    unix int64
}
func NewFormat(t string) (*theTime, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    tm, err := datetime.NewFormat("2022-03-18 17:04:05")
    fmt.Println(tm)

    // Output:
    // &{1647594245}
}
```

### <span id="NewISO8601">NewISO8601</span>

<p>创建一个iso8601格式时间字符串的unix时间戳。</p>

<b>函数签名:</b>

```go
type theTime struct {
    unix int64
}
func NewISO8601(iso8601 string) (*theTime, error)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    tm, err := datetime.NewISO8601("2006-01-02T15:04:05.999Z")
    fmt.Println(tm)

    // Output:
    // &{1136214245}
}
```

### <span id="ToUnix">ToUnix</span>

<p>返回unix时间戳。</p>

<b>函数签名:</b>

```go
func (t *theTime) ToUnix() int64
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    tm := datetime.NewUnixNow()
    fmt.Println(tm.ToUnix())

    // Output:
    // 1647597438
}
```

### <span id="ToFormat">ToFormat</span>

<p>返回格式'yyyy-mm-dd hh:mm:ss'的日期字符串。</p>

<b>函数签名:</b>

```go
func (t *theTime) ToFormat() string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    tm, _ := datetime.NewFormat("2022-03-18 17:04:05")
    fmt.Println(tm.ToFormat())

    // Output:
    // 2022-03-18 17:04:05
}
```

### <span id="ToFormatForTpl">ToFormatForTpl</span>

<p>返回tpl格式指定的日期字符串。</p>

<b>函数签名:</b>

```go
func (t *theTime) ToFormatForTpl(tpl string) string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    tm, _ := datetime.NewFormat("2022-03-18 17:04:05")
    ts := tm.ToFormatForTpl("2006/01/02 15:04:05")
    fmt.Println(ts)

    // Output:
    // 2022/03/18 17:04:05
}
```

### <span id="ToIso8601">ToIso8601</span>

<p>返回iso8601日期字符串。</p>

<b>函数签名:</b>

```go
func (t *theTime) ToIso8601() string
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    tm, _ := datetime.NewISO8601("2006-01-02T15:04:05.999Z")
    ts := tm.ToIso8601()
    fmt.Println(ts)

    // Output:
    // 2006-01-02T23:04:05+08:00
}
```

### <span id="IsLeapYear">IsLeapYear</span>

<p>验证是否是闰年。</p>

<b>函数签名:</b>

```go
func IsLeapYear(year int) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    result1 := datetime.IsLeapYear(2000)
    result2 := datetime.IsLeapYear(2001)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // true
    // false
}
```

### <span id="BetweenSeconds">BetweenSeconds</span>

<p>返回两个时间的间隔秒数。</p>

<b>函数签名:</b>

```go
func BetweenSeconds(t1 time.Time, t2 time.Time) int64
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    today := time.Now()
    tomorrow := datetime.AddDay(today, 1)
    yesterday := datetime.AddDay(today, -1)

    result1 := datetime.BetweenSeconds(today, tomorrow)
    result2 := datetime.BetweenSeconds(today, yesterday)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 86400
    // -86400
}
```

### <span id="DayOfYear">DayOfYear</span>

<p>返回参数日期是一年中的第几天。</p>

<b>函数签名:</b>

```go
func DayOfYear(t time.Time) int
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    date1 := time.Date(2023, 02, 01, 1, 1, 1, 0, time.Local)
    result1 := datetime.DayOfYear(date1)

    date2 := time.Date(2023, 01, 02, 1, 1, 1, 0, time.Local)
    result2 := datetime.DayOfYear(date2)

    date3 := time.Date(2023, 01, 01, 1, 1, 1, 0, time.Local)
    result3 := datetime.DayOfYear(date3)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 31
    // 1
    // 0
}
```

### <span id="IsWeekend">IsWeekend</span>

<p>判断日期是否是周末。</p>

<b>函数签名:</b>

```go
func IsWeekend(t time.Time) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    date1 := time.Date(2023, 06, 03, 0, 0, 0, 0, time.Local)
    date2 := time.Date(2023, 06, 04, 0, 0, 0, 0, time.Local)
    date3 := time.Date(2023, 06, 02, 0, 0, 0, 0, time.Local)

    result1 := datetime.IsWeekend(date1)
    result2 := datetime.IsWeekend(date2)
    result3 := datetime.IsWeekend(date3)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // true
    // true
    // false
}
```
