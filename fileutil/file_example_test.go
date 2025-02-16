package fileutil

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
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
	dirPath := pwd + "/createdir/a/b"

	result1 := IsExist(dirPath)

	err := CreateDir(dirPath)
	if err != nil {
		return
	}

	result2 := IsExist(pwd + "/createdir/")
	result3 := IsExist(pwd + "/createdir/a")
	result4 := IsExist(pwd + "/createdir/a/b")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	os.RemoveAll(pwd + "/createdir/")

	// Output:
	// false
	// true
	// true
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
	fileList, _ := ListFileNames("../internal")
	fmt.Println(fileList)

	// Output:
	// [assert.go assert_test.go error_join.go]
}

func ExampleMiMeType() {
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
	zipFile := "./testdata/file.go.zip"

	err := UnZip(zipFile, "./testdata")
	if err != nil {
		return
	}

	exist := IsExist("./testdata/file.go")
	fmt.Println(exist)

	os.Remove("./testdata/file.go")

	// Output:
	// true
}

func ExampleZipAppendEntry() {
	zipFile := "./test.zip"
	CopyFile("./testdata/file.go.zip", zipFile)

	ZipAppendEntry("./testdata", zipFile)

	unZipPath := "./unzip"
	UnZip(zipFile, unZipPath)

	fmt.Println(IsExist("./unzip/file.go"))
	fmt.Println(IsExist("./unzip/testdata/file.go.zip"))
	fmt.Println(IsExist("./unzip/testdata/test.txt"))

	os.Remove(zipFile)
	os.RemoveAll(unZipPath)

	// Output:
	// true
	// true
	// true
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
	content, err := ReadCsvFile("./testdata/demo.csv")

	fmt.Println(content)
	fmt.Println(err)

	// Output:
	// [[Bob  12  male] [Duke  14  male] [Lucy  16  female]]
	// <nil>
}

func ExampleWriteCsvFile() {
	data := [][]string{
		{"Lili", "22", "female"},
		{"Jim", "21", "male"},
	}
	err := WriteCsvFile("./testdata/test2.csv", data, false)
	fmt.Println(err)

	content, _ := ReadCsvFile("./testdata/test2.csv")
	fmt.Println(content)

	// Output:
	// <nil>
	// [[Lili 22 female] [Jim 21 male]]
}

func ExampleWriteMapsToCsv() {
	csvFilePath := "./testdata/test3.csv"
	records := []map[string]any{
		{"Name": "Lili", "Age": "22", "Gender": "female"},
		{"Name": "Jim", "Age": "21", "Gender": "male"},
	}

	headers := []string{"Name", "Age", "Gender"}
	err := WriteMapsToCsv(csvFilePath, records, false, ';', headers)

	if err != nil {
		log.Fatal(err)
	}

	content, err := ReadCsvFile(csvFilePath, ';')

	fmt.Println(content)

	// Output:
	// [[Name Age Gender] [Lili 22 female] [Jim 21 male]]
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

func ExampleReadFile() {
	reader, fn, err := ReadFile("https://httpbin.org/robots.txt")
	if err != nil {
		return
	}
	defer fn()

	dat, err := io.ReadAll(reader)
	if err != nil {
		return
	}

	fmt.Println(string(dat))

	// Output:
	// User-agent: *
	// Disallow: /deny
}

func ExampleChunkRead() {
	const mb = 1024 * 1024
	const defaultChunkSizeMB = 100

	filePath := "./testdata/test1.csv"
	f, err := os.Open(filePath)
	if err != nil {
		return
	}

	defer f.Close()

	var bufPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 0, defaultChunkSizeMB*mb)
		},
	}

	lines, err := ChunkRead(f, 0, 100, &bufPool)
	if err != nil {
		return
	}

	fmt.Println(lines[0])
	fmt.Println(lines[1])

	// Output:
	// Lili,22,female
	// Jim,21,male
}

func ExampleParallelChunkRead() {
	const mb = 1024 * 1024
	const defaultChunkSizeMB = 100 // 默认值

	numParsers := runtime.NumCPU()

	linesCh := make(chan []string, numParsers)
	filePath := "./testdata/test1.csv"

	go ParallelChunkRead(filePath, linesCh, defaultChunkSizeMB, numParsers)

	var totalLines int
	for lines := range linesCh {
		totalLines += len(lines)

		for _, line := range lines {
			fmt.Println(line)
		}
	}

	fmt.Println(totalLines)

	// Output:
	// Lili,22,female
	// Jim,21,male
	// 2
}
