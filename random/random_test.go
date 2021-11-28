package random

import (
	"fmt"
	"github.com/duke-git/lancet/utils"
	"reflect"
	"regexp"
	"testing"
)

func TestRandString(t *testing.T) {
	randStr := RandString(6)
	fmt.Println(randStr)
	pattern := `^[a-zA-Z]+$`
	reg := regexp.MustCompile(pattern)

	if len(randStr) != 6 || !reg.MatchString(randStr) {
		utils.LogFailedTestInfo(t, "RandString", "RandString(6)", "RandString(6) should be 6 letters ", randStr)
		t.FailNow()
	}
}

func TestRandInt(t *testing.T) {
	randInt := RandInt(1, 10)

	if randInt < 1 || randInt >= 10 {
		utils.LogFailedTestInfo(t, "RandInt", "RandInt(1, 10)", "RandInt(1, 10) should between [1, 10) ", randInt)
		t.FailNow()
	}
}

func TestRandBytes(t *testing.T) {
	randBytes := RandBytes(4)

	if len(randBytes) != 4 {
		utils.LogFailedTestInfo(t, "RandBytes", "RandBytes(4)", "RandBytes(4) should return 4 element of []bytes", randBytes)
		t.FailNow()
	}

	v := reflect.ValueOf(randBytes)
	et := v.Type().Elem()
	if v.Kind() != reflect.Slice || et.Kind() != reflect.Uint8 {
		utils.LogFailedTestInfo(t, "RandBytes", "RandBytes(4)", "RandBytes(4) should return 4 element of []bytes", randBytes)
		t.FailNow()
	}
}
