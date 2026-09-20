package md5

import (
	"bytes"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
	"testing"

	_ "embed"
)

var (
	//go:embed testdata/empty.txt
	benchmarkEmptyTxt []byte

	//go:embed testdata/hello.txt
	benchmarkHelloTxt []byte

	//go:embed testdata/random_4kb.bin
	benchmarkRandom4KBBin []byte
)

func BenchmarkHash(b *testing.B) {
	type tcase struct {
		Filename string
		data     []byte
	}

	tests := []tcase{
		{data: benchmarkEmptyTxt, Filename: "empty.txt"},
		{data: benchmarkHelloTxt, Filename: "hello.txt"},
		{data: benchmarkRandom4KBBin, Filename: "random_4kb.bin"},
	}

	for _, tc := range tests {
		b.Run(tc.Filename, func(b *testing.B) {

			buf := make([]byte, blockSize*4*1024)
			r := bytes.NewReader(tc.data)

			b.ResetTimer()
			for b.Loop() {
				Hash(r, &Options{Buf: buf}) // ← In-memory reader
			}
		})
	}
}

func TestHash(t *testing.T) {

	type tcase struct {
		Code     string
		Filename string
	}

	fn := func(tc tcase) func(*testing.T) {
		fd, err := os.Open(filepath.Join("testdata", tc.Filename))
		if err != nil {
			log.Fatalf("Failed to open test file: %v, %v", tc.Filename, err)
			return nil
		}
		return func(t *testing.T) {
			defer fd.Close()
			code, err := Hash(fd, nil)
			if err != nil {
				t.Errorf("error, expected nil, got %v", err)
				return
			}
			t.Logf("Got Code: '%v'", code)
			gotCode := hex.EncodeToString(code[:])
			if gotCode != tc.Code {
				t.Errorf("code, expected '%v' got '%v'", tc.Code, code)
				return
			}
		}
	}

	tests := []tcase{
		{Code: "d41d8cd98f00b204e9800998ecf8427e", Filename: "empty.txt"},
		{Code: "b1946ac92492d2347c6235b4d2611184", Filename: "hello.txt"},
		{Code: "9f40725ceb305d49f294c89fd36a808a", Filename: "random_4kb.bin"},
	}

	for _, tc := range tests {
		t.Run(tc.Filename, fn(tc))
	}
}
