package datetime

import (
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestAddYear(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAddDay")

	now := time.Now()
	after2Years := AddYear(now, 1)
	diff1 := after2Years.Sub(now)
	assert.Equal(float64(8760), diff1.Hours())

	before2Years := AddYear(now, -1)
	diff2 := before2Years.Sub(now)
	assert.Equal(float64(-8760), diff2.Hours())
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

func TestAddDay(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAddDay")

	now := time.Now()
	after2Days := AddDay(now, 2)
	diff1 := after2Days.Sub(now)
	assert.Equal(float64(48), diff1.Hours())

	before2Days := AddDay(now, -2)
	diff2 := before2Days.Sub(now)
	assert.Equal(float64(-48), diff2.Hours())
}

func TestAddHour(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAddHour")

	now := time.Now()
	after2Hours := AddHour(now, 2)
	diff1 := after2Hours.Sub(now)
	assert.Equal(float64(2), diff1.Hours())

	before2Hours := AddHour(now, -2)
	diff2 := before2Hours.Sub(now)
	assert.Equal(float64(-2), diff2.Hours())
}

func TestAddMinute(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAddMinute")

	now := time.Now()
	after2Minutes := AddMinute(now, 2)
	diff1 := after2Minutes.Sub(now)
	assert.Equal(float64(2), diff1.Minutes())

	before2Minutes := AddMinute(now, -2)
	diff2 := before2Minutes.Sub(now)
	assert.Equal(float64(-2), diff2.Minutes())
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
