package fileutil

import (
	"os"
	"reflect"
	"testing"

	"github.com/duke-git/lancet/utils"
)

func TestIsExist(t *testing.T) {
	cases := []string{"./", "./a.txt"}
	expected := []bool{true, false}

	for i := 0; i < len(cases); i++ {
		res := IsExist(cases[i])
		if res != expected[i] {
			utils.LogFailedTestInfo(t, "IsExist", cases[i], expected[i], res)
			t.FailNow()
		}
	}
}

func TestCreateFile(t *testing.T) {
	f := "./text.txt"
	if CreateFile(f) {
		file, err := os.Open(f)
		if err != nil {
			utils.LogFailedTestInfo(t, "CreateFile", f, f, "create file error: "+err.Error())
			t.FailNow()
		}
		if file.Name() != f {
			utils.LogFailedTestInfo(t, "CreateFile", f, f, file.Name())
			t.FailNow()
		}
	} else {
		utils.LogFailedTestInfo(t, "CreateFile", f, f, "create file error")
		t.FailNow()
	}
}

func TestIsDir(t *testing.T) {
	cases := []string{"./", "./a.txt"}
	expected := []bool{true, false}

	for i := 0; i < len(cases); i++ {
		res := IsDir(cases[i])
		if res != expected[i] {
			utils.LogFailedTestInfo(t, "IsDir", cases[i], expected[i], res)
			t.FailNow()
		}
	}
}

func TestRemoveFile(t *testing.T) {
	f := "./text.txt"
	if CreateFile(f) {
		err := RemoveFile(f)
		if err != nil {
			utils.LogFailedTestInfo(t, "RemoveFile", f, f, err.Error())
			t.FailNow()
		}
	} else {
		utils.LogFailedTestInfo(t, "RemoveFile", f, f, "create file error")
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
			utils.LogFailedTestInfo(t, "CopyFile", srcFile, dstFile, "create file error: "+err.Error())
			t.FailNow()
		}
		if file.Name() != dstFile {
			utils.LogFailedTestInfo(t, "CopyFile", srcFile, dstFile, file.Name())
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
		utils.LogFailedTestInfo(t, "ToChar", "./", expected, filesInCurrentPath)
		t.FailNow()
	}

}
