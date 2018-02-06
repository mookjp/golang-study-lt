package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
)

const prefix string = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			fmt.Fprintf(os.Stdout, "url `%s` has no prefix. addding prefix %s...\n", url, prefix)
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout,"status: %s\n", resp.Status)

		// https://golang.org/doc/effective_go.html#if
		// statement initialization ができる
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}