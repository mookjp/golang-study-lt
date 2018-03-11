package main

import (
	"os"
	"strconv"
	"fmt"
	"bufio"

	"../tempconv"
)

// TODO: その他の変換

func main() {
	println("main")
	fmt.Printf("args len: %v\n", len(os.Args))
	if len(os.Args) < 2 {
		fmt.Print("Enter value or exit to type 'exit': ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			if text == "exit" {
				os.Exit(0)
			}
			v := parseToFloat64(text)
			conv(v)
			fmt.Print("Enter value or exit to type 'exit': ")
		}
	} else {
		for _, arg := range os.Args[1:] {
			v := parseToFloat64(arg)
			conv(v)
		}
	}
}

func parseToFloat64(value string) float64 {
	t, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return t
}

func conv(value float64) {
	f := tempconv.Fahrenheit(value)
	c := tempconv.Celsius(value)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}
