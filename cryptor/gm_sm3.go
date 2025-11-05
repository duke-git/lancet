package cryptor

import (
	"encoding/binary"
	"hash"
)

// SM3 implements the Chinese SM3 cryptographic hash algorithm.
// SM3 produces a 256-bit (32-byte) hash value.

const (
	sm3BlockSize = 64
	sm3Size      = 32
	sm3T1        = 0x79cc4519
	sm3T2        = 0x7a879d8a
)

var sm3IV = [8]uint32{
	0x7380166f, 0x4914b2b9, 0x172442d7, 0xda8a0600,
	0xa96f30bc, 0x163138aa, 0xe38dee4d, 0xb0fb0e4e,
}

type sm3Digest struct {
	h   [8]uint32
	x   [sm3BlockSize]byte
	nx  int
	len uint64
}

// Sm3 returns a new hash.Hash computing the SM3 checksum.
// Play: https://go.dev/play/p/zDAQpteAiOc
func Sm3(data []byte) []byte {
	h := newSm3()
	h.Write(data)
	return h.Sum(nil)
}

func newSm3() hash.Hash {
	d := new(sm3Digest)
	d.Reset()
	return d
}

func (d *sm3Digest) Reset() {
	d.h = sm3IV
	d.nx = 0
	d.len = 0
}

func (d *sm3Digest) Size() int {
	return sm3Size
}

func (d *sm3Digest) BlockSize() int {
	return sm3BlockSize
}

func (d *sm3Digest) Write(p []byte) (nn int, err error) {
	nn = len(p)
	d.len += uint64(nn)

	if d.nx > 0 {
		n := copy(d.x[d.nx:], p)
		d.nx += n
		if d.nx == sm3BlockSize {
			sm3Block(d, d.x[:])
			d.nx = 0
		}
		p = p[n:]
	}

	if len(p) >= sm3BlockSize {
		n := len(p) &^ (sm3BlockSize - 1)
		sm3Block(d, p[:n])
		p = p[n:]
	}

	if len(p) > 0 {
		d.nx = copy(d.x[:], p)
	}

	return
}

func (d *sm3Digest) Sum(in []byte) []byte {
	d0 := *d
	hash := d0.checkSum()
	return append(in, hash[:]...)
}

func (d *sm3Digest) checkSum() [sm3Size]byte {
	len := d.len
	var tmp [64]byte
	tmp[0] = 0x80

	if len%64 < 56 {
		d.Write(tmp[0 : 56-len%64])
	} else {
		d.Write(tmp[0 : 64+56-len%64])
	}

	len <<= 3
	binary.BigEndian.PutUint64(tmp[:], len)
	d.Write(tmp[0:8])

	if d.nx != 0 {
		panic("d.nx != 0")
	}

	var digest [sm3Size]byte
	for i := 0; i < 8; i++ {
		binary.BigEndian.PutUint32(digest[i*4:], d.h[i])
	}

	return digest
}

func sm3Block(dig *sm3Digest, p []byte) {
	var w [68]uint32
	var w1 [64]uint32

	h0, h1, h2, h3, h4, h5, h6, h7 := dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7]

	for len(p) >= sm3BlockSize {
		for i := 0; i < 16; i++ {
			j := i * 4
			w[i] = binary.BigEndian.Uint32(p[j : j+4])
		}

		for i := 16; i < 68; i++ {
			w[i] = sm3P1(w[i-16]^w[i-9]^sm3RotateLeft(w[i-3], 15)) ^ sm3RotateLeft(w[i-13], 7) ^ w[i-6]
		}

		for i := 0; i < 64; i++ {
			w1[i] = w[i] ^ w[i+4]
		}

		A, B, C, D, E, F, G, H := h0, h1, h2, h3, h4, h5, h6, h7

		for j := 0; j < 64; j++ {
			var ss1, ss2, tt1, tt2, t uint32

			if j < 16 {
				t = sm3T1
			} else {
				t = sm3T2
			}

			ss1 = sm3RotateLeft(sm3RotateLeft(A, 12)+E+sm3RotateLeft(t, uint32(j%32)), 7)
			ss2 = ss1 ^ sm3RotateLeft(A, 12)

			if j < 16 {
				tt1 = sm3FF0(A, B, C) + D + ss2 + w1[j]
				tt2 = sm3GG0(E, F, G) + H + ss1 + w[j]
			} else {
				tt1 = sm3FF1(A, B, C) + D + ss2 + w1[j]
				tt2 = sm3GG1(E, F, G) + H + ss1 + w[j]
			}

			D = C
			C = sm3RotateLeft(B, 9)
			B = A
			A = tt1
			H = G
			G = sm3RotateLeft(F, 19)
			F = E
			E = sm3P0(tt2)
		}

		h0 ^= A
		h1 ^= B
		h2 ^= C
		h3 ^= D
		h4 ^= E
		h5 ^= F
		h6 ^= G
		h7 ^= H

		p = p[sm3BlockSize:]
	}

	dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7] = h0, h1, h2, h3, h4, h5, h6, h7
}

func sm3RotateLeft(x, n uint32) uint32 {
	return (x << n) | (x >> (32 - n))
}

func sm3P0(x uint32) uint32 {
	return x ^ sm3RotateLeft(x, 9) ^ sm3RotateLeft(x, 17)
}

func sm3P1(x uint32) uint32 {
	return x ^ sm3RotateLeft(x, 15) ^ sm3RotateLeft(x, 23)
}

func sm3FF0(x, y, z uint32) uint32 {
	return x ^ y ^ z
}

func sm3FF1(x, y, z uint32) uint32 {
	return (x & y) | (x & z) | (y & z)
}

func sm3GG0(x, y, z uint32) uint32 {
	return x ^ y ^ z
}

func sm3GG1(x, y, z uint32) uint32 {
	return (x & y) | (^x & z)
}
