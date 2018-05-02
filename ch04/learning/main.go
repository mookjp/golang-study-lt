package main

import (
	"fmt"
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
}
