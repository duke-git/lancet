package datetime

import (
	"fmt"
	"reflect"
	"time"
)

func ExampleAddDay() {
	date, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 00:00:00")

	after1Day := AddDay(date, 1)
	before1Day := AddDay(date, -1)

	fmt.Println(after1Day.Format("2006-01-02 15:04:05"))
	fmt.Println(before1Day.Format("2006-01-02 15:04:05"))

	// Output:
	// 2021-01-02 00:00:00
	// 2020-12-31 00:00:00
}

func ExampleAddWeek() {
	date, _ := time.Parse("2006-01-02", "2021-01-01")

	after2Weeks := AddWeek(date, 2)
	before2Weeks := AddWeek(date, -2)

	fmt.Println(after2Weeks.Format("2006-01-02"))
	fmt.Println(before2Weeks.Format("2006-01-02"))

	// Output:
	// 2021-01-15
	// 2020-12-18
}

func ExampleAddMonth() {
	date, _ := time.Parse("2006-01-02", "2021-01-01")

	after2Months := AddMonth(date, 2)
	before2Months := AddMonth(date, -2)

	fmt.Println(after2Months.Format("2006-01-02"))
	fmt.Println(before2Months.Format("2006-01-02"))

	// Output:
	// 2021-03-01
	// 2020-11-01
}

func ExampleAddHour() {
	date, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 00:00:00")

	after2Hours := AddHour(date, 2)
	before2Hours := AddHour(date, -2)

	fmt.Println(after2Hours.Format("2006-01-02 15:04:05"))
	fmt.Println(before2Hours.Format("2006-01-02 15:04:05"))

	// Output:
	// 2021-01-01 02:00:00
	// 2020-12-31 22:00:00
}

func ExampleAddMinute() {
	date, _ := time.Parse("2006-01-02 15:04:05", "2021-01-01 00:00:00")

	after2Minutes := AddMinute(date, 2)
	before2Minutes := AddMinute(date, -2)

	fmt.Println(after2Minutes.Format("2006-01-02 15:04:05"))
	fmt.Println(before2Minutes.Format("2006-01-02 15:04:05"))

	// Output:
	// 2021-01-01 00:02:00
	// 2020-12-31 23:58:00
}

func ExampleAddYear() {
	date, _ := time.Parse("2006-01-02", "2021-01-01")

	after2Years := AddYear(date, 2)
	before2Years := AddYear(date, -2)

	fmt.Println(after2Years.Format("2006-01-02"))
	fmt.Println(before2Years.Format("2006-01-02"))

	// Output:
	// 2023-01-01
	// 2019-01-01
}

func ExampleAddDaySafe() {
	leapYearDate1, _ := time.Parse("2006-01-02", "2024-02-29")
	result1 := AddDaySafe(leapYearDate1, 1)

	leapYearDate2, _ := time.Parse("2006-01-02", "2024-03-01")
	result2 := AddDaySafe(leapYearDate2, -1)

	nonLeapYearDate1, _ := time.Parse("2006-01-02", "2025-02-28")
	result3 := AddDaySafe(nonLeapYearDate1, 1)

	nonLeaYearDate2, _ := time.Parse("2006-01-02", "2025-03-01")
	result4 := AddDaySafe(nonLeaYearDate2, -1)

	fmt.Println(result1.Format("2006-01-02"))
	fmt.Println(result2.Format("2006-01-02"))
	fmt.Println(result3.Format("2006-01-02"))
	fmt.Println(result4.Format("2006-01-02"))

	// Output:
	// 2024-03-01
	// 2024-02-29
	// 2025-03-01
	// 2025-02-28
}

func ExampleAddMonthSafe() {
	date1, _ := time.Parse("2006-01-02", "2025-01-31")
	result1 := AddMonthSafe(date1, 1)

	date2, _ := time.Parse("2006-01-02", "2024-02-29")
	result2 := AddMonthSafe(date2, -1)

	fmt.Println(result1.Format("2006-01-02"))
	fmt.Println(result2.Format("2006-01-02"))

	// Output:
	// 2025-02-28
	// 2024-01-29
}

func ExampleAddYearSafe() {
	date, _ := time.Parse("2006-01-02", "2020-02-29")

	result1 := AddYearSafe(date, 1)
	result2 := AddYearSafe(date, -1)

	fmt.Println(result1.Format("2006-01-02"))
	fmt.Println(result2.Format("2006-01-02"))

	// Output:
	// 2021-02-28
	// 2019-02-28
}

func ExampleGetNowDate() {
	result := GetNowDate()

	expected := time.Now().Format("2006-01-02")

	fmt.Println(result == expected)

	// Output:
	// true
}

func ExampleGetNowTime() {
	result := GetNowTime()

	expected := time.Now().Format("15:04:05")

	fmt.Println(result == expected)

	// Output:
	// true
}

func ExampleGetNowDateTime() {
	result := GetNowDateTime()

	expected := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(result == expected)

	// Output:
	// true
}

// func ExampleGetZeroHourTimestamp() {
// 	ts := GetZeroHourTimestamp()

// 	fmt.Println(ts)

// 	// Output:
// 	// 1673107200
// }

// func ExampleGetNightTimestamp() {
// 	ts := GetNightTimestamp()

// 	fmt.Println(ts)

// 	// Output:
// 	// 1673193599
// }

func ExampleFormatTimeToStr() {
	datetime, _ := time.Parse("2006-01-02 15:04:05", "2021-01-02 16:04:08")

	result1 := FormatTimeToStr(datetime, "yyyy-mm-dd hh:mm:ss")
	result2 := FormatTimeToStr(datetime, "yyyy-mm-dd")
	result3 := FormatTimeToStr(datetime, "dd-mm-yy hh:mm:ss")
	result4 := FormatTimeToStr(datetime, "yyyy-mm-dd hh")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 2021-01-02 16:04:08
	// 2021-01-02
	// 02-01-21 16:04:08
	// 2021-01-02 16
}

func ExampleFormatStrToTime() {
	result1, _ := FormatStrToTime("2021-01-02 16:04:08", "yyyy-mm-dd hh:mm:ss")
	result2, _ := FormatStrToTime("2021-01-02", "yyyy-mm-dd")
	result3, _ := FormatStrToTime("02-01-21 16:04:08", "dd-mm-yy hh:mm:ss")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 2021-01-02 16:04:08 +0000 UTC
	// 2021-01-02 00:00:00 +0000 UTC
	// 2021-01-02 16:04:08 +0000 UTC
}

func ExampleBeginOfMinute() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := BeginOfMinute(input)

	fmt.Println(result)

	// Output:
	// 2023-01-08 18:50:00 +0000 UTC
}

func ExampleEndOfMinute() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := EndOfMinute(input)

	fmt.Println(result)

	// Output:
	// 2023-01-08 18:50:59.999999999 +0000 UTC
}

func ExampleBeginOfHour() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := BeginOfHour(input)

	fmt.Println(result)

	// Output:
	// 2023-01-08 18:00:00 +0000 UTC
}

func ExampleEndOfHour() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := EndOfHour(input)

	fmt.Println(result)

	// Output:
	// 2023-01-08 18:59:59.999999999 +0000 UTC
}

func ExampleBeginOfDay() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := BeginOfDay(input)

	fmt.Println(result)

	// Output:
	// 2023-01-08 00:00:00 +0000 UTC
}

func ExampleEndOfDay() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := EndOfDay(input)

	fmt.Println(result)

	// Output:
	// 2023-01-08 23:59:59.999999999 +0000 UTC
}

func ExampleBeginOfWeek() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := BeginOfWeek(input)

	fmt.Println(result)

	// Output:
	// 2023-01-08 00:00:00 +0000 UTC
}

func ExampleEndOfWeek() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := EndOfWeek(input)

	fmt.Println(result)

	// Output:
	// 2023-01-14 23:59:59.999999999 +0000 UTC
}

func ExampleBeginOfMonth() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := BeginOfMonth(input)

	fmt.Println(result)

	// Output:
	// 2023-01-01 00:00:00 +0000 UTC
}

func ExampleEndOfMonth() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := EndOfMonth(input)

	fmt.Println(result)

	// Output:
	// 2023-01-31 23:59:59.999999999 +0000 UTC
}

func ExampleBeginOfYear() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := BeginOfYear(input)

	fmt.Println(result)

	// Output:
	// 2023-01-01 00:00:00 +0000 UTC
}

func ExampleEndOfYear() {
	input := time.Date(2023, 1, 8, 18, 50, 10, 100, time.UTC)

	result := EndOfYear(input)

	fmt.Println(result)

	// Output:
	// 2023-12-31 23:59:59.999999999 +0000 UTC
}

func ExampleNewUnix() {
	result := NewUnix(1647597438)

	fmt.Println(result)

	// Output:
	// &{1647597438}
}

func ExampleNewUnixNow() {
	tm1 := NewUnixNow()

	unixTimestamp := tm1.ToUnix()

	tm2 := NewUnix(unixTimestamp)

	fmt.Println(reflect.DeepEqual(tm1, tm2))

	// Output:
	// true
}

// func ExampleNewFormat() {
// 	tm, err := NewFormat("2022-03-18 17:04:05")
// 	if err != nil {
// 		return
// 	}

// 	result := tm.ToFormat()

// 	fmt.Println(result)

// 	// Output:
// 	// 2022-03-18 17:04:05
// }

// func ExampleNewISO8601() {
// 	tm, err := NewISO8601("2006-01-02T15:04:05.999Z")
// 	if err != nil {
// 		return
// 	}

// 	result := tm.ToIso8601()

// 	fmt.Println(result)

// 	// Output:
// 	// 2006-01-02T23:04:05+08:00
// }

func ExampleIsLeapYear() {
	result1 := IsLeapYear(2000)
	result2 := IsLeapYear(2001)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleBetweenSeconds() {
	today := time.Now()
	tomorrow := AddDay(today, 1)
	yesterday := AddDay(today, -1)

	result1 := BetweenSeconds(today, tomorrow)
	result2 := BetweenSeconds(today, yesterday)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 86400
	// -86400
}

func ExampleDayOfYear() {
	date1 := time.Date(2023, 02, 01, 1, 1, 1, 0, time.Local)
	result1 := DayOfYear(date1)

	date2 := time.Date(2023, 01, 02, 1, 1, 1, 0, time.Local)
	result2 := DayOfYear(date2)

	date3 := time.Date(2023, 01, 01, 1, 1, 1, 0, time.Local)
	result3 := DayOfYear(date3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 31
	// 1
	// 0
}

func ExampleIsWeekend() {
	date1 := time.Date(2023, 06, 03, 0, 0, 0, 0, time.Local)
	date2 := time.Date(2023, 06, 04, 0, 0, 0, 0, time.Local)
	date3 := time.Date(2023, 06, 02, 0, 0, 0, 0, time.Local)

	result1 := IsWeekend(date1)
	result2 := IsWeekend(date2)
	result3 := IsWeekend(date3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// true
	// false
}

func ExampleDaysBetween() {
	start := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, time.September, 10, 0, 0, 0, 0, time.UTC)

	result := DaysBetween(start, end)

	fmt.Println(result)

	// Output:
	// 9
}

func ExampleGenerateDatetimesBetween() {
	start := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, time.September, 1, 2, 0, 0, 0, time.UTC)

	layout := "2006-01-02 15:04:05"
	interval := "1h"

	result, err := GenerateDatetimesBetween(start, end, layout, interval)

	fmt.Println(result)
	fmt.Println(err)

	// Output:
	// [2024-09-01 00:00:00 2024-09-01 01:00:00 2024-09-01 02:00:00]
	// <nil>
}

func ExampleMin() {
	result := Min(time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.September, 2, 0, 0, 0, 0, time.UTC))

	fmt.Println(result)

	// Output:
	// 2024-09-01 00:00:00 +0000 UTC
}

func ExampleMax() {
	result := Max(time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.September, 2, 0, 0, 0, 0, time.UTC))

	fmt.Println(result)

	// Output:
	// 2024-09-02 00:00:00 +0000 UTC
}

func ExampleMaxMin() {
	max, min := MaxMin(time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, time.September, 2, 0, 0, 0, 0, time.UTC), time.Date(2024, time.September, 3, 0, 0, 0, 0, time.UTC))

	fmt.Println(max)
	fmt.Println(min)

	// Output:
	// 2024-09-03 00:00:00 +0000 UTC
	// 2024-09-01 00:00:00 +0000 UTC
}
