package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	md5 "github.com/gdey/md5-go"
)

func main() {
	var (
		fd      = os.Stdin
		err     error
		memMult int
	)

	flag.IntVar(&memMult, "memory", 4, "multiple of 1024 to allocate for memory")

	flag.Parse()

	filename := flag.Arg(0)

	if filename != "" && filename != "-" {
		fd, err = os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Got error trying to open file: %s\n\t%v\n", filename, err)
			os.Exit(1)
			return
		}
		defer fd.Close()
	} else {
		filename = "<<StandardIn>>"
	}

	code, err := md5.Hash(fd, &md5.Options{
		Buf: make([]byte, 64*memMult*1024),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to hash file: %s\n\t%v\n", filename, err)
		os.Exit(2)
	}

	fmt.Println(hex.EncodeToString(code[:]), filename)
}
