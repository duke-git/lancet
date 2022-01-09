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
	"reflect"
	"regexp"
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestRandString(t *testing.T) {
	pattern := `^[a-zA-Z]+$`
	reg := regexp.MustCompile(pattern)

	randStr := RandString(6)

	assert := internal.NewAssert(t, "TestRandString")
	assert.Equal(6, len(randStr))
	assert.Equal(true, reg.MatchString(randStr))
}

func TestRandInt(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandInt")

	r1 := RandInt(1, 10)
	assert.GreaterOrEqual(r1, 1)
	assert.Less(r1, 10)

	r2 := RandInt(1, 1)
	assert.Equal(1, r2)

	r3 := RandInt(10, 1)
	assert.GreaterOrEqual(r1, 1)
	assert.Less(r3, 10)
}

func TestRandBytes(t *testing.T) {
	assert := internal.NewAssert(t, "TestRandBytes")

	randBytes := RandBytes(4)
	assert.Equal(4, len(randBytes))

	v := reflect.ValueOf(randBytes)
	elemType := v.Type().Elem()
	assert.Equal(reflect.Slice, v.Kind())
	assert.Equal(reflect.Uint8, elemType.Kind())

	assert.Equal([]byte{}, RandBytes(0))
}
