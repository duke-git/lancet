package function

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/duke-git/lancet/internal"
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
	assert := internal.NewAssert(t, "TestBefore")

	arr := []string{"a", "b", "c", "d", "e"}
	f := Before(3, func(i int) int {
		return i
	})

	var res []int64
	type cb func(args ...interface{}) []reflect.Value
	appendStr := func(i int, _ string, fn cb) {
		v := fn(i)
		res = append(res, v[0].Int())
	}

	for i := 0; i < len(arr); i++ {
		appendStr(i, arr[i], f)
	}

	expected := []int64{0, 1, 2, 2, 2}
	assert.Equal(expected, res)
}

func TestCurry(t *testing.T) {
	assert := internal.NewAssert(t, "TestCurry")

	add := func(a, b int) int {
		return a + b
	}
	var addCurry Fn = func(values ...interface{}) interface{} {
		return add(values[0].(int), values[1].(int))
	}
	add1 := addCurry.Curry(1)
	assert.Equal(3, add1(2))
}

func TestCompose(t *testing.T) {
	assert := internal.NewAssert(t, "TestCompose")

	toUpper := func(a ...interface{}) interface{} {
		return strings.ToUpper(a[0].(string))
	}
	toLower := func(a ...interface{}) interface{} {
		return strings.ToLower(a[0].(string))
	}

	expected := toUpper(toLower("aBCde"))
	cf := Compose(toUpper, toLower)
	res := cf("aBCde")

	assert.Equal(expected, res)
}

func TestDelay(t *testing.T) {
	var print = func(s string) {
		t.Log(s)
	}
	Delay(2*time.Second, print, "test delay")
}

func TestDebounced(t *testing.T) {
	assert := internal.NewAssert(t, "TestDebounced")

	count := 0
	add := func() {
		count++
	}

	debouncedAdd := Debounced(add, 50*time.Microsecond)
	debouncedAdd()
	debouncedAdd()
	debouncedAdd()
	debouncedAdd()

	time.Sleep(100 * time.Millisecond)
	assert.Equal(1, count)

	debouncedAdd()
	time.Sleep(100 * time.Millisecond)
	assert.Equal(2, count)
}

func TestSchedule(t *testing.T) {
	assert := internal.NewAssert(t, "TestSchedule")

	var res []string
	appendStr := func(s string) {
		res = append(res, s)
	}

	stop := Schedule(1*time.Second, appendStr, "*")
	time.Sleep(5 * time.Second)
	close(stop)

	expected := []string{"*", "*", "*", "*", "*"}
	assert.Equal(expected, res)
}

func TestPipeline(t *testing.T) {
	assert := internal.NewAssert(t, "TestPipeline")

	addOne := func(x interface{}) interface{} {
		return x.(int) + 1
	}
	double := func(x interface{}) interface{} {
		return 2 * x.(int)
	}
	square := func(x interface{}) interface{} {
		return x.(int) * x.(int)
	}

	f := Pipeline(addOne, double, square)

	assert.Equal(36, f(2))
}

func TestDebounce(t *testing.T) {
	assert := internal.NewAssert(t, "TestDebounce")

	t.Run("Single call", func(t *testing.T) {
		callCount := 0

		debouncedFn, _ := Debounce(func() {
			callCount++
		}, 500*time.Millisecond)

		debouncedFn()

		time.Sleep(1 * time.Second)

		assert.Equal(1, callCount)
	})

	t.Run("Multiple calls within debounce interval", func(t *testing.T) {
		callCount := 0

		debouncedFn, _ := Debounce(func() {
			callCount++
		}, 1*time.Second)

		for i := 0; i < 5; i++ {
			go func(index int) {
				time.Sleep(time.Duration(index) * 200 * time.Millisecond)
				debouncedFn()
			}(i)
		}

		time.Sleep(2 * time.Second)

		assert.Equal(1, callCount)
	})

	t.Run("Immediate consecutive calls", func(t *testing.T) {
		callCount := 0

		debouncedFn, _ := Debounce(func() {
			callCount++
		}, 500*time.Millisecond)

		for i := 0; i < 10; i++ {
			debouncedFn()
			time.Sleep(50 * time.Millisecond)
		}

		time.Sleep(1 * time.Second)

		assert.Equal(1, callCount)
	})

	t.Run("Cancel calls", func(t *testing.T) {
		callCount := 0

		debouncedFn, cancelFn := Debounce(func() {
			callCount++
		}, 500*time.Millisecond)

		debouncedFn()

		cancelFn()

		time.Sleep(1 * time.Second)

		assert.Equal(0, callCount)
	})

}

func TestThrottle(t *testing.T) {
	assert := internal.NewAssert(t, "TestThrottle")

	t.Run("Single call", func(t *testing.T) {
		callCount := 0

		throttledFn := Throttle(func() {
			callCount++
		}, 1*time.Second)

		throttledFn()

		time.Sleep(100 * time.Millisecond)

		assert.Equal(1, callCount)
	})

	t.Run("Multiple calls within throttle interval", func(t *testing.T) {
		callCount := 0

		throttledFn := Throttle(func() {
			callCount++
		}, 1*time.Second)

		for i := 0; i < 5; i++ {
			throttledFn()
		}

		time.Sleep(1 * time.Second)

		assert.Equal(1, callCount)
	})

	t.Run("Mutiple calls space out throttle interval", func(t *testing.T) {
		callCount := 0

		throttledFn := Throttle(func() {
			callCount++
		}, 500*time.Millisecond)

		throttledFn()
		time.Sleep(600 * time.Millisecond)

		throttledFn()
		time.Sleep(600 * time.Millisecond)

		throttledFn()

		time.Sleep(1 * time.Second)

		assert.Equal(3, callCount)
	})

	t.Run("Call function near the end of the interval", func(t *testing.T) {
		callCount := 0

		throttledFn := Throttle(func() {
			callCount++
		}, 1*time.Second)

		throttledFn()
		time.Sleep(900 * time.Millisecond)

		throttledFn()
		time.Sleep(200 * time.Millisecond)

		assert.Equal(2, callCount)
	})

}
