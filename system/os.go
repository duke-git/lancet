// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package system contain some functions about os, runtime, shell command.
package system

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
	"unicode/utf8"

	"github.com/duke-git/lancet/validator"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// IsWindows check if current os is windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux check if current os is linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMac check if current os is macos
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// GetOsEnv gets the value of the environment variable named by the key.
func GetOsEnv(key string) string {
	return os.Getenv(key)
}

// SetOsEnv sets the value of the environment variable named by the key.
func SetOsEnv(key, value string) error {
	return os.Setenv(key, value)
}

// RemoveOsEnv remove a single environment variable.
func RemoveOsEnv(key string) error {
	return os.Unsetenv(key)
}

// CompareOsEnv gets env named by the key and compare it with comparedEnv
func CompareOsEnv(key, comparedEnv string) bool {
	env := GetOsEnv(key)
	if env == "" {
		return false
	}
	return env == comparedEnv
}

// ExecCommand execute command, return the stdout and stderr string of command, and error if error occur
// param `command` is a complete command string, like, ls -a (linux), dir(windows), ping 127.0.0.1
// in linux,  use /bin/bash -c to execute command
// in windows, use powershell.exe to execute command
func ExecCommand(command string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	if IsWindows() {
		cmd = exec.Command("powershell.exe", command)
	}
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err = cmd.Run()

	if err != nil {
		if utf8.Valid(errOut.Bytes()) {
			stderr = byteToString(errOut.Bytes(), "UTF8")
		} else if validator.IsGBK(errOut.Bytes()) {
			stderr = byteToString(errOut.Bytes(), "GBK")
		}
		return
	}

	data := out.Bytes()
	if utf8.Valid(data) {
		stdout = byteToString(data, "UTF8")
	} else if validator.IsGBK(data) {
		stdout = byteToString(data, "GBK")
	}

	return
}

func byteToString(data []byte, charset string) string {
	var result string

	switch charset {
	case "GBK":
		decodeBytes, _ := simplifiedchinese.GBK.NewDecoder().Bytes(data)
		result = string(decodeBytes)
	case "GB18030":
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(data)
		result = string(decodeBytes)
	case "UTF8":
		fallthrough
	default:
		result = string(data)
	}

	return result
}

// GetOsBits get this system bits 32bit or 64bit
// return bit int (32/64)
func GetOsBits() int {
	return 32 << (^uint(0) >> 63)
}
