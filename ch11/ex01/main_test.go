package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

type pair struct {
	key   string
	value string
}

func TestNorMalScenario(t *testing.T) {
	var tests = []struct {
		input                string
		expectedCharCount    map[string]string
		expectedUnicodeCount map[string]string
	}{
		{"abc",
			map[string]string{"'a'": "1", "'b'": "1", "'c'": "1"},
			map[string]string{"1": "3", "2": "0", "3": "0", "4": "0"}},
		{"ああいううえおおお",
			map[string]string{"'あ'": "2", "'い'": "1", "'う'": "2", "'え'": "1", "'お'": "3"},
			map[string]string{"1": "0", "2": "0", "3": "9", "4": "0"}},
		{"",
			map[string]string{},
			map[string]string{"1": "0", "2": "0", "3": "0", "4": "0"}},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expectedCharCount: %v, expectedUnicodeCount: %v\n",
			test.input, test.expectedCharCount, test.expectedUnicodeCount)

		reader := strings.NewReader(test.input)
		writer := &bytes.Buffer{}
		errWriter := &bytes.Buffer{}

		// Execution
		charCount(reader, writer, errWriter)

		headerScanner := bufio.NewScanner(writer)

		// "rune	count" ヘッダの検証
		if headerScanner.Scan() {
			if headerScanner.Text() != "rune\tcount" {
				t.Errorf("it doesnt show expected header")
			}
		} else {
			t.Errorf("it doesnt show expected header")
		}

		// rune, count の値
		for headerScanner.Scan() {
			line := headerScanner.Text()
			if line == "" {
				break
			}
			splitted := strings.Split(line, "\t")
			charVal := test.expectedCharCount[splitted[0]]
			if charVal == "" {
				t.Errorf("target char was not in returns. target line: %s", line)
			}
			if line != fmt.Sprintf("%s\t%s", splitted[0], charVal) {
				t.Errorf("charPair was '%v\t%v', actual was %v", splitted[0], charVal, line)
			}
			delete(test.expectedCharCount, splitted[0])
		}
		if len(test.expectedCharCount) != 0 {
			t.Errorf("actual lines are not enough")
		}

		// validate len count header
		if headerScanner.Scan() {
			header := headerScanner.Text()
			if header != "len\tcount" {
				t.Errorf("it doesnt show expected header: %s, actual: %s", "len\tcount", header)
			}
		} else {
			t.Errorf("it doesnt show expected header")
		}
		// len count
		for headerScanner.Scan() {
			line := headerScanner.Text()
			splitted := strings.Split(line, "\t")
			charVal := test.expectedUnicodeCount[splitted[0]]
			if charVal == "" {
				t.Errorf("target char was not in returns. target line: %s", line)
			}
			if line != fmt.Sprintf("%s\t%s", splitted[0], charVal) {
				t.Errorf("charPair was '%v\t%v', actual was %v", splitted[0], charVal, line)
			}
			delete(test.expectedUnicodeCount, splitted[0])
		}
		if len(test.expectedUnicodeCount) != 0 {
			t.Errorf("actual lines are not enough")
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
