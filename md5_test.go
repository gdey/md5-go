package md5

import (
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkHash(b *testing.B) {
	type tcase struct {
		Code     string
		Filename string
	}

	tests := []tcase{
		{Code: "d41d8cd98f00b204e9800998ecf8427e", Filename: "empty.txt"},
		{Code: "b1946ac92492d2347c6235b4d2611184", Filename: "hello.txt"},
		{Code: "9f40725ceb305d49f294c89fd36a808a", Filename: "random_4kb.bin"},
	}

	for _, tc := range tests {
		b.Run(tc.Filename, func(b *testing.B) {
			var filename string
			if tc.Filename[0] != '/' {
				filename = filepath.Join("testdata", tc.Filename)
			} else {
				filename = tc.Filename
			}
			fd, err := os.Open(filename)
			if err != nil {
				b.Fatalf("Failed to open test file: %v", err)
			}
			defer fd.Close()
			buf := make([]byte, blockSize*4*1024)

			b.ResetTimer()
			for b.Loop() {
				fd.Seek(0, 0)
				Hash(fd, &Options{
					Buf: buf,
				})
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
