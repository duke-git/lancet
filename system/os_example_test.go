package system

import (
	"fmt"
	"time"
)

func ExampleSetOsEnv() {
	err := SetOsEnv("foo", "abc")
	result := GetOsEnv("foo")

	fmt.Println(err)
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
	err1 := SetOsEnv("foo", "abc")
	result1 := GetOsEnv("foo")

	err2 := RemoveOsEnv("foo")
	result2 := GetOsEnv("foo")

	fmt.Println(err1)
	fmt.Println(err2)
	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// <nil>
	// <nil>
	// abc
	//
}

func ExampleCompareOsEnv() {
	err := SetOsEnv("foo", "abc")
	if err != nil {
		return
	}

	result := CompareOsEnv("foo", "abc")

	fmt.Println(result)

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

func ExampleStartProcess() {
	pid, err := StartProcess("sleep", "2")
	if err != nil {
		return
	}

	fmt.Println(pid)
}

func ExampleStopProcess() {
	pid, err := StartProcess("sleep", "10")
	if err != nil {
		return
	}
	time.Sleep(1 * time.Second)

	err = StopProcess(pid)

	fmt.Println(err)

	// Output:
	// <nil>
}

func ExampleKillProcess() {
	pid, err := StartProcess("sleep", "3")
	if err != nil {
		return
	}
	time.Sleep(1 * time.Second)

	err = KillProcess(pid)

	fmt.Println(err)

	// Output:
	// <nil>
}
