package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"

	md5 "github.com/gdey/md5-go"
)

func main() {
	var (
		debug uint8
		fd    = os.Stdin
		err   error
	)

	flag.BoolFunc("debug", "set the amount of debug message to log", func(s string) error {
		if s == "true" {
			debug++
			return nil
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		if v+int(debug) > 255 {
			debug = 255
		} else {
			debug += uint8(v)
		}
		return nil
	})

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
		Debug: md5.INFO,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to hash file: %s\n\t%v\n", filename, err)
		os.Exit(2)
	}

	fmt.Println(hex.EncodeToString(code[:]), filename)
}
