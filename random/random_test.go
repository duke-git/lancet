/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: dudaodong@kingsoft.com
 * @Date: 2021-11-29 11:43:28
 * @LastEditors: dudaodong@kingsoft.com
 * @LastEditTime: 2021-12-01 18:05:29
 */
package random

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestRandString(t *testing.T) {
	randStr := RandString(6)
	fmt.Println(randStr)
	pattern := `^[a-zA-Z]+$`
	reg := regexp.MustCompile(pattern)

	if len(randStr) != 6 || !reg.MatchString(randStr) {
		internal.LogFailedTestInfo(t, "RandString", "RandString(6)", "RandString(6) should be 6 letters ", randStr)
		t.FailNow()
	}
}

func TestRandInt(t *testing.T) {
	res1 := RandInt(1, 10)
	if res1 < 1 || res1 >= 10 {
		internal.LogFailedTestInfo(t, "RandInt", "RandInt(1, 10)", "RandInt(1, 10) should between [1, 10) ", res1)
		t.FailNow()
	}

	res2 := RandInt(1, 1)
	if res2 != 1 {
		internal.LogFailedTestInfo(t, "RandInt", "RandInt(1, 1)", "RandInt(1, 1) should be 1 ", res2)
		t.FailNow()
	}

	res3 := RandInt(10, 1)
	if res3 < 1 || res3 >= 10 {
		internal.LogFailedTestInfo(t, "RandInt", "RandInt(10, 1)", "RandInt(10, 1) should between [1, 10) ", res3)
		t.FailNow()
	}
}

func TestRandBytes(t *testing.T) {
	randBytes := RandBytes(4)
	if len(randBytes) != 4 {
		internal.LogFailedTestInfo(t, "RandBytes", "RandBytes(4)", "RandBytes(4) should return 4 element of []bytes", randBytes)
		t.FailNow()
	}
	v := reflect.ValueOf(randBytes)
	et := v.Type().Elem()
	if v.Kind() != reflect.Slice || et.Kind() != reflect.Uint8 {
		internal.LogFailedTestInfo(t, "RandBytes", "RandBytes(4)", "RandBytes(4) should return 4 element of []bytes", randBytes)
		t.FailNow()
	}

	randErr := RandBytes(0)
	if randErr != nil {
		internal.LogFailedTestInfo(t, "RandBytes", "RandBytes(0)", "RandBytes(0) should return nil", randErr)
		t.FailNow()
	}

}
