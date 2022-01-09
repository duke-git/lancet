package datetime

import (
	"testing"
	"time"

	"github.com/duke-git/lancet/internal"
)

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

func TestGetNotTime(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetNotTime")
	expected := time.Now().Format("15:04:05")
	assert.Equal(expected, GetNowTime())
}

func TestGetNowDateTime(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetNowDateTime")
	expected := time.Now().Format("2006-01-02 15:04:05")
	assert.Equal(expected, GetNowDateTime())
}

//todo
//func TestGetZeroHourTimestamp(t *testing.T) {
//	ts := GetZeroHourTimestamp()
//	expected := time.Now().UTC().Unix() - 8*3600
//	if ts != expected {
//		utils.LogFailedTestInfo(t, "GetZeroHourTimestamp", "", expected, ts)
//		t.FailNow()
//	}
//}

func TestFormatTimeToStr(t *testing.T) {
	assert := internal.NewAssert(t, "TestFormatTimeToStr")

	datetime, _ := time.Parse("2006-01-02 15:04:05", "2021-01-02 16:04:08")
	cases := []string{
		"yyyy-mm-dd hh:mm:ss", "yyyy-mm-dd",
		"dd-mm-yy hh:mm:ss", "yyyy/mm/dd hh:mm:ss",
		"hh:mm:ss", "yyyy/mm"}

	expected := []string{
		"2021-01-02 16:04:08", "2021-01-02",
		"02-01-21 16:04:08", "2021/01/02 16:04:08",
		"16:04:08", "2021/01"}

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
