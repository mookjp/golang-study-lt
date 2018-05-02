package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"io"
	"os"
)

type mode string

const (
	SHA254 mode = "sha254"
	SHA384 mode = "sha384"
	SHA512 mode = "sha512"
)

// TODO typeとして定義されているかを検証するにはどうすればよいか？
func main() {
	var m mode
	if len(os.Args) > 1 {
		opt := os.Args[1]
		fmt.Fprintf(os.Stdout, "option: %v\n", opt)
		m = mode(opt)
	} else {
		m = mode("sha254")
	}
	print("Input characters to convert to SHA hash, then type Enter-key:\n> ")
	printSha(os.Stdin, os.Stdout, m)
}

func printSha(r io.Reader, w io.Writer, m mode) error {
	reader := bufio.NewReader(r)
	input, isPrefix, err := reader.ReadLine()
	if err != nil {
		return fmt.Errorf("error occurred: %v", err)
	}
	if isPrefix == true {
		return errors.New("input was too long")
	}
	var out []byte
	switch m {
	case SHA384:
		out = calcSha384(input)
	case SHA512:
		out = calcSha512(input)
	default: // Including SHA254
		out = calcSha256(input)
	}
	fmt.Fprintf(w, "%x", out)
	return nil
}

func calcSha256(bytes []byte) []byte {
	res := sha256.Sum256(bytes)
	return res[:]
}

func calcSha384(bytes []byte) []byte {
	res := sha512.Sum384(bytes)
	return res[:]
}

func calcSha512(bytes []byte) []byte {
	res := sha512.Sum512(bytes)
	return res[:]
}
