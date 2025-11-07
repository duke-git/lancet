package cryptor

import (
	"encoding/hex"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestSm3(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSm3")

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "abc",
			expected: "66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0",
		},
		{
			input:    "abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd",
			expected: "debe9ff92275b8a138604889c18e5a4d6fdb70e5387e5765293dcba39c0c5732",
		},
		{
			input:    "",
			expected: "1ab21d8355cfa17f8e61194831e81a8f22bec8c728fefb747ed035eb5082aa2b",
		},
	}

	for _, tt := range tests {
		result := Sm3([]byte(tt.input))
		resultHex := hex.EncodeToString(result)
		assert.Equal(tt.expected, resultHex)
	}
}

func TestSm4EcbEncryptDecrypt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSm4EcbEncryptDecrypt")

	key := []byte("1234567890abcdef") // 16 bytes
	plaintext := []byte("Hello, SM4!")

	// Encrypt
	encrypted := Sm4EcbEncrypt(plaintext, key)
	assert.IsNotNil(encrypted)

	// Decrypt
	decrypted := Sm4EcbDecrypt(encrypted, key)
	assert.Equal(plaintext, decrypted)
}

func TestSm4CbcEncryptDecrypt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSm4CbcEncryptDecrypt")

	key := []byte("1234567890abcdef") // 16 bytes
	plaintext := []byte("Hello, SM4 CBC mode!")

	// Encrypt
	encrypted := Sm4CbcEncrypt(plaintext, key)
	assert.IsNotNil(encrypted)

	// Decrypt
	decrypted := Sm4CbcDecrypt(encrypted, key)
	assert.Equal(plaintext, decrypted)
}

func TestSm4EcbWithLongData(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSm4EcbWithLongData")

	key := []byte("1234567890abcdef")
	plaintext := []byte("This is a longer message that spans multiple blocks for SM4 encryption testing.")

	encrypted := Sm4EcbEncrypt(plaintext, key)
	decrypted := Sm4EcbDecrypt(encrypted, key)

	assert.Equal(plaintext, decrypted)
}

func TestSm2EncryptDecrypt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSm2EncryptDecrypt")

	// Generate key pair
	privateKey, err := GenerateSm2Key()
	assert.IsNil(err)
	assert.IsNotNil(privateKey)

	plaintext := []byte("Hello, SM2!")

	// Encrypt with public key
	ciphertext, err := Sm2Encrypt(&privateKey.PublicKey, plaintext)
	assert.IsNil(err)
	assert.IsNotNil(ciphertext)

	// Decrypt with private key
	decrypted, err := Sm2Decrypt(privateKey, ciphertext)
	assert.IsNil(err)
	assert.Equal(plaintext, decrypted)
}

func TestSm2WithLongData(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSm2WithLongData")

	privateKey, err := GenerateSm2Key()
	assert.IsNil(err)

	plaintext := []byte("This is a longer message for SM2 encryption testing. " +
		"SM2 is an elliptic curve public key cryptography algorithm.")

	ciphertext, err := Sm2Encrypt(&privateKey.PublicKey, plaintext)
	assert.IsNil(err)

	decrypted, err := Sm2Decrypt(privateKey, ciphertext)
	assert.IsNil(err)
	assert.Equal(plaintext, decrypted)
}

func TestSm4InvalidKeyLength(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSm4InvalidKeyLength")

	defer func() {
		if r := recover(); r != nil {
			assert.IsNotNil(r)
		}
	}()

	key := []byte("short")
	plaintext := []byte("test")
	Sm4EcbEncrypt(plaintext, key) // Should panic
}

func TestSm2InvalidInput(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSm2InvalidInput")

	// Test with nil public key
	_, err := Sm2Encrypt(nil, []byte("test"))
	assert.IsNotNil(err)

	// Test with nil private key
	_, err = Sm2Decrypt(nil, []byte("test"))
	assert.IsNotNil(err)

	// Test with invalid ciphertext
	privateKey, _ := GenerateSm2Key()
	_, err = Sm2Decrypt(privateKey, []byte("short"))
	assert.IsNotNil(err)
}
