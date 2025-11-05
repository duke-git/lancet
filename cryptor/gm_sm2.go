package cryptor

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"math/big"
)

// SM2 implements the Chinese SM2 elliptic curve public key algorithm.
// SM2 is based on elliptic curve cryptography and provides encryption, decryption, signing and verification.
//
// Note: This implementation uses crypto/elliptic package methods (GenerateKey, ScalarBaseMult, ScalarMult, IsOnCurve)
// which are marked as deprecated in Go 1.20+. These methods still work correctly and are widely used.
// The //nolint:staticcheck directive suppresses deprecation warnings.
// A future version may replace these with a custom elliptic curve implementation.

var (
	sm2P256       *sm2Curve
	sm2P256Params = &elliptic.CurveParams{Name: "sm2p256v1"}
)

func init() {
	// SM2 curve parameters
	sm2P256Params.P, _ = new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000FFFFFFFFFFFFFFFF", 16)
	sm2P256Params.N, _ = new(big.Int).SetString("FFFFFFFEFFFFFFFFFFFFFFFFFFFFFFFF7203DF6B21C6052B53BBF40939D54123", 16)
	sm2P256Params.B, _ = new(big.Int).SetString("28E9FA9E9D9F5E344D5A9E4BCF6509A7F39789F515AB8F92DDBCBD414D940E93", 16)
	sm2P256Params.Gx, _ = new(big.Int).SetString("32C4AE2C1F1981195F9904466A39C9948FE30BBFF2660BE1715A4589334C74C7", 16)
	sm2P256Params.Gy, _ = new(big.Int).SetString("BC3736A2F4F6779C59BDCEE36B692153D0A9877CC62A474002DF32E52139F0A0", 16)
	sm2P256Params.BitSize = 256

	sm2P256 = &sm2Curve{sm2P256Params}
}

type sm2Curve struct {
	*elliptic.CurveParams
}

// Sm2PrivateKey represents an SM2 private key.
type Sm2PrivateKey struct {
	D         *big.Int
	PublicKey Sm2PublicKey
}

// Sm2PublicKey represents an SM2 public key.
type Sm2PublicKey struct {
	X, Y *big.Int
}

// GenerateSm2Key generates a new SM2 private/public key pair.
// Play: https://go.dev/play/p/bKYMqRLvIx3
func GenerateSm2Key() (*Sm2PrivateKey, error) {
	priv, x, y, err := elliptic.GenerateKey(sm2P256, rand.Reader)
	if err != nil {
		return nil, err
	}

	privateKey := &Sm2PrivateKey{
		D: new(big.Int).SetBytes(priv),
		PublicKey: Sm2PublicKey{
			X: x,
			Y: y,
		},
	}

	return privateKey, nil
}

// Sm2Encrypt encrypts plaintext using SM2 public key.
// Returns ciphertext in the format: C1 || C3 || C2
// C1 = kG (65 bytes in uncompressed format)
// C3 = Hash(x2 || M || y2) (32 bytes for SM3)
// C2 = M xor t (same length as plaintext)
// Play: https://go.dev/play/p/bKYMqRLvIx3
func Sm2Encrypt(pub *Sm2PublicKey, plaintext []byte) ([]byte, error) {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil, errors.New("sm2: invalid public key")
	}

	for {
		// Generate random k
		k, err := randFieldElement(sm2P256, rand.Reader)
		if err != nil {
			return nil, err
		}

		// C1 = kG
		c1x, c1y := sm2P256.ScalarBaseMult(k.Bytes())

		// kP = (x2, y2)
		x2, y2 := sm2P256.ScalarMult(pub.X, pub.Y, k.Bytes())

		// Derive key using KDF
		kdfLen := len(plaintext)
		t := sm2KDF(append(toBytes(sm2P256, x2), toBytes(sm2P256, y2)...), kdfLen)

		// Check if t is all zeros
		allZero := true
		for _, b := range t {
			if b != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			continue
		}

		// C2 = M xor t
		c2 := make([]byte, len(plaintext))
		for i := 0; i < len(plaintext); i++ {
			c2[i] = plaintext[i] ^ t[i]
		}

		// C3 = Hash(x2 || M || y2)
		c3Input := append(toBytes(sm2P256, x2), plaintext...)
		c3Input = append(c3Input, toBytes(sm2P256, y2)...)
		c3 := Sm3(c3Input)

		// Return C1 || C3 || C2
		c1 := sm2MarshalUncompressed(sm2P256, c1x, c1y)
		result := append(c1, c3...)
		result = append(result, c2...)

		return result, nil
	}
}

// Sm2Decrypt decrypts ciphertext using SM2 private key.
// Expects ciphertext in the format: C1 || C3 || C2
// Play: https://go.dev/play/p/bKYMqRLvIx3
func Sm2Decrypt(priv *Sm2PrivateKey, ciphertext []byte) ([]byte, error) {
	if priv == nil || priv.D == nil {
		return nil, errors.New("sm2: invalid private key")
	}

	// Parse C1 (65 bytes), C3 (32 bytes), C2 (remaining)
	if len(ciphertext) < 97 {
		return nil, errors.New("sm2: ciphertext too short")
	}

	c1 := ciphertext[:65]
	c3 := ciphertext[65:97]
	c2 := ciphertext[97:]

	// Parse C1
	c1x, c1y := sm2UnmarshalUncompressed(sm2P256, c1)
	if c1x == nil {
		return nil, errors.New("sm2: invalid C1 point")
	}

	// Verify C1 is on curve
	if !sm2P256.IsOnCurve(c1x, c1y) {
		return nil, errors.New("sm2: C1 not on curve")
	}

	// dC1 = (x2, y2)
	x2, y2 := sm2P256.ScalarMult(c1x, c1y, priv.D.Bytes())

	// Derive key using KDF
	kdfLen := len(c2)
	t := sm2KDF(append(toBytes(sm2P256, x2), toBytes(sm2P256, y2)...), kdfLen)

	// M = C2 xor t
	plaintext := make([]byte, len(c2))
	for i := 0; i < len(c2); i++ {
		plaintext[i] = c2[i] ^ t[i]
	}

	// Verify C3 = Hash(x2 || M || y2)
	u := append(toBytes(sm2P256, x2), plaintext...)
	u = append(u, toBytes(sm2P256, y2)...)
	hash := Sm3(u)

	for i := 0; i < len(c3); i++ {
		if c3[i] != hash[i] {
			return nil, errors.New("sm2: hash verification failed")
		}
	}

	return plaintext, nil
}

// SM2 KDF (Key Derivation Function)
func sm2KDF(z []byte, klen int) []byte {
	limit := (klen + 31) / 32
	result := make([]byte, 0, limit*32)

	for i := 1; i <= limit; i++ {
		counter := make([]byte, 4)
		binary.BigEndian.PutUint32(counter, uint32(i))
		hash := Sm3(append(z, counter...))
		result = append(result, hash...)
	}

	return result[:klen]
}

func toBytes(curve elliptic.Curve, value *big.Int) []byte {
	byteLen := (curve.Params().BitSize + 7) / 8
	buf := make([]byte, byteLen)
	b := value.Bytes()
	copy(buf[byteLen-len(b):], b)
	return buf
}

func sm2MarshalUncompressed(curve *sm2Curve, x, y *big.Int) []byte {
	byteLen := (curve.BitSize + 7) / 8
	ret := make([]byte, 1+2*byteLen)
	ret[0] = 4 // uncompressed point

	xBytes := x.Bytes()
	copy(ret[1+byteLen-len(xBytes):], xBytes)
	yBytes := y.Bytes()
	copy(ret[1+2*byteLen-len(yBytes):], yBytes)

	return ret
}

func sm2UnmarshalUncompressed(curve *sm2Curve, data []byte) (*big.Int, *big.Int) {
	byteLen := (curve.BitSize + 7) / 8
	if len(data) != 1+2*byteLen {
		return nil, nil
	}
	if data[0] != 4 {
		return nil, nil
	}

	x := new(big.Int).SetBytes(data[1 : 1+byteLen])
	y := new(big.Int).SetBytes(data[1+byteLen:])

	return x, y
}

func randFieldElement(c elliptic.Curve, rand io.Reader) (*big.Int, error) {
	params := c.Params()
	b := make([]byte, params.BitSize/8+8)
	_, err := io.ReadFull(rand, b)
	if err != nil {
		return nil, err
	}

	k := new(big.Int).SetBytes(b)
	n := new(big.Int).Sub(params.N, big.NewInt(1))
	k.Mod(k, n)
	k.Add(k, big.NewInt(1))

	return k, nil
}
