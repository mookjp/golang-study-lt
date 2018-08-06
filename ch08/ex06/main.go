// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"fmt"
	"log"
	"os"

	"flag"

	"gopl.io/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

type workList struct {
	url   string
	depth int
}

//!+
func main() {
	depthLimit := flag.Int("depth", 3, "depth to crawl")
	flag.Parse()
	fmt.Fprintf(os.Stdout, "========================= depthLimit: %v\n", *depthLimit)
	worklist := make(chan []workList)  // lists of URLs, may have duplicates
	unseenLinks := make(chan workList) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {

		var wls []workList
		for _, url := range os.Args[1:] {
			wl := workList{url, 1}
			wls = append(wls, wl)
		}
		worklist <- wls
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				fmt.Fprintf(os.Stdout, "========================= link: %v\n", link)
				newDepth := link.depth + 1
				foundLinks := crawl(link.url)
				if newDepth <= *depthLimit {
					var targets []workList
					for _, foundLink := range foundLinks {
						targets = append(targets, workList{foundLink, newDepth})
					}
					go func() { worklist <- targets }()

				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

//!-
