package cryptor

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestAesEcbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesEcbEncrypt := AesEcbEncrypt([]byte(data), []byte(key))
	aesEcbDecrypt := AesEcbDecrypt(aesEcbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesEcbEncrypt")
	assert.Equal(data, string(aesEcbDecrypt))
}

func TestAesCbcEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesCbcEncrypt := AesCbcEncrypt([]byte(data), []byte(key))
	aesCbcDecrypt := AesCbcDecrypt(aesCbcEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesCbcEncrypt")
	assert.Equal(data, string(aesCbcDecrypt))
}

func TestAesCtrCrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesCtrCrypt := AesCtrCrypt([]byte(data), []byte(key))
	aesCtrDeCrypt := AesCtrCrypt(aesCtrCrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesCtrCrypt")
	assert.Equal(data, string(aesCtrDeCrypt))
}

func TestAesCfbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesCfbEncrypt := AesCfbEncrypt([]byte(data), []byte(key))
	aesCfbDecrypt := AesCfbDecrypt(aesCfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesCfbEncrypt")
	assert.Equal(data, string(aesCfbDecrypt))
}

func TestAesOfbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesOfbEncrypt := AesOfbEncrypt([]byte(data), []byte(key))
	aesOfbDecrypt := AesOfbDecrypt(aesOfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesOfbEncrypt")
	assert.Equal(data, string(aesOfbDecrypt))
}
