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
	os.Remove(f)
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

	destFile := "./text_copy.txt"

	err := CopyFile(srcFile, destFile)
	if err != nil {
		file, err := os.Open(destFile)
		if err != nil {
			internal.LogFailedTestInfo(t, "CopyFile", srcFile, destFile, "create file error: "+err.Error())
			t.FailNow()
		}
		if file.Name() != destFile {
			internal.LogFailedTestInfo(t, "CopyFile", srcFile, destFile, file.Name())
			t.FailNow()
		}
	}
	os.Remove(srcFile)
	os.Remove(destFile)
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
		t.FailNow()
	}
	os.Remove(path)
}

func TestClearFile(t *testing.T) {
	path := "./text.txt"
	CreateFile(path)

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	f.WriteString("hello world")

	err := ClearFile(path)
	if err != nil {
		t.Error("Clear file error: ", err)
	}
	fileContent, _ := ReadFileToString(path)
	if fileContent != "" {
		internal.LogFailedTestInfo(t, "ClearFile", path, "", fileContent)
		t.FailNow()
	}
	os.Remove(path)
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
		t.FailNow()
	}
	os.Remove(path)
}

func TestZipAndUnZip(t *testing.T) {
	srcFile := "./text.txt"
	CreateFile(srcFile)

	file, _ := os.OpenFile(srcFile, os.O_WRONLY|os.O_TRUNC, 0777)
	file.WriteString("hello\nworld")

	zipFile := "./text.zip"
	err := Zip(srcFile, zipFile)
	if err != nil {
		internal.LogFailedTestInfo(t, "Zip", srcFile, zipFile, err)
		t.FailNow()
	}

	unZipPath := "./unzip"
	err = UnZip(zipFile, unZipPath)
	if err != nil {
		internal.LogFailedTestInfo(t, "UnZip", srcFile, unZipPath, err)
		t.FailNow()
	}

	unZipFile := "./unzip/text.txt"
	if !IsExist(unZipFile) {
		internal.LogFailedTestInfo(t, "UnZip", zipFile, zipFile, err)
		t.FailNow()
	}
	os.Remove(srcFile)
	os.Remove(zipFile)
	os.RemoveAll(unZipPath)
}

func TestFileMode(t *testing.T) {
	srcFile := "./text.txt"
	CreateFile(srcFile)

	mode, err := FileMode(srcFile)
	if err != nil {
		t.Fail()
	}
	t.Log(mode)

	os.Remove(srcFile)
}

func TestIsLink(t *testing.T) {
	srcFile := "./text.txt"
	CreateFile(srcFile)

	linkFile := "./text.link"
	if !IsExist(linkFile) {
		_ = os.Symlink(srcFile, linkFile)
	}
	if !IsLink(linkFile) {
		internal.LogFailedTestInfo(t, "IsLink", linkFile, "true", "false")
		t.FailNow()
	}

	if IsLink("./file.go") {
		internal.LogFailedTestInfo(t, "IsLink", "./file.go", "false", "true")
		t.FailNow()
	}
	os.Remove(srcFile)
	os.Remove(linkFile)
}

func TestMiMeType(t *testing.T) {
	mt1 := MiMeType("./file.go")
	expected := "text/plain; charset=utf-8"

	if mt1 != expected {
		internal.LogFailedTestInfo(t, "MiMeType", "./file.go", expected, mt1)
		t.FailNow()
	}
	f, _ := os.Open("./file.go")
	mt2 := MiMeType(f)
	if mt2 != expected {
		internal.LogFailedTestInfo(t, "MiMeType", "./file.go", expected, mt2)
		t.FailNow()
	}
}
