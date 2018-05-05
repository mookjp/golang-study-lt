package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	var tests = []struct {
		input    []byte
		expected []byte
	}{
		{[]byte{}, []byte{}},
		{[]byte("  "), []byte(" ")},
		{[]byte("   "), []byte(" ")},
		{[]byte("    "), []byte(" ")},
		{[]byte(" \r \n  "), []byte(" ")},
		{[]byte("  abc"), []byte(" abc")},
		{[]byte("  a  b  c  "), []byte(" a b c ")},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n", test.input, test.expected)
		res := compressSpace(test.input)
		fmt.Fprintf(os.Stdout, "res: %v\n", res)
		if len(res) != len(test.expected) {
			t.Errorf("func(%v) = %v", test.input, res)
			fmt.Println("==================================================")
			continue
		}
		for i, v := range res {
			if v != test.expected[i] {
				t.Errorf("result[%v](%v) != expected[%v](%v)", i, v, i, test.expected[i])
			}
		}
		fmt.Println("==================================================")
	}

}
