package fileutil

import (
	"io"
	"os"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestIsExist(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsExist")

	cases := []string{"./", "./file.go", "./a.txt"}
	expected := []bool{true, true, false}

	for i := 0; i < len(cases); i++ {
		actual := IsExist(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestCreateFile(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCreateFile")

	f := "./testdata/text.txt"
	if CreateFile(f) {
		file, err := os.Open(f)
		assert.IsNil(err)
		assert.Equal(f, file.Name())

		defer file.Close()
	} else {
		t.FailNow()
	}
	os.Remove(f)
}

func TestCreateDir(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCreateDir")

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	dirPath := pwd + "/a/b"
	err = CreateDir(dirPath)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(true, IsExist(dirPath))

	os.RemoveAll(pwd + "/a")

	assert.Equal(false, IsExist(dirPath))
}

func TestIsDir(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsDir")

	cases := []string{"./", "./a.txt"}
	expected := []bool{true, false}

	for i := 0; i < len(cases); i++ {
		actual := IsDir(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestRemoveFile(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRemoveFile")
	f := "./text.txt"
	if !IsExist(f) {
		CreateFile(f)
		err := RemoveFile(f)
		assert.IsNil(err)
	}
}

func TestCopyFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestCopyFile")

	srcFile := "./text.txt"
	CreateFile(srcFile)

	destFile := "./text_copy.txt"

	err := CopyFile(srcFile, destFile)
	if err != nil {
		file, err := os.Open(destFile)
		assert.IsNil(err)
		assert.Equal(destFile, file.Name())
	}
	os.Remove(srcFile)
	os.Remove(destFile)
}

func TestReadFileToString(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadFileToString")

	path := "./text.txt"
	CreateFile(path)

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()

	_, err := f.WriteString("hello world")
	if err != nil {
		t.Error(err)
	}

	content, _ := ReadFileToString(path)
	assert.Equal("hello world", content)

	os.Remove(path)
}

func TestClearFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestClearFile")

	path := "./text.txt"
	CreateFile(path)

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	defer f.Close()

	_, err := f.WriteString("hello world")
	if err != nil {
		t.Error(err)
	}

	err = ClearFile(path)
	assert.IsNil(err)

	content, _ := ReadFileToString(path)
	assert.Equal("", content)

	os.Remove(path)
}

func TestReadFileByLine(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadFileByLine")

	path := "./text.txt"
	CreateFile(path)

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)

	defer f.Close()

	_, err := f.WriteString("hello\nworld")
	if err != nil {
		t.Error(err)
	}

	expected := []string{"hello", "world"}
	actual, _ := ReadFileByLine(path)
	assert.Equal(expected, actual)

	os.Remove(path)
}

func TestZipAndUnZip(t *testing.T) {
	assert := internal.NewAssert(t, "TestZipAndUnZip")

	srcFile := "./text.txt"
	CreateFile(srcFile)

	file, _ := os.OpenFile(srcFile, os.O_WRONLY|os.O_TRUNC, 0777)
	defer file.Close()

	_, err := file.WriteString("hello\nworld")
	if err != nil {
		t.Fail()
	}

	zipFile := "./text.zip"
	err = Zip(srcFile, zipFile)
	assert.IsNil(err)

	unZipPath := "./unzip"
	err = UnZip(zipFile, unZipPath)
	assert.IsNil(err)

	unZipFile := "./unzip/text.txt"
	assert.Equal(true, IsExist(unZipFile))

	os.Remove(srcFile)
	os.Remove(zipFile)
	os.RemoveAll(unZipPath)
}

func TestZipAppendEntry(t *testing.T) {
	assert := internal.NewAssert(t, "TestZipAppendEntry")

	zipFile := "./text.zip"
	err := CopyFile("./testdata/file.go.zip", zipFile)
	assert.IsNil(err)

	srcFile := "./text.txt"
	CreateFile(srcFile)

	file, _ := os.OpenFile(srcFile, os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	_, err = file.WriteString("hello\nworld")
	if err != nil {
		t.Fail()
	}
	file.Close()

	err = ZipAppendEntry(srcFile, zipFile)
	assert.IsNil(err)

	err = ZipAppendEntry("./testdata", zipFile)
	assert.IsNil(err)

	unZipPath := "./unzip"
	err = UnZip(zipFile, unZipPath)
	assert.IsNil(err)

	assert.Equal(true, IsExist("./unzip/text.txt"))
	assert.Equal(true, IsExist("./unzip/file.go"))
	assert.Equal(true, IsExist("./unzip/testdata/file.go.zip"))
	assert.Equal(true, IsExist("./unzip/testdata/test.txt"))

	os.Remove(srcFile)
	os.Remove(zipFile)
	os.RemoveAll(unZipPath)
}

func TestZipFolder(t *testing.T) {
	// assert := internal.NewAssert(t, "TestZipFolder")

	// toZipFolder := "./tempdir/a/b"
	// zipFolder := "./tempdir/a/b.zip"

	// err := Zip(toZipFolder, zipFolder)
	// assert.IsNil(err)
	// assert.Equal(true, IsExist(zipFolder))

	// os.Remove(zipFolder)
}

func TestFileMode(t *testing.T) {
	assert := internal.NewAssert(t, "TestFileMode")

	srcFile := "./text.txt"
	CreateFile(srcFile)

	mode, err := FileMode(srcFile)

	assert.IsNotNil(mode)
	assert.IsNil(err)

	os.Remove(srcFile)
}

func TestIsLink(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsLink")

	srcFile := "./text.txt"
	CreateFile(srcFile)

	linkFile := "./text.link"
	if !IsExist(linkFile) {
		_ = os.Symlink(srcFile, linkFile)
	}
	assert.Equal(true, IsLink(linkFile))

	assert.Equal(false, IsLink("./file.go"))

	os.Remove(srcFile)
	os.Remove(linkFile)
}

func TestMiMeType(t *testing.T) {
	assert := internal.NewAssert(t, "TestMiMeType")

	f, _ := os.Open("./file.go")
	defer f.Close()
	assert.Equal("text/plain; charset=utf-8", MiMeType(f))
	assert.Equal("text/plain; charset=utf-8", MiMeType("./file.go"))
}

func TestListFileNames(t *testing.T) {
	assert := internal.NewAssert(t, "TestListFileNames")

	filesInPath, err := ListFileNames("../internal")
	assert.IsNil(err)

	expected := []string{"assert.go", "assert_test.go", "error_join.go"}
	assert.Equal(expected, filesInPath)
}

func TestCurrentPath(t *testing.T) {
	absPath := CurrentPath()
	t.Log(absPath)
}

func TestIsZipFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsZipFile")

	assert.Equal(false, IsZipFile("./file.go"))
	assert.Equal(true, IsZipFile("./testdata/file.go.zip"))
}

func TestFileSize(t *testing.T) {
	assert := internal.NewAssert(t, "TestFileSize")

	size, err := FileSize("./testdata/test.txt")

	assert.IsNil(err)
	assert.Equal(int64(20), size)
}

func TestMTime(t *testing.T) {
	assert := internal.NewAssert(t, "TestMTime")

	mtime, err := MTime("./testdata/test.txt")
	t.Log("TestMTime", mtime)
	assert.IsNil(err)
	// assert.Equal(int64(1682478195), mtime)
}

func TestSha(t *testing.T) {
	assert := internal.NewAssert(t, "TestSha")

	sha1, err := Sha("./testdata/test.txt", 1)
	assert.IsNil(err)

	sha256, err := Sha("./testdata/test.txt", 256)
	assert.IsNil(err)

	sha512, err := Sha("./testdata/test.txt", 512)
	assert.IsNil(err)

	assert.Equal("dda3cf10c5a6ff6c6659a497bf7261b287af2bc7", sha1)
	assert.Equal("aa6d0a3fbc3442c228d606da09e0c1dc98c69a1cac3da1909199e0266171df35", sha256)
	assert.Equal("d22aba2a1b7a2e2f512756255cc1c3708905646920cb1eb95e45b531ba74774dbbb89baebf1f716220eb9cf4908f1cfc5b2a01267704d9a59f59d77cab609870", sha512)
}

func TestReadCsvFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadCsvFile")

	content, err := ReadCsvFile("./testdata/demo.csv")

	assert.IsNil(err)

	assert.Equal(3, len(content))
	assert.Equal(3, len(content[0]))
	assert.Equal("Bob", content[0][0])
}

func TestWriteCsvFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestWriteCsvFile")

	csvFilePath := "./testdata/test1.csv"
	content := [][]string{
		{"Lili", "22", "female"},
		{"Jim", "21", "male"},
	}

	err := WriteCsvFile(csvFilePath, content, false)
	assert.IsNil(err)

	readContent, err := ReadCsvFile(csvFilePath)

	assert.IsNil(err)

	assert.Equal(2, len(readContent))
	assert.Equal(3, len(readContent[0]))
	assert.Equal("Lili", readContent[0][0])

	// RemoveFile(csvFilePath)
}

func TestWriteMapsToCsv(t *testing.T) {
	assert := internal.NewAssert(t, "TestWriteMapsToCSV")

	csvFilePath := "./testdata/test4.csv"
	records := []map[string]any{
		{"Name": "Lili", "Age": "22", "Gender": "female"},
		{"Name": "Jim", "Age": "21", "Gender": "male"},
	}

	headers := []string{"Name", "Age", "Gender"}
	err := WriteMapsToCsv(csvFilePath, records, false, ';', headers)

	assert.IsNil(err)

	content, err := ReadCsvFile(csvFilePath, ';')

	assert.IsNil(err)

	assert.Equal(3, len(content))
	assert.Equal(3, len(content[0]))
	assert.Equal("Lili", content[1][0])
	assert.Equal("22", content[1][1])
	assert.Equal("female", content[1][2])
}

func TestWriteStringToFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestWriteStringToFile")

	filepath := "./test.txt"

	file, err := os.Create(filepath)
	if err != nil {
		t.Fail()
	}

	err = WriteStringToFile(filepath, "hello world", false)
	if err != nil {
		t.Fail()
	}

	content1, err := ReadFileToString(filepath)
	if err != nil {
		t.Fail()
	}

	err = WriteStringToFile(filepath, "hello", false)
	if err != nil {
		t.Fail()
	}

	content2, err := ReadFileToString(filepath)
	if err != nil {
		t.Fail()
	}

	err = WriteStringToFile(filepath, " world", true)
	if err != nil {
		t.Fail()
	}

	content3, err := os.ReadFile(filepath)
	if err != nil {
		t.Fail()
	}

	assert.Equal("hello world", content1)
	assert.Equal("hello", content2)
	assert.Equal("hello world", string(content3))

	_ = file.Close()
	_ = os.Remove(filepath)
}

func TestWriteBytesToFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestWriteBytesToFile")

	filepath := "./bytes.txt"

	file, err := os.Create(filepath)
	if err != nil {
		t.Fail()
	}

	defer file.Close()

	err = WriteBytesToFile(filepath, []byte("hello"))
	if err != nil {
		t.Fail()
	}

	content, err := os.ReadFile(filepath)
	if err != nil {
		t.Fail()
	}

	assert.Equal("hello", string(content))

	os.Remove(filepath)
}

func TestReadFile(t *testing.T) {
	reader, close, err := ReadFile("https://httpbin.org/robots.txt")
	if err != nil {
		t.Fail()
	}
	defer close()

	dat, err := io.ReadAll(reader)
	if err != nil {
		t.Fail()
	}

	want := `User-agent: *
Disallow: /deny
`
	internal.NewAssert(t, "TestReadFile").Equal(want, string(dat))
}

func TestReadlineFile(t *testing.T) {
	path := "./testdata/demo.csv"
	reader, err := NewFileReader(path)
	if err != nil {
		t.Fail()
	}
	defer reader.Close()

	indexMap := make(map[string]int64)
	defer reader.Close()
	for {
		offset := reader.Offset()
		line, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		indexMap[line] = offset
	}

	lines, err := ReadFileByLine(path)
	if err != nil {
		t.Fail()
	}
	for _, line := range lines {
		offset, ok := indexMap[line]
		if !ok {
			t.Fail()
		}
		if err = reader.Seek(offset); err != nil {
			t.Fail()
		}
		lineRead, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		internal.NewAssert(t, "TestReadlineFile").Equal(line, lineRead)
	}
}
