// charcountはUnicode文字の数を計算します。
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
	counts := make(map[rune]int)    // Unicode文字の数
	types := make(map[string]int)   // Unicode文字種別の文字の数
	var utflen [utf8.UTFMax + 1]int // UTF-8エンコーディングの長さの数
	invalid := 0                    // 不正なUTF-8文字の数

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // rune, nbytes(rune1件に対するbyte数), error を返す
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		// 文字種
		if unicode.IsLetter(r) {
			types["Letter"]++
		}
		if unicode.IsGraphic(r) {
			types["Graphic"]++
		}
		if unicode.IsMark(r) {
			types["Mark"]++
		}
		if unicode.IsNumber(r) {
			types["Number"]++
		}
		if unicode.IsDigit(r) {
			types["Digit"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts { // TODO: cはrune的な名前にしたい
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ntypes\tcount\n")
	for t, n := range types {
		fmt.Printf("%s\t%d\n", t, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
