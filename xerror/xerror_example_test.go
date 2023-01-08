package xerror

import (
	"fmt"
	"reflect"
	"strconv"
)

func ExampleUnwrap() {
	result1 := Unwrap(strconv.Atoi("42"))
	fmt.Println(result1)

	_, err := strconv.Atoi("4o2")
	defer func() {
		v := recover()
		result2 := reflect.DeepEqual(err.Error(), v.(*strconv.NumError).Error())
		fmt.Println(result2)
	}()

	Unwrap(strconv.Atoi("4o2"))

	// Output:
	// 42
	// true
}
