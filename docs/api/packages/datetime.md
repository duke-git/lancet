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
-   [AddWeek](#AddWeek)
-   [AddMonth](#AddMonth)
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
-   [GetTodayStartTime](#GetTodayStartTime)
-   [GetTodayEndTime](#GetTodayEndTime)
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
-   [IsWeekend<sup>deprecated</sup>](#IsWeekend)
-   [NowDateOrTime](#NowDateOrTime)
-   [Timestamp](#Timestamp)
-   [TimestampMilli](#TimestampMilli)
-   [TimestampMicro](#TimestampMicro)
-   [TimestampNano](#TimestampNano)
-   [TrackFuncTime](#TrackFuncTime)
-   [DaysBetween](#DaysBetween)
-   [GenerateDatetimesBetween](#GenerateDatetimesBetween)
-   [Min](#Min)
-   [Max](#Max)
-   [MaxMin](#MaxMin)

<div STYLE="page-break-after: always;"></div>

## 文档

## 注:

1. 函数中`format`参数值需要传以下值之一 (忽略大小写):

-   yyyy-mm-dd hh:mm:ss
-   yyyy-mm-dd hh:mm
-   yyyy-mm-dd hh
-   yyyy-mm-dd
-   yyyy-mm
-   mm-dd
-   dd-mm-yy hh:mm:ss
-   yyyy/mm/dd hh:mm:ss
-   yyyy/mm/dd hh:mm
-   yyyy/mm/dd hh
-   yyyy/mm/dd
-   yyyy/mm
-   mm/dd
-   dd/mm/yy hh:mm:ss
-   yyyymmdd
-   mmddyy
-   yyyy
-   yy
-   mm
-   hh:mm:ss
-   hh:mm
-   mm:ss

### <span id="AddDay">AddDay</span>

<p>将日期加/减天数。</p>

<b>函数签名:</b>

```go
func AddDay(t time.Time, days int64) time.Time
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/dIGbs_uTdFa)</span></b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    date, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 00:00:00")

    after1Day := datetime.AddDay(date, 1)
    before1Day := datetime.AddDay(date, -1)

    fmt.Println(after1Day.Format("2006-01-02 15:04:05"))
    fmt.Println(before1Day.Format("2006-01-02 15:04:05"))

    // Output:
    // 2021-01-02 00:00:00
    // 2020-12-31 00:00:00
}
```

### <span id="AddWeek">AddWeek</span>

<p>将日期加/减星期数。</p>

<b>函数签名:</b>

```go
func AddWeek(t time.Time, weeks int64) time.Time
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    date, _ := time.Parse("2006-01-02", "2021-01-01")

    after2Weeks := datetime.AddWeek(date, 2)
    before2Weeks := datetime.AddWeek(date, -2)

    fmt.Println(after2Weeks.Format("2006-01-02"))
    fmt.Println(before2Weeks.Format("2006-01-02"))

    // Output:
    // 2021-01-15
    // 2020-12-18
}
```

### <span id="AddMonth">AddMonth</span>

<p>将日期加/减月数。</p>

<b>函数签名:</b>

```go
func AddMonth(t time.Time, months int64) time.Time
```

<b>示例:<span style="float:right;display:inline-block;">[运行](todo)</span></b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    date, _ := time.Parse("2006-01-02", "2021-01-01")

    after2Months := datetime.AddMonth(date, 2)
    before2Months := datetime.AddMonth(date, -2)

    fmt.Println(after2Months.Format("2006-01-02"))
    fmt.Println(before2Months.Format("2006-01-02"))

    // Output:
    // 2021-03-01
    // 2020-11-01
}
```

### <span id="AddHour">AddHour</span>

<p>将日期加/减小时数。</p>

<b>函数签名:</b>

```go
func AddHour(t time.Time, hours int64) time.Time
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/rcMjd7OCsi5)</span></b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    date, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 00:00:00")

    after2Hours := datetime.AddHour(date, 2)
    before2Hours := datetime.AddHour(date, -2)

    fmt.Println(after2Hours.Format("2006-01-02 15:04:05"))
    fmt.Println(before2Hours.Format("2006-01-02 15:04:05"))

    // Output:
    // 2021-01-01 02:00:00
    // 2020-12-31 22:00:00
}
```

### <span id="AddMinute">AddMinute</span>

<p>将日期加/减分钟数。</p>

<b>函数签名:</b>

```go
func AddMinute(t time.Time, minutes int64) time.Time
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/nT1heB1KUUK)</span></b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    date, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 00:00:00")

    after2Minutes := datetime.AddMinute(date, 2)
    before2Minutes := datetime.AddMinute(date, -2)

    fmt.Println(after2Minutes.Format("2006-01-02 15:04:05"))
    fmt.Println(before2Minutes.Format("2006-01-02 15:04:05"))

    // Output:
    // 2021-01-01 00:02:00
    // 2020-12-31 23:58:00
}
```

### <span id="AddYear">AddYear</span>

<p>将日期加/减年数。</p>

<b>函数签名:</b>

```go
func AddYear(t time.Time, years int64) time.Time
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/MqW2ujnBx10)</span></b>

```go
package main

import (
    "fmt"
    "time"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    date, _ := time.Parse("2006-01-02", "2021-01-01")

    after2Years := AddYear(date, 2)
    before2Years := AddYear(date, -2)

    fmt.Println(after2Years.Format("2006-01-02"))
    fmt.Println(before2Years.Format("2006-01-02"))

    // Output:
    // 2023-01-01
    // 2019-01-01
}
```

### <span id="BeginOfMinute">BeginOfMinute</span>

<p>返回指定时间的分钟开始时间。</p>

<b>函数签名:</b>

```go
func BeginOfMinute(t time.Time) time.Time
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ieOLVJ9CiFT)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/GhdGFnDWpYs)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/94m_UT6cWs9)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ynjoJPz7VNV)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/bWXVFsmmzwL)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/i326DSwLnV8)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/yrL5wGzPj4z)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/6ce3j_6cVqN)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/eMBOvmq5Ih1)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/i08qKXD9flf)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/_GWh10B3Nqi)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/G01cKlMCvNm)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/PvfkPpcpBBf)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/l7BNxCkTmJS)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    currentTime := datetime.GetNowTime()
    fmt.Println(currentTime)

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/pI4AqngD0al)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    current := datetime.GetNowDateTime()
    fmt.Println(current)

    // Output:
    // 2022-01-28 15:59:33
}
```

### <span id="GetTodayStartTime">GetTodayStartTime</span>

<p>返回当天开始时间， 格式: yyyy-mm-dd 00:00:00。</p>

<b>函数签名:</b>

```go
func GetTodayStartTime() string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/84siyYF7t99)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    startTime := datetime.GetTodayStartTime()
    fmt.Println(startTime)

    // Output:
    // 2023-06-29 00:00:00
}
```

### <span id="GetTodayEndTime">GetTodayEndTime</span>

<p>返回当天结束时间，格式: yyyy-mm-dd 23:59:59。</p>

<b>函数签名:</b>

```go
func GetTodayEndTime() string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/jjrLnfoqgn3)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    endTime := datetime.GetTodayEndTime()
    fmt.Println(endTime)

    // Output:
    // 2023-06-29 23:59:59
}
```

### <span id="GetZeroHourTimestamp">GetZeroHourTimestamp</span>

<p>获取零点时间戳(timestamp of 00:00)</p>

<b>函数签名:</b>

```go
func GetZeroHourTimestamp() int64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/QmL2oIaGE3q)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/UolysR3MYP1)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
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
func FormatTimeToStr(t time.Time, format string, timezone ...string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/_Ia7M8H_OvE)</span></b>

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
func FormatStrToTime(str, format string, timezone ...string) (time.Time, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/1h9FwdU8ql4)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/U4PPx-9D0oz)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/psoSuh_kLRt)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/VkW08ZOaXPZ)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/mkhOHQkdeA2)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/_LUiwAdocjy)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/VkW08ZOaXPZ)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/nyXxXcQJ8L5)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/mkhOHQkdeA2)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/xS1eS2ejGew)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/n3YDRyfyXJu)</span></b>

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

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/0hjqhTwFNlH)</span></b>

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

### <span id="IsWeekend">IsWeekend（已废弃, 使用 '== Weekday'）</span>

<p>判断日期是否是周末。</p>

<b>函数签名:</b>

```go
func IsWeekend(t time.Time) bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/cupRM5aZOIY)</span></b>

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

### <span id="NowDateOrTime">NowDateOrTime</span>

<p>根据指定的格式和时区返回当前时间字符串。</p>

<b>函数签名:</b>

```go
func NowDateOrTime(format string, timezone ...string) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/EZ-begEjtT0)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    result1 := datetime.NowDateOrTime("yyyy-mm-dd hh:mm:ss")

    result2 := datetime.NowDateOrTime("yyyy-mm-dd hh:mm:ss", "EST")

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 2023-07-26 15:01:30
    // 2023-07-26 02:01:30
}
```

### <span id="Timestamp">Timestamp</span>

<p>返回当前秒级时间戳。</p>

<b>函数签名:</b>

```go
func Timestamp(timezone ...string) int64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/iU5b7Vvjx6x)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    ts := datetime.Timestamp()

    fmt.Println(ts)

    // Output:
    // 1690363051
}
```


### <span id="TimestampMilli">TimestampMilli</span>

<p>返回当前毫秒级时间戳。</p>

<b>函数签名:</b>

```go
func TimestampMilli(timezone ...string) int64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/4gvEusOTu1T)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    ts := datetime.TimestampMilli()

    fmt.Println(ts)

    // Output:
    // 1690363051331
}
```

### <span id="TimestampMicro">TimestampMicro</span>

<p>返回当前微秒级时间戳。</p>

<b>函数签名:</b>

```go
func TimestampMicro(timezone ...string) int64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/2maANglKHQE)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    ts := datetime.TimestampMicro()

    fmt.Println(ts)

    // Output:
    // 1690363051331784
}
```

### <span id="TimestampNano">TimestampNano</span>

<p>返回当前纳秒级时间戳。</p>

<b>函数签名:</b>

```go
func TimestampNano(timezone ...string) int64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/A9Oq_COrcCF)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    ts := datetime.TimestampNano()

    fmt.Println(ts)

    // Output:
    // 1690363051331788000
}
```

### <span id="TrackFuncTime">TrackFuncTime</span>

<p>测试函数执行时间。</p>

<b>函数签名:</b>

```go
func TrackFuncTime(pre time.Time) func()
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/QBSEdfXHPTp)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    defer datetime.TrackFuncTime(time.Now())()

    var n int
    for i := 0; i < 5000000; i++ {
        n++
    }

    fmt.Println(1) // Function main execution time:     1.460287ms
}
```

### <span id="DaysBetween">DaysBetween</span>

<p>返回两个日期之间的天数差。</p>

<b>函数签名:</b>

```go
func DaysBetween(start, end time.Time) int
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/qD6qGb3TbOy)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    start := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2024, time.September, 10, 0, 0, 0, 0, time.UTC)

    result := datetime.DaysBetween(start, end)

    fmt.Println(result)

    // Output:
    // 9
}
```

### <span id="GenerateDatetimesBetween">GenerateDatetimesBetween</span>

<p>生成从start到end的所有日期时间的字符串列表。layout参数表示时间格式，例如"2006-01-02 15:04:05"，interval参数表示时间间隔，例如"1h"表示1小时，"30m"表示30分钟。</p>

<b>函数签名:</b>

```go
func GenerateDatetimesBetween(start, end time.Time, layout string, interval string) ([]string, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/6kHBpAxD9ZC)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    start := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
    end := time.Date(2024, time.September, 1, 2, 0, 0, 0, time.UTC)

    layout := "2006-01-02 15:04:05"
    interval := "1h"

    result, err := datetime.GenerateDatetimesBetween(start, end, layout, interval)

    fmt.Println(result)
    fmt.Println(err)

    // Output:
    // [2024-09-01 00:00:00 2024-09-01 01:00:00 2024-09-01 02:00:00]
    // <nil>
}
```

### <span id="Min">Min</span>

<p>返回最早时间。</p>

<b>函数签名:</b>

```go
func Min(t1 time.Time, times ...time.Time) time.Time
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/MCIDvHNOGGb)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    minTime := datetime.Min(time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.September, 2, 0, 0, 0, 0, time.UTC))

	fmt.Println(minTime)

	// Output:
	// 2024-09-01 00:00:00 +0000 UTC
}
```

### <span id="Max">Max</span>

<p>返回最晚时间。</p>

<b>函数签名:</b>

```go
func Max(t1 time.Time, times ...time.Time) time.Time
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/9m6JMk1LB7-)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    maxTime := datetime.Min(time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.September, 2, 0, 0, 0, 0, time.UTC))

	fmt.Println(maxTime)

	// Output:
	// 2024-09-02 00:00:00 +0000 UTC
}
```

### <span id="MaxMin">MaxMin</span>

<p>返回最早和最晚时间。</p>

<b>函数签名:</b>

```go
func MaxMin(t1 time.Time, times ...time.Time) (maxTime time.Time, minTime time.Time)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/rbW51cDtM_2)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    max, min := datetime.MaxMin(time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.September, 2, 0, 0, 0, 0, time.UTC), time.Date(2024, time.September, 3, 0, 0, 0, 0, time.UTC))

	fmt.Println(max)
	fmt.Println(min)

	// Output:
	// 2024-09-03 00:00:00 +0000 UTC
	// 2024-09-01 00:00:00 +0000 UTC
}
```