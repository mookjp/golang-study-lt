package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		fmt.Printf("index: %v, value: %v\n", index, arg)
	}
}

