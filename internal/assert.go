package internal

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

const (
	compareNotEqual int = iota - 2
	compareLess
	compareEqual
	compareGreater
)

// Assert is a simple implementation of assertion, only for internal useage
type Assert struct {
	T        *testing.T
	CaseName string
}

// NewAssert return instance of Assert
func NewAssert(t *testing.T, caseName string) *Assert {
	return &Assert{T: t, CaseName: caseName}
}

// Equal check if expected is equal with actual
func (a *Assert) Equal(expected, actual interface{}) {
	if compare(expected, actual) != compareEqual {
		logFailedInfo(a.T, a.CaseName, expected, actual)
	}
}

// NotEqual check if expected is not equal with actual
func (a *Assert) NotEqual(expected, actual interface{}) {
	if compare(expected, actual) == compareEqual {
		expectedInfo := fmt.Sprintf("not %v", expected)
		logFailedInfo(a.T, a.CaseName, expectedInfo, actual)
	}
}

// Greater check if expected is greate than actual
func (a *Assert) Greater(expected, actual interface{}) {
	if compare(expected, actual) != compareGreater {
		expectedInfo := fmt.Sprintf("> %v", expected)
		logFailedInfo(a.T, a.CaseName, expectedInfo, actual)
	}
}

// Less check if expected is less than actual
func (a *Assert) Less(expected, actual interface{}) {
	if compare(expected, actual) != compareLess {
		expectedInfo := fmt.Sprintf("< %v", expected)
		logFailedInfo(a.T, a.CaseName, expectedInfo, actual)
	}
}

// IsNil check if actual is nil
func (a *Assert) IsNil(actual interface{}) {
	if actual != nil {
		logFailedInfo(a.T, a.CaseName, nil, actual)
	}
}

// IsNil check if actual is not nil
func (a *Assert) IsNotNil(actual interface{}) {
	if actual == nil {
		logFailedInfo(a.T, a.CaseName, "not nil", actual)
	}
}

// compare x and y retun :
// x > y -> 1, x < y -> -1, x == y -> 0, x != y -> -2
func compare(x, y interface{}) int {
	vx := reflect.ValueOf(x)
	vy := reflect.ValueOf(y)

	if vx.Type() != vy.Type() {
		return compareNotEqual
	}

	switch vx.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		xInt := vx.Int()
		yInt := vy.Int()
		if xInt > yInt {
			return compareGreater
		}
		if xInt == yInt {
			return compareEqual
		}
		if xInt < yInt {
			return compareLess
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		xUint := vx.Uint()
		yUint := vy.Uint()
		if xUint > yUint {
			return compareGreater
		}
		if xUint == yUint {
			return compareEqual
		}
		if xUint < yUint {
			return compareLess
		}
	case reflect.Float32, reflect.Float64:
		xFloat := vx.Float()
		yFloat := vy.Float()
		if xFloat > yFloat {
			return compareGreater
		}
		if xFloat == yFloat {
			return compareEqual
		}
		if xFloat < yFloat {
			return compareLess
		}
	case reflect.String:
		xString := vx.String()
		yString := vy.String()
		if xString > yString {
			return compareGreater
		}
		if xString == yString {
			return compareEqual
		}
		if xString < yString {
			return compareLess
		}
	default:
		if reflect.DeepEqual(x, y) {
			return compareEqual
		} else {
			return compareNotEqual
		}
	}

	return compareNotEqual

}

// logFailedInfo make test failed and log error info
func logFailedInfo(t *testing.T, caseName string, expected, actual interface{}) {
	_, file, line, _ := runtime.Caller(2)
	errInfo := fmt.Sprintf("Case %v failed. file: %v, line: %v, expected: %v, actual: %v.", caseName, file, line, expected, actual)
	t.Error(errInfo)
	t.FailNow()
}
