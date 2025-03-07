// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package datetime implements some functions to format date and time.
// Note:
// 1. `format` param in FormatTimeToStr function should be as flow (case no sensitive):
// "yyyy-mm-dd hh:mm:ss"
// "yyyy-mm-dd hh:mm"
// "yyyy-mm-dd hh"
// "yyyy-mm-dd"
// "yyyy-mm"
// "mm-dd"
// "dd-mm-yy hh:mm:ss"
// "yyyy/mm/dd hh:mm:ss"
// "yyyy/mm/dd hh:mm"
// "yyyy/mm/dd hh"
// "yyyy/mm/dd"
// "yyyy/mm"
// "mm/dd"
// "dd/mm/yy hh:mm:ss"
// "yyyymmdd"
// "mmddyy"
// "yyyy"
// "yy"
// "mm"
// "hh:mm:ss"
// "hh:mm"
// "mm:ss"
package datetime

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

var timeFormat map[string]string

func init() {
	timeFormat = map[string]string{
		"yyyy-mm-dd hh:mm:ss": "2006-01-02 15:04:05",
		"yyyy-mm-dd hh:mm":    "2006-01-02 15:04",
		"yyyy-mm-dd hh":       "2006-01-02 15",
		"yyyy-mm-dd":          "2006-01-02",
		"yyyy-mm":             "2006-01",
		"mm-dd":               "01-02",
		"dd-mm-yy hh:mm:ss":   "02-01-06 15:04:05",
		"yyyy/mm/dd hh:mm:ss": "2006/01/02 15:04:05",
		"yyyy/mm/dd hh:mm":    "2006/01/02 15:04",
		"yyyy/mm/dd hh":       "2006/01/02 15",
		"yyyy/mm/dd":          "2006/01/02",
		"yyyy/mm":             "2006/01",
		"mm/dd":               "01/02",
		"dd/mm/yy hh:mm:ss":   "02/01/06 15:04:05",
		"yyyymmdd":            "20060102",
		"mmddyy":              "010206",
		"yyyy":                "2006",
		"yy":                  "06",
		"mm":                  "01",
		"hh:mm:ss":            "15:04:05",
		"hh:mm":               "15:04",
		"mm:ss":               "04:05",
	}
}

// AddMinute add or sub minutes to the time.
// Play: https://go.dev/play/p/nT1heB1KUUK
func AddMinute(t time.Time, minutes int64) time.Time {
	return t.Add(time.Minute * time.Duration(minutes))
}

// AddHour add or sub hours to the time.
// Play: https://go.dev/play/p/rcMjd7OCsi5
func AddHour(t time.Time, hours int64) time.Time {
	return t.Add(time.Hour * time.Duration(hours))
}

// AddDay add or sub days to the time.
// Play: https://go.dev/play/p/dIGbs_uTdFa
func AddDay(t time.Time, days int64) time.Time {
	return t.Add(24 * time.Hour * time.Duration(days))
}

// AddWeek add or sub weeks to the time.
// play: https://go.dev/play/p/M9TqdMiaA2p
func AddWeek(t time.Time, weeks int64) time.Time {
	return t.Add(7 * 24 * time.Hour * time.Duration(weeks))
}

// AddMonth add or sub months to the time.
// Play: https://go.dev/play/p/DLoiOnpLvsN
func AddMonth(t time.Time, months int64) time.Time {
	return t.AddDate(0, int(months), 0)
}

// AddYear add or sub year to the time.
// Play: https://go.dev/play/p/MqW2ujnBx10
func AddYear(t time.Time, year int64) time.Time {
	return t.AddDate(int(year), 0, 0)
}

// AddDaySafe add or sub days to the time and ensure that the returned date does not exceed the valid date of the target year and month.
// Play: https://go.dev/play/p/JTohZFpoDJ3
func AddDaySafe(t time.Time, days int) time.Time {
	t = t.AddDate(0, 0, days)
	year, month, day := t.Date()

	lastDayOfMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, t.Location()).Day()

	if day > lastDayOfMonth {
		t = time.Date(year, month, lastDayOfMonth, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	return t
}

// AddMonthSafe add or sub months to the time and ensure that the returned date does not exceed the valid date of the target year and month.
// Play: https://go.dev/play/p/KLw0lo6mbVW
func AddMonthSafe(t time.Time, months int) time.Time {
	year := t.Year()
	month := int(t.Month()) + months

	for month > 12 {
		month -= 12
		year++
	}
	for month < 1 {
		month += 12
		year--
	}

	daysInMonth := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()

	day := t.Day()
	if day > daysInMonth {
		day = daysInMonth
	}

	return time.Date(year, time.Month(month), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// AddYearSafe add or sub years to the time and ensure that the returned date does not exceed the valid date of the target year and month.
// Play: https://go.dev/play/p/KVGXWZZ54ZH
func AddYearSafe(t time.Time, years int) time.Time {
	year, month, day := t.Date()
	year += years

	if month == time.February && day == 29 {
		if !IsLeapYear(year) {
			day = 28
		}
	}

	return time.Date(year, month, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// GetNowDate return format yyyy-mm-dd of current date.
// Play: https://go.dev/play/p/PvfkPpcpBBf
func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

// GetNowTime return format hh-mm-ss of current time.
// Play: https://go.dev/play/p/l7BNxCkTmJS
func GetNowTime() string {
	return time.Now().Format("15:04:05")
}

// GetNowDateTime return format yyyy-mm-dd hh-mm-ss of current datetime.
// Play: https://go.dev/play/p/pI4AqngD0al
func GetNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetTodayStartTime return the start time of today, format: yyyy-mm-dd 00:00:00.
// Play: https://go.dev/play/p/84siyYF7t99
func GetTodayStartTime() string {
	return time.Now().Format("2006-01-02") + " 00:00:00"
}

// GetTodayEndTime return the end time of today, format: yyyy-mm-dd 23:59:59.
// Play: https://go.dev/play/p/jjrLnfoqgn3
func GetTodayEndTime() string {
	return time.Now().Format("2006-01-02") + " 23:59:59"
}

// GetZeroHourTimestamp return timestamp of zero hour (timestamp of 00:00).
// Play: https://go.dev/play/p/QmL2oIaGE3q
func GetZeroHourTimestamp() int64 {
	ts := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", ts)
	return t.UTC().Unix() - 8*3600
}

// GetNightTimestamp return timestamp of zero hour (timestamp of 23:59).
// Play: https://go.dev/play/p/UolysR3MYP1
func GetNightTimestamp() int64 {
	return GetZeroHourTimestamp() + 86400 - 1
}

// FormatTimeToStr convert time to string.
// Play: https://go.dev/play/p/_Ia7M8H_OvE
func FormatTimeToStr(t time.Time, format string, timezone ...string) string {
	tf, ok := timeFormat[strings.ToLower(format)]
	if !ok {
		return ""
	}

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return ""
		}
		return t.In(loc).Format(tf)
	}
	return t.Format(tf)
}

// FormatStrToTime convert string to time.
// Play: https://go.dev/play/p/1h9FwdU8ql4
func FormatStrToTime(str, format string, timezone ...string) (time.Time, error) {
	tf, ok := timeFormat[strings.ToLower(format)]
	if !ok {
		return time.Time{}, fmt.Errorf("format %s not support", format)
	}

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return time.Time{}, err
		}

		return time.ParseInLocation(tf, str, loc)
	}

	return time.Parse(tf, str)
}

// BeginOfMinute return beginning minute time of day.
// Play: https://go.dev/play/p/ieOLVJ9CiFT
func BeginOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute return end minute time of day.
// Play: https://go.dev/play/p/yrL5wGzPj4z
func EndOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfHour return beginning hour time of day.
// Play: https://go.dev/play/p/GhdGFnDWpYs
func BeginOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour return end hour time of day.
// Play: https://go.dev/play/p/6ce3j_6cVqN
func EndOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfDay return beginning hour time of day.
// Play: https://go.dev/play/p/94m_UT6cWs9
func BeginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay return end time of day.
// Play: https://go.dev/play/p/eMBOvmq5Ih1
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfWeek return beginning week, default week begin from Sunday.
// Play: https://go.dev/play/p/ynjoJPz7VNV
func BeginOfWeek(t time.Time, beginFrom ...time.Weekday) time.Time {
	var beginFromWeekday = time.Sunday
	if len(beginFrom) > 0 {
		beginFromWeekday = beginFrom[0]
	}
	y, m, d := t.AddDate(0, 0, int(beginFromWeekday-t.Weekday())).Date()
	beginOfWeek := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	if beginOfWeek.After(t) {
		return beginOfWeek.AddDate(0, 0, -7)
	}
	return beginOfWeek
}

// EndOfWeek return end week time, default week end with Saturday.
// Play: https://go.dev/play/p/i08qKXD9flf
func EndOfWeek(t time.Time, endWith ...time.Weekday) time.Time {
	var endWithWeekday = time.Saturday
	if len(endWith) > 0 {
		endWithWeekday = endWith[0]
	}
	y, m, d := t.AddDate(0, 0, int(endWithWeekday-t.Weekday())).Date()
	var endWithWeek = time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
	if endWithWeek.Before(t) {
		endWithWeek = endWithWeek.AddDate(0, 0, 7)
	}
	return endWithWeek
}

// BeginOfMonth return beginning of month.
// Play: https://go.dev/play/p/bWXVFsmmzwL
func BeginOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth return end of month.
// Play: https://go.dev/play/p/_GWh10B3Nqi
func EndOfMonth(t time.Time) time.Time {
	return BeginOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfYear return the date time at the begin of year.
// Play: https://go.dev/play/p/i326DSwLnV8
func BeginOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear return the date time at the end of year.
// Play: https://go.dev/play/p/G01cKlMCvNm
func EndOfYear(t time.Time) time.Time {
	return BeginOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// IsLeapYear check if param year is leap year or not.
// Play: https://go.dev/play/p/xS1eS2ejGew
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// BetweenSeconds returns the number of seconds between two times.
// Play: https://go.dev/play/p/n3YDRyfyXJu
func BetweenSeconds(t1 time.Time, t2 time.Time) int64 {
	index := t2.Unix() - t1.Unix()
	return index
}

// DayOfYear returns which day of the year the parameter date `t` is.
// Play: https://go.dev/play/p/0hjqhTwFNlH
func DayOfYear(t time.Time) int {
	y, m, d := t.Date()
	firstDay := time.Date(y, 1, 1, 0, 0, 0, 0, t.Location())
	nowDate := time.Date(y, m, d, 0, 0, 0, 0, t.Location())

	return int(nowDate.Sub(firstDay).Hours() / 24)
}

// IsWeekend checks if passed time is weekend or not.
// Play: https://go.dev/play/p/cupRM5aZOIY
// Deprecated Use '== Weekday' instead
func IsWeekend(t time.Time) bool {
	return time.Saturday == t.Weekday() || time.Sunday == t.Weekday()
}

// NowDateOrTime return current datetime with specific format and timezone.
// Play: https://go.dev/play/p/EZ-begEjtT0
func NowDateOrTime(format string, timezone ...string) string {
	tf, ok := timeFormat[strings.ToLower(format)]
	if !ok {
		return ""
	}

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return ""
		}

		return time.Now().In(loc).Format(tf)
	}

	return time.Now().Format(tf)
}

// Timestamp return current second timestamp.
// Play: https://go.dev/play/p/iU5b7Vvjx6x
func Timestamp(timezone ...string) int64 {
	t := time.Now()

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}

		t = t.In(loc)
	}

	return t.Unix()
}

// TimestampMilli return current mill second timestamp.
// Play: https://go.dev/play/p/4gvEusOTu1T
func TimestampMilli(timezone ...string) int64 {
	t := time.Now()

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}
		t = t.In(loc)
	}

	return int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
}

// TimestampMicro return current micro second timestamp.
// Play: https://go.dev/play/p/2maANglKHQE
func TimestampMicro(timezone ...string) int64 {
	t := time.Now()

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}
		t = t.In(loc)
	}

	return int64(time.Nanosecond) * t.UnixNano() / int64(time.Microsecond)
}

// TimestampNano return current nano second timestamp.
// Play: https://go.dev/play/p/A9Oq_COrcCF
func TimestampNano(timezone ...string) int64 {
	t := time.Now()

	if timezone != nil && timezone[0] != "" {
		loc, err := time.LoadLocation(timezone[0])
		if err != nil {
			return 0
		}
		t = t.In(loc)
	}

	return t.UnixNano()
}

// TrackFuncTime track the time of function execution.
// call it at top of the func like `defer TrackFuncTime(time.Now())()`
// Play: https://go.dev/play/p/QBSEdfXHPTp
func TrackFuncTime(pre time.Time) func() {
	callerName := getCallerName()
	return func() {
		elapsed := time.Since(pre)
		fmt.Printf("Function %s execution time:\t %v", callerName, elapsed)
	}
}

func getCallerName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "Unknown"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "Unknown"
	}

	fullName := fn.Name()
	if lastDot := strings.LastIndex(fullName, "."); lastDot != -1 {
		return fullName[lastDot+1:]
	}

	return fullName
}

// DaysBetween returns the number of days between two times.
// Play: https://go.dev/play/p/qD6qGb3TbOy
func DaysBetween(start, end time.Time) int {
	duration := end.Sub(start)
	days := int(duration.Hours() / 24)

	return days
}

// GenerateDatetimesBetween returns a slice of strings between two times.
// layout: the format of the datetime string
// interval: the interval between two datetimes
// Play: https://go.dev/play/p/6kHBpAxD9ZC
func GenerateDatetimesBetween(start, end time.Time, layout string, interval string) ([]string, error) {
	var result []string

	if start.After(end) {
		start, end = end, start
	}

	duration, err := time.ParseDuration(interval)
	if err != nil {
		return nil, err
	}

	for current := start; !current.After(end); current = current.Add(duration) {
		result = append(result, current.Format(layout))
	}

	return result, nil
}

// Min returns the earliest time among the given times.
// Play: https://go.dev/play/p/MCIDvHNOGGb
func Min(t1 time.Time, times ...time.Time) time.Time {
	minTime := t1

	for _, t := range times {
		if t.Before(minTime) {
			minTime = t
		}
	}

	return minTime
}

// Max returns the latest time among the given times.
// Play: https://go.dev/play/p/9m6JMk1LB7-
func Max(t1 time.Time, times ...time.Time) time.Time {
	maxTime := t1

	for _, t := range times {
		if t.After(maxTime) {
			maxTime = t
		}
	}

	return maxTime
}

// MaxMin returns the latest and earliest time among the given times.
// Play: https://go.dev/play/p/rbW51cDtM_2
func MaxMin(t1 time.Time, times ...time.Time) (maxTime time.Time, minTime time.Time) {
	maxTime = t1
	minTime = t1

	for _, t := range times {
		if t.Before(minTime) {
			minTime = t
		}

		if t.After(maxTime) {
			maxTime = t
		}
	}

	return maxTime, minTime
}
