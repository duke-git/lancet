package cryptor

import (
	"fmt"
	"os"
	"testing"

	"github.com/duke-git/lancet/utils"
)

func TestBase64StdEncode(t *testing.T) {
	s := "hello world"
	bs := Base64StdEncode(s)

	if bs != "aGVsbG8gd29ybGQ=" {
		utils.LogFailedTestInfo(t, "Base64StdEncode", s, "aGVsbG8gd29ybGQ=", bs)
		t.FailNow()
	}
}

func TestBase64StdDecode(t *testing.T) {
	bs := "aGVsbG8gd29ybGQ="
	s := Base64StdDecode(bs)

	if s != "hello world" {
		utils.LogFailedTestInfo(t, "Base64StdDecode", bs, "hello world=", s)
		t.FailNow()
	}
}

func TestMd5String(t *testing.T) {
	s := "hello"
	smd5 := Md5String(s)
	expected := "5d41402abc4b2a76b9719d911017c592"

	if smd5 != expected {
		utils.LogFailedTestInfo(t, "Md5String", s, expected, smd5)
		t.FailNow()
	}
}

func TestMd5File(t *testing.T) {
	file, _ := os.Create("./hello.txt")
	defer file.Close()
	file.WriteString("hello\n")

	fileMd5, err := Md5File("./hello.txt")
	if err != nil {
		t.FailNow()
	}
	fmt.Println(fileMd5)
}

func TestHmacMd5(t *testing.T) {
	s := "hello world"
	key := "12345"
	hmacMd5 := HmacMd5(s, key)
	expected := "5f4c9faaff0a1ad3007d9ddc06abe36d"

	if hmacMd5 != expected {
		utils.LogFailedTestInfo(t, "HmacMd5", s, expected, hmacMd5)
		t.FailNow()
	}
}

func TestHmacSha1(t *testing.T) {
	s := "hello world"
	key := "12345"
	hmacSha1 := HmacSha1(s, key)
	expected := "3826f812255d8683f051ee97346d1359234d5dbd"

	if hmacSha1 != expected {
		utils.LogFailedTestInfo(t, "HmacSha1", s, expected, hmacSha1)
		t.FailNow()
	}
}

func TestHmacSha256(t *testing.T) {
	s := "hello world"
	key := "12345"
	hmacSha256 := HmacSha256(s, key)
	expected := "9dce2609f2d67d41f74c7f9efc8ccd44370d41ad2de52982627588dfe7289ab8"

	if hmacSha256 != expected {
		utils.LogFailedTestInfo(t, "HmacSha256", s, expected, hmacSha256)
		t.FailNow()
	}
}

func TestHmacSha512(t *testing.T) {
	s := "hello world"
	key := "12345"
	hmacSha512 := HmacSha512(s, key)
	expected := "5b1563ac4e9b49c9ada8ccb232588fc4f0c30fd12f756b3a0b95af4985c236ca60925253bae10ce2c6bf9af1c1679b51e5395ff3d2826c0a2c7c0d72225d4175"

	if hmacSha512 != expected {
		utils.LogFailedTestInfo(t, "HmacSha512", s, expected, hmacSha512)
		t.FailNow()
	}
}

func TestSha1(t *testing.T) {
	s := "hello world"
	sha1 := Sha1(s)
	expected := "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"

	if sha1 != expected {
		utils.LogFailedTestInfo(t, "Sha1", s, expected, sha1)
		t.FailNow()
	}
}

func TestSha256(t *testing.T) {
	s := "hello world"
	sha256 := Sha256(s)
	expected := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"

	if sha256 != expected {
		utils.LogFailedTestInfo(t, "Sha256", s, expected, sha256)
		t.FailNow()
	}
}

func TestSha512(t *testing.T) {
	s := "hello world"
	sha512 := Sha512(s)
	expected := "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"

	if sha512 != expected {
		utils.LogFailedTestInfo(t, "Sha512", s, expected, sha512)
		t.FailNow()
	}
}
