package fileutil

import (
	"os"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestIsExist(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsExist")

	cases := []string{"./", "./file.go", "./a.txt"}
	expected := []bool{true, true, false}

	for i := 0; i < len(cases); i++ {
		actual := IsExist(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestCreateFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestCreateFile")

	f := "./text.txt"
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
	assert := internal.NewAssert(t, "TestCreateDir")

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	dirPath := pwd + "/a/"
	err = CreateDir(dirPath)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	assert.Equal(true, IsExist(dirPath))
	os.Remove(dirPath)
	assert.Equal(false, IsExist(dirPath))
}

func TestIsDir(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsDir")

	cases := []string{"./", "./a.txt"}
	expected := []bool{true, false}

	for i := 0; i < len(cases); i++ {
		actual := IsDir(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestRemoveFile(t *testing.T) {
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
	assert.Equal(true, IsExist("./unzip/testdata/test.csv"))
	assert.Equal(true, IsExist("./unzip/testdata/test.txt"))

	os.Remove(srcFile)
	os.Remove(zipFile)
	os.RemoveAll(unZipPath)
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

	filesInPath, err := ListFileNames("./")
	assert.IsNil(err)

	expected := []string{"file.go", "file_example_test.go", "file_test.go"}
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
	sha256, err := Sha("./testdata/test.txt", 256)
	sha512, err := Sha("./testdata/test.txt", 512)

	assert.IsNil(err)
	assert.Equal("dda3cf10c5a6ff6c6659a497bf7261b287af2bc7", sha1)
	assert.Equal("aa6d0a3fbc3442c228d606da09e0c1dc98c69a1cac3da1909199e0266171df35", sha256)
	assert.Equal("d22aba2a1b7a2e2f512756255cc1c3708905646920cb1eb95e45b531ba74774dbbb89baebf1f716220eb9cf4908f1cfc5b2a01267704d9a59f59d77cab609870", sha512)
}

func TestReadCsvFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestReadCsvFile")

	content, err := ReadCsvFile("./testdata/test.csv")

	assert.IsNil(err)

	assert.Equal(3, len(content))
	assert.Equal(3, len(content[0]))
	assert.Equal("Bob", content[0][0])
}

func TestWriteStringToFile(t *testing.T) {
	assert := internal.NewAssert(t, "TestWriteStringToFile")

	filepath := "./test.txt"

	file, err := os.Create(filepath)
	if err != nil {
		t.Fail()
	}

	defer file.Close()

	err = WriteStringToFile(filepath, "hello", false)
	if err != nil {
		t.Fail()
	}

	content1, err := ReadFileToString(filepath)
	if err != nil {
		t.Fail()
	}

	err = WriteStringToFile(filepath, " world", true)
	if err != nil {
		t.Fail()
	}

	content2, err := os.ReadFile(filepath)
	if err != nil {
		t.Fail()
	}

	assert.Equal("hello", content1)
	assert.Equal("hello world", string(content2))

	os.Remove(filepath)
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
