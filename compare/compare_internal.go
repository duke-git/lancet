package compare

import (
	"bytes"
	"encoding/json"
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
	case json.Number:
		if left, err := leftVal.Float64(); err == nil {
			switch rightVal := rightValue.(type) {
			case json.Number:
				if right, err := rightVal.Float64(); err == nil {
					switch operator {
					case equal:
						if left == right {
							return true
						}
					case lessThan:
						if left < right {
							return true
						}
					case greaterThan:
						if left > right {
							return true
						}
					case lessOrEqual:
						if left <= right {
							return true
						}
					case greaterOrEqual:
						if left >= right {
							return true
						}
					}

				}

			case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
				right, err := convertor.ToFloat(rightValue)
				if err != nil {
					return false
				}
				switch operator {
				case equal:
					if left == right {
						return true
					}
				case lessThan:
					if left < right {
						return true
					}
				case greaterThan:
					if left > right {
						return true
					}
				case lessOrEqual:
					if left <= right {
						return true
					}
				case greaterOrEqual:
					if left >= right {
						return true
					}
				}
			}

		}

	case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		left, err := convertor.ToFloat(leftValue)
		if err != nil {
			return false
		}
		switch rightVal := rightValue.(type) {
		case json.Number:
			if right, err := rightVal.Float64(); err == nil {
				switch operator {
				case equal:
					if left == right {
						return true
					}
				case lessThan:
					if left < right {
						return true
					}
				case greaterThan:
					if left > right {
						return true
					}
				case lessOrEqual:
					if left <= right {
						return true
					}
				case greaterOrEqual:
					if left >= right {
						return true
					}
				}
			}
		case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
			right, err := convertor.ToFloat(rightValue)
			if err != nil {
				return false
			}

			switch operator {
			case equal:
				if left == right {
					return true
				}
			case lessThan:
				if left < right {
					return true
				}
			case greaterThan:
				if left > right {
					return true
				}
			case lessOrEqual:
				if left <= right {
					return true
				}
			case greaterOrEqual:
				if left >= right {
					return true
				}
			}
		}

	case string:
		left := leftVal
		switch right := rightValue.(type) {
		case string:
			switch operator {
			case equal:
				if left == right {
					return true
				}
			case lessThan:
				if left < right {
					return true
				}
			case greaterThan:
				if left > right {
					return true
				}
			case lessOrEqual:
				if left <= right {
					return true
				}
			case greaterOrEqual:
				if left >= right {
					return true
				}
			}
		}

	case bool:
		left := leftVal
		switch right := rightValue.(type) {
		case bool:
			switch operator {
			case equal:
				if left == right {
					return true
				}
			}
		}

	}

	return false
}
