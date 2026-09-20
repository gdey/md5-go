//go:build !unrolled

package md5

import "math/bits"

//go:generate go run generators/unroll_block_function.go

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
