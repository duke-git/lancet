package cryptor_test

import (
	"encoding/hex"
	"fmt"

	"github.com/duke-git/lancet/v2/cryptor"
)

func ExampleSm3() {
	data := []byte("hello world")
	hash := cryptor.Sm3(data)

	fmt.Println(hex.EncodeToString(hash))

	// Output:
	// 44f0061e69fa6fdfc290c494654a05dc0c053da7e5c52b84ef93a9d67d3fff88
}

func ExampleSm4EcbEncrypt() {
	key := []byte("1234567890abcdef") // 16 bytes key
	plaintext := []byte("hello world")

	encrypted := cryptor.Sm4EcbEncrypt(plaintext, key)
	decrypted := cryptor.Sm4EcbDecrypt(encrypted, key)

	fmt.Println(string(decrypted))

	// Output:
	// hello world
}

func ExampleSm4CbcEncrypt() {
	key := []byte("1234567890abcdef") // 16 bytes key
	plaintext := []byte("hello world")

	encrypted := cryptor.Sm4CbcEncrypt(plaintext, key)
	decrypted := cryptor.Sm4CbcDecrypt(encrypted, key)

	fmt.Println(string(decrypted))

	// Output:
	// hello world
}

func ExampleGenerateSm2Key() {
	// Generate SM2 key pair
	privateKey, err := cryptor.GenerateSm2Key()
	if err != nil {
		return
	}

	plaintext := []byte("hello world")

	// Encrypt with public key
	ciphertext, err := cryptor.Sm2Encrypt(&privateKey.PublicKey, plaintext)
	if err != nil {
		return
	}

	// Decrypt with private key
	decrypted, err := cryptor.Sm2Decrypt(privateKey, ciphertext)
	if err != nil {
		return
	}

	fmt.Println(string(decrypted))

	// Output:
	// hello world
}
