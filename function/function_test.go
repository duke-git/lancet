package function

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestAfter(t *testing.T) {
	arr := []string{"a", "b"}
	f := After(len(arr), func(i int) int {
		fmt.Println("print done")
		return i
	})
	type cb func(args ...interface{}) []reflect.Value
	print := func(i int, s string, fn cb) {
		fmt.Printf("print: arr[%d] is %s \n", i, s)
		v := fn(i)
		if v != nil {
			vv := v[0].Int()
			if vv != 1 {
				t.FailNow()
			}
		}
	}
	fmt.Println("print: arr is", arr)
	for i := 0; i < len(arr); i++ {
		print(i, arr[i], f)
	}
}

func TestBefore(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e"}
	f := Before(3, func(i int) int {
		return i
	})

	var res []int64
	type cb func(args ...interface{}) []reflect.Value
	appendStr := func(i int, s string, fn cb) {
		fmt.Printf("appendStr: arr[%d] is %s \n", i, s)
		v := fn(i)
		res = append(res, v[0].Int())
	}

	for i := 0; i < len(arr); i++ {
		appendStr(i, arr[i], f)
	}

	expect := []int64{0, 1, 2, 2, 2}
	if !reflect.DeepEqual(expect, res) {
		t.FailNow()
	}
}

func TestCurry(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}
	var addCurry Fn = func(values ...interface{}) interface{} {
		return add(values[0].(int), values[1].(int))
	}

	add1 := addCurry.Curry(1)
	v := add1(2)
	if v != 3 {
		t.FailNow()
	}
}

func TestCompose(t *testing.T) {
	toUpper := func(a... string) string {
		return strings.ToUpper(a[0])
	}

	toLower := func(a... string) string {
		return strings.ToLower(a[0])
	}

	expect := toUpper(toLower("aBCde"))
	cf := Compose(toUpper, toLower)
	res := cf("aBCde")

	fmt.Println(res, expect)
	if res != expect {
		t.FailNow()
	}

}

func TestDelay(t *testing.T) {
	var print = func(s string) {
		fmt.Println(s)
	}
	Delay(2*time.Second, print, "test delay")
}

func TestSchedule(t *testing.T) {
	var res []string
	appendStr := func(s string) {
		fmt.Println(s)
		res = append(res, s)
	}

	stop := Schedule(1*time.Second, appendStr, "*")
	time.Sleep(5 * time.Second)
	close(stop)

	expect := []string{"*", "*", "*", "*", "*"}
	if !reflect.DeepEqual(expect, res) {
		t.FailNow()
	}
	fmt.Println("done")
}
