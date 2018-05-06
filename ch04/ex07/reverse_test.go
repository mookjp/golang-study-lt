package rev

import (
	"fmt"
	"os"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		input    []byte
		expected []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte("a"), []byte("a")},
		{[]byte("ab"), []byte("ba")},
		{[]byte("abc"), []byte("cba")},
		{[]byte("abcd"), []byte("dcba")},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n", test.input, test.expected)
		reverse(test.input)
		fmt.Fprintf(os.Stdout, "after reverse input: %v\n", test.input)
		for i, v := range test.input {
			if v != test.expected[i] {
				t.Errorf("assertion error. input: input[%v] = %v, expected[%v] = %v", i, v, i, test.expected[i])
			}
		}
	}
}
