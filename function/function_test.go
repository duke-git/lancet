package function

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestAfter(t *testing.T) {
	t.Parallel()

	arr := []string{"a", "b"}
	f := After(len(arr), func(i int) int {
		fmt.Println("print done")
		return i
	})
	type cb func(args ...any) []reflect.Value
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
	t.Parallel()

	assert := internal.NewAssert(t, "TestBefore")

	arr := []string{"a", "b", "c", "d", "e"}
	f := Before(3, func(i int) int {
		return i
	})

	var res []int64
	type cb func(args ...any) []reflect.Value
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
	t.Parallel()

	assert := internal.NewAssert(t, "TestCurry")

	add := func(a, b int) int {
		return a + b
	}
	var addCurry CurryFn[int] = func(values ...int) int {
		return add(values[0], values[1])
	}
	add1 := addCurry.New(1)

	assert.Equal(3, add1(2))
}

func TestCompose(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCompose")

	toUpper := func(strs ...string) string {
		return strings.ToUpper(strs[0])
	}
	toLower := func(strs ...string) string {
		return strings.ToLower(strs[0])
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

func TestSchedule(t *testing.T) {
	assert := internal.NewAssert(t, "TestSchedule")

	t.Run("Single call", func(t *testing.T) {
		res := []string{}
		appendStr := func(s string) {
			res = append(res, s)
		}
		stop := Schedule(200*time.Millisecond, appendStr, "*")
		close(stop)

		time.Sleep(400 * time.Millisecond)

		assert.Equal([]string{"*"}, res)
	})

	t.Run("Multiple calls", func(t *testing.T) {
		res := []string{}
		appendStr := func(s string) {
			res = append(res, s)
		}
		stop := Schedule(200*time.Millisecond, appendStr, "*")

		time.Sleep(800 * time.Millisecond)

		close(stop)

		assert.Equal([]string{"*", "*", "*", "*"}, res)
	})

}

func TestPipeline(t *testing.T) {
	assert := internal.NewAssert(t, "TestPipeline")

	addOne := func(x int) int {
		return x + 1
	}
	double := func(x int) int {
		return 2 * x
	}
	square := func(x int) int {
		return x * x
	}

	f := Pipeline(addOne, double, square)

	assert.Equal(36, f(2))
}

func TestAcceptIf(t *testing.T) {
	assert := internal.NewAssert(t, "AcceptIf")

	adder := AcceptIf(
		And(
			func(x int) bool {
				return x > 10
			}, func(x int) bool {
				return x%2 == 0
			}),
		func(x int) int {
			return x + 1
		},
	)

	result, ok := adder(20)
	assert.Equal(21, result)
	assert.Equal(true, ok)

	result, ok = adder(21)
	assert.Equal(0, result)
	assert.Equal(false, ok)
}

func TestAcceptIfPanicMissingPredicate(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAcceptIfPanicMissingPredicate")

	defer func() {
		v := recover()
		assert.Equal("programming error: predicate must be not nil", v)
	}()

	AcceptIf(
		nil,
		func(x int) int {
			return x
		},
	)
}

func TestAcceptIfPanicMissingApply(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAcceptIfPanicMissingApply")

	defer func() {
		v := recover()
		assert.Equal("programming error: apply must be not nil", v)
	}()

	AcceptIf(
		func(i int) bool {
			return false
		},
		nil,
	)
}
