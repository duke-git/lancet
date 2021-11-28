// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package cryptor implements some util functions to encrypt and decrypt.
package cryptor

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"io"
)

// DesEcbEncrypt encrypt data with key use DES ECB algorithm
// len(key) should be 8
func DesEcbEncrypt(data, key []byte) []byte {
	cipher, _ := des.NewCipher(generateDesKey(key))
	length := (len(data) + des.BlockSize) / des.BlockSize
	plain := make([]byte, length*des.BlockSize)
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

// DesEcbDecrypt decrypt data with key use DES ECB algorithm
// len(key) should be 8
func DesEcbDecrypt(encrypted, key []byte) []byte {
	cipher, _ := des.NewCipher(generateDesKey(key))
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

// DesCbcEncrypt encrypt data with key use DES CBC algorithm
// len(key) should be 8
func DesCbcEncrypt(data, key []byte) []byte {
	block, _ := des.NewCipher(key)
	blockSize := block.BlockSize()
	data = pkcs7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])

	encrypted := make([]byte, len(data))
	blockMode.CryptBlocks(encrypted, data)
	return encrypted
}

// DesCbcDecrypt decrypt data with key use DES CBC algorithm
// len(key) should be 8
func DesCbcDecrypt(encrypted, key []byte) []byte {
	block, _ := des.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])

	decrypted := make([]byte, len(encrypted))
	blockMode.CryptBlocks(decrypted, encrypted)
	decrypted = pkcs7UnPadding(decrypted)
	return decrypted
}

// DesCtrCrypt encrypt data with key use DES CTR algorithm
// len(key) should be 8
func DesCtrCrypt(data, key []byte) []byte {
	block, _ := des.NewCipher(key)

	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	dst := make([]byte, len(data))
	stream.XORKeyStream(dst, data)

	return dst
}

// DesCfbEncrypt encrypt data with key use DES CFB algorithm
// len(key) should be 8
func DesCfbEncrypt(data, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	encrypted := make([]byte, des.BlockSize+len(data))
	iv := encrypted[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[des.BlockSize:], data)
	return encrypted
}

// DesCfbDecrypt decrypt data with key use DES CFB algorithm
// len(encrypted) should be great than 16, len(key) should be 8
func DesCfbDecrypt(encrypted, key []byte) []byte {
	block, _ := des.NewCipher(key)
	if len(encrypted) < des.BlockSize {
		panic("encrypted data is too short")
	}
	iv := encrypted[:des.BlockSize]
	encrypted = encrypted[des.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}

// DesOfbEncrypt encrypt data with key use DES OFB algorithm
// len(key) should be 16, 24 or 32
func DesOfbEncrypt(data, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	data = pkcs7Padding(data, des.BlockSize)
	encrypted := make([]byte, des.BlockSize+len(data))
	iv := encrypted[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(encrypted[des.BlockSize:], data)
	return encrypted
}

// DesOfbDecrypt decrypt data with key use DES OFB algorithm
// len(key) should be 8
func DesOfbDecrypt(data, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := data[:des.BlockSize]
	data = data[des.BlockSize:]
	if len(data)%des.BlockSize != 0 {
		return nil
	}

	decrypted := make([]byte, len(data))
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(decrypted, data)

	decrypted = pkcs7UnPadding(decrypted)
	return decrypted

}
