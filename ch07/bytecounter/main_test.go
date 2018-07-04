package bytecounter_test

import (
	"fmt"
	"os"
	"testing"

	"bytes"

	"github.com/mookjp/golang-study-lt/ch07/bytecounter"
)

func Test(t *testing.T) {
	var tests = []struct {
		input    []byte
		expected string
	}{
		{[]byte("a"), "1"},
		{[]byte("abc"), "3"},
		{[]byte(""), "0"},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n", test.input, test.expected)
		var c bytecounter.ByteCounter
		c.Write(test.input)
		actual := &bytes.Buffer{}

		fmt.Fprint(actual, c)
		if actual.String() != test.expected {
			t.Errorf("assertion error. input: = %v, expected = %v",
				test.input, test.expected)
		}
	}
}
