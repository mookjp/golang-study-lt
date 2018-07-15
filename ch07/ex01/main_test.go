package ex01_test

import (
	"fmt"
	"os"
	"testing"

	"bytes"

	"github.com/mookjp/golang-study-lt/ch07/ex01"
)

func TestWordCounter(t *testing.T) {
	var tests = []struct {
		input    []byte
		expected string
	}{
		{[]byte("a dog"), "2"},
		{[]byte("my name is John"), "4"},
		{[]byte(""), "0"},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n", test.input, test.expected)
		var c ex01.WordCounter
		c.Write(test.input)
		actual := &bytes.Buffer{}

		fmt.Fprint(actual, c)
		if actual.String() != test.expected {
			t.Errorf("assertion error. input = %v, actual = %v, expected = %v",
				test.input, actual.String(), test.expected)
		}
	}
}

func TestLineCounter(t *testing.T) {
	var tests = []struct {
		input    []byte
		expected string
	}{
		{[]byte("The dog is mine.\n That dog is not mine."), "2"},
		{[]byte("Hello,\nmy name is John.\nI'm software engineer."), "3"},
		{[]byte("\n\n"), "2"},
		{[]byte(""), "0"},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n", test.input, test.expected)
		var c ex01.LineCounter
		c.Write(test.input)
		actual := &bytes.Buffer{}

		fmt.Fprint(actual, c)
		if actual.String() != test.expected {
			t.Errorf("assertion error. input = %v, actual = %v, expected = %v",
				test.input, actual.String(), test.expected)
		}
	}
}
