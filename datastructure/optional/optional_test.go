package optional

import (
	"errors"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestDefault(t *testing.T) {
	assert := internal.NewAssert(t, "TestEmpty")
	opt := Default[int]()

	assert.ShouldBeTrue(opt.IsNil())
}

func TestOf(t *testing.T) {
	assert := internal.NewAssert(t, "TestOf")
	value := 42
	opt := Of(value)

	assert.ShouldBeTrue(opt.IsNotNil())
	assert.Equal(opt.Unwarp(), value)
}

func TestFromNillable(t *testing.T) {
	assert := internal.NewAssert(t, "TestOfNullable")
	var value *int = nil
	opt := FromNillable(value)

	assert.ShouldBeFalse(opt.IsNotNil())

	value = new(int)
	*value = 42
	opt = FromNillable(value)

	assert.ShouldBeTrue(opt.IsNotNil())
}

func TestOrElse(t *testing.T) {
	assert := internal.NewAssert(t, "TestOrElse")
	optDefault := Default[int]()
	defaultValue := 100

	val := optDefault.OrElse(defaultValue)
	assert.Equal(val, defaultValue)

	optWithValue := Of(42)
	val = optWithValue.OrElse(defaultValue)
	assert.Equal(val, 42)
}

func TestOrElseGetHappyPath(t *testing.T) {
	assert := internal.NewAssert(t, "TestOrElseGetHappyPath")
	optWithValue := Of(42)
	action := func() int { return 100 }

	val := optWithValue.OrElseGet(action)
	assert.Equal(val, 42)
}

func TestOrElseGet(t *testing.T) {
	assert := internal.NewAssert(t, "TestOrElseGet")
	optDefault := Default[int]()
	action := func() int { return 100 }

	val := optDefault.OrElseGet(action)
	assert.Equal(val, action())
}

func TestOrElseTrigger(t *testing.T) {
	assert := internal.NewAssert(t, "OrElseTrigger")
	optDefault := Default[int]()
	_, err := optDefault.OrElseTrigger(func() error { return errors.New("no value") })

	assert.Equal(err.Error(), "no value")

	optWithValue := Of(42)
	val, err := optWithValue.OrElseTrigger(func() error { return errors.New("no value") })

	assert.IsNil(err)
	assert.Equal(val, 42)
}

func TestIfNotNil(t *testing.T) {
	assert := internal.NewAssert(t, "IfNotNil")
	called := false
	action := func(value int) { called = true }

	optDefault := Default[int]()
	optDefault.IfNotNil(action)

	assert.ShouldBeFalse(called)

	called = false // Reset for next test
	optWithValue := Of(42)
	optWithValue.IfNotNil(action)

	assert.ShouldBeTrue(called)
}

func TestIfNotNilOrElse(t *testing.T) {
	assert := internal.NewAssert(t, "TestIfNotNilOrElse")

	// Test when value is present
	calledWithValue := false
	valueAction := func(value int) { calledWithValue = true }
	fallbackAction := func() { t.Errorf("Empty action should not be called when value is present") }

	optWithValue := Of(42)
	optWithValue.IfNotNilOrElse(valueAction, fallbackAction)

	assert.ShouldBeTrue(calledWithValue)

	// Test when value is not present
	calledWithEmpty := false
	valueAction = func(value int) { t.Errorf("Value action should not be called when value is not present") }
	fallbackAction = func() { calledWithEmpty = true }

	optDefault := Default[int]()
	optDefault.IfNotNilOrElse(valueAction, fallbackAction)

	assert.ShouldBeTrue(calledWithEmpty)
}

func TestGetWithPanicStandard(t *testing.T) {
	assert := internal.NewAssert(t, "TestGetWithPanicStandard")

	// Test when value is present
	optWithValue := Of(42)
	func() {
		defer func() {
			r := recover()
			assert.IsNil(r)
		}()
		val := optWithValue.Unwarp()
		if val != 42 {
			t.Errorf("Expected Unwarp to return 42, got %v", val)
		}
	}()

	// Test when value is not present
	optDefault := Default[int]()
	func() {
		defer func() {
			r := recover()
			assert.IsNotNil(r)
		}()
		_ = optDefault.Unwarp()
	}()
}
