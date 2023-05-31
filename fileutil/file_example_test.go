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
	fileList, _ := ListFileNames("./")
	fmt.Println(fileList)

	// Output:
	// [file.go file_example_test.go file_test.go]
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

func ExampleIsZipFile() {
	result1 := IsZipFile("./file.go")
	result2 := IsZipFile("./testdata/file.go.zip")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// false
	// true
}

func ExampleFileSize() {
	size, err := FileSize("./testdata/test.txt")

	fmt.Println(size)
	fmt.Println(err)

	// Output:
	// 20
	// <nil>
}

// func ExampleMTime() {
// 	mtime, err := MTime("./testdata/test.txt")

// 	fmt.Println(mtime) // 1682478195 (unix timestamp)
// 	fmt.Println(err)

// 	// Output:
// 	// 1682478195
// 	// <nil>
// }

func ExampleSha() {
	sha1, err := Sha("./testdata/test.txt", 1)
	sha256, _ := Sha("./testdata/test.txt", 256)
	sha512, _ := Sha("./testdata/test.txt", 512)

	fmt.Println(sha1)
	fmt.Println(sha256)
	fmt.Println(sha512)
	fmt.Println(err)

	// Output:
	// dda3cf10c5a6ff6c6659a497bf7261b287af2bc7
	// aa6d0a3fbc3442c228d606da09e0c1dc98c69a1cac3da1909199e0266171df35
	// d22aba2a1b7a2e2f512756255cc1c3708905646920cb1eb95e45b531ba74774dbbb89baebf1f716220eb9cf4908f1cfc5b2a01267704d9a59f59d77cab609870
	// <nil>
}

func ExampleReadCsvFile() {
	content, err := ReadCsvFile("./testdata/test.csv")

	fmt.Println(content)
	fmt.Println(err)

	// Output:
	// [[Bob  12  male] [Duke  14  male] [Lucy  16  female]]
	// <nil>
}

func ExampleWriteStringToFile() {
	filepath := "./test.txt"

	file, err := os.Create(filepath)
	if err != nil {
		return
	}

	defer file.Close()

	err = WriteStringToFile(filepath, "hello", true)
	if err != nil {
		return
	}

	content, err := ReadFileToString(filepath)
	if err != nil {
		return
	}

	os.Remove(filepath)

	fmt.Println(content)

	// Output:
	// hello
}

func ExampleWriteBytesToFile() {
	filepath := "./bytes.txt"

	file, err := os.Create(filepath)
	if err != nil {
		return
	}

	defer file.Close()

	err = WriteBytesToFile(filepath, []byte("hello"))
	if err != nil {
		return
	}

	content, err := ReadFileToString(filepath)
	if err != nil {
		return
	}

	os.Remove(filepath)

	fmt.Println(content)

	// Output:
	// hello
}
