//go:build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
)

var (
	filename    = flag.String("filename", "process_block.generated.go", "filename to write output to; '-' for standard out")
	packageName = flag.String("package", "md5", "package name")
)

func main() {

	flag.Parse()
	var (
		err error
		src = generateSource()
	)

	if *filename == "-" {
		fmt.Println(string(src))
		return
	}
	err = os.WriteFile(*filename, src, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write file %s: %v", *filename, err)
		os.Exit(2)
	}
}

func generateSource() []byte {
	var (
		out bytes.Buffer
		// Define the four round types
		rounds = []struct {
			name    string
			start   int
			count   int
			formula string
			indexFn func(int) int
		}{
			{
				name:    "Round 1 - F rounds (ii = 0..15)",
				start:   0,
				count:   16,
				formula: "(B & C) | ((^B) & D)",
				indexFn: func(ii int) int { return ii },
			},
			{
				name:    "Round 2 - G rounds (ii = 16..31)",
				start:   16,
				count:   16,
				formula: "(B & D) | (C & (^D))",
				indexFn: func(ii int) int { return (5*ii + 1) % 16 },
			},
			{
				name:    "Round 3 - H rounds (ii = 32..47)",
				start:   32,
				count:   16,
				formula: "B ^ C ^ D",
				indexFn: func(ii int) int { return (3*ii + 5) % 16 },
			},
			{
				name:    "Round 4 - I rounds (ii = 48..63)",
				start:   48,
				count:   16,
				formula: "C ^ (B | (^D))",
				indexFn: func(ii int) int { return (7 * ii) % 16 },
			},
		}
	)

	fmt.Fprintf(&out, preambleFmt, *packageName)
	for _, round := range rounds {
		fmt.Fprintf(&out, "\t// %s\n", round.name)
		for i := range round.count {
			ii := round.start + i
			mIdx := round.indexFn(ii)
			fmt.Fprintf(&out, roundFmt, round.formula, ii, mIdx)
		}
	}
	fmt.Fprintf(&out, postFmt)

	// Format the code
	formattedCode, err := format.Source(out.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to format code with error: %v", err)
		os.Exit(3)
		return nil
	}

	return formattedCode
}

const (
	preambleFmt = `//go:build unrolled

// AUTO-GENERATED: Full unrolled Block() function

package %s;

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
`
	roundFmt = `
	F = %[1]s
	F = F + A + k[%[2]d] + M[%[3]d]
	A = D
	D = C
	C = B
	B += bits.RotateLeft32(F, int(s[%[2]d]))
`
	postFmt = `
	p.a0 += A
	p.b0 += B
	p.c0 += C
	p.d0 += D
}
`
)
