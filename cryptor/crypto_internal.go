package cryptor

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"strings"
)

func generateAesKey(key []byte, size int) []byte {
	genKey := make([]byte, size)
	copy(genKey, key)
	for i := size; i < len(key); {
		for j := 0; j < size && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func generateDesKey(key []byte) []byte {
	genKey := make([]byte, 8)
	copy(genKey, key)
	for i := 8; i < len(key); {
		for j := 0; j < 8 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func pkcs7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func pkcs7UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}

func pkcs5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs5UnPadding(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	padLen := int(data[length-1])
	if padLen == 0 || padLen > length {
		return nil
	}
	return data[:length-padLen]
}

func isAesKeyLengthValid(n int) bool {
	return n == 16 || n == 24 || n == 32
}

// loadRsaPrivateKey loads and parses a PEM encoded private key file.
func loadRsaPublicKey(filename string) (*rsa.PublicKey, error) {
	pubKeyData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pubKeyData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the public key")
	}

	var pubKey *rsa.PublicKey
	blockType := strings.ToUpper(block.Type)

	if blockType == "RSA PUBLIC KEY" {
		pubKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			key, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return nil, err
			}

			var ok bool
			pubKey, ok = key.(*rsa.PublicKey)
			if !ok {
				return nil, errors.New("failed to parse RSA private key")
			}
		}
	} else if blockType == "PUBLIC KEY" {
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}

		var ok bool
		pubKey, ok = key.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("failed to parse RSA private key")
		}

	} else {
		return nil, errors.New("unsupported key type")
	}

	return pubKey, nil
}

// loadRsaPrivateKey loads and parses a PEM encoded private key file.
func loadRasPrivateKey(filename string) (*rsa.PrivateKey, error) {
	priKeyData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(priKeyData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the private key")
	}

	var privateKey *rsa.PrivateKey
	blockType := strings.ToUpper(block.Type)

	// PKCS#1 format
	if blockType == "RSA PRIVATE KEY" {
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	} else if blockType == "PRIVATE KEY" { // PKCS#8 format
		priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		var ok bool
		privateKey, ok = priKey.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("failed to parse RSA private key")
		}
	} else {
		return nil, errors.New("unsupported key type")
	}

	return privateKey, nil
}

// hashData returns the hash value of the data, using the specified hash function
func hashData(hash crypto.Hash, data []byte) ([]byte, error) {
	if !hash.Available() {
		return nil, errors.New("unsupported hash algorithm")
	}

	var hashed []byte

	switch hash {
	case crypto.SHA224:
		h := sha256.Sum224(data)
		hashed = h[:]
	case crypto.SHA256:
		h := sha256.Sum256(data)
		hashed = h[:]
	case crypto.SHA384:
		h := sha512.Sum384(data)
		hashed = h[:]
	case crypto.SHA512:
		h := sha512.Sum512(data)
		hashed = h[:]
	default:
		return nil, errors.New("unsupported hash algorithm")
	}

	return hashed, nil
}
