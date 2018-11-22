// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	charCount(os.Stdin, os.Stdout, os.Stderr)
}

// テストのためにcharCountを分離
// 分離してよかったかはわからない…
func charCount(reader io.Reader, writer io.Writer, errWriter io.Writer) {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	in := bufio.NewReader(reader)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(errWriter, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Fprintf(writer, "rune\tcount\n")
	for c, n := range counts {
		fmt.Fprintf(writer, "%q\t%d\n", c, n)
	}
	fmt.Fprintf(writer, "\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Fprintf(writer, "%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Fprintf(writer, "\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
