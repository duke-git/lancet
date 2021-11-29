// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package cryptor implements some util functions to encrypt and decrypt.
// Note:
// 1. for aes crypt function, the `key` param length should be 16, 24 or 32. if not, will panic.
package cryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// AesEcbEncrypt encrypt data with key use AES ECB algorithm
// len(key) should be 16, 24 or 32
func AesEcbEncrypt(data, key []byte) []byte {
	cipher, _ := aes.NewCipher(generateAesKey(key))
	length := (len(data) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, data)
	pad := byte(len(plain) - len(data))
	for i := len(data); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, cipher.BlockSize(); bs <= len(data); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

// AesEcbDecrypt decrypt data with key use AES ECB algorithm
// len(key) should be 16, 24 or 32
func AesEcbDecrypt(encrypted, key []byte) []byte {
	cipher, _ := aes.NewCipher(generateAesKey(key))
	decrypted := make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

// AesCbcEncrypt encrypt data with key use AES CBC algorithm
// len(key) should be 16, 24 or 32
func AesCbcEncrypt(data, key []byte) []byte {
	// len(key) should be 16, 24 or 32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	data = pkcs7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])

	encrypted := make([]byte, len(data))
	blockMode.CryptBlocks(encrypted, data)
	return encrypted
}

// AesCbcDecrypt decrypt data with key use AES CBC algorithm
// len(key) should be 16, 24 or 32
func AesCbcDecrypt(encrypted, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])

	decrypted := make([]byte, len(encrypted))
	blockMode.CryptBlocks(decrypted, encrypted)
	decrypted = pkcs7UnPadding(decrypted)
	return decrypted
}

// AesCtrCrypt encrypt data with key use AES CTR algorithm
// len(key) should be 16, 24 or 32
func AesCtrCrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)

	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	dst := make([]byte, len(data))
	stream.XORKeyStream(dst, data)

	return dst
}

// AesCfbEncrypt encrypt data with key use AES CFB algorithm
// len(key) should be 16, 24 or 32
func AesCfbEncrypt(data, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	encrypted := make([]byte, aes.BlockSize+len(data))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], data)
	return encrypted
}

// AesCfbDecrypt decrypt data with key use AES CFB algorithm
// len(encrypted) should be great than 16, len(key) should be 16, 24 or 32
func AesCfbDecrypt(encrypted, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("encrypted data is too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}

// AesOfbEncrypt encrypt data with key use AES OFB algorithm
// len(key) should be 16, 24 or 32
func AesOfbEncrypt(data, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	data = pkcs7Padding(data, aes.BlockSize)
	encrypted := make([]byte, aes.BlockSize+len(data))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], data)
	return encrypted
}

// AesOfbDecrypt decrypt data with key use AES OFB algorithm
// len(key) should be 16, 24 or 32
func AesOfbDecrypt(data, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	if len(data)%aes.BlockSize != 0 {
		return nil
	}

	decrypted := make([]byte, len(data))
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(decrypted, data)

	decrypted = pkcs7UnPadding(decrypted)
	return decrypted
}
