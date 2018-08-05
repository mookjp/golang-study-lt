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

	"net/http"

	"bytes"

	"regexp"

	"path/filepath"

	"gopl.io/ch5/links"
)

const outputDir string = "/output"

func getOutputBase() string {
	baseDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return baseDir + outputDir
}

func crawl(url string) []string {
	fmt.Fprintf(os.Stdout, "Getting %s...\n", url)
	save(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func save(url string) {
	body, err := getPageBody(url)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error while getting the page: %v, skipping to save the page as a file...\n", err)
		return
	}
	formatted := formatUrl(body)
	fileNameRe := regexp.MustCompile(`^(http://|https://)`)
	filename := getOutputBase() + string(fileNameRe.ReplaceAll([]byte(url), []byte("/")))
	dirNameRe := regexp.MustCompile("[^/]+?$")
	dirname := string(dirNameRe.ReplaceAll([]byte(filename), []byte("")))
	fmt.Fprintf(os.Stdout, "=========== filename: %s, dirname: %s\n", filename, dirname)

	// ファイルを保持するディレクトリがなかったら作成
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		fmt.Fprintf(os.Stdout, "=========== creating dir: %s\n", dirname)
		os.Mkdir(dirname, 0700)
	}
	// ファイル作成
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error while creating file: %v, skipping to save the page as a file...\n", err)
		return
	}
	file.WriteString(formatted)
}

func getPageBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String(), nil
}

func formatUrl(body string) string {
	re := regexp.MustCompile(`href="(http://|https://)`)
	return string(re.ReplaceAll([]byte(body), []byte("href=\"file://")))
}

type workList struct {
	url   string
	depth int
}

//!+
func main() {
	targetSite := flag.String("target", "https://gopl.io", "target to crawl")
	depthLimit := flag.Int("depth", 3, "depth to crawl")
	flag.Parse()
	fmt.Fprintf(os.Stdout, "========================= depthLimit: %v\n", *depthLimit)
	fmt.Fprintf(os.Stdout, "========================= target: %v\n", *targetSite)

	// outputディレクトリの作成
	outputDir := getOutputBase()
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stdout, "Creating output dir %s ...", outputDir)
		os.Mkdir(outputDir, 0700)
	}

	worklist := make(chan []workList)  // lists of URLs, may have duplicates
	unseenLinks := make(chan workList) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {

		var wls []workList
		wls = append(wls, workList{*targetSite, 1})
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
