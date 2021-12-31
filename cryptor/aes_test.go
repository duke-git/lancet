package cryptor

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestAesEcbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesEcbEncrypt := AesEcbEncrypt([]byte(data), []byte(key))
	aesEcbDecrypt := AesEcbDecrypt(aesEcbEncrypt, []byte(key))

	if string(aesEcbDecrypt) != data {
		internal.LogFailedTestInfo(t, "AesEcbEncrypt/AesEcbDecrypt", data, data, string(aesEcbDecrypt))
		t.FailNow()
	}
}

func TestAesCbcEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesCbcEncrypt := AesCbcEncrypt([]byte(data), []byte(key))
	aesCbcDecrypt := AesCbcDecrypt(aesCbcEncrypt, []byte(key))

	if string(aesCbcDecrypt) != data {
		internal.LogFailedTestInfo(t, "AesCbcEncrypt/AesCbcDecrypt", data, data, string(aesCbcDecrypt))
		t.FailNow()
	}
}

func TestAesCtrCrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesCtrCrypt := AesCtrCrypt([]byte(data), []byte(key))
	aesCtrDeCrypt := AesCtrCrypt(aesCtrCrypt, []byte(key))

	if string(aesCtrDeCrypt) != data {
		internal.LogFailedTestInfo(t, "AesCtrCrypt", data, data, string(aesCtrDeCrypt))
		t.FailNow()
	}
}

func TestAesCfbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesCfbEncrypt := AesCfbEncrypt([]byte(data), []byte(key))
	aesCfbDecrypt := AesCfbDecrypt(aesCfbEncrypt, []byte(key))

	if string(aesCfbDecrypt) != data {
		internal.LogFailedTestInfo(t, "AesCfbEncrypt/AesCfbDecrypt", data, data, string(aesCfbDecrypt))
		t.FailNow()
	}
}

func TestAesOfbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefghijklmnop"

	aesOfbEncrypt := AesOfbEncrypt([]byte(data), []byte(key))
	aesOfbDecrypt := AesOfbDecrypt(aesOfbEncrypt, []byte(key))

	if string(aesOfbDecrypt) != data {
		internal.LogFailedTestInfo(t, "AesOfbEncrypt/AesOfbDecrypt", data, data, string(aesOfbDecrypt))
		t.FailNow()
	}
}
