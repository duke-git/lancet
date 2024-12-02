package compare

import (
	"bytes"
	"encoding/json"
	"math/big"
	"reflect"
	"time"

	"github.com/duke-git/lancet/v2/convertor"
)

func compareValue(operator string, left, right any) bool {
	leftType, rightType := reflect.TypeOf(left), reflect.TypeOf(right)

	if leftType.Kind() != rightType.Kind() {
		return false
	}

	switch leftType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
		return compareBasicValue(operator, left, right)

	case reflect.Struct, reflect.Slice, reflect.Map:
		return compareRefValue(operator, left, right, leftType.Kind())

	case reflect.Ptr:
		if leftVal, ok := left.(*big.Int); ok {
			if rightVal, ok := right.(*big.Int); ok {
				return compareBigInt(operator, leftVal, rightVal)
			}
		}
	}

	return false
}

func compareRefValue(operator string, leftObj, rightObj any, kind reflect.Kind) bool {
	leftVal, rightVal := reflect.ValueOf(leftObj), reflect.ValueOf(rightObj)

	switch kind {
	case reflect.Struct:

		// compare time
		if leftVal.CanConvert(timeType) {
			timeObj1, ok := leftObj.(time.Time)
			if !ok {
				timeObj1 = leftVal.Convert(timeType).Interface().(time.Time)
			}

			timeObj2, ok := rightObj.(time.Time)
			if !ok {
				timeObj2 = rightVal.Convert(timeType).Interface().(time.Time)
			}

			return compareBasicValue(operator, timeObj1.UnixNano(), timeObj2.UnixNano())
		}

		// for other struct type, only process equal operator
		switch operator {
		case equal:
			return objectsAreEqualValues(leftObj, rightObj)
		}

	case reflect.Slice:
		// compare []byte
		if leftVal.CanConvert(bytesType) {
			bytesObj1, ok := leftObj.([]byte)
			if !ok {
				bytesObj1 = leftVal.Convert(bytesType).Interface().([]byte)
			}
			bytesObj2, ok := rightObj.([]byte)
			if !ok {
				bytesObj2 = rightVal.Convert(bytesType).Interface().([]byte)
			}

			switch operator {
			case equal:
				if bytes.Equal(bytesObj1, bytesObj2) {
					return true
				}
			case lessThan:
				if bytes.Compare(bytesObj1, bytesObj2) == -1 {
					return true
				}
			case greaterThan:
				if bytes.Compare(bytesObj1, bytesObj2) == 1 {
					return true
				}
			case lessOrEqual:
				if bytes.Compare(bytesObj1, bytesObj2) <= 0 {
					return true
				}
			case greaterOrEqual:
				if bytes.Compare(bytesObj1, bytesObj2) >= 0 {
					return true
				}
			}

		}

		// for other type slice, only process equal operator
		switch operator {
		case equal:
			return reflect.DeepEqual(leftObj, rightObj)
		}

	case reflect.Map:
		// only process equal operator
		switch operator {
		case equal:
			return reflect.DeepEqual(leftObj, rightObj)
		}
	}

	return false
}

func objectsAreEqualValues(expected, actual interface{}) bool {
	if objectsAreEqual(expected, actual) {
		return true
	}

	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
		// Attempt comparison after type conversion
		return reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
	}

	return false
}

func objectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

// compareBasic compare basic value: integer, float, string, bool
func compareBasicValue(operator string, leftValue, rightValue any) bool {
	if leftValue == nil && rightValue == nil && operator == equal {
		return true
	}

	switch leftVal := leftValue.(type) {
	case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		left, err := convertor.ToBigInt(leftValue)
		if err != nil {
			return false
		}

		right, err := convertor.ToBigInt(rightValue)
		if err != nil {
			return false
		}

		return compareBigInt(operator, left, right)

	case float32, float64:
		left, err := convertor.ToFloat(leftValue)
		if err != nil {
			return false
		}
		right, err := convertor.ToFloat(rightValue)
		if err != nil {
			return false
		}

		return compareFloats(operator, left, right)

	case string:
		left := leftVal
		switch right := rightValue.(type) {
		case string:
			return compareStrings(operator, left, right)
		}

	case bool:
		left := leftVal
		switch right := rightValue.(type) {
		case bool:
			return compareBools(operator, left, right)
		}

	case json.Number:
		if left, err := leftVal.Float64(); err == nil {
			switch rightVal := rightValue.(type) {
			case json.Number:
				if right, err := rightVal.Float64(); err == nil {
					return compareFloats(operator, left, right)
				}
			case float32, float64:
				right, err := convertor.ToFloat(rightValue)
				if err != nil {
					return false
				}
				return compareFloats(operator, left, right)

			case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
				right, err := convertor.ToBigInt(rightValue)
				if err != nil {
					return false
				}
				left, err := convertor.ToBigInt(left)
				return compareBigInt(operator, left, right)
			}
		}
	}

	return false
}

// compareBigInt compares two big.Int values based on the operator
func compareBigInt(operator string, left, right *big.Int) bool {
	switch operator {
	case equal:
		return left.Cmp(right) == 0
	case lessThan:
		return left.Cmp(right) < 0
	case greaterThan:
		return left.Cmp(right) > 0
	case lessOrEqual:
		return left.Cmp(right) <= 0
	case greaterOrEqual:
		return left.Cmp(right) >= 0
	}
	return false
}

// compareFloats compares two float64 values based on the operator
func compareFloats(operator string, left, right float64) bool {
	switch operator {
	case equal:
		return left == right
	case lessThan:
		return left < right
	case greaterThan:
		return left > right
	case lessOrEqual:
		return left <= right
	case greaterOrEqual:
		return left >= right
	}
	return false
}

// compareStrings compares two string values based on the operator
func compareStrings(operator string, left, right string) bool {
	switch operator {
	case equal:
		return left == right
	case lessThan:
		return left < right
	case greaterThan:
		return left > right
	case lessOrEqual:
		return left <= right
	case greaterOrEqual:
		return left >= right
	}
	return false
}

// compareBools compares two boolean values based on the operator
func compareBools(operator string, left, right bool) bool {
	switch operator {
	case equal:
		return left == right
	}
	return false
}
