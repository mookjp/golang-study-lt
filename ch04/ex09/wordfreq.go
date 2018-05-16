package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordMap := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		word := in.Text()
		wordMap[word]++
	}

	fmt.Print("\nword\tcount\n")

	for w, c := range wordMap {
		fmt.Printf("%s\t%d\n", w, c)
	}
}
