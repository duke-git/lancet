package cryptor

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestBase64StdEncode(t *testing.T) {
	assert := internal.NewAssert(t, "TestBase64StdEncode")
	assert.Equal("aGVsbG8gd29ybGQ=", Base64StdEncode("hello world"))
}

func TestBase64StdDecode(t *testing.T) {
	assert := internal.NewAssert(t, "TestBase64StdDecode")
	assert.Equal("hello world", Base64StdDecode("aGVsbG8gd29ybGQ="))
}

func TestMd5String(t *testing.T) {
	assert := internal.NewAssert(t, "TestMd5String")
	assert.Equal("5d41402abc4b2a76b9719d911017c592", Md5String("hello"))
}

func TestMd5File(t *testing.T) {
	fileMd5, err := Md5File("./basic.go")
	assert := internal.NewAssert(t, "TestMd5File")
	assert.IsNotNil(fileMd5)
	assert.IsNil(err)
}

func TestHmacMd5(t *testing.T) {
	assert := internal.NewAssert(t, "TestHmacMd5")
	assert.Equal("5f4c9faaff0a1ad3007d9ddc06abe36d", HmacMd5("hello world", "12345"))
}

func TestHmacSha1(t *testing.T) {
	s := "hello world"
	key := "12345"
	hmacSha1 := HmacSha1(s, key)
	expected := "3826f812255d8683f051ee97346d1359234d5dbd"

	assert := internal.NewAssert(t, "TestHmacSha1")
	assert.Equal(expected, hmacSha1)
}

func TestHmacSha256(t *testing.T) {
	s := "hello world"
	key := "12345"
	hmacSha256 := HmacSha256(s, key)
	expected := "9dce2609f2d67d41f74c7f9efc8ccd44370d41ad2de52982627588dfe7289ab8"

	assert := internal.NewAssert(t, "TestHmacSha256")
	assert.Equal(expected, hmacSha256)
}

func TestHmacSha512(t *testing.T) {
	s := "hello world"
	key := "12345"
	hmacSha512 := HmacSha512(s, key)
	expected := "5b1563ac4e9b49c9ada8ccb232588fc4f0c30fd12f756b3a0b95af4985c236ca60925253bae10ce2c6bf9af1c1679b51e5395ff3d2826c0a2c7c0d72225d4175"

	assert := internal.NewAssert(t, "TestHmacSha512")
	assert.Equal(expected, hmacSha512)
}

func TestSha1(t *testing.T) {
	s := "hello world"
	sha1 := Sha1(s)
	expected := "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"

	assert := internal.NewAssert(t, "TestSha1")
	assert.Equal(expected, sha1)
}

func TestSha256(t *testing.T) {
	s := "hello world"
	sha256 := Sha256(s)
	expected := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"

	assert := internal.NewAssert(t, "TestSha256")
	assert.Equal(expected, sha256)
}

func TestSha512(t *testing.T) {
	s := "hello world"
	sha512 := Sha512(s)
	expected := "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"

	assert := internal.NewAssert(t, "TestSha512")
	assert.Equal(expected, sha512)
}
