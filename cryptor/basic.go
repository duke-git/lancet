// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package cryptor implements some util functions to encrypt and decrypt.
// Contain base64, hmac, sha, aes, des, and rsa
package cryptor

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
)

// Base64StdEncode encode string with base64 encoding
func Base64StdEncode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64StdEncode decode a base64 encoded string
func Base64StdDecode(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

// Md5Str return the md5 value of string
func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5File return the md5 value of file
func Md5File(filename string) (string, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	h := md5.New()
	h.Write(f)
	return hex.EncodeToString(h.Sum(nil)), nil
}

// HmacMd5 return the hmac hash of string use md5
func HmacMd5(data, key string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha1 return the hmac hash of string use sha1
func HmacSha1(data, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha256 return the hmac hash of string use sha256
func HmacSha256(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// HmacSha512 return the hmac hash of string use sha512
func HmacSha512(data, key string) string {
	h := hmac.New(sha512.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum([]byte("")))
}

// Sha1 return the sha1 value (SHA-1 hash algorithm) of string
func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

// Sha256 return the sha256 value (SHA256 hash algorithm) of string
func Sha256(data string) string {
	sha256 := sha256.New()
	sha256.Write([]byte(data))
	return hex.EncodeToString(sha256.Sum([]byte("")))
}

// Sha512 return the sha512 value (SHA512 hash algorithm) of string
func Sha512(data string) string {
	sha512 := sha512.New()
	sha512.Write([]byte(data))
	return hex.EncodeToString(sha512.Sum([]byte("")))
}
