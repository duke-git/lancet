// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package datetime implements some functions to format date and time.

// Note:
// 1. `format` param in FormatTimeToStr function should be as flow:
//"yyyy-mm-dd hh:mm:ss"
//"yyyy-mm-dd hh:mm"
//"yyyy-mm-dd hh"
//"yyyy-mm-dd"
//"yyyy-mm"
//"mm-dd"
//"dd-mm-yy hh:mm:ss"
//"yyyy/mm/dd hh:mm:ss"
//"yyyy/mm/dd hh:mm"
//"yyyy/mm/dd hh"
//"yyyy/mm/dd"
//"yyyy/mm"
//"mm/dd"
//"dd/mm/yy hh:mm:ss"
//"yyyy"
//"mm"
//"hh:mm:ss"
//"mm:ss"

package datetime

import (
	"strconv"
	"time"
)

var timeFormat map[string]string

func init() {
	timeFormat = map[string]string{
		"yyyy-mm-dd hh:mm:ss": "2006-01-02 15:04:05",
		"yyyy-mm-dd hh:mm":    "2006-01-02 15:04",
		"yyyy-mm-dd hh":       "2006-01-02 15:04",
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

// AddMinute add or sub minute to the time
func AddMinute(t time.Time, minute int64) time.Time {
	s := strconv.FormatInt(minute, 10)
	m, _ := time.ParseDuration(s + "m")
	return t.Add(m)
}

// AddHour add or sub hour to the time
func AddHour(t time.Time, hour int64) time.Time {
	s := strconv.FormatInt(hour, 10)
	h, _ := time.ParseDuration(s + "h")
	return t.Add(h)
}

// AddDay add or sub day to the time
func AddDay(t time.Time, day int64) time.Time {
	dayHours := day * 24
	d := strconv.FormatInt(dayHours, 10)
	h, _ := time.ParseDuration(d + "h")
	return t.Add(h)
}

// GetNowDate return format yyyy-mm-dd of current date
func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

// GetNowTime return format hh-mm-ss of current time
func GetNowTime() string {
	return time.Now().Format("15:04:05")
}

// GetNowDateTime return format yyyy-mm-dd hh-mm-ss of current datetime
func GetNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetZeroHourTimestamp return timestamp of zero hour (timestamp of 00:00)
func GetZeroHourTimestamp() int64 {
	ts := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", ts)
	return t.UTC().Unix() - 8*3600
}

// GetNightTimestamp return timestamp of zero hour (timestamp of 23:59)
func GetNightTimestamp() int64 {
	return GetZeroHourTimestamp() + 86400 - 1
}

// FormatTimeToStr convert time to string
func FormatTimeToStr(t time.Time, format string) string {
	return t.Format(timeFormat[format])
}

// FormatStrToTime convert string to time
func FormatStrToTime(str, format string) time.Time {
	t, _ := time.Parse(timeFormat[format], str)
	return t
}
