// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package cryptor implements some util functions to encrypt and decrypt.
// Note:
// 1. for aes crypt function, the `key` param length should be 16, 24 or 32. if not, will panic.
package cryptor

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io"
	"os"
)

// AesEcbEncrypt encrypt data with key use AES ECB algorithm
// len(key) should be 16, 24 or 32.
// Play: https://go.dev/play/p/jT5irszHx-j
func AesEcbEncrypt(data, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	blockSize := aes.BlockSize
	dataLen := len(data)
	padding := blockSize - (dataLen % blockSize)
	paddedLen := dataLen + padding

	paddedData := make([]byte, paddedLen)
	copy(paddedData, data)

	for i := dataLen; i < paddedLen; i++ {
		paddedData[i] = byte(padding)
	}

	cipher, err := aes.NewCipher(generateAesKey(key, len(key)))
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	encrypted := make([]byte, paddedLen)
	for bs := 0; bs < paddedLen; bs += blockSize {
		cipher.Encrypt(encrypted[bs:], paddedData[bs:])
	}

	return encrypted
}

// AesEcbDecrypt decrypt data with key use AES ECB algorithm
// len(key) should be 16, 24 or 32.
// Play: https://go.dev/play/p/jT5irszHx-j
func AesEcbDecrypt(encrypted, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	blockSize := aes.BlockSize
	if len(encrypted)%blockSize != 0 {
		panic("aes: encrypted data length is not a multiple of block size")
	}

	cipher, err := aes.NewCipher(generateAesKey(key, len(key)))
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	decrypted := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i += blockSize {
		cipher.Decrypt(decrypted[i:], encrypted[i:])
	}

	if len(decrypted) == 0 {
		return nil
	}
	padding := int(decrypted[len(decrypted)-1])
	if padding == 0 || padding > blockSize {
		panic("aes: invalid PKCS#7 padding")
	}
	for i := len(decrypted) - padding; i < len(decrypted); i++ {
		if decrypted[i] != byte(padding) {
			panic("aes: invalid PKCS#7 padding content")
		}
	}

	return decrypted[:len(decrypted)-padding]
}

// AesCbcEncrypt encrypt data with key use AES CBC algorithm
// len(key) should be 16, 24 or 32.
// Play: https://go.dev/play/p/IOq_g8_lKZD
func AesCbcEncrypt(data, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	padding := aes.BlockSize - len(data)%aes.BlockSize
	padded := append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("aes: failed to generate IV: " + err.Error())
	}

	encrypted := make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, padded)

	return append(iv, encrypted...)
}

// AesCbcDecrypt decrypt data with key use AES CBC algorithm
// len(key) should be 16, 24 or 32.
// Play: https://go.dev/play/p/IOq_g8_lKZD
func AesCbcDecrypt(encrypted, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	if len(encrypted) < aes.BlockSize {
		panic("aes: ciphertext too short")
	}

	if len(encrypted)%aes.BlockSize != 0 {
		panic("aes: ciphertext is not a multiple of the block size")
	}

	iv := encrypted[:aes.BlockSize]
	ciphertext := encrypted[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	decrypted := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, ciphertext)

	return pkcs7UnPadding(decrypted)
}

// AesCtrCrypt encrypt data with key use AES CTR algorithm
// len(key) should be 16, 24 or 32.
// Play: https://go.dev/play/p/SpaZO0-5Nsp
// deprecated: use AesCtrEncrypt and AesCtrDecrypt instead.
func AesCtrCrypt(data, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	block, _ := aes.NewCipher(key)

	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	dst := make([]byte, len(data))
	stream.XORKeyStream(dst, data)

	return dst
}

// AesCtrEncrypt encrypt data with key use AES CTR algorithm
// len(key) should be 16, 24 or 32.
// Play: todo
func AesCtrEncrypt(data, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("aes: failed to generate IV: " + err.Error())
	}

	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(data))
	stream.XORKeyStream(ciphertext, data)

	return append(iv, ciphertext...)
}

// AesCtrDecrypt decrypt data with key use AES CTR algorithm
// len(key) should be 16, 24 or 32.
// Play: todo
func AesCtrDecrypt(encrypted, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}
	if len(encrypted) < aes.BlockSize {
		panic("aes: invalid ciphertext length")
	}

	iv := encrypted[:aes.BlockSize]
	ciphertext := encrypted[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext
}

// AesCfbEncrypt encrypt data with key use AES CFB algorithm
// len(key) should be 16, 24 or 32.
// Play: https://go.dev/play/p/tfkF10B13kH
func AesCfbEncrypt(data, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("aes: failed to generate IV: " + err.Error())
	}

	ciphertext := make([]byte, len(data))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, data)

	return append(iv, ciphertext...)
}

// AesCfbDecrypt decrypt data with key use AES CFB algorithm
// len(encrypted) should be great than 16, len(key) should be 16, 24 or 32.
// Play: https://go.dev/play/p/tfkF10B13kH
func AesCfbDecrypt(encrypted, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	if len(encrypted) < aes.BlockSize {
		panic("aes: encrypted data too short")
	}

	iv := encrypted[:aes.BlockSize]
	ciphertext := encrypted[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext
}

// AesOfbEncrypt encrypt data with key use AES OFB algorithm
// len(key) should be 16, 24 or 32.
// Play: https://go.dev/play/p/VtHxtkUj-3F
func AesOfbEncrypt(data, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("aes: failed to generate IV: " + err.Error())
	}

	ciphertext := make([]byte, len(data))
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertext, data)

	return append(iv, ciphertext...)
}

// AesOfbDecrypt decrypt data with key use AES OFB algorithm
// len(key) should be 16, 24 or 32.
// Play: https://go.dev/play/p/VtHxtkUj-3F
func AesOfbDecrypt(data, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	if len(data) < aes.BlockSize {
		panic("aes: encrypted data too short")
	}

	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext
}

// AesGcmEncrypt encrypt data with key use AES GCM algorithm
// Play: https://go.dev/play/p/rUt0-DmsPCs
func AesGcmEncrypt(data, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic("aes: failed to create GCM: " + err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic("aes: failed to generate nonce: " + err.Error())
	}

	ciphertext := gcm.Seal(nil, nonce, data, nil)

	return append(nonce, ciphertext...)
}

// AesGcmDecrypt decrypt data with key use AES GCM algorithm
// Play: https://go.dev/play/p/rUt0-DmsPCs
func AesGcmDecrypt(data, key []byte) []byte {
	if !isAesKeyLengthValid(len(key)) {
		panic("aes: invalid key length (must be 16, 24, or 32 bytes)")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic("aes: failed to create cipher: " + err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic("aes: failed to create GCM: " + err.Error())
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		panic("aes: ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic("aes: decryption failed: " + err.Error())
	}

	return plaintext
}

// DesEcbEncrypt encrypt data with key use DES ECB algorithm
// len(key) should be 8.
// Play: https://go.dev/play/p/8qivmPeZy4P
func DesEcbEncrypt(data, key []byte) []byte {
	cipher, err := des.NewCipher(generateDesKey(key))
	if err != nil {
		panic("des: failed to create cipher: " + err.Error())
	}

	blockSize := cipher.BlockSize()
	padded := pkcs5Padding(data, blockSize)
	encrypted := make([]byte, len(padded))

	for i := 0; i < len(padded); i += blockSize {
		cipher.Encrypt(encrypted[i:], padded[i:])
	}

	return encrypted
}

// DesEcbDecrypt decrypt data with key use DES ECB algorithm
// len(key) should be 8.
// Play: https://go.dev/play/p/8qivmPeZy4P
func DesEcbDecrypt(encrypted, key []byte) []byte {
	cipher, err := des.NewCipher(generateDesKey(key))
	if err != nil {
		panic("des: failed to create cipher: " + err.Error())
	}

	blockSize := cipher.BlockSize()
	if len(encrypted)%blockSize != 0 {
		panic("des: invalid encrypted data length")
	}

	decrypted := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i += blockSize {
		cipher.Decrypt(decrypted[i:], encrypted[i:])
	}

	// Remove padding
	return pkcs5UnPadding(decrypted)
}

// DesCbcEncrypt encrypt data with key use DES CBC algorithm
// len(key) should be 8.
// Play: https://go.dev/play/p/4cC4QvWfe3_1
func DesCbcEncrypt(data, key []byte) []byte {
	if len(key) != 8 {
		panic("des: key length must be 8 bytes")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		panic("des: failed to create cipher: " + err.Error())
	}

	blockSize := block.BlockSize()
	data = pkcs7Padding(data, blockSize)

	encrypted := make([]byte, blockSize+len(data))
	iv := encrypted[:blockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("des: failed to generate IV: " + err.Error())
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted[blockSize:], data)

	return encrypted
}

// DesCbcDecrypt decrypt data with key use DES CBC algorithm
// len(key) should be 8.
// Play: https://go.dev/play/p/4cC4QvWfe3_1
func DesCbcDecrypt(encrypted, key []byte) []byte {
	if len(key) != 8 {
		panic("des: key length must be 8 bytes")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		panic("des: failed to create cipher: " + err.Error())
	}

	blockSize := block.BlockSize()
	if len(encrypted) < blockSize || len(encrypted)%blockSize != 0 {
		panic("des: invalid encrypted data length")
	}

	iv := encrypted[:blockSize]
	ciphertext := encrypted[blockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return pkcs7UnPadding(ciphertext)
}

// DesCtrCrypt encrypt data with key use DES CTR algorithm
// len(key) should be 8.
// Play: https://go.dev/play/p/9-T6OjKpcdw
// deprecated: use DesCtrEncrypt and DesCtrDecrypt instead.
func DesCtrCrypt(data, key []byte) []byte {
	size := len(key)
	if size != 8 {
		panic("key length shoud be 8")
	}

	block, _ := des.NewCipher(key)

	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	dst := make([]byte, len(data))
	stream.XORKeyStream(dst, data)

	return dst
}

// DesCtrEncrypt encrypt data with key use DES CTR algorithm
// len(key) should be 8.
// Play: todo
func DesCtrEncrypt(data, key []byte) []byte {
	if len(key) != 8 {
		panic("des: key length must be 8 bytes")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		panic("des: failed to create cipher: " + err.Error())
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("des: failed to generate IV: " + err.Error())
	}

	stream := cipher.NewCTR(block, iv)

	encrypted := make([]byte, len(data))
	stream.XORKeyStream(encrypted, data)

	// 返回前缀包含 IV，便于解密
	return append(iv, encrypted...)
}

// DesCtrDecrypt decrypt data with key use DES CTR algorithm
// len(key) should be 8.
// Play: todo
func DesCtrDecrypt(encrypted, key []byte) []byte {
	if len(key) != 8 {
		panic("des: key length must be 8 bytes")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		panic("des: failed to create cipher: " + err.Error())
	}

	blockSize := block.BlockSize()
	if len(encrypted) < blockSize {
		panic("des: ciphertext too short")
	}

	iv := encrypted[:blockSize]
	ciphertext := encrypted[blockSize:]

	stream := cipher.NewCTR(block, iv)

	decrypted := make([]byte, len(ciphertext))
	stream.XORKeyStream(decrypted, ciphertext)

	return decrypted
}

// DesCfbEncrypt encrypt data with key use DES CFB algorithm
// len(key) should be 8.
// Play: https://go.dev/play/p/y-eNxcFBlxL
func DesCfbEncrypt(data, key []byte) []byte {
	if len(key) != 8 {
		panic("des: key length must be 8 bytes")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		panic("des: failed to create cipher: " + err.Error())
	}

	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("des: failed to generate IV: " + err.Error())
	}

	encrypted := make([]byte, des.BlockSize+len(data))

	copy(encrypted[:des.BlockSize], iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[des.BlockSize:], data)

	return encrypted
}

// DesCfbDecrypt decrypt data with key use DES CFB algorithm
// len(encrypted) should be great than 16, len(key) should be 8.
// Play: https://go.dev/play/p/y-eNxcFBlxL
func DesCfbDecrypt(encrypted, key []byte) []byte {
	if len(key) != 8 {
		panic("des: key length must be 8 bytes")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		panic("des: failed to create cipher: " + err.Error())
	}

	if len(encrypted) < des.BlockSize {
		panic("des: encrypted data too short")
	}

	iv := encrypted[:des.BlockSize]
	ciphertext := encrypted[des.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}

// DesOfbEncrypt encrypt data with key use DES OFB algorithm
// len(key) should be 8.
// Play: https://go.dev/play/p/74KmNadjN1J
func DesOfbEncrypt(data, key []byte) []byte {
	if len(key) != 8 {
		panic("des: key length must be 8 bytes")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		panic("des: failed to create cipher: " + err.Error())
	}

	data = pkcs7Padding(data, des.BlockSize)

	iv := make([]byte, des.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("des: failed to generate IV: " + err.Error())
	}

	encrypted := make([]byte, des.BlockSize+len(data))
	copy(encrypted[:des.BlockSize], iv)

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(encrypted[des.BlockSize:], data)

	return encrypted
}

// DesOfbDecrypt decrypt data with key use DES OFB algorithm
// len(key) should be 8.
// Play: https://go.dev/play/p/74KmNadjN1J
func DesOfbDecrypt(data, key []byte) []byte {
	if len(key) != 8 {
		panic("des: key length must be 8 bytes")
	}

	block, err := des.NewCipher(key)
	if err != nil {
		panic("des: failed to create cipher: " + err.Error())
	}

	if len(data) < des.BlockSize {
		panic("des: encrypted data too short")
	}

	iv := data[:des.BlockSize]
	ciphertext := data[des.BlockSize:]

	stream := cipher.NewOFB(block, iv)
	decrypted := make([]byte, len(ciphertext))
	stream.XORKeyStream(decrypted, ciphertext)

	decrypted = pkcs7UnPadding(decrypted)

	return decrypted
}

// GenerateRsaKey create rsa private and public pemo file.
// Play: https://go.dev/play/p/zutRHrDqs0X
func GenerateRsaKey(keySize int, priKeyFile, pubKeyFile string) error {
	// private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return err
	}

	derText := x509.MarshalPKCS1PrivateKey(privateKey)

	block := pem.Block{
		Type:  "rsa private key",
		Bytes: derText,
	}

	file, err := os.Create(priKeyFile)
	if err != nil {
		panic(err)
	}
	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}

	file.Close()

	// public key
	publicKey := privateKey.PublicKey

	derpText, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}

	block = pem.Block{
		Type:  "rsa public key",
		Bytes: derpText,
	}

	file, err = os.Create(pubKeyFile)
	if err != nil {
		return err
	}

	err = pem.Encode(file, &block)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

// RsaEncrypt encrypt data with ras algorithm.
// Play: https://go.dev/play/p/7_zo6mrx-eX
func RsaEncrypt(data []byte, pubKeyFileName string) []byte {
	file, err := os.Open(pubKeyFileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := make([]byte, fileInfo.Size())

	_, err = file.Read(buf)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(buf)

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	pubKey := pubInterface.(*rsa.PublicKey)

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, data)
	if err != nil {
		panic(err)
	}

	return cipherText
}

// RsaDecrypt decrypt data with ras algorithm.
// Play: https://go.dev/play/p/7_zo6mrx-eX
func RsaDecrypt(data []byte, privateKeyFileName string) []byte {
	file, err := os.Open(privateKeyFileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	defer file.Close()

	_, err = file.Read(buf)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(buf)

	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, data)
	if err != nil {
		panic(err)
	}

	return plainText
}

// GenerateRsaKeyPair create rsa private and public key.
// Play: https://go.dev/play/p/sSVmkfENKMz
func GenerateRsaKeyPair(keySize int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, keySize)
	return privateKey, &privateKey.PublicKey
}

// RsaEncryptOAEP encrypts the given data with RSA-OAEP.
// Play: https://go.dev/play/p/sSVmkfENKMz
func RsaEncryptOAEP(data []byte, label []byte, key rsa.PublicKey) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &key, data, label)
	if err != nil {
		return nil, err
	}

	return encryptedBytes, nil
}

// RsaDecryptOAEP decrypts the data with RSA-OAEP.
// Play: https://go.dev/play/p/sSVmkfENKMz
func RsaDecryptOAEP(ciphertext []byte, label []byte, key rsa.PrivateKey) ([]byte, error) {
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, &key, ciphertext, label)
	if err != nil {
		return nil, err
	}

	return decryptedBytes, nil
}

// RsaSign signs the data with RSA.
// Play: https://go.dev/play/p/qhsbf8BJ6Mf
func RsaSign(hash crypto.Hash, data []byte, privateKeyFileName string) ([]byte, error) {
	privateKey, err := loadRasPrivateKey(privateKeyFileName)
	if err != nil {
		return nil, err
	}

	hashed, err := hashData(hash, data)
	if err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, privateKey, hash, hashed)
}

// RsaVerifySign verifies the signature of the data with RSA.
// Play: https://go.dev/play/p/qhsbf8BJ6Mf
func RsaVerifySign(hash crypto.Hash, data, signature []byte, pubKeyFileName string) error {
	publicKey, err := loadRsaPublicKey(pubKeyFileName)
	if err != nil {
		return err
	}

	hashed, err := hashData(hash, data)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(publicKey, hash, hashed, signature)
}
