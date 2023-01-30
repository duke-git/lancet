package cryptor

import (
	"encoding/base64"

	"github.com/duke-git/lancet/v2/random"
	"golang.org/x/crypto/scrypt"
)

const (
	_saltLen = 8
	_keyLen  = 32

	_rn = 1 << 15
	_fn = 1 << 2
	_r  = 8 // _r * _p < 2^30
	_p  = 1 // _r * _p < 2^30
)

func RandSalt(length int) []byte {
	return random.RandBytes(length)
}

// CustomCryptography is a custom cryptography function.
func CustomCryptography(password, salt []byte, N, r, p, keyLen int) (string, error) {
	dk, err := scrypt.Key(password, salt, N, r, p, keyLen)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(dk), nil
}

// FastCryptography is a fast cryptography function. 8 bytes is a good length for salt.
func FastCryptography(salt []byte, key string, length int) (string, error) {
	return CustomCryptography([]byte(key), salt, _fn, _r, _p, length)
}

// RecommendedCryptography is a recommended cryptography function. 8 bytes is a good length for salt. It is slower than FastCryptography.
func RecommendedCryptography(salt []byte, key string, length int) (string, error) {
	return CustomCryptography([]byte(key), salt, _rn, _r, _p, length)
}

// DefaultFastCryptography is a default fast cryptography function. return res, salt, error.
func DefaultFastCryptography(key string) (string, string, error) {
	return cryptographyDefault(key, FastCryptography)
}

// DefaultRecommendedCryptography is a default recommended cryptography function. return res, salt, error.
func DefaultRecommendedCryptography(key string) (string, string, error) {
	return cryptographyDefault(key, RecommendedCryptography)
}

func cryptographyDefault(key string, do func(salt []byte, key string, length int) (string, error)) (string, string, error) {
	salt := RandSalt(_saltLen)
	dk, err := do(salt, key, _keyLen)
	if err != nil {
		return "", "", err
	}
	return dk, base64.StdEncoding.EncodeToString(salt), nil
}
