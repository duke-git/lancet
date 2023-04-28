package datetime

import (
	"fmt"
	"reflect"
	"time"
)

func ExampleAddDay() {
	now := time.Now()

	tomorrow := AddDay(now, 1)
	diff1 := tomorrow.Sub(now)

	yesterday := AddDay(now, -1)
	diff2 := yesterday.Sub(now)

	fmt.Println(diff1)
	fmt.Println(diff2)

	// Output:
	// 24h0m0s
	// -24h0m0s
}

func ExampleAddHour() {
	now := time.Now()

	after2Hours := AddHour(now, 2)
	diff1 := after2Hours.Sub(now)

	before2Hours := AddHour(now, -2)
	diff2 := before2Hours.Sub(now)

	fmt.Println(diff1)
	fmt.Println(diff2)

	// Output:
	// 2h0m0s
	// -2h0m0s
}

func ExampleAddMinute() {
	now := time.Now()

	after2Minutes := AddMinute(now, 2)
	diff1 := after2Minutes.Sub(now)

	before2Minutes := AddMinute(now, -2)
	diff2 := before2Minutes.Sub(now)

	fmt.Println(diff1)
	fmt.Println(diff2)

	// Output:
	// 2m0s
	// -2m0s
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

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 2021-01-02 16:04:08
	// 2021-01-02
	// 02-01-21 16:04:08
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
