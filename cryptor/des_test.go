package cryptor

import (
	"testing"

	"github.com/duke-git/lancet/utils"
)

func TestDesEcbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefgh"

	desEcbEncrypt := DesEcbEncrypt([]byte(data), []byte(key))
	desEcbDecrypt := DesEcbDecrypt(desEcbEncrypt, []byte(key))

	if string(desEcbDecrypt) != data {
		utils.LogFailedTestInfo(t, "DesEcbEncrypt/DesEcbDecrypt", data, data, string(desEcbDecrypt))
		t.FailNow()
	}
}

func TestDesCbcEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefgh"

	desCbcEncrypt := DesCbcEncrypt([]byte(data), []byte(key))
	desCbcDecrypt := DesCbcDecrypt(desCbcEncrypt, []byte(key))

	if string(desCbcDecrypt) != data {
		utils.LogFailedTestInfo(t, "DesCbcEncrypt/DesCbcDecrypt", data, data, string(desCbcDecrypt))
		t.FailNow()
	}
}

func TestDesCtrCrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefgh"

	desCtrCrypt := DesCtrCrypt([]byte(data), []byte(key))
	desCtrDeCrypt := DesCtrCrypt(desCtrCrypt, []byte(key))

	if string(desCtrDeCrypt) != data {
		utils.LogFailedTestInfo(t, "DesCtrCrypt", data, data, string(desCtrDeCrypt))
		t.FailNow()
	}
}

func TestDesCfbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefgh"

	desCfbEncrypt := DesCfbEncrypt([]byte(data), []byte(key))
	desCfbDecrypt := DesCfbDecrypt(desCfbEncrypt, []byte(key))

	if string(desCfbDecrypt) != data {
		utils.LogFailedTestInfo(t, "DesCfbEncrypt/DesCfbDecrypt", data, data, string(desCfbDecrypt))
		t.FailNow()
	}
}

func TestDesOfbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefgh"

	desOfbEncrypt := DesOfbEncrypt([]byte(data), []byte(key))
	desOfbDecrypt := DesOfbDecrypt(desOfbEncrypt, []byte(key))

	if string(desOfbDecrypt) != data {
		utils.LogFailedTestInfo(t, "DesOfbEncrypt/DesOfbDecrypt", data, data, string(desOfbDecrypt))
		t.FailNow()
	}
}
