package md5

import (
	"encoding/binary"
	"errors"
	"io"
	"math/bits"
)

type Options struct {
	Buf []byte
}

func (o *Options) buf() []byte {
	if o == nil || o.Buf == nil {
		return make([]byte, 2*blockSize)
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
	start int
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

	p.start = 0
	for {
		n, err := r.Read(p._buf[p.idx:])
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		if n == 0 {
			panic("did not expect to read zero")
		}

		p.AmendBytes(n)
		p.count += int64(n)
	}

	// we've read the entire message - now we have to finalize it by padding
	// it and appending the length

	// pad it with 1 and then a bunch of 0s
	p.Byte(0x80)

	if p.idx >= (blockSize - 8) {
		// Make sure there is space for the 8 we need to append
		copy(p._buf[p.idx:(2*blockSize-p.idx)], zeros)
		p.Block()
		p.start = blockSize
	} else {
		copy(p._buf[p.idx:blockSize], zeros)
	}
	p.idx = (blockSize - 8)

	// append original length in bits mod 2^64 to message
	// turn the byte length into 2 numbers - high bits and low bits
	lowBits := uint32(p.count << 3)
	highBits := uint32(p.count >> (32 - 3))

	// Write 8 bytes directly to buffer without going through Byte()
	for i := range 4 {
		p._buf[p.start+p.idx] = byte(lowBits >> (i * 8) & 0xff)
		p.idx++
	}
	for i := range 4 {
		p._buf[p.start+p.idx] = byte(highBits >> (i * 8) & 0xff)
		p.idx++
	}

	p.Block()

	return nil
}

func (p *process) AmendBytes(n int) {
	if n == 0 {
		return
	}
	p.start = 0

	if n == 1 {
		p.idx++
		if p.idx == blockSize {
			p.Block()
			p.idx = 0
		}
		return
	}

	// lbIdx by construction has to be 2+
	lbIdx := n + p.idx
	// fits cannot be (blockSize - 1)+
	fits := blockSize - lbIdx

	// fits is only ever zero if there is only blocksize of bytes left in bs
	if fits == 0 {
		p.Block()
		p.idx = 0
		return
	}

	if fits > 0 {
		p.idx = lbIdx
		return
	}

	rBlkLen := blockSize - p.idx

	p.Block()
	p.start = p.idx
	p.idx = 0

	n -= rBlkLen

	for n >= blockSize {
		p.start += blockSize
		p.Block()
		n -= blockSize
	}

	// need to copy the reest to the begining
	copy(p._buf[0:blockSize], p._buf[p.start:])

	p.idx = n
	p.start = 0
}

func (p *process) Byte(b byte) {

	p._buf[p.start+p.idx] = b
	p.idx++
	if p.idx == blockSize {
		p.Block()
		p.idx = 0
		p.start = 0
	}
}

func (p *process) Block() {

	var (
		M [16]uint32

		// Initialize hash value for this chunk:
		A    = p.a0
		B    = p.b0
		C    = p.c0
		D    = p.d0
		_buf = p._buf[p.start:]
	)

	// break chunk into sixteen 32-bit words M[j], 0 ≤ j ≤ 15

	for i := range len(M) {
		j := i * 4
		M[i] = uint32(_buf[j]) | uint32(_buf[j+1])<<8 |
			uint32(_buf[j+2])<<16 | uint32(_buf[j+3])<<24
	}

	// Round 1
	for ii := range 16 {
		F := (B & C) | ((^B) & D)
		F = F + A + k[ii] + M[ii]
		A = D
		D = C
		C = B
		B += bits.RotateLeft32(F, int(s[ii]))
	}

	// Round 2
	for ii := 16; ii < 32; ii++ {
		g := uint32((5*ii + 1) & 15)
		F := (D & B) | ((^D) & C)
		F = F + A + k[ii] + M[g]
		A = D
		D = C
		C = B
		B += bits.RotateLeft32(F, int(s[ii]))
	}

	// Round 3
	for ii := 32; ii < 48; ii++ {
		g := uint32((3*ii + 5) & 15)
		F := B ^ C ^ D
		F = F + A + k[ii] + M[g]
		A = D
		D = C
		C = B
		B += bits.RotateLeft32(F, int(s[ii]))
	}

	// Round 4
	for ii := 48; ii < 64; ii++ {
		g := uint32((7 * ii) & 15)
		F := C ^ (B | ^D)
		F = F + A + k[ii] + M[g]
		A = D
		D = C
		C = B
		B += bits.RotateLeft32(F, int(s[ii]))
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
	zeros = make([]byte, 2*blockSize)
)

const (
	a0 = uint32(0x67452301)
	b0 = uint32(0xefcdab89)
	c0 = uint32(0x98badcfe)
	d0 = uint32(0x10325476)

	blockSize = 64
)
