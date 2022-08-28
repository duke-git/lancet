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

func TestAnd(t *testing.T) {
	assert := internal.NewAssert(t, "TestAnd")
	assert.Equal(false, And(0, 1))
	assert.Equal(false, And(0, ""))
	assert.Equal(false, And(0, "0"))
	assert.Equal(true, And(1, "0"))
}

func TestOr(t *testing.T) {
	assert := internal.NewAssert(t, "TestOr")
	assert.Equal(false, Or(0, ""))
	assert.Equal(true, Or(0, 1))
	assert.Equal(true, Or(0, "0"))
	assert.Equal(true, Or(1, "0"))
}

func TestXor(t *testing.T) {
	assert := internal.NewAssert(t, "TestOr")
	assert.Equal(false, Xor(0, 0))
	assert.Equal(true, Xor(0, 1))
	assert.Equal(true, Xor(1, 0))
	assert.Equal(false, Xor(1, 1))
}

func TestNor(t *testing.T) {
	assert := internal.NewAssert(t, "TestNor")
	assert.Equal(true, Nor(0, 0))
	assert.Equal(false, Nor(0, 1))
	assert.Equal(false, Nor(1, 0))
	assert.Equal(false, Nor(1, 1))
}

func TestXnor(t *testing.T) {
	assert := internal.NewAssert(t, "TestXnor")
	assert.Equal(true, Xnor(0, 0))
	assert.Equal(false, Xnor(0, 1))
	assert.Equal(false, Xnor(1, 0))
	assert.Equal(true, Xnor(1, 1))
}

func TestNand(t *testing.T) {
	assert := internal.NewAssert(t, "TestNand")
	assert.Equal(true, Nand(0, 0))
	assert.Equal(true, Nand(0, 1))
	assert.Equal(true, Nand(1, 0))
	assert.Equal(false, Nand(1, 1))
}

func TestTernaryOperator(t *testing.T) {
	assert := internal.NewAssert(t, "TernaryOperator")
	trueValue := "1"
	falseValue := "0"

	assert.Equal(trueValue, TernaryOperator(true, trueValue, falseValue))
}
