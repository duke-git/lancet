package cryptor

import (
	"crypto"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestAesEcbEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefghijklmnop"

	aesEcbEncrypt := AesEcbEncrypt([]byte(data), []byte(key))
	aesEcbDecrypt := AesEcbDecrypt(aesEcbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesEcbEncrypt")
	assert.Equal(data, string(aesEcbDecrypt))
}

func TestAesCbcEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefghijklmnop"

	aesCbcEncrypt := AesCbcEncrypt([]byte(data), []byte(key))
	aesCbcDecrypt := AesCbcDecrypt(aesCbcEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesCbcEncrypt")
	assert.Equal(data, string(aesCbcDecrypt))
}

func TestAesCtrCrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefghijklmnop"

	aesCtrCrypt := AesCtrCrypt([]byte(data), []byte(key))
	aesCtrDeCrypt := AesCtrCrypt(aesCtrCrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesCtrCrypt")
	assert.Equal(data, string(aesCtrDeCrypt))
}

func TestAesCfbEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefghijklmnop"

	aesCfbEncrypt := AesCfbEncrypt([]byte(data), []byte(key))
	aesCfbDecrypt := AesCfbDecrypt(aesCfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesCfbEncrypt")
	assert.Equal(data, string(aesCfbDecrypt))
}

func TestAesOfbEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefghijklmnop"

	aesOfbEncrypt := AesOfbEncrypt([]byte(data), []byte(key))
	aesOfbDecrypt := AesOfbDecrypt(aesOfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesOfbEncrypt")
	assert.Equal(data, string(aesOfbDecrypt))
}

func TestDesEcbEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefgh"

	desEcbEncrypt := DesEcbEncrypt([]byte(data), []byte(key))
	desEcbDecrypt := DesEcbDecrypt(desEcbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestDesEcbEncrypt")
	assert.Equal(data, string(desEcbDecrypt))
}

func TestDesCbcEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefgh"

	desCbcEncrypt := DesCbcEncrypt([]byte(data), []byte(key))
	desCbcDecrypt := DesCbcDecrypt(desCbcEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestDesCbcEncrypt")
	assert.Equal(data, string(desCbcDecrypt))
}

func TestDesCtrCrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefgh"

	desCtrCrypt := DesCtrCrypt([]byte(data), []byte(key))
	desCtrDeCrypt := DesCtrCrypt(desCtrCrypt, []byte(key))

	assert := internal.NewAssert(t, "TestDesCtrCrypt")
	assert.Equal(data, string(desCtrDeCrypt))
}

func TestDesCfbEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefgh"

	desCfbEncrypt := DesCfbEncrypt([]byte(data), []byte(key))
	desCfbDecrypt := DesCfbDecrypt(desCfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestDesCfbEncrypt")
	assert.Equal(data, string(desCfbDecrypt))
}

func TestDesOfbEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefgh"

	desOfbEncrypt := DesOfbEncrypt([]byte(data), []byte(key))
	desOfbDecrypt := DesOfbDecrypt(desOfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestDesOfbEncrypt")
	assert.Equal(data, string(desOfbDecrypt))
}

func TestRsaEncrypt(t *testing.T) {
	t.Parallel()

	err := GenerateRsaKey(4096, "./rsa_private_example.pem", "./rsa_public_example.pem")
	if err != nil {
		t.FailNow()
	}
	data := []byte("hello world")
	encrypted := RsaEncrypt(data, "./rsa_public_example.pem")
	decrypted := RsaDecrypt(encrypted, "./rsa_private_example.pem")

	assert := internal.NewAssert(t, "TestRsaEncrypt")
	assert.Equal(string(data), string(decrypted))
}

func TestRsaEncryptOAEP(t *testing.T) {
	assert := internal.NewAssert(t, "TestRsaEncrypt")
	t.Parallel()

	pri, pub := GenerateRsaKeyPair(1024)

	data := []byte("hello world")
	label := []byte("123456")

	encrypted, err := RsaEncryptOAEP(data, label, *pub)
	assert.IsNil(err)

	decrypted, err := RsaDecryptOAEP([]byte(encrypted), label, *pri)

	assert.IsNil(err)
	assert.Equal("hello world", string(decrypted))
}

func TestAesGcmEncrypt(t *testing.T) {
	t.Parallel()

	data := "hello world"
	key := "abcdefghijklmnop"

	encrypted := AesGcmEncrypt([]byte(data), []byte(key))
	decrypted := AesGcmDecrypt(encrypted, []byte(key))

	assert := internal.NewAssert(t, "TestAesGcmEncrypt")
	assert.Equal(data, string(decrypted))
}

func TestRsaSignAndVerify(t *testing.T) {
	t.Parallel()

	data := []byte("This is a test data for RSA signing")
	hash := crypto.SHA256

	t.Run("RSA Sign and Verify", func(t *testing.T) {
		privateKey := "./rsa_private_example.pem"
		publicKey := "./rsa_public_example.pem"

		signature, err := RsaSign(hash, data, privateKey)
		if err != nil {
			t.Fatalf("RsaSign failed: %v", err)
		}

		err = RsaVerifySign(hash, data, signature, publicKey)
		if err != nil {
			t.Fatalf("RsaVerifySign failed: %v", err)
		}
	})

	t.Run("RSA Sign and Verify Invalid Signature", func(t *testing.T) {
		publicKey := "./rsa_public_example.pem"

		invalidSig := []byte("InvalidSignature")

		err := RsaVerifySign(hash, data, invalidSig, publicKey)
		if err == nil {
			t.Fatalf("RsaVerifySign failed: %v", err)
		}
	})

	t.Run("RSA Sign and Verify With Different Hash", func(t *testing.T) {
		publicKey := "./rsa_public_example.pem"
		privateKey := "./rsa_private_example.pem"
		hashSign := crypto.SHA256
		hashVerify := crypto.SHA512

		signature, err := RsaSign(hashSign, data, privateKey)
		if err != nil {
			t.Fatalf("RsaSign failed: %v", err)
		}

		err = RsaVerifySign(hashVerify, data, signature, publicKey)
		if err == nil {
			t.Fatalf("RsaVerifySign failed: %v", err)
		}
	})
}
