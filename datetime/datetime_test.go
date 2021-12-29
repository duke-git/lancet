package datetime

import (
	"github.com/duke-git/lancet/utils"
	"testing"
	"time"
)

func TestAddDay(t *testing.T) {
	now := time.Now()

	after2Days := AddDay(now, 2)
	diff1 := after2Days.Sub(now)
	if diff1.Hours() != 48 {
		utils.LogFailedTestInfo(t, "AddDay", now, 48, diff1.Hours())
		t.FailNow()
	}

	before2Days := AddDay(now, -2)
	diff2 := before2Days.Sub(now)
	if diff2.Hours() != -48 {
		utils.LogFailedTestInfo(t, "AddDay", now, -48, diff2.Hours())
		t.FailNow()
	}

}
func TestAddHour(t *testing.T) {
	now := time.Now()

	after2Hours := AddHour(now, 2)
	diff1 := after2Hours.Sub(now)
	if diff1.Hours() != 2 {
		utils.LogFailedTestInfo(t, "AddHour", now, 2, diff1.Hours())
		t.FailNow()
	}

	before2Hours := AddHour(now, -2)
	diff2 := before2Hours.Sub(now)
	if diff2.Hours() != -2 {
		utils.LogFailedTestInfo(t, "AddHour", now, -2, diff2.Hours())
		t.FailNow()
	}
}

func TestAddMinute(t *testing.T) {
	now := time.Now()

	after2Minutes := AddMinute(now, 2)
	diff1 := after2Minutes.Sub(now)
	if diff1.Minutes() != 2 {
		utils.LogFailedTestInfo(t, "AddMinute", now, 2, diff1.Minutes())
		t.FailNow()
	}

	before2Minutes := AddMinute(now, -2)
	diff2 := before2Minutes.Sub(now)
	if diff2.Minutes() != -2 {
		utils.LogFailedTestInfo(t, "AddMinute", now, -2, diff2.Minutes())
		t.FailNow()
	}
}

func TestGetNowDate(t *testing.T) {
	date := GetNowDate()
	expected := time.Now().Format("2006-01-02")
	if date != expected {
		utils.LogFailedTestInfo(t, "GetNowDate", "", expected, date)
		t.FailNow()
	}
}

func TestGetNotTime(t *testing.T) {
	ts := GetNowTime()
	expected := time.Now().Format("15:04:05")
	if ts != expected {
		utils.LogFailedTestInfo(t, "GetNowTime", "", expected, ts)
		t.FailNow()
	}
}

func TestGetNowDateTime(t *testing.T) {
	ts := GetNowDateTime()
	expected := time.Now().Format("2006-01-02 15:04:05")
	if ts != expected {
		utils.LogFailedTestInfo(t, "GetNowDateTime", "", expected, ts)
		t.FailNow()
	}
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
		res := FormatTimeToStr(datetime, cases[i])
		if res != expected[i] {
			utils.LogFailedTestInfo(t, "FormatTimeToStr", cases[i], expected[i], res)
			t.FailNow()
		}
	}

}

func TestFormatStrToTime(t *testing.T) {
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
		res, err := FormatStrToTime(datetimeStr[i], cases[i])
		if err != nil {
			t.Fatal(err)
		}
		expected, _ := time.Parse(formats[i], datetimeStr[i])
		if res != expected {
			utils.LogFailedTestInfo(t, "FormatTimeToStr", cases[i], expected, res)
			t.FailNow()
		}
	}
}
