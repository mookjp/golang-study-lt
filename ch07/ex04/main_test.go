package ex04_test

import (
	"fmt"
	"os"
	"testing"

	"bytes"

	"strings"

	"github.com/mookjp/golang-study-lt/ch07/ex04"
)

func Test(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"<a href=\"localhost\">link1</a>", "localhost"},
		{"<a href=\"localhost\">link1</a><a href=\"github.com\">github</a>", "localhostgithub.com"},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n", test.input, test.expected)
		byteWriter := &bytes.Buffer{}

		// execute
		ex04.FromReader(strings.NewReader(test.input), byteWriter)

		// actual
		actual := byteWriter.String()

		if actual != test.expected {
			t.Errorf("assertion error. input = %v, actual = %v, expected = %v",
				test.input, actual, test.expected)
		}
	}
}
