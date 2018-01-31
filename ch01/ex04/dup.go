package main

import (
	"os"
	"bufio"
	"fmt"
)

// TODO: refactor
func main() {
	counts_by_file := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts)
		counts_by_file["no_file"] = counts
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			counts_by_file[arg] = counts
			f.Close()
		}
	}

	// Print the result
	for file_name, counts := range counts_by_file {
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", file_name, n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
