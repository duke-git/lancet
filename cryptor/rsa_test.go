package cryptor

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestRsaEncrypt(t *testing.T) {
	err := GenerateRsaKey(4096, "rsa_private.pem", "rsa_public.pem")
	if err != nil {
		t.FailNow()
	}
	data := []byte("hello world")
	encrypted := RsaEncrypt(data, "rsa_public.pem")
	decrypted := RsaDecrypt(encrypted, "rsa_private.pem")

	assert := internal.NewAssert(t, "TestRsaEncrypt")
	assert.Equal(string(data), string(decrypted))
}
