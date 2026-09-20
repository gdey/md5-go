//go:build unrolled

// AUTO-GENERATED: Full unrolled Block() function

package md5

import (
	"math/bits"
)

func (p *process) Block() {
	var (
		M [16]uint32

		// Initialize hash value for this chunk:
		A    = p.a0
		B    = p.b0
		C    = p.c0
		D    = p.d0
		F    uint32
		_buf = p._buf[p.start:]
	)

	// break chunk into sixteen 32-bit words M[j], 0 ≤ j ≤ 15

	for i := range len(M) {
		j := i * 4
		M[i] = uint32(_buf[j]) | uint32(_buf[j+1])<<8 |
			uint32(_buf[j+2])<<16 | uint32(_buf[j+3])<<24
	}
	// Round 1 - F rounds (ii = 0..15)

	F = (B & C) | ((^B) & D)
	F = F + A + k[0] + M[0]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[0]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[1] + M[1]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[1]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[2] + M[2]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[2]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[3] + M[3]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[3]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[4] + M[4]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[4]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[5] + M[5]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[5]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[6] + M[6]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[6]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[7] + M[7]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[7]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[8] + M[8]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[8]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[9] + M[9]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[9]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[10] + M[10]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[10]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[11] + M[11]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[11]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[12] + M[12]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[12]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[13] + M[13]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[13]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[14] + M[14]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[14]))

	F = (B & C) | ((^B) & D)
	F = F + A + k[15] + M[15]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[15]))
	// Round 2 - G rounds (ii = 16..31)

	F = (B & D) | (C & (^D))
	F = F + A + k[16] + M[1]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[16]))

	F = (B & D) | (C & (^D))
	F = F + A + k[17] + M[6]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[17]))

	F = (B & D) | (C & (^D))
	F = F + A + k[18] + M[11]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[18]))

	F = (B & D) | (C & (^D))
	F = F + A + k[19] + M[0]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[19]))

	F = (B & D) | (C & (^D))
	F = F + A + k[20] + M[5]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[20]))

	F = (B & D) | (C & (^D))
	F = F + A + k[21] + M[10]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[21]))

	F = (B & D) | (C & (^D))
	F = F + A + k[22] + M[15]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[22]))

	F = (B & D) | (C & (^D))
	F = F + A + k[23] + M[4]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[23]))

	F = (B & D) | (C & (^D))
	F = F + A + k[24] + M[9]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[24]))

	F = (B & D) | (C & (^D))
	F = F + A + k[25] + M[14]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[25]))

	F = (B & D) | (C & (^D))
	F = F + A + k[26] + M[3]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[26]))

	F = (B & D) | (C & (^D))
	F = F + A + k[27] + M[8]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[27]))

	F = (B & D) | (C & (^D))
	F = F + A + k[28] + M[13]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[28]))

	F = (B & D) | (C & (^D))
	F = F + A + k[29] + M[2]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[29]))

	F = (B & D) | (C & (^D))
	F = F + A + k[30] + M[7]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[30]))

	F = (B & D) | (C & (^D))
	F = F + A + k[31] + M[12]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[31]))
	// Round 3 - H rounds (ii = 32..47)

	F = B ^ C ^ D
	F = F + A + k[32] + M[5]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[32]))

	F = B ^ C ^ D
	F = F + A + k[33] + M[8]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[33]))

	F = B ^ C ^ D
	F = F + A + k[34] + M[11]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[34]))

	F = B ^ C ^ D
	F = F + A + k[35] + M[14]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[35]))

	F = B ^ C ^ D
	F = F + A + k[36] + M[1]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[36]))

	F = B ^ C ^ D
	F = F + A + k[37] + M[4]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[37]))

	F = B ^ C ^ D
	F = F + A + k[38] + M[7]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[38]))

	F = B ^ C ^ D
	F = F + A + k[39] + M[10]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[39]))

	F = B ^ C ^ D
	F = F + A + k[40] + M[13]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[40]))

	F = B ^ C ^ D
	F = F + A + k[41] + M[0]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[41]))

	F = B ^ C ^ D
	F = F + A + k[42] + M[3]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[42]))

	F = B ^ C ^ D
	F = F + A + k[43] + M[6]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[43]))

	F = B ^ C ^ D
	F = F + A + k[44] + M[9]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[44]))

	F = B ^ C ^ D
	F = F + A + k[45] + M[12]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[45]))

	F = B ^ C ^ D
	F = F + A + k[46] + M[15]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[46]))

	F = B ^ C ^ D
	F = F + A + k[47] + M[2]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[47]))
	// Round 4 - I rounds (ii = 48..63)

	F = C ^ (B | (^D))
	F = F + A + k[48] + M[0]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[48]))

	F = C ^ (B | (^D))
	F = F + A + k[49] + M[7]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[49]))

	F = C ^ (B | (^D))
	F = F + A + k[50] + M[14]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[50]))

	F = C ^ (B | (^D))
	F = F + A + k[51] + M[5]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[51]))

	F = C ^ (B | (^D))
	F = F + A + k[52] + M[12]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[52]))

	F = C ^ (B | (^D))
	F = F + A + k[53] + M[3]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[53]))

	F = C ^ (B | (^D))
	F = F + A + k[54] + M[10]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[54]))

	F = C ^ (B | (^D))
	F = F + A + k[55] + M[1]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[55]))

	F = C ^ (B | (^D))
	F = F + A + k[56] + M[8]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[56]))

	F = C ^ (B | (^D))
	F = F + A + k[57] + M[15]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[57]))

	F = C ^ (B | (^D))
	F = F + A + k[58] + M[6]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[58]))

	F = C ^ (B | (^D))
	F = F + A + k[59] + M[13]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[59]))

	F = C ^ (B | (^D))
	F = F + A + k[60] + M[4]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[60]))

	F = C ^ (B | (^D))
	F = F + A + k[61] + M[11]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[61]))

	F = C ^ (B | (^D))
	F = F + A + k[62] + M[2]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[62]))

	F = C ^ (B | (^D))
	F = F + A + k[63] + M[9]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[63]))

	p.a0 += A
	p.b0 += B
	p.c0 += C
	p.d0 += D
}
