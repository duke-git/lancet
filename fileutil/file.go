// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package fileutil implements some basic functions for file operations
package fileutil

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

// IsExist checks if a file or directory exists
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

// CreateFile create a file in path
func CreateFile(path string) bool {
	file, err := os.Create(path)
	if err != nil {
		return false
	}

	defer file.Close()
	return true
}

// IsDir checks if the path is directory or not
func IsDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return file.IsDir()
}

// RemoveFile remove the path file
func RemoveFile(path string) error {
	return os.Remove(path)
}

// CopyFile copy src file to dest file
func CopyFile(srcFilePath string, dstFilePath string) error {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	distFile, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer distFile.Close()

	var tmp = make([]byte, 1024*4)
	for {
		n, err := srcFile.Read(tmp)
		distFile.Write(tmp[:n])
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

//ClearFile write empty string to path file
func ClearFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("")
	return err
}

//ReadFileToString return string of file content
func ReadFileToString(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ReadFileByLine
func ReadFileByLine(path string)([]string, error)  {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]string, 0)
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		l := string(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		res = append(res, l)
	}

	return res, nil
}

// ListFileNames return all file names in the path
func ListFileNames(path string) ([]string, error) {
	if !IsExist(path) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	res := []string{}
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			res = append(res, fs[i].Name())
		}
	}

	return res, nil
}
