// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package datetime implements some functions to format date and time.
// Note:
// 1. `format` param in FormatTimeToStr function should be as flow:
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
// "yyyy"
// "mm"
// "hh:mm:ss"
// "mm:ss"
package datetime

import (
	"fmt"
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
		"yyyy":                "2006",
		"mm":                  "01",
		"hh:mm:ss":            "15:04:05",
		"mm:ss":               "04:05",
	}
}

// AddMinute add or sub minute to the time.
// Play: https://go.dev/play/p/nT1heB1KUUK
func AddMinute(t time.Time, minute int64) time.Time {
	return t.Add(time.Minute * time.Duration(minute))
}

// AddHour add or sub hour to the time.
// Play: https://go.dev/play/p/rcMjd7OCsi5
func AddHour(t time.Time, hour int64) time.Time {
	return t.Add(time.Hour * time.Duration(hour))
}

// AddDay add or sub day to the time.
// Play: https://go.dev/play/p/dIGbs_uTdFa
func AddDay(t time.Time, day int64) time.Time {
	return t.Add(24 * time.Hour * time.Duration(day))
}

// AddYear add or sub year to the time.
// Play: https://go.dev/play/p/MqW2ujnBx10
func AddYear(t time.Time, year int64) time.Time {
	return t.Add(365 * 24 * time.Hour * time.Duration(year))
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
func FormatTimeToStr(t time.Time, format string) string {
	return t.Format(timeFormat[format])
}

// FormatStrToTime convert string to time.
// Play: https://go.dev/play/p/1h9FwdU8ql4
func FormatStrToTime(str, format string) (time.Time, error) {
	v, ok := timeFormat[format]
	if !ok {
		return time.Time{}, fmt.Errorf("format %s not found", format)
	}

	return time.Parse(v, str)
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
func IsWeekend(t time.Time) bool {
	return time.Saturday == t.Weekday() || time.Sunday == t.Weekday()
}
