// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package system contain some functions about os, runtime, shell command.
package system

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/duke-git/lancet/v2/validator"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type (
	Option func(*exec.Cmd)
)

// IsWindows check if current os is windows.
// Play: https://go.dev/play/p/XzJULbzmf9m
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux check if current os is linux.
// Play: https://go.dev/play/p/zIflQgZNuxD
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMac check if current os is macos.
// Play: https://go.dev/play/p/Mg4Hjtyq7Zc
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// GetOsEnv gets the value of the environment variable named by the key.
// Play: https://go.dev/play/p/D88OYVCyjO-
func GetOsEnv(key string) string {
	return os.Getenv(key)
}

// SetOsEnv sets the value of the environment variable named by the key.
// Play: https://go.dev/play/p/D88OYVCyjO-
func SetOsEnv(key, value string) error {
	return os.Setenv(key, value)
}

// RemoveOsEnv remove a single environment variable.
// Play: https://go.dev/play/p/fqyq4b3xUFQ
func RemoveOsEnv(key string) error {
	return os.Unsetenv(key)
}

// CompareOsEnv gets env named by the key and compare it with comparedEnv.
// Play: https://go.dev/play/p/BciHrKYOHbp
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
// Play: https://go.dev/play/p/n-2fLyZef-4
func ExecCommand(command string, opts ...Option) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	if IsWindows() {
		cmd = exec.Command("powershell.exe", command)
	}

	for _, opt := range opts {
		if opt != nil {
			opt(cmd)
		}
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

// GetOsBits return current os bits (32 or 64).
// Play: https://go.dev/play/p/ml-_XH3gJbW
func GetOsBits() int {
	return 32 << (^uint(0) >> 63)
}

// StartProcess start a new process with the specified name and arguments.
// Play: todo
func StartProcess(command string, args ...string) (int, error) {
	cmd := exec.Command(command, args...)

	if err := cmd.Start(); err != nil {
		return 0, err
	}

	return cmd.Process.Pid, nil
}

// StopProcess stop a process by pid.
// Play: todo
func StopProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Signal(os.Kill)
}

// KillProcess kill a process by pid.
// Play: todo
func KillProcess(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Kill()
}

// ProcessInfo contains detailed information about a process.
type ProcessInfo struct {
	PID                int
	CPU                string
	Memory             string
	State              string
	User               string
	Cmd                string
	Threads            []string
	IOStats            string
	StartTime          string
	ParentPID          int
	NetworkConnections string
}

// GetProcessInfo retrieves detailed process information by pid.
// Play: todo
func GetProcessInfo(pid int) (*ProcessInfo, error) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid), "/FO", "CSV", "/V")
	} else {
		cmd = exec.Command("ps", "-p", strconv.Itoa(pid), "-o", "pid,%cpu,%mem,state,user,comm")
	}

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	processInfo, err := parseProcessInfo(output, pid)
	if err != nil {
		return nil, err
	}

	if runtime.GOOS != "windows" {
		processInfo.Threads, _ = getThreadsInfo(pid)
		processInfo.IOStats, _ = getIOStats(pid)
		processInfo.StartTime, _ = getProcessStartTime(pid)
		processInfo.ParentPID, _ = getParentProcess(pid)
		processInfo.NetworkConnections, _ = getNetworkConnections(pid)
	}

	return processInfo, nil
}

// parseProcessInfo parses the output of `ps` or `tasklist` to fill the ProcessInfo structure.
func parseProcessInfo(output []byte, pid int) (*ProcessInfo, error) {
	lines := strings.Split(string(output), "\n")

	if len(lines) < 2 {
		return nil, fmt.Errorf("no process found with PID %d", pid)
	}

	var processInfo ProcessInfo
	if runtime.GOOS == "windows" {
		fields := strings.Split(lines[1], "\",\"")
		if len(fields) < 9 {
			return nil, fmt.Errorf("unexpected tasklist output format")
		}

		processInfo = ProcessInfo{
			PID:    pid,
			CPU:    "N/A",
			Memory: fields[4], // Memory usage in K
			State:  fields[5],
			User:   "N/A",
			Cmd:    fields[8],
		}
	} else {
		fields := strings.Fields(lines[1])
		if len(fields) < 6 {
			return nil, fmt.Errorf("unexpected ps output format")
		}

		processInfo = ProcessInfo{
			PID:    pid,
			CPU:    fields[1],
			Memory: fields[2],
			State:  fields[3],
			User:   fields[4],
			Cmd:    fields[5],
		}
	}

	return &processInfo, nil
}

func getThreadsInfo(pid int) ([]string, error) {
	cmd := exec.Command("ps", "-T", "-p", strconv.Itoa(pid))
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")

	var threads []string
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) != "" {
			threads = append(threads, line)
		}
	}

	return threads, nil
}

func getIOStats(pid int) (string, error) {
	filePath := fmt.Sprintf("/proc/%d/io", pid)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func getProcessStartTime(pid int) (string, error) {
	cmd := exec.Command("ps", "-p", strconv.Itoa(pid), "-o", "lstart=")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func getParentProcess(pid int) (int, error) {
	cmd := exec.Command("ps", "-o", "ppid=", "-p", strconv.Itoa(pid))
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	ppid, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		return 0, err
	}

	return ppid, nil
}

func getNetworkConnections(pid int) (string, error) {
	cmd := exec.Command("lsof", "-p", strconv.Itoa(pid), "-i")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
