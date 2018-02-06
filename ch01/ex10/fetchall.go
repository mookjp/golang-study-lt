package main

import (
	"time"
	"os"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"path"
)

func main() {
	dirPath, err := getOutputDirPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "parameter error: %s", err)
		os.Exit(1)
	}

	outputFile, err := createOutputFile(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create file, error: %s", err)
		os.Exit(1)
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[2:] {
		go fetch(url, ch) // start goroutine
	}
	for range os.Args[2:] {
		res := <- ch
		fmt.Println(res)
		outputFile.WriteString(fmt.Sprintf("%s\n", res))
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Printf("The result was written to `%s`", outputFile.Name())
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


// TODO: how to write returning error
func getOutputDirPath() (string, error) {
	outputDirPath := os.Args[1]
	_, err := os.Stat(outputDirPath)
	if err != nil {
		return "", err // TODO: これでいいのか？
	} else {
		return outputDirPath, nil
	}
}

func createOutputFile(dirPath string) (*os.File, error) {
	fileName := fmt.Sprintf("output_%s.txt", time.Now().Format("2006-01-02T15:04:05"))
	filePath := path.Join(dirPath, fileName)
	outputFile, err := os.Create(filePath)
	if err != nil {
		return nil, err
	} else {
		return outputFile, nil
	}
}
