package datetime

import (
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestAddYear(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestAddDay")

	tests := []struct {
		inputDate string
		years     int
		expected  string
	}{
		{
			inputDate: "2021-01-01 00:00:00",
			years:     1,
			expected:  "2022-01-01 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			years:     -1,
			expected:  "2020-01-01 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			years:     0,
			expected:  "2021-01-01 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			years:     2,
			expected:  "2023-01-01 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			years:     3,
			expected:  "2024-01-01 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			years:     4,
			expected:  "2025-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		date, _ := time.Parse("2006-01-02 15:04:05", tt.inputDate)
		result := AddYear(date, int64(tt.years))
		assert.Equal(tt.expected, result.Format("2006-01-02 15:04:05"))
	}
}

func TestAddDay(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestAddDay")

	tests := []struct {
		inputDate string
		days      int
		expected  string
	}{
		{
			inputDate: "2021-01-01 00:00:00",
			days:      1,
			expected:  "2021-01-02 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			days:      -1,
			expected:  "2020-12-31 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			days:      0,
			expected:  "2021-01-01 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			days:      2,
			expected:  "2021-01-03 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			days:      3,
			expected:  "2021-01-04 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			days:      4,
			expected:  "2021-01-05 00:00:00",
		},
	}

	for _, tt := range tests {
		date, _ := time.Parse("2006-01-02 15:04:05", tt.inputDate)
		result := AddDay(date, int64(tt.days))
		assert.Equal(tt.expected, result.Format("2006-01-02 15:04:05"))
	}
}

func TestAddHour(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestAddHour")

	tests := []struct {
		inputDate string
		hours     int
		expected  string
	}{
		{
			inputDate: "2021-01-01 00:00:00",
			hours:     1,
			expected:  "2021-01-01 01:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			hours:     -1,
			expected:  "2020-12-31 23:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			hours:     0,
			expected:  "2021-01-01 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			hours:     24,
			expected:  "2021-01-02 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			hours:     25,
			expected:  "2021-01-02 01:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			hours:     48,
			expected:  "2021-01-03 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			hours:     49,
			expected:  "2021-01-03 01:00:00",
		},
	}

	for _, tt := range tests {
		date, _ := time.Parse("2006-01-02 15:04:05", tt.inputDate)
		result := AddHour(date, int64(tt.hours))
		assert.Equal(tt.expected, result.Format("2006-01-02 15:04:05"))
	}

}

func TestAddMinute(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestAddMinute")

	tests := []struct {
		inputDate string
		minutes   int
		expected  string
	}{
		{
			inputDate: "2021-01-01 00:00:00",
			minutes:   1,
			expected:  "2021-01-01 00:01:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			minutes:   -1,
			expected:  "2020-12-31 23:59:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			minutes:   0,
			expected:  "2021-01-01 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			minutes:   60,
			expected:  "2021-01-01 01:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			minutes:   61,
			expected:  "2021-01-01 01:01:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			minutes:   1440,
			expected:  "2021-01-02 00:00:00",
		},
		{
			inputDate: "2021-01-01 00:00:00",
			minutes:   1441,
			expected:  "2021-01-02 00:01:00",
		},
	}

	for _, tt := range tests {
		date, _ := time.Parse("2006-01-02 15:04:05", tt.inputDate)
		result := AddMinute(date, int64(tt.minutes))
		assert.Equal(tt.expected, result.Format("2006-01-02 15:04:05"))
	}
}

func TestAddWeek(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAddWeek")

	tests := []struct {
		inputDate string
		weeks     int
		expected  string
	}{
		{
			inputDate: "2021-01-01",
			weeks:     1,
			expected:  "2021-01-08",
		},
		{
			inputDate: "2021-01-01",
			weeks:     -1,
			expected:  "2020-12-25",
		},
		{
			inputDate: "2021-01-01",
			weeks:     0,
			expected:  "2021-01-01",
		},
		{
			inputDate: "2021-01-01",
			weeks:     52,
			expected:  "2021-12-31",
		},
		{
			inputDate: "2021-01-01",
			weeks:     53,
			expected:  "2022-01-07",
		},
		{
			inputDate: "2021-01-01",
			weeks:     104,
			expected:  "2022-12-30",
		},
	}

	for _, tt := range tests {
		date, _ := time.Parse("2006-01-02", tt.inputDate)
		result := AddWeek(date, int64(tt.weeks))
		assert.Equal(tt.expected, result.Format("2006-01-02"))
	}
}

func TestAddMonth(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAddMonth")

	tests := []struct {
		inputDate string
		months    int
		expected  string
	}{
		{
			inputDate: "2021-01-01",
			months:    1,
			expected:  "2021-02-01",
		},
		{
			inputDate: "2021-01-01",
			months:    -1,
			expected:  "2020-12-01",
		},
		{
			inputDate: "2021-01-01",
			months:    0,
			expected:  "2021-01-01",
		},
		{
			inputDate: "2021-01-01",
			months:    12,
			expected:  "2022-01-01",
		},
		{
			inputDate: "2021-01-01",
			months:    13,
			expected:  "2022-02-01",
		},
		{
			inputDate: "2021-01-01",
			months:    24,
			expected:  "2023-01-01",
		},
	}

	for _, tt := range tests {
		date, _ := time.Parse("2006-01-02", tt.inputDate)
		result := AddMonth(date, int64(tt.months))
		assert.Equal(tt.expected, result.Format("2006-01-02"))
	}

}

func TestAddDaySafe(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAddDaySafe")

	tests := []struct {
		inputDate string
		days      int
		expected  string
	}{
		{"2025-01-31", 10, "2025-02-10"},
		{"2025-01-01", 30, "2025-01-31"},
		{"2025-01-31", 1, "2025-02-01"},
		{"2025-02-28", 1, "2025-03-01"},
		{"2024-02-29", 1, "2024-03-01"},
		{"2024-02-29", 365, "2025-02-28"},

		{"2025-01-31", -10, "2025-01-21"},
		{"2025-01-01", -30, "2024-12-02"},
		{"2025-02-01", -1, "2025-01-31"},
		{"2025-03-01", -1, "2025-02-28"},
		{"2024-03-01", -1, "2024-02-29"},

		{"2025-01-31", -31, "2024-12-31"},
		{"2025-12-31", 1, "2026-01-01"},
	}

	for _, tt := range tests {
		date, _ := time.Parse("2006-01-02", tt.inputDate)
		result := AddDaySafe(date, tt.days)
		assert.Equal(tt.expected, result.Format("2006-01-02"))
	}
}

func TestAddMonthSafe(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAddMonthSafe")

	tests := []struct {
		inputDate string
		months    int
		expected  string
	}{
		{
			inputDate: "2025-01-31",
			months:    1,
			expected:  "2025-02-28",
		},
		{
			inputDate: "2025-01-31",
			months:    -1,
			expected:  "2024-12-31",
		},
		{
			inputDate: "2025-12-31",
			months:    1,
			expected:  "2026-01-31",
		},
		{
			inputDate: "2025-01-31",
			months:    -1,
			expected:  "2024-12-31",
		},
		{
			inputDate: "2024-02-29",
			months:    1,
			expected:  "2024-03-29",
		},
		{
			inputDate: "2024-02-29",
			months:    -1,
			expected:  "2024-01-29",
		},
	}

	for _, tt := range tests {
		date, _ := time.Parse("2006-01-02", tt.inputDate)
		result := AddMonthSafe(date, tt.months)
		assert.Equal(tt.expected, result.Format("2006-01-02"))
	}

}

func TestAddYearSafe(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAddYearSafe")

	tests := []struct {
		inputDate string
		years     int
		expected  string
	}{
		{
			inputDate: "2020-02-29",
			years:     1,
			expected:  "2021-02-28",
		},
		{
			inputDate: "2020-02-29",
			years:     2,
			expected:  "2022-02-28",
		},
		{
			inputDate: "2020-02-29",
			years:     -1,
			expected:  "2019-02-28",
		},
		{
			inputDate: "2020-02-29",
			years:     -2,
			expected:  "2018-02-28",
		},
	}

	for _, tt := range tests {
		date, _ := time.Parse("2006-01-02", tt.inputDate)
		result := AddYearSafe(date, tt.years)
		assert.Equal(tt.expected, result.Format("2006-01-02"))
	}
}

func TestGetNowDate(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestGetNowDate")
	expected := time.Now().Format("2006-01-02")
	assert.Equal(expected, GetNowDate())
}

func TestGetNowTime(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestGetNowTime")
	expected := time.Now().Format("15:04:05")
	assert.Equal(expected, GetNowTime())
}

func TestGetNowDateTime(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestGetNowDateTime")
	expected := time.Now().Format("2006-01-02 15:04:05")
	assert.Equal(expected, GetNowDateTime())
}

func TestGetTodayStartTime(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestGetTodayStartTime")
	expected := time.Now().Format("2006-01-02") + " 00:00:00"
	assert.Equal(expected, GetTodayStartTime())
}

func TestGetTodayEndTime(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestGetTodayEndTime")
	expected := time.Now().Format("2006-01-02") + " 23:59:59"
	assert.Equal(expected, GetTodayEndTime())
}

func TestFormatTimeToStr(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFormatTimeToStr")

	datetime, _ := time.Parse("2006-01-02 15:04:05", "2021-01-02 16:04:08")
	cases := []string{
		"yyyy-mm-dd hh:mm:ss", "yyyy-mm-dd",
		"dd-mm-yy hh:mm:ss", "yyyy/mm/dd hh:mm:ss",
		"hh:mm:ss", "yyyy/mm",
		"yyyy-mm-dd hh",
	}

	expected := []string{
		"2021-01-02 16:04:08", "2021-01-02",
		"02-01-21 16:04:08", "2021/01/02 16:04:08",
		"16:04:08", "2021/01",
		"2021-01-02 16",
	}

	for i := 0; i < len(cases); i++ {
		actual := FormatTimeToStr(datetime, cases[i])
		assert.Equal(expected[i], actual)
	}

	ds := FormatTimeToStr(datetime, "yyyy-mm-dd hh:mm:ss", "EST")
	t.Log(ds)
}

func TestFormatStrToTime(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFormatStrToTime")

	formats := []string{
		"2006-01-02 15:04:05", "2006-01-02",
		"02-01-06 15:04:05", "2006/01/02 15:04:05",
		"2006/01"}
	cases := []string{
		"yyyy-mm-dd hh:mm:ss", "yyyy-mm-dd",
		"dd-mm-yy hh:mm:ss", "yyyy/mm/dd hh:mm:ss",
		"yyyy/mm"}

	expected := []string{
		"2021-01-02 16:04:08", "2021-01-02",
		"02-01-21 16:04:08", "2021/01/02 16:04:08",
		"2021/01"}

	for i := 0; i < len(cases); i++ {
		actual, err := FormatStrToTime(expected[i], cases[i])
		if err != nil {
			t.Fatal(err)
		}
		expected, _ := time.Parse(formats[i], expected[i])
		assert.Equal(expected, actual)
	}

	estTime, err := FormatStrToTime("2021-01-02 16:04:08", "yyyy-mm-dd hh:mm:ss", "EST")
	t.Log(estTime)
	assert.IsNil(err)
}

func TestBeginOfMinute(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBeginOfMinute")

	expected := time.Date(2022, 2, 15, 15, 48, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfMinute(td)

	assert.Equal(expected, actual)
}

func TestEndOfMinute(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEndOfMinute")

	expected := time.Date(2022, 2, 15, 15, 48, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfMinute(td)

	assert.Equal(expected, actual)
}

func TestBeginOfHour(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBeginOfHour")

	expected := time.Date(2022, 2, 15, 15, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfHour(td)

	assert.Equal(expected, actual)
}

func TestEndOfHour(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEndOfHour")

	expected := time.Date(2022, 2, 15, 15, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfHour(td)

	assert.Equal(expected, actual)
}

func TestBeginOfDay(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBeginOfDay")

	expected := time.Date(2022, 2, 15, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfDay(td)

	assert.Equal(expected, actual)
}

func TestEndOfDay(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEndOfDay")

	expected := time.Date(2022, 2, 15, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfDay(td)

	assert.Equal(expected, actual)
}

func TestBeginOfWeek(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBeginOfWeek")

	expected := time.Date(2022, 2, 13, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfWeek(td)

	assert.Equal(expected, actual)
}

func TestEndOfWeek(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEndOfWeek")

	expected := time.Date(2022, 2, 19, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfWeek(td)

	assert.Equal(expected, actual)
}

func TestBeginOfMonth(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBeginOfMonth")

	expected := time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfMonth(td)

	assert.Equal(expected, actual)
}

func TestEndOfMonth(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEndOfMonth")

	expected := time.Date(2022, 2, 28, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfMonth(td)

	assert.Equal(expected, actual)
}

func TestBeginOfYear(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBeginOfYear")

	expected := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfYear(td)

	assert.Equal(expected, actual)
}

func TestEndOfYear(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEndOfYear")

	expected := time.Date(2022, 12, 31, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfYear(td)

	assert.Equal(expected, actual)
}

func TestIsLeapYear(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEndOfYear")

	result1 := IsLeapYear(2000)
	result2 := IsLeapYear(2001)

	assert.Equal(true, result1)
	assert.Equal(false, result2)
}

func TestDayOfYear(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDayOfYear")
	date1 := time.Date(2023, 02, 01, 1, 1, 1, 0, time.Local)
	result1 := DayOfYear(date1)
	assert.Equal(31, result1)

	date2 := time.Date(2023, 01, 02, 1, 1, 1, 0, time.Local)
	result2 := DayOfYear(date2)
	assert.Equal(1, result2)

	date3 := time.Date(2023, 01, 01, 1, 1, 1, 0, time.Local)
	result3 := DayOfYear(date3)
	assert.Equal(0, result3)
}

func TestIsWeekend(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsWeekend")

	date := time.Date(2023, 06, 03, 0, 0, 0, 0, time.Local)
	result := IsWeekend(date)
	assert.Equal(true, result)

	date1 := time.Date(2023, 06, 04, 0, 0, 0, 0, time.Local)
	result1 := IsWeekend(date1)
	assert.Equal(true, result1)

	date2 := time.Date(2023, 06, 02, 0, 0, 0, 0, time.Local)
	result2 := IsWeekend(date2)
	assert.Equal(false, result2)
}

func TestNowDateOrTime(t *testing.T) {
	t.Parallel()

	formats := []string{
		"yyyy-mm-dd hh:mm:ss",
		"yyyy-mm-dd",
		"dd-mm-yy hh:mm:ss",
		"yyyy/mm/dd hh:mm:ss",
		"hh:mm:ss",
		"yyyy/mm",
		"yyyy-mm-dd hh",
	}

	for i := 0; i < len(formats); i++ {
		result := NowDateOrTime(formats[i], "UTC")
		t.Log(result)
	}
}

func TestTimestamp(t *testing.T) {
	t.Parallel()

	ts1 := Timestamp()
	t.Log(ts1)

	ts2 := TimestampMilli()
	t.Log(ts2)

	ts3 := TimestampMicro()
	t.Log(ts3)

	ts4 := TimestampNano()
	t.Log(ts4)
}

func TestTrackFuncTime(t *testing.T) {
	defer TrackFuncTime(time.Now())()

	var n int
	for i := 0; i < 5000000; i++ {
		n++
	}
}

func TestDaysBetween(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDaysBetween")

	tests := []struct {
		start    time.Time
		end      time.Time
		expected int
	}{
		{
			start:    time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2024, time.September, 10, 0, 0, 0, 0, time.UTC),
			expected: 9,
		},
		{
			start:    time.Date(2024, time.September, 10, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC),
			expected: -9,
		},
		{
			start:    time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC),
			expected: 0,
		},
		{
			start:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2024, time.December, 31, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			start:    time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2024, time.March, 31, 0, 0, 0, 0, time.UTC),
			expected: 30,
		},
	}

	for _, tt := range tests {
		result := DaysBetween(tt.start, tt.end)
		assert.Equal(tt.expected, result)
	}
}

func TestGenerateDatetimesBetween(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestGenerateDatetimesBetween")

	tests := []struct {
		start    time.Time
		end      time.Time
		layout   string
		interval string
		expected []string
	}{
		{
			start:    time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2024, time.September, 1, 2, 0, 0, 0, time.UTC),
			layout:   "2006-01-02 15:04:05",
			interval: "30m",
			expected: []string{
				"2024-09-01 00:00:00",
				"2024-09-01 00:30:00",
				"2024-09-01 01:00:00",
				"2024-09-01 01:30:00",
				"2024-09-01 02:00:00",
			},
		},
		{
			start:    time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC),
			layout:   "2006-01-02 15:04:05",
			interval: "1h",
			expected: []string{"2024-09-01 00:00:00"},
		},
		{
			start:    time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2024, time.September, 1, 3, 0, 0, 0, time.UTC),
			layout:   "2006-01-02 15:04:05",
			interval: "2h",
			expected: []string{
				"2024-09-01 00:00:00",
				"2024-09-01 02:00:00",
			},
		},
	}

	for _, tt := range tests {
		result, err := GenerateDatetimesBetween(tt.start, tt.end, tt.layout, tt.interval)

		assert.Equal(tt.expected, result)
		assert.IsNil(err)
	}

	t.Run("Invalid interval", func(t *testing.T) {
		_, err := GenerateDatetimesBetween(time.Now(), time.Now(), "2006-01-02 15:04:05", "invalid")
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}

func TestMin(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMin")

	zeroTime := time.Time{}
	now := time.Now()
	oneMinuteAgo := now.Add(-time.Minute)
	oneMinuteAfter := now.Add(time.Minute)

	assert.Equal(zeroTime, Min(zeroTime, now, oneMinuteAgo, oneMinuteAfter))

	assert.Equal(zeroTime, Min(now, zeroTime))

	assert.Equal(oneMinuteAgo, Min(oneMinuteAgo, now, oneMinuteAfter))
}

func TestMax(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMax")

	zeroTime := time.Time{}
	now := time.Now()
	oneMinuteAgo := now.Add(-time.Minute)
	oneMinuteAfter := now.Add(time.Minute)

	assert.Equal(oneMinuteAfter, Max(zeroTime, now, oneMinuteAgo, oneMinuteAfter))

	assert.Equal(now, Max(now, zeroTime))

	assert.Equal(oneMinuteAfter, Max(oneMinuteAgo, now, oneMinuteAfter))
}

func TestMaxMin(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMinMax")

	zeroTime := time.Time{}
	now := time.Now()
	oneMinuteAgo := now.Add(-time.Minute)
	oneMinuteAfter := now.Add(time.Minute)

	max, min := MaxMin(zeroTime, now, oneMinuteAgo, oneMinuteAfter)
	assert.Equal(zeroTime, min)
	assert.Equal(oneMinuteAfter, max)

	max, min = MaxMin(now, zeroTime)
	assert.Equal(zeroTime, min)
	assert.Equal(now, max)

	max, min = MaxMin(oneMinuteAgo, now, oneMinuteAfter)
	assert.Equal(oneMinuteAgo, min)
	assert.Equal(oneMinuteAfter, max)
}

func TestBetweenSeconds(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBetweenSeconds")

	today := time.Now()
	tomorrow := AddDay(today, 1)
	yesterday := AddDay(today, -1)

	result1 := BetweenSeconds(today, tomorrow)
	result2 := BetweenSeconds(today, yesterday)

	assert.Equal(int64(86400), result1)
	assert.Equal(int64(-86400), result2)
}
