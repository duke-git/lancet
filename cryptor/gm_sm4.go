package cryptor

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"io"
)

// SM4 implements the Chinese SM4 block cipher.
// SM4 is a 128-bit block cipher with 128-bit keys.
// This implementation uses pre-computed lookup tables for optimal performance.

const sm4BlockSize = 16

// Pre-computed T-transformation lookup tables for performance optimization
var sm4T1Table [256][4]uint32 // S-box + L1 transformation
var sm4T2Table [256][4]uint32 // S-box + L2 transformation

var sm4Sbox = [256]byte{
	0xd6, 0x90, 0xe9, 0xfe, 0xcc, 0xe1, 0x3d, 0xb7, 0x16, 0xb6, 0x14, 0xc2, 0x28, 0xfb, 0x2c, 0x05,
	0x2b, 0x67, 0x9a, 0x76, 0x2a, 0xbe, 0x04, 0xc3, 0xaa, 0x44, 0x13, 0x26, 0x49, 0x86, 0x06, 0x99,
	0x9c, 0x42, 0x50, 0xf4, 0x91, 0xef, 0x98, 0x7a, 0x33, 0x54, 0x0b, 0x43, 0xed, 0xcf, 0xac, 0x62,
	0xe4, 0xb3, 0x1c, 0xa9, 0xc9, 0x08, 0xe8, 0x95, 0x80, 0xdf, 0x94, 0xfa, 0x75, 0x8f, 0x3f, 0xa6,
	0x47, 0x07, 0xa7, 0xfc, 0xf3, 0x73, 0x17, 0xba, 0x83, 0x59, 0x3c, 0x19, 0xe6, 0x85, 0x4f, 0xa8,
	0x68, 0x6b, 0x81, 0xb2, 0x71, 0x64, 0xda, 0x8b, 0xf8, 0xeb, 0x0f, 0x4b, 0x70, 0x56, 0x9d, 0x35,
	0x1e, 0x24, 0x0e, 0x5e, 0x63, 0x58, 0xd1, 0xa2, 0x25, 0x22, 0x7c, 0x3b, 0x01, 0x21, 0x78, 0x87,
	0xd4, 0x00, 0x46, 0x57, 0x9f, 0xd3, 0x27, 0x52, 0x4c, 0x36, 0x02, 0xe7, 0xa0, 0xc4, 0xc8, 0x9e,
	0xea, 0xbf, 0x8a, 0xd2, 0x40, 0xc7, 0x38, 0xb5, 0xa3, 0xf7, 0xf2, 0xce, 0xf9, 0x61, 0x15, 0xa1,
	0xe0, 0xae, 0x5d, 0xa4, 0x9b, 0x34, 0x1a, 0x55, 0xad, 0x93, 0x32, 0x30, 0xf5, 0x8c, 0xb1, 0xe3,
	0x1d, 0xf6, 0xe2, 0x2e, 0x82, 0x66, 0xca, 0x60, 0xc0, 0x29, 0x23, 0xab, 0x0d, 0x53, 0x4e, 0x6f,
	0xd5, 0xdb, 0x37, 0x45, 0xde, 0xfd, 0x8e, 0x2f, 0x03, 0xff, 0x6a, 0x72, 0x6d, 0x6c, 0x5b, 0x51,
	0x8d, 0x1b, 0xaf, 0x92, 0xbb, 0xdd, 0xbc, 0x7f, 0x11, 0xd9, 0x5c, 0x41, 0x1f, 0x10, 0x5a, 0xd8,
	0x0a, 0xc1, 0x31, 0x88, 0xa5, 0xcd, 0x7b, 0xbd, 0x2d, 0x74, 0xd0, 0x12, 0xb8, 0xe5, 0xb4, 0xb0,
	0x89, 0x69, 0x97, 0x4a, 0x0c, 0x96, 0x77, 0x7e, 0x65, 0xb9, 0xf1, 0x09, 0xc5, 0x6e, 0xc6, 0x84,
	0x18, 0xf0, 0x7d, 0xec, 0x3a, 0xdc, 0x4d, 0x20, 0x79, 0xee, 0x5f, 0x3e, 0xd7, 0xcb, 0x39, 0x48,
}

var sm4FK = [4]uint32{0xa3b1bac6, 0x56aa3350, 0x677d9197, 0xb27022dc}

var sm4CK = [32]uint32{
	0x00070e15, 0x1c232a31, 0x383f464d, 0x545b6269,
	0x70777e85, 0x8c939aa1, 0xa8afb6bd, 0xc4cbd2d9,
	0xe0e7eef5, 0xfc030a11, 0x181f262d, 0x343b4249,
	0x50575e65, 0x6c737a81, 0x888f969d, 0xa4abb2b9,
	0xc0c7ced5, 0xdce3eaf1, 0xf8ff060d, 0x141b2229,
	0x30373e45, 0x4c535a61, 0x686f767d, 0x848b9299,
	0xa0a7aeb5, 0xbcc3cad1, 0xd8dfe6ed, 0xf4fb0209,
	0x10171e25, 0x2c333a41, 0x484f565d, 0x646b7279,
}

// 初始化预计算查找表
func init() {
	// Pre-compute all possible T1 and T2 transformations
	for pos := 0; pos < 4; pos++ {
		for i := 0; i < 256; i++ {
			// S-box 替换
			sboxVal := sm4Sbox[i]
			
			// 根据字节位置计算偏移
			shift := uint32((3 - pos) * 8)
			b := uint32(sboxVal) << shift
			
			// L1 变换：b ^ ROL(b,2) ^ ROL(b,10) ^ ROL(b,18) ^ ROL(b,24)
			sm4T1Table[i][pos] = b ^ sm4RotateLeft(b, 2) ^ sm4RotateLeft(b, 10) ^ sm4RotateLeft(b, 18) ^ sm4RotateLeft(b, 24)
			
			// L2 变换：b ^ ROL(b,13) ^ ROL(b,23)
			sm4T2Table[i][pos] = b ^ sm4RotateLeft(b, 13) ^ sm4RotateLeft(b, 23)
		}
	}
}

type sm4Cipher struct {
	enc [32]uint32
	dec [32]uint32
}

// Sm4EcbEncrypt encrypts data using SM4 in ECB mode.
// key must be 16 bytes.
// Play: https://go.dev/play/p/l5IQxYuuaED
func Sm4EcbEncrypt(data, key []byte) []byte {
	if len(key) != 16 {
		panic("sm4: key length must be 16 bytes")
	}

	c := newSm4Cipher(key)
	padded := pkcs7Padding(data, sm4BlockSize)
	encrypted := make([]byte, len(padded))

	for i := 0; i < len(padded); i += sm4BlockSize {
		c.Encrypt(encrypted[i:i+sm4BlockSize], padded[i:i+sm4BlockSize])
	}

	return encrypted
}

// Sm4EcbDecrypt decrypts data using SM4 in ECB mode.
// key must be 16 bytes.
// Play: https://go.dev/play/p/l5IQxYuuaED
func Sm4EcbDecrypt(encrypted, key []byte) []byte {
	if len(key) != 16 {
		panic("sm4: key length must be 16 bytes")
	}

	if len(encrypted)%sm4BlockSize != 0 {
		panic("sm4: encrypted data length must be multiple of block size")
	}

	c := newSm4Cipher(key)
	decrypted := make([]byte, len(encrypted))

	for i := 0; i < len(encrypted); i += sm4BlockSize {
		c.Decrypt(decrypted[i:i+sm4BlockSize], encrypted[i:i+sm4BlockSize])
	}

	return pkcs7UnPadding(decrypted)
}

// Sm4CbcEncrypt encrypts data using SM4 in CBC mode.
// key must be 16 bytes.
// Play: https://go.dev/play/p/65Q6iYhLRTa
func Sm4CbcEncrypt(data, key []byte) []byte {
	if len(key) != 16 {
		panic("sm4: key length must be 16 bytes")
	}

	c := newSm4Cipher(key)
	padded := pkcs7Padding(data, sm4BlockSize)

	iv := make([]byte, sm4BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("sm4: failed to generate IV: " + err.Error())
	}

	encrypted := make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(c, iv)
	mode.CryptBlocks(encrypted, padded)

	return append(iv, encrypted...)
}

// Sm4CbcDecrypt decrypts data using SM4 in CBC mode.
// key must be 16 bytes.
// Play: https://go.dev/play/p/65Q6iYhLRTa
func Sm4CbcDecrypt(encrypted, key []byte) []byte {
	if len(key) != 16 {
		panic("sm4: key length must be 16 bytes")
	}

	if len(encrypted) < sm4BlockSize {
		panic("sm4: encrypted data too short")
	}

	if len(encrypted)%sm4BlockSize != 0 {
		panic("sm4: encrypted data length must be multiple of block size")
	}

	c := newSm4Cipher(key)
	iv := encrypted[:sm4BlockSize]
	ciphertext := encrypted[sm4BlockSize:]

	decrypted := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(c, iv)
	mode.CryptBlocks(decrypted, ciphertext)

	return pkcs7UnPadding(decrypted)
}

func newSm4Cipher(key []byte) *sm4Cipher {
	c := &sm4Cipher{}

	var mk [4]uint32
	for i := 0; i < 4; i++ {
		mk[i] = binary.BigEndian.Uint32(key[i*4 : (i+1)*4])
	}

	var k [36]uint32
	k[0] = mk[0] ^ sm4FK[0]
	k[1] = mk[1] ^ sm4FK[1]
	k[2] = mk[2] ^ sm4FK[2]
	k[3] = mk[3] ^ sm4FK[3]

	for i := 0; i < 32; i++ {
		k[i+4] = k[i] ^ sm4T2Fast(k[i+1]^k[i+2]^k[i+3]^sm4CK[i])
		c.enc[i] = k[i+4]
	}

	for i := 0; i < 32; i++ {
		c.dec[i] = c.enc[31-i]
	}

	return c
}

func (c *sm4Cipher) BlockSize() int {
	return sm4BlockSize
}

func (c *sm4Cipher) Encrypt(dst, src []byte) {
	if len(src) < sm4BlockSize {
		panic("sm4: input not full block")
	}
	if len(dst) < sm4BlockSize {
		panic("sm4: output not full block")
	}

	// 使用局部变量避免数组分配，提升性能
	x0 := binary.BigEndian.Uint32(src[0:4])
	x1 := binary.BigEndian.Uint32(src[4:8])
	x2 := binary.BigEndian.Uint32(src[8:12])
	x3 := binary.BigEndian.Uint32(src[12:16])

	// 32 轮加密
	for i := 0; i < 32; i++ {
		t := x1 ^ x2 ^ x3 ^ c.enc[i]
		x0 ^= sm4T1Fast(t)
		x0, x1, x2, x3 = x1, x2, x3, x0
	}

	binary.BigEndian.PutUint32(dst[0:4], x3)
	binary.BigEndian.PutUint32(dst[4:8], x2)
	binary.BigEndian.PutUint32(dst[8:12], x1)
	binary.BigEndian.PutUint32(dst[12:16], x0)
}

func (c *sm4Cipher) Decrypt(dst, src []byte) {
	if len(src) < sm4BlockSize {
		panic("sm4: input not full block")
	}
	if len(dst) < sm4BlockSize {
		panic("sm4: output not full block")
	}

	x0 := binary.BigEndian.Uint32(src[0:4])
	x1 := binary.BigEndian.Uint32(src[4:8])
	x2 := binary.BigEndian.Uint32(src[8:12])
	x3 := binary.BigEndian.Uint32(src[12:16])

	// 32 轮解密
	for i := 0; i < 32; i++ {
		t := x1 ^ x2 ^ x3 ^ c.dec[i]
		x0 ^= sm4T1Fast(t)
		x0, x1, x2, x3 = x1, x2, x3, x0
	}

	binary.BigEndian.PutUint32(dst[0:4], x3)
	binary.BigEndian.PutUint32(dst[4:8], x2)
	binary.BigEndian.PutUint32(dst[8:12], x1)
	binary.BigEndian.PutUint32(dst[12:16], x0)
}

// 使用预计算查找表的快速 T1 变换（用于加密轮函数）
func sm4T1Fast(a uint32) uint32 {
	return sm4T1Table[byte(a>>24)][0] ^
		sm4T1Table[byte(a>>16)][1] ^
		sm4T1Table[byte(a>>8)][2] ^
		sm4T1Table[byte(a)][3]
}

// 使用预计算查找表的快速 T2 变换（用于密钥扩展）
func sm4T2Fast(a uint32) uint32 {
	return sm4T2Table[byte(a>>24)][0] ^
		sm4T2Table[byte(a>>16)][1] ^
		sm4T2Table[byte(a>>8)][2] ^
		sm4T2Table[byte(a)][3]
}

func sm4RotateLeft(x uint32, n uint32) uint32 {
	return (x << n) | (x >> (32 - n))
}
