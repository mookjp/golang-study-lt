package main

import "fmt"

func main() {
	multiByteText := "こんにちは！123"
	fmt.Println("[0]:", multiByteText[0])
	fmt.Println("[:3]:", multiByteText[:3])
	fmt.Println("[len(multiByteText) - 1:]:", multiByteText[len(multiByteText)-1:])
	fmt.Println("[len(multiByteText) - 1:]", multiByteText[len(multiByteText)-1])
	fmt.Println("[:]:", multiByteText[:])

	a := "aaa"
	b := a
	a += "aaa"
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
