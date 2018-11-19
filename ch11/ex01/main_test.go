package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestNorMalScenario(t *testing.T) {
	var tests = []struct {
		input                string
		expectedCharCount    map[string]int
		expectedUnicodeCount map[int]int
	}{
		{"abc", map[string]int{"a": 1, "b": 1, "c": 1}, map[int]int{1: 3, 2: 0, 3: 0, 4: 0}},
		{"ああいううえおおお", map[string]int{"あ": 2, "い": 1, "う": 2, "え": 1, "お": 3}, map[int]int{1: 0, 2: 0, 3: 9, 4: 0}},
		{"", map[string]int{}, map[int]int{1: 0, 2: 0, 3: 0, 4: 0}},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expectedCharCount: %v, expectedUnicodeCount: %v\n",
			test.input, test.expectedCharCount, test.expectedUnicodeCount)

		reader := strings.NewReader(test.input)
		writer := &bytes.Buffer{}
		errWriter := &bytes.Buffer{}

		// Execution
		charCount(reader, writer, errWriter)

		// "rune	count" ヘッダの検証
		scanCounter := 0
		headerScanner := bufio.NewScanner(writer)
		for headerScanner.Scan() {
			line := headerScanner.Text()
			if scanCounter == 0 {
				if line != "rune\tcount" {
					t.Errorf("header line was incorrect")
				}
				scanCounter++
				continue
			}
			// rune, countの表示
			for charKey, charCount := range test.expectedCharCount {
				if line != "'"+charKey+"'\t"+strconv.Itoa(charCount) {
					t.Errorf("charKey was %v", charKey)
				}
			}
			// TODO: len, countの表示
			if line == "" {
				for headerScanner.Scan() {

				}
			}
		}
	}
}

/**
[git][* master]:~/go/src/github.com/mookjp/golang-study-lt/ch11/ex01/ go run main.go                                  [18-11-19 20:44]
abc

rune    count
'a'     1
'b'     1
'c'     1
'\n'    2

len     count
1       5
2       0
3       0
4       0
[git][* master]:~/go/src/github.com/mookjp/golang-study-lt/ch11/ex01/ go run main.go                                  [18-11-19 21:00]
rune    count

len     count
1       0
2       0
3       0
4       0
*/
