package ex02_test

import (
	"fmt"
	"os"
	"testing"

	"bytes"

	"github.com/mookjp/golang-study-lt/ch07/ex02"
)

func Test(t *testing.T) {
	var tests = []struct {
		input    []byte
		expected int64
	}{
		{[]byte("dog"), 3},
		{[]byte("software"), 8},
		{[]byte(""), 0},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n", test.input, test.expected)
		byteWriter := &bytes.Buffer{}
		countingWriter, count := ex02.CountingWriter(byteWriter)
		countingWriter.Write(test.input)

		if *count != test.expected {
			t.Errorf("assertion error. input = %v, actual = %v, expected = %v",
				test.input, count, test.expected)
		}
	}
}
