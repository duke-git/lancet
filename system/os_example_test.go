package system

import "fmt"

func ExampleSetOsEnv() {
	ok := SetOsEnv("foo", "abc")
	result := GetOsEnv("foo")

	fmt.Println(ok)
	fmt.Println(result)
	// Output:
	// <nil>
	// abc
}

func ExampleGetOsEnv() {
	ok := SetOsEnv("foo", "abc")
	result := GetOsEnv("foo")

	fmt.Println(ok)
	fmt.Println(result)
	// Output:
	// <nil>
	// abc
}

func ExampleRemoveOsEnv() {
	ok1 := SetOsEnv("foo", "abc")
	result1 := GetOsEnv("foo")

	ok2 := RemoveOsEnv("foo")
	result2 := GetOsEnv("foo")

	fmt.Println(ok1)
	fmt.Println(ok2)
	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// <nil>
	// <nil>
	// abc
	//
}

func ExampleCompareOsEnv() {
	SetOsEnv("foo", "abc")
	result1 := CompareOsEnv("foo", "abc")

	fmt.Println(result1)
	// Output:
	// true
}

func ExampleExecCommand() {
	_, stderr, err := ExecCommand("ls")
	// fmt.Println(stdout)
	fmt.Println(stderr)
	fmt.Println(err)

	// Output:
	//
	// <nil>
}

func ExampleGetOsBits() {
	osBits := GetOsBits()

	fmt.Println(osBits)
	// Output:
	// 64
}
