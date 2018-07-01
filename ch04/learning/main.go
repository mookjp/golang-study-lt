package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	array := [12]byte{}
	slice := []byte{}
	fmt.Fprintf(os.Stdout, "cap(slice): %v, len(slice): %v\n", cap(slice), len(slice))
	fmt.Fprintf(os.Stdout, "cap(array): %v, len(array): %v\n", cap(array), len(array))
	extended := array[:10]
	fmt.Fprintf(os.Stdout, "cap(extended): %v, len(extended): %v\n", cap(extended), len(extended))

	oneToFive := [5]int{1, 2, 3, 4, 5}
	sliceOfOneToFive := oneToFive[:]
	fmt.Fprintf(os.Stdout, "val: %v, cap(oneToFive): %v, len(oneToFive): %v\n", oneToFive, cap(oneToFive), len(oneToFive))
	fmt.Fprintf(os.Stdout, "val: %v, cap(sliceOfOneToFive): %v, len(sliceOfOneToFive): %v\n", sliceOfOneToFive, cap(sliceOfOneToFive), len(sliceOfOneToFive))
	sliceOfOneToFive[1] = 100
	fmt.Println("after sliceOfOneToFive[1] = 100")
	fmt.Fprintf(os.Stdout, "val: %v, cap(oneToFive): %v, len(oneToFive): %v\n", oneToFive, cap(oneToFive), len(oneToFive))
	fmt.Fprintf(os.Stdout, "val: %v, cap(sliceOfOneToFive): %v, len(sliceOfOneToFive): %v\n", sliceOfOneToFive, cap(sliceOfOneToFive), len(sliceOfOneToFive))

	made := make([]int, 3)
	fmt.Fprintf(os.Stdout, "val: %v, cap(made): %v, len(made): %v\n", made, cap(made), len(made))
	intS := []int{}
	fmt.Fprintf(os.Stdout, "val: %v, cap(intS): %v, len(intS): %v\n", intS, cap(intS), len(intS))

	appended := []rune("こんにちは")
	fmt.Fprintf(os.Stdout, "val: %v, cap(rune[]): %v, len(rune[]): %v\n", append(appended), cap(appended), len(appended))

	multiBytes := []byte("こんにちは")
	fmt.Fprintf(os.Stdout, "multiBytes: %v\n", multiBytes)

	orig := []byte("abcde")
	fmt.Fprintf(os.Stdout, "orig: %v\n", orig)
	copied := make([]byte, len(orig))
	copy(copied, orig)
	fmt.Fprintf(os.Stdout, "copied: %v\n", copied)
	orig[0] = 100
	fmt.Fprintf(os.Stdout, "after updated orig: %v\n", orig)
	fmt.Fprintf(os.Stdout, "after updated copied: %v\n", copied)

	myMap := make(map[float64]int, 0)
	myMap[math.NaN()] = 10

	type a struct {
		name string
		names string[]
	}

	aa := a{"john", []string{}}
	aaa := a{"john", []string{}}

	fmt.Printf("is equal: %v", aa == aaa)
}
