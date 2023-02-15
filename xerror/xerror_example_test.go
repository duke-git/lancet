package xerror

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ExampleNew() {
	err := New("error")
	fmt.Println(err.Error())

	// Output:
	// error
}

func ExampleWrap() {
	err := New("wrong password")
	wrapErr := Wrap(err, "error")

	fmt.Println(wrapErr.Error())

	// Output:
	// error: wrong password
}

func ExampleXError_Wrap() {
	err1 := New("error").With("level", "high")
	err2 := err1.Wrap(errors.New("invalid username"))

	fmt.Println(err2.Error())

	// Output:
	// error: invalid username
}

func ExampleXError_Unwrap() {
	err1 := New("error").With("level", "high")
	err2 := err1.Wrap(errors.New("invalid username"))

	err := err2.Unwrap()

	fmt.Println(err.Error())

	// Output:
	// invalid username
}

func ExampleXError_StackTrace() {
	err := New("error")

	stacks := err.Stacks()

	fmt.Println(stacks[0].Func)
	fmt.Println(stacks[0].Line)

	containFile := strings.Contains(stacks[0].File, "xerror_example_test.go")
	fmt.Println(containFile)

	// Output:
	// github.com/duke-git/lancet/v2/xerror.ExampleXError_StackTrace
	// 52
	// true
}

func ExampleXError_With() {
	err := New("error").With("level", "high")

	errLevel := err.Values()["level"]

	fmt.Println(errLevel)

	// Output:
	// high
}

func ExampleXError_Id() {
	err1 := New("error").Id("e001")
	err2 := New("error").Id("e001")
	err3 := New("error").Id("e003")

	equal := err1.Is(err2)
	notEqual := err1.Is(err3)

	fmt.Println(equal)
	fmt.Println(notEqual)

	// Output:
	// true
	// false
}

func ExampleXError_Is() {
	err1 := New("error").Id("e001")
	err2 := New("error").Id("e001")
	err3 := New("error").Id("e003")

	equal := err1.Is(err2)
	notEqual := err1.Is(err3)

	fmt.Println(equal)
	fmt.Println(notEqual)

	// Output:
	// true
	// false
}

func ExampleXError_Values() {
	err := New("error").With("level", "high")

	errLevel := err.Values()["level"]

	fmt.Println(errLevel)

	// Output:
	// high
}

func ExampleXError_Info() {
	cause := errors.New("error")
	err := Wrap(cause, "invalid username").Id("e001").With("level", "high")

	errInfo := err.Info()

	fmt.Println(errInfo.Id)
	fmt.Println(errInfo.Cause)
	fmt.Println(errInfo.Values["level"])
	fmt.Println(errInfo.Message)

	// Output:
	// e001
	// error
	// high
	// invalid username
}

func ExampleTryUnwrap() {
	result1 := TryUnwrap(strconv.Atoi("42"))
	fmt.Println(result1)

	_, err := strconv.Atoi("4o2")
	defer func() {
		v := recover()
		result2 := reflect.DeepEqual(err.Error(), v.(*strconv.NumError).Error())
		fmt.Println(result2)
	}()

	TryUnwrap(strconv.Atoi("4o2"))

	// Output:
	// 42
	// true
}
