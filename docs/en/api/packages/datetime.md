# Datetime

Package datetime supports date and time format and compare.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/datetime/datetime.go](https://github.com/duke-git/lancet/blob/main/datetime/datetime.go)
-   [https://github.com/duke-git/lancet/blob/main/datetime/conversion.go](https://github.com/duke-git/lancet/blob/main/datetime/conversion.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/datetime"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

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
-   [GenerateDatetimesBetween](#GenerateDatetimesBetween)


<div STYLE="page-break-after: always;"></div>

## Documentation

## Note:

1. In below functions, the `format` string param should be one of flows value (case no sensitive):

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

<p>Add or sub days to time.</p>

<b>Signature:</b>

```go
func AddDay(t time.Time, day int64) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/dIGbs_uTdFa)</span></b>

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

<p>Add or sub hours to time.</p>

<b>Signature:</b>

```go
func AddHour(t time.Time, hour int64) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/rcMjd7OCsi5)</span></b>

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

<p>Add or sub minutes to time.</p>

<b>Signature:</b>

```go
func AddMinute(t time.Time, minute int64) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/nT1heB1KUUK)</span></b>

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

<p>Add or sub year to the time.</p>

<b>Signature:</b>

```go
func AddYear(t time.Time, year int64) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/MqW2ujnBx10)</span></b>

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

<p>Return beginning minute time of day.</p>

<b>Signature:</b>

```go
func BeginOfMinute(t time.Time) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/ieOLVJ9CiFT)</span></b>

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

<p>Return zero time of day.</p>

<b>Signature:</b>

```go
func BeginOfHour(t time.Time) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/GhdGFnDWpYs)</span></b>

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

<p>Return begin time of day.</p>

<b>Signature:</b>

```go
func BeginOfDay(t time.Time) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/94m_UT6cWs9)</span></b>

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

<p>Return beginning time of week, week begin from Sunday.</p>

<b>Signature:</b>

```go
func BeginOfWeek(t time.Time, beginFrom ...time.Weekday) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/ynjoJPz7VNV)</span></b>

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

<p>Return beginning time of month</p>

<b>Signature:</b>

```go
func BeginOfMonth(t time.Time) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/bWXVFsmmzwL)</span></b>

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

<p>Return beginning time of year.</p>

<b>Signature:</b>

```go
func BeginOfYear(t time.Time) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/i326DSwLnV8)</span></b>

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

<p>Return end time minute of day.</p>

<b>Signature:</b>

```go
func EndOfMinute(t time.Time) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/yrL5wGzPj4z)</span></b>

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

<p>Return end time hour of day.</p>

<b>Signature:</b>

```go
func EndOfHour(t time.Time) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/6ce3j_6cVqN)</span></b>

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

<p>Return end time hour of day.</p>

<b>Signature:</b>

```go
func EndOfDay(t time.Time) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/eMBOvmq5Ih1)</span></b>

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

<p>Return end time of week, week end with Saturday.</p>

<b>Signature:</b>

```go
func EndOfWeek(t time.Time, endWith ...time.Weekday) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/i08qKXD9flf)</span></b>

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

<p>Return end time of month</p>

<b>Signature:</b>

```go
func EndOfMonth(t time.Time) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/_GWh10B3Nqi)</span></b>

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

<p>Return beginning time of year.</p>

<b>Signature:</b>

```go
func EndOfYear(t time.Time) time.Time
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/G01cKlMCvNm)</span></b>

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

<p>Get current date string, format is yyyy-mm-dd.</p>

<b>Signature:</b>

```go
func GetNowDate() string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/PvfkPpcpBBf)</span></b>

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

<p>Get current time string, format is hh:mm:ss.</p>

<b>Signature:</b>

```go
func GetNowTime() string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/l7BNxCkTmJS)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datetime"
)

func main() {
    currentTime := datetime.GetNowTime()
    fmt.Println(currentTime) // 15:57:33

    // Output:
    // 15:57:33
}
```

### <span id="GetNowDateTime">GetNowDateTime</span>

<p>Get current date time string, format is yyyy-mm-dd hh:mm:ss.</p>

<b>Signature:</b>

```go
func GetNowDateTime() string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/pI4AqngD0al)</span></b>

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

<p>Return the start time of today, format: yyyy-mm-dd 00:00:00.</p>

<b>Signature:</b>

```go
func GetTodayStartTime() string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/84siyYF7t99)</span></b>

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

<p>Return the end time of today, format: yyyy-mm-dd 23:59:59.</p>

<b>Signature:</b>

```go
func GetTodayEndTime() string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/jjrLnfoqgn3)</span></b>

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

<p>Return timestamp of zero hour (timestamp of 00:00).</p>

<b>Signature:</b>

```go
func GetZeroHourTimestamp() int64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/QmL2oIaGE3q)</span></b>

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

<p>Return timestamp of zero hour (timestamp of 23:59).</p>

<b>Signature:</b>

```go
func GetNightTimestamp() int64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/UolysR3MYP1)</span></b>

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

<p>Format time to string, `format` param specification see note 1.</p>

<b>Signature:</b>

```go
func FormatTimeToStr(t time.Time, format string, timezone ...string) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/_Ia7M8H_OvE)</span></b>

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

<p>Format string to time, `format` param specification see note 1.</p>

<b>Signature:</b>

```go
func FormatStrToTime(str, format string, timezone ...string) (time.Time, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/1h9FwdU8ql4)</span></b>

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

<p>Return unix timestamp of current time</p>

<b>Signature:</b>

```go
type theTime struct {
    unix int64
}
func NewUnixNow() *theTime
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/U4PPx-9D0oz)</span></b>

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

<p>Return unix timestamp of specified int64 value.</p>

<b>Signature:</b>

```go
type theTime struct {
    unix int64
}
func NewUnix(unix int64) *theTime
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/psoSuh_kLRt)</span></b>

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

<p>Return unix timestamp of specified time string, t should be "yyyy-mm-dd hh:mm:ss".</p>

<b>Signature:</b>

```go
type theTime struct {
    unix int64
}
func NewFormat(t string) (*theTime, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/VkW08ZOaXPZ)</span></b>

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

<p>Return unix timestamp of specified iso8601 time string.</p>

<b>Signature:</b>

```go
type theTime struct {
    unix int64
}
func NewISO8601(iso8601 string) (*theTime, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/mkhOHQkdeA2)</span></b>

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

<p>Return unix timestamp.</p>

<b>Signature:</b>

```go
func (t *theTime) ToUnix() int64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/_LUiwAdocjy)</span></b>

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

<p>Return time string 'yyyy-mm-dd hh:mm:ss'.</p>

<b>Signature:</b>

```go
func (t *theTime) ToFormat() string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/VkW08ZOaXPZ)</span></b>

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

<p>Return the time string which format is specified tpl.</p>

<b>Signature:</b>

```go
func (t *theTime) ToFormatForTpl(tpl string) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/nyXxXcQJ8L5)</span></b>

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

<p>Return iso8601 time string.</p>

<b>Signature:</b>

```go
func (t *theTime) ToIso8601() string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/mkhOHQkdeA2)</span></b>

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

<p>check if param `year` is leap year or not.</p>

<b>Signature:</b>

```go
func IsLeapYear(year int) bool
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/xS1eS2ejGew)</span></b>

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

<p>Return the number of seconds between two times.</p>

<b>Signature:</b>

```go
func BetweenSeconds(t1 time.Time, t2 time.Time) int64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/n3YDRyfyXJu)</span></b>

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

<p>Returns which day of the year the parameter date `t` is.</p>

<b>Signature:</b>

```go
func DayOfYear(t time.Time) int
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/0hjqhTwFNlH)</span></b>

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

### <span id="IsWeekend">IsWeekend(Deprecated, Use '== Weekday' instead)</span>

<p>Checks if passed time is weekend or not.</p>

<b>Signature:</b>

```go
func IsWeekend(t time.Time) bool
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/cupRM5aZOIY)</span></b>

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

<p>Return current datetime with specific format and timezone.</p>

<b>Signature:</b>

```go
func NowDateOrTime(format string, timezone ...string) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/EZ-begEjtT0)</span></b>

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

<p>Return current second timestamp.</p>

<b>Signature:</b>

```go
func Timestamp(timezone ...string) int64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/iU5b7Vvjx6x)</span></b>

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

<p>Return current mill second timestamp.</p>

<b>Signature:</b>

```go
func TimestampMilli(timezone ...string) int64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/4gvEusOTu1T)</span></b>

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

<p>Return current micro second timestamp.</p>

<b>Signature:</b>

```go
func TimestampMicro(timezone ...string) int64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/2maANglKHQE)</span></b>

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

<p>Return current nano second timestamp.</p>

<b>Signature:</b>

```go
func TimestampNano(timezone ...string) int64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/A9Oq_COrcCF)</span></b>

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

<p>Tracks function execution time.</p>

<b>Signature:</b>

```go
func TrackFuncTime(pre time.Time) func()
```

<b>Example:<span style="float:right;display:inline-block;">[Run]()</span></b>

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

<p>Returns the number of days between two times.</p>

<b>Signature:</b>

```go
func DaysBetween(start, end time.Time) int
```

<b>Example:<span style="float:right;display:inline-block;">[Run]()</span></b>

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

<p>Returns a slice of strings between two times. `layout`: the format of the datetime string.`interval`: the interval between two datetimes.</p>

<b>Signature:</b>

```go
func GenerateDatetimesBetween(start, end time.Time, layout string, interval string) ([]string, error)
```

<b>Example:<span style="float:right;display:inline-block;">[Run]()</span></b>

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