package fileutil

import (
	"fmt"
	"os"
)

func ExampleIsExist() {

	result1 := IsExist("./")
	result2 := IsExist("./xxx.go")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleCreateFile() {
	fname := "./a.txt"

	result1 := IsExist(fname)

	CreateFile(fname)

	result2 := IsExist(fname)

	os.Remove(fname)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// false
	// true
}

func ExampleCreateDir() {
	pwd, _ := os.Getwd()
	dirPath := pwd + "/test_xxx/"

	result1 := IsExist(dirPath)

	err := CreateDir(dirPath)
	if err != nil {
		return
	}

	result2 := IsExist(dirPath)

	os.Remove(dirPath)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// false
	// true
}

func ExampleIsDir() {

	result1 := IsDir("./")
	result2 := IsDir("./xxx.go")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleRemoveFile() {
	srcFile := "./text.txt"
	CreateFile(srcFile)

	copyFile := "./text_copy.txt"
	err := CopyFile(srcFile, copyFile)
	if err != nil {
		return
	}
	file, err := os.Open(copyFile)
	if err != nil {
		return
	}
	result1 := IsExist(copyFile)
	result2 := file.Name()

	os.Remove(srcFile)
	os.Remove(copyFile)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// ./text_copy.txt
}

func ExampleReadFileToString() {
	fname := "./test.txt"
	CreateFile(fname)

	f, _ := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()

	_, err := f.WriteString("hello world")
	if err != nil {
		return
	}

	content, _ := ReadFileToString(fname)

	os.Remove(fname)

	fmt.Println(content)

	// Output:
	// hello world
}

func ExampleClearFile() {
	fname := "./test.txt"
	CreateFile(fname)

	f, _ := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()

	_, err := f.WriteString("hello world")
	if err != nil {
		return
	}

	content1, _ := ReadFileToString(fname)

	err = ClearFile(fname)
	if err != nil {
		return
	}
	content2, _ := ReadFileToString(fname)

	os.Remove(fname)

	fmt.Println(content1)
	fmt.Println(content2)

	// Output:
	// hello world
	//
}

func ExampleReadFileByLine() {
	fname := "./test.txt"
	CreateFile(fname)

	f, _ := os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()

	_, err := f.WriteString("hello\nworld")
	if err != nil {
		return
	}

	content, _ := ReadFileByLine(fname)

	os.Remove(fname)

	fmt.Println(content)

	// Output:
	// [hello world]
}

func ExampleListFileNames() {
	fileList, _ := ListFileNames("../formatter/")
	fmt.Println(fileList)

	// Output:
	// [formatter.go formatter_example_test.go formatter_test.go]
}

func ExampleZip() {
	srcFile := "./test.txt"
	CreateFile(srcFile)

	zipFile := "./test.zip"
	err := Zip(srcFile, zipFile)
	if err != nil {
		return
	}

	result := IsExist(zipFile)

	os.Remove(srcFile)
	os.Remove(zipFile)

	fmt.Println(result)

	// Output:
	// true
}

func ExampleUnZip() {
	fname := "./test.txt"
	file, _ := os.Create(fname)

	_, err := file.WriteString("hello\nworld")
	if err != nil {
		return
	}

	f, _ := os.Open(fname)
	defer f.Close()

	mimeType := MiMeType(f)
	fmt.Println(mimeType)

	os.Remove(fname)

	// Output:
	// application/octet-stream
}
