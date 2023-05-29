package datetime

import (
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestAddYear(t *testing.T) {
	assert := internal.NewAssert(t, "TestAddDay")

	now := time.Now()
	after2Years := AddYear(now, 1)
	diff1 := after2Years.Sub(now)
	assert.Equal(float64(8760), diff1.Hours())

	before2Years := AddYear(now, -1)
	diff2 := before2Years.Sub(now)
	assert.Equal(float64(-8760), diff2.Hours())
}

func TestAddDay(t *testing.T) {
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
	assert := internal.NewAssert(t, "TestGetNowDate")
	expected := time.Now().Format("2006-01-02")
	assert.Equal(expected, GetNowDate())
}

func TestGetNowTime(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetNowTime")
	expected := time.Now().Format("15:04:05")
	assert.Equal(expected, GetNowTime())
}

func TestGetNowDateTime(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetNowDateTime")
	expected := time.Now().Format("2006-01-02 15:04:05")
	assert.Equal(expected, GetNowDateTime())
}

func TestFormatTimeToStr(t *testing.T) {
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
}

func TestFormatStrToTime(t *testing.T) {
	assert := internal.NewAssert(t, "TestFormatStrToTime")

	formats := []string{
		"2006-01-02 15:04:05", "2006-01-02",
		"02-01-06 15:04:05", "2006/01/02 15:04:05",
		"2006/01"}
	cases := []string{
		"yyyy-mm-dd hh:mm:ss", "yyyy-mm-dd",
		"dd-mm-yy hh:mm:ss", "yyyy/mm/dd hh:mm:ss",
		"yyyy/mm"}

	datetimeStr := []string{
		"2021-01-02 16:04:08", "2021-01-02",
		"02-01-21 16:04:08", "2021/01/02 16:04:08",
		"2021/01"}

	for i := 0; i < len(cases); i++ {
		actual, err := FormatStrToTime(datetimeStr[i], cases[i])
		if err != nil {
			t.Fatal(err)
		}
		expected, _ := time.Parse(formats[i], datetimeStr[i])
		assert.Equal(expected, actual)
	}
}

func TestBeginOfMinute(t *testing.T) {
	assert := internal.NewAssert(t, "TestBeginOfMinute")

	expected := time.Date(2022, 2, 15, 15, 48, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfMinute(td)

	assert.Equal(expected, actual)
}

func TestEndOfMinute(t *testing.T) {
	assert := internal.NewAssert(t, "TestEndOfMinute")

	expected := time.Date(2022, 2, 15, 15, 48, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfMinute(td)

	assert.Equal(expected, actual)
}

func TestBeginOfHour(t *testing.T) {
	assert := internal.NewAssert(t, "TestBeginOfHour")

	expected := time.Date(2022, 2, 15, 15, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfHour(td)

	assert.Equal(expected, actual)
}

func TestEndOfHour(t *testing.T) {
	assert := internal.NewAssert(t, "TestEndOfHour")

	expected := time.Date(2022, 2, 15, 15, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfHour(td)

	assert.Equal(expected, actual)
}

func TestBeginOfDay(t *testing.T) {
	assert := internal.NewAssert(t, "TestBeginOfDay")

	expected := time.Date(2022, 2, 15, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfDay(td)

	assert.Equal(expected, actual)
}

func TestEndOfDay(t *testing.T) {
	assert := internal.NewAssert(t, "TestEndOfDay")

	expected := time.Date(2022, 2, 15, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfDay(td)

	assert.Equal(expected, actual)
}

func TestBeginOfWeek(t *testing.T) {
	assert := internal.NewAssert(t, "TestBeginOfWeek")

	expected := time.Date(2022, 2, 13, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfWeek(td)

	assert.Equal(expected, actual)
}

func TestEndOfWeek(t *testing.T) {
	assert := internal.NewAssert(t, "TestEndOfWeek")

	expected := time.Date(2022, 2, 19, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfWeek(td)

	assert.Equal(expected, actual)
}

func TestBeginOfMonth(t *testing.T) {
	assert := internal.NewAssert(t, "TestBeginOfMonth")

	expected := time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfMonth(td)

	assert.Equal(expected, actual)
}

func TestEndOfMonth(t *testing.T) {
	assert := internal.NewAssert(t, "TestEndOfMonth")

	expected := time.Date(2022, 2, 28, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfMonth(td)

	assert.Equal(expected, actual)
}

func TestBeginOfYear(t *testing.T) {
	assert := internal.NewAssert(t, "TestBeginOfYear")

	expected := time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := BeginOfYear(td)

	assert.Equal(expected, actual)
}

func TestEndOfYear(t *testing.T) {
	assert := internal.NewAssert(t, "TestEndOfYear")

	expected := time.Date(2022, 12, 31, 23, 59, 59, 999999999, time.Local)
	td := time.Date(2022, 2, 15, 15, 48, 40, 112, time.Local)
	actual := EndOfYear(td)

	assert.Equal(expected, actual)
}

func TestIsLeapYear(t *testing.T) {
	assert := internal.NewAssert(t, "TestEndOfYear")

	result1 := IsLeapYear(2000)
	result2 := IsLeapYear(2001)

	assert.Equal(true, result1)
	assert.Equal(false, result2)
}
