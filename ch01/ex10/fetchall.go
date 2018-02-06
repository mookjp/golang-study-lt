package main

import (
	"time"
	"os"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from ch channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send err to ch channel
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // ioutil.Discard does nothing
	resp.Body.Close() // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // send err to ch channel
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}