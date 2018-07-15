package ex01

import (
	"bufio"
	"bytes"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	var counter int
	for scanner.Scan() {
		counter++
	}
	*c += WordCounter(counter)
	return counter, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	var counter int
	for scanner.Scan() {
		counter++
	}
	*c += LineCounter(counter)
	return counter, nil
}
