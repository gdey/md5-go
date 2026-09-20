package md5

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

type Options struct {
	Buf []byte
}

func (o *Options) buf() []byte {
	if o == nil || o.Buf == nil {
		return make([]byte, blockSize)
	}
	return o.Buf
}

func Hash(r io.Reader, opts *Options) (code [16]byte, err error) {
	p := process{
		a0:   a0,
		b0:   b0,
		c0:   c0,
		d0:   d0,
		_buf: opts.buf(),
	}

	err = p.Input(r)
	if err != nil {
		return code, err
	}

	return p.Code(), nil
}

// //////  Process functions ////////////////////////////////////////

type process struct {
	_buf  []byte
	block [blockSize]byte
	idx   int
	count int64
	a0    uint32
	b0    uint32
	c0    uint32
	d0    uint32
}

func (p *process) Code() (buf [16]byte) {

	binary.LittleEndian.PutUint32(buf[0:4], p.a0)
	binary.LittleEndian.PutUint32(buf[4:8], p.b0)
	binary.LittleEndian.PutUint32(buf[8:12], p.c0)
	binary.LittleEndian.PutUint32(buf[12:16], p.d0)
	return buf
}

func (p *process) Input(r io.Reader) error {

	for {
		n, err := r.Read(p._buf[:])
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		if n == 0 {
			panic("did not expect to read zero")
		}

		p.Byte(p._buf[:n]...)
		p.count += int64(n)
	}

	// we've read the entire message - now we have to finalize it by padding
	// it and appending the length

	// pad it with 1 and then a bunch of 0s
	p.Byte(0x80)

	// Make sure there is space for the 8 we need to append
	for p.idx != (blockSize - 8) {
		p.Byte(0)
	}

	// append original length in bits mod 2^64 to message
	// turn the byte length into 2 numbers - high bits and low bits
	lowBits := uint32(p.count << 3)
	highBits := uint32(p.count >> (32 - 3))

	// encode the full 64bit int (bits length) as little endian
	for i := range 4 {
		b := byte(lowBits >> (i * 8) & 0xff)
		p.Byte(b)
	}
	for i := range 4 {
		b := byte(highBits >> (i * 8) & 0xff)
		p.Byte(b)
	}
	if p.idx != 0 {
		fmt.Fprintf(os.Stderr, "uh oh lol bad block index: %v\n", p.count)
		os.Exit(3)
	}
	return nil
}

func (p *process) Byte(bs ...byte) {
	lnBS := len(bs)
	if lnBS == 0 {
		return
	}

	if lnBS == 1 {
		p.block[p.idx] = bs[0]
		p.idx++
		if p.idx == blockSize {
			p.Block()
			p.idx = 0
		}
		return
	}

	// lbIdx by construction has to be 2+
	lbIdx := lnBS + p.idx
	// fits cannot be (blockSize - 1)+
	fits := blockSize - lbIdx

	// fits is only ever zero if there is only blocksize of bytes left in bs
	if fits == 0 {
		copy(p.block[p.idx:], bs)
		p.Block()
		p.idx = 0
		return
	}

	if fits > 0 {
		copy(p.block[p.idx:], bs)
		p.idx = lbIdx
		return
	}

	rBlkLen := blockSize - p.idx

	copy(p.block[p.idx:], bs)
	p.Block()
	p.idx = 0
	bs = bs[rBlkLen:]
	fits += rBlkLen
	lnBS = len(bs)

	for lnBS >= blockSize {
		copy(p.block[:], bs)
		p.Block()
		fits += blockSize
		bs = bs[blockSize:]
		lnBS -= blockSize
	}

	copy(p.block[:], bs)
	p.idx = lnBS

}

func (p *process) Block() {

	var (
		M [16]uint32

		// Initialize hash value for this chunk:
		A = p.a0
		B = p.b0
		C = p.c0
		D = p.d0
	)

	// break chunk into sixteen 32-bit words M[j], 0 ≤ j ≤ 15

	for i := range len(M) {
		j := i * 4
		M[i] = uint32(p.block[j]) | uint32(p.block[j+1])<<8 |
			uint32(p.block[j+2])<<16 | uint32(p.block[j+3])<<24
	}

	for ii := range blockSize {
		var (
			i = uint32(ii)
			F uint32
			g uint32
		)

		if i < 16 {
			F = (B & C) | ((^B) & D)
			g = (i)
		} else if i < 32 {
			F = (D & B) | ((^D) & C)
			g = (5*i + 1) % 16
		} else if i < 48 {
			F = B ^ C ^ D
			g = (3*i + 5) % 16
		} else {
			F = C ^ (B | ^D)
			g = (7 * i) % 16
		}

		F = F + A + k[i] + M[g]
		A = D
		D = C
		C = B

		B += (F << s[i]) | (F >> (32 - s[i]))
	}

	p.a0 += A
	p.b0 += B
	p.c0 += C
	p.d0 += D
}

// //////  CONSTANTS ////////////////////////////////////////
var (
	s = [...]uint32{
		7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
		5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
		4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
		6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
	}

	k = [...]uint32{
		0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee,
		0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
		0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be,
		0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
		0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa,
		0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
		0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed,
		0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
		0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c,
		0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
		0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05,
		0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
		0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039,
		0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
		0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1,
		0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
	}
)

const (
	a0 = uint32(0x67452301)
	b0 = uint32(0xefcdab89)
	c0 = uint32(0x98badcfe)
	d0 = uint32(0x10325476)

	blockSize = 64
)
