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
		t.Log(err)
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
		t.Log(err)
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
		t.Log(err)
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

func TestFileMode(t *testing.T) {
	assert := internal.NewAssert(t, "TestFileMode")

	srcFile := "./text.txt"
	CreateFile(srcFile)

	mode, err := FileMode(srcFile)
	assert.IsNil(err)

	t.Log(mode)

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
