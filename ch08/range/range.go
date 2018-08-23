package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	ch := make(chan string)
	go func() {
		t := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-t.C:
				ch <- "hello"
			}
		}
		t.Stop()
	}()

	for message := range ch {
		fmt.Fprintf(os.Stdout, "message: %s\n", message)
	}

}
