// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package cryptor implements some util functions to encrypt and decrypt.
package cryptor

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// GenerateRsaKey make  a rsa private key, and return key file name
// Generated key file is `rsa_private.pem` and `rsa_public.pem` in current path
func GenerateRsaKey(keySize int, priKeyFile, pubKeyFile string) {
	// private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}

	derText := x509.MarshalPKCS1PrivateKey(privateKey)

	block := pem.Block{
		Type:  "rsa private key",
		Bytes: derText,
	}

	//file,err := os.Create("rsa_private.pem")
	file, err := os.Create(priKeyFile)
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()

	// public key
	publicKey := privateKey.PublicKey

	derpText, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}

	block = pem.Block{
		Type:  "rsa public key",
		Bytes: derpText,
	}

	//file,err = os.Create("rsa_public.pem")
	file, err = os.Create(pubKeyFile)
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()
}

// RsaEncrypt encrypt data with ras algorithm
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
	file.Read(buf)

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

// RsaDecrypt decrypt data with ras algorithm
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
	file.Read(buf)

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
