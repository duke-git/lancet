package condition

import (
	"errors"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

type TestStruct struct{}

func TestBool(t *testing.T) {
	assert := internal.NewAssert(t, "TestBool")

	// bool
	assert.Equal(false, Bool(false))
	assert.Equal(true, Bool(true))

	// integer
	assert.Equal(false, Bool(0))
	assert.Equal(true, Bool(1))

	// float
	assert.Equal(false, Bool(0.0))
	assert.Equal(true, Bool(0.1))

	// string
	assert.Equal(false, Bool(""))
	assert.Equal(true, Bool(" "))
	assert.Equal(true, Bool("0"))

	// slice
	var nums [2]int
	assert.Equal(false, Bool(nums))
	nums = [2]int{0, 1}
	assert.Equal(true, Bool(nums))

	// map
	assert.Equal(false, Bool(map[string]string{}))
	assert.Equal(true, Bool(map[string]string{"a": "a"}))

	// channel
	var ch chan int
	assert.Equal(false, Bool(ch))
	ch = make(chan int)
	assert.Equal(true, Bool(ch))

	//  interface
	var err error
	assert.Equal(false, Bool(err))
	err = errors.New("error message")
	assert.Equal(true, Bool(err))

	// struct
	assert.Equal(false, Bool(struct{}{}))
	assert.Equal(true, Bool(time.Now()))

	// struct pointer
	ts := TestStruct{}
	assert.Equal(false, Bool(ts))
	assert.Equal(true, Bool(&ts))
}

func TestTernaryOperator(t *testing.T) {
	assert := internal.NewAssert(t, "TernaryOperator")
	trueValue := "1"
	falseValue := "0"

	assert.Equal(trueValue, TernaryOperator(true, trueValue, falseValue))
}
