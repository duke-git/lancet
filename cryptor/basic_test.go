package cryptor

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestBase64StdEncode(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBase64StdEncode")
	assert.Equal("aGVsbG8gd29ybGQ=", Base64StdEncode("hello world"))
}

func TestBase64StdDecode(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBase64StdDecode")
	assert.Equal("hello world", Base64StdDecode("aGVsbG8gd29ybGQ="))
}

func TestMd5String(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMd5String")
	assert.Equal("5d41402abc4b2a76b9719d911017c592", Md5String("hello"))
}

func TestMd5StringWithBase64(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMd5StringWithBase64")
	assert.Equal("XUFAKrxLKna5cZ2REBfFkg==", Md5StringWithBase64("hello"))
}

func TestMd5Byte(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMd5Byte")
	data := []byte{'a'}
	assert.Equal("0cc175b9c0f1b6a831c399e269772661", Md5Byte(data))
}

func TestMd5ByteWithBase64(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMd5ByteWithBase64")
	assert.Equal("XUFAKrxLKna5cZ2REBfFkg==", Md5ByteWithBase64([]byte("hello")))
}

func TestMd5File(t *testing.T) {
	t.Parallel()

	fileMd5, err := Md5File("./basic.go")
	assert := internal.NewAssert(t, "TestMd5File")
	assert.IsNotNil(fileMd5)
	assert.IsNil(err)
}

func TestHmacMd5(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestHmacMd5")
	assert.Equal("5f4c9faaff0a1ad3007d9ddc06abe36d", HmacMd5("hello world", "12345"))
}

func TestHmacMd5WithBase64(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestHmacMd5WithBase64")
	assert.Equal("6DQwbquJLYclJdSRinpjmg==", HmacMd5WithBase64("hello", "12345"))
}

func TestHmacSha1(t *testing.T) {
	t.Parallel()

	s := "hello world"
	key := "12345"
	hmacSha1 := HmacSha1(s, key)
	expected := "3826f812255d8683f051ee97346d1359234d5dbd"

	assert := internal.NewAssert(t, "TestHmacSha1")
	assert.Equal(expected, hmacSha1)
}

func TestHmacSha1WithBase64(t *testing.T) {
	t.Parallel()

	s := "hello"
	key := "12345"
	hmacSha1 := HmacSha1WithBase64(s, key)
	expected := "XGqdsMzLkuNu0DI/0Jt/k23prOA="

	assert := internal.NewAssert(t, "TestHmacSha1")
	assert.Equal(expected, hmacSha1)
}

func TestHmacSha256(t *testing.T) {
	t.Parallel()

	str := "hello world"
	key := "12345"
	hmacSha256 := HmacSha256(str, key)
	expected := "9dce2609f2d67d41f74c7f9efc8ccd44370d41ad2de52982627588dfe7289ab8"

	assert := internal.NewAssert(t, "TestHmacSha256")
	assert.Equal(expected, hmacSha256)
}

func TestHmacSha256WithBase64(t *testing.T) {
	t.Parallel()

	str := "hello"
	key := "12345"
	hms := HmacSha256WithBase64(str, key)
	expected := "MVu5PE6YmGK6Ccti4F1zpfN2yzbw14btqwwyDQWf3nU="

	assert := internal.NewAssert(t, "TestHmacSha256WithBase64")
	assert.Equal(expected, hms)
}

func TestHmacSha512(t *testing.T) {
	t.Parallel()

	s := "hello world"
	key := "12345"
	hmacSha512 := HmacSha512(s, key)
	expected := "5b1563ac4e9b49c9ada8ccb232588fc4f0c30fd12f756b3a0b95af4985c236ca60925253bae10ce2c6bf9af1c1679b51e5395ff3d2826c0a2c7c0d72225d4175"

	assert := internal.NewAssert(t, "TestHmacSha512")
	assert.Equal(expected, hmacSha512)
}

func TestHmacSha512WithBase64(t *testing.T) {
	t.Parallel()

	str := "hello"
	key := "12345"
	hms := HmacSha512WithBase64(str, key)
	expected := "3Y8SkKndI9NU4lJtmi6c6M///dN8syCADRxsE9Lvw2Mog3ahlsVFja9T+OGqa0Wm2FYwPVwKIGS/+XhYYdSM/A=="

	assert := internal.NewAssert(t, "TestHmacSha512WithBase64")
	assert.Equal(expected, hms)
}

func TestSha1(t *testing.T) {
	t.Parallel()

	s := "hello world"
	sha1 := Sha1(s)
	expected := "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed"

	assert := internal.NewAssert(t, "TestSha1")
	assert.Equal(expected, sha1)
}

func TestSha1WithBase64(t *testing.T) {
	t.Parallel()

	str := Sha1WithBase64("hello")
	expected := "qvTGHdzF6KLavt4PO0gs2a6pQ00="

	assert := internal.NewAssert(t, "TestSha1WithBase64")
	assert.Equal(expected, str)
}

func TestSha256(t *testing.T) {
	t.Parallel()

	s := "hello world"
	sha256 := Sha256(s)
	expected := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"

	assert := internal.NewAssert(t, "TestSha256")
	assert.Equal(expected, sha256)
}

func TestSha256WithBase64(t *testing.T) {
	t.Parallel()

	str := Sha256WithBase64("hello")
	expected := "LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ="

	assert := internal.NewAssert(t, "TestSha256WithBase64")
	assert.Equal(expected, str)
}

func TestSha512(t *testing.T) {
	t.Parallel()

	s := "hello world"
	sha512 := Sha512(s)
	expected := "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"

	assert := internal.NewAssert(t, "TestSha512")
	assert.Equal(expected, sha512)
}

func TestSha512WithBase64(t *testing.T) {
	t.Parallel()

	str := Sha512WithBase64("hello")
	expected := "m3HSJL1i83hdltRq0+o9czGb+8KJDKra4t/3JRlnPKcjI8PZm6XBHXx6zG4UuMXaDEZjR1wuXDre9G9zvN7AQw=="

	assert := internal.NewAssert(t, "TestSha512WithBase64")
	assert.Equal(expected, str)
}
