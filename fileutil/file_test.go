package fileutil

import (
	"os"
	"reflect"
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestIsExist(t *testing.T) {
	cases := []string{"./", "./file.go", "./a.txt"}
	expected := []bool{true, true, false}

	for i := 0; i < len(cases); i++ {
		res := IsExist(cases[i])
		if res != expected[i] {
			internal.LogFailedTestInfo(t, "IsExist", cases[i], expected[i], res)
			t.FailNow()
		}
	}
}

func TestCreateFile(t *testing.T) {
	f := "./text.txt"
	if CreateFile(f) {
		file, err := os.Open(f)
		if err != nil {
			internal.LogFailedTestInfo(t, "CreateFile", f, f, "create file error: "+err.Error())
			t.FailNow()
		}
		if file.Name() != f {
			internal.LogFailedTestInfo(t, "CreateFile", f, f, file.Name())
			t.FailNow()
		}
	} else {
		internal.LogFailedTestInfo(t, "CreateFile", f, f, "create file error")
		t.FailNow()
	}
}

func TestIsDir(t *testing.T) {
	cases := []string{"./", "./a.txt"}
	expected := []bool{true, false}

	for i := 0; i < len(cases); i++ {
		res := IsDir(cases[i])
		if res != expected[i] {
			internal.LogFailedTestInfo(t, "IsDir", cases[i], expected[i], res)
			t.FailNow()
		}
	}
}

func TestRemoveFile(t *testing.T) {
	f := "./text.txt"
	if CreateFile(f) {
		err := RemoveFile(f)
		if err != nil {
			internal.LogFailedTestInfo(t, "RemoveFile", f, f, err.Error())
			t.FailNow()
		}
	} else {
		internal.LogFailedTestInfo(t, "RemoveFile", f, f, "create file error")
		t.FailNow()
	}
}

func TestCopyFile(t *testing.T) {
	srcFile := "./text.txt"
	CreateFile(srcFile)

	dstFile := "./text_copy.txt"

	err := CopyFile(srcFile, dstFile)
	if err != nil {
		file, err := os.Open(dstFile)
		if err != nil {
			internal.LogFailedTestInfo(t, "CopyFile", srcFile, dstFile, "create file error: "+err.Error())
			t.FailNow()
		}
		if file.Name() != dstFile {
			internal.LogFailedTestInfo(t, "CopyFile", srcFile, dstFile, file.Name())
			t.FailNow()
		}
	}
}

func TestListFileNames(t *testing.T) {
	filesInCurrentPath, err := ListFileNames("../datetime/")
	if err != nil {
		t.FailNow()
	}
	expected := []string{"datetime.go", "datetime_test.go"}
	if !reflect.DeepEqual(filesInCurrentPath, expected) {
		internal.LogFailedTestInfo(t, "ToChar", "./", expected, filesInCurrentPath)
		t.FailNow()
	}
}

func TestReadFileToString(t *testing.T) {
	path := "./text.txt"
	CreateFile(path)
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	f.WriteString("hello world")

	res, _ := ReadFileToString(path)
	if res != "hello world" {
		internal.LogFailedTestInfo(t, "ReadFileToString", path, "hello world", res)
	}
}

func TestClearFile(t *testing.T) {
	path := "./text.txt"
	CreateFile(path)
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	f.WriteString("hello world")

	CreateFile(path)

	res, _ := ReadFileToString(path)
	if res != "" {
		internal.LogFailedTestInfo(t, "CreateFile", path, "", res)
	}
}

func TestReadFileByLine(t *testing.T) {
	path := "./text.txt"
	CreateFile(path)
	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	f.WriteString("hello\nworld")

	expected := []string{"hello", "world"}
	res, _ := ReadFileByLine(path)
	if !reflect.DeepEqual(res, expected) {
		internal.LogFailedTestInfo(t, "ReadFileByLine", path, expected, res)
	}
}
