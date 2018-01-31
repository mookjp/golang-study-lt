package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	non_efficient_start := time.Now()
	non_efficient_print()
	non_efficient_end := time.Now()
	fmt.Printf("Non Efficient ver time: %v nanosec\n", non_efficient_end.Sub(non_efficient_start).Nanoseconds())

	efficient_start := time.Now()
	non_efficient_print()
	efficient_end := time.Now()
	fmt.Printf("Non Efficient ver time: %v nanosec\n", efficient_end.Sub(efficient_start).Nanoseconds())
}

func non_efficient_print() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
func efficient_print() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
