package formatter

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func commaString(s string) string {
	if len(s) <= 3 {
		return s
	}
	return commaString(s[:len(s)-3]) + "," + commaString(s[len(s)-3:])
}

func numberToString(value any) (string, error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%v", value), nil

		// todo: need to handle 12345678.9 => 1.23456789e+07
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%v", value), nil

	case reflect.String:
		{
			sv := fmt.Sprintf("%v", value)
			if strings.Contains(sv, ".") {
				_, err := strconv.ParseFloat(sv, 64)
				if err != nil {
					return "", err
				}
				return sv, nil
			} else {
				_, err := strconv.ParseInt(sv, 10, 64)
				if err != nil {
					return "", nil
				}
				return sv, nil
			}
		}
	default:
		return "", nil
	}
}
