package cryptor

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestDesEcbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefgh"

	desEcbEncrypt := DesEcbEncrypt([]byte(data), []byte(key))
	desEcbDecrypt := DesEcbDecrypt(desEcbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestDesEcbEncrypt")
	assert.Equal(data, string(desEcbDecrypt))
}

func TestDesCbcEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefgh"

	desCbcEncrypt := DesCbcEncrypt([]byte(data), []byte(key))
	desCbcDecrypt := DesCbcDecrypt(desCbcEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestDesCbcEncrypt")
	assert.Equal(data, string(desCbcDecrypt))
}

func TestDesCtrCrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefgh"

	desCtrCrypt := DesCtrCrypt([]byte(data), []byte(key))
	desCtrDeCrypt := DesCtrCrypt(desCtrCrypt, []byte(key))

	assert := internal.NewAssert(t, "TestDesCtrCrypt")
	assert.Equal(data, string(desCtrDeCrypt))
}

func TestDesCfbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefgh"

	desCfbEncrypt := DesCfbEncrypt([]byte(data), []byte(key))
	desCfbDecrypt := DesCfbDecrypt(desCfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestDesCfbEncrypt")
	assert.Equal(data, string(desCfbDecrypt))
}

func TestDesOfbEncrypt(t *testing.T) {
	data := "hello world"
	key := "abcdefgh"

	desOfbEncrypt := DesOfbEncrypt([]byte(data), []byte(key))
	desOfbDecrypt := DesOfbDecrypt(desOfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestDesOfbEncrypt")
	assert.Equal(data, string(desOfbDecrypt))
}
