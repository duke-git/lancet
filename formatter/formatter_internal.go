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

func numString(value any) string {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int, reflect.Int64, reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%v", value)
	case reflect.String:
		{
			sv := fmt.Sprintf("%v", value)
			if strings.Contains(sv, ".") {
				_, err := strconv.ParseFloat(sv, 64)
				if err == nil {
					return sv
				}
			} else {
				_, err := strconv.ParseInt(sv, 10, 64)
				if err == nil {
					return sv
				}
			}
		}
	default:
		return ""
	}
	return ""
}
