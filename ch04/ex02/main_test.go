package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMainFunc(t *testing.T) {
	var tests = []struct {
		input    string
		m        mode
		expected string
	}{
		{"hoge\n", SHA254, getStringHash("hoge", SHA254)},
		{"hoge\n", SHA384, getStringHash("hoge", SHA384)},
		{"hoge\n", SHA512, getStringHash("hoge", SHA512)},
		{"hoge\n", "aaa", getStringHash("hoge", SHA254)},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v mode: %v, expected: %v\n", test.input, test.m, test.expected)
		out := new(bytes.Buffer)
		err := printSha(strings.NewReader(test.input), out, test.m)
		if err != nil {
			t.Errorf("err: %v in func(%v, %v)", err, test.input, test.m)
		}
		res := out.String()
		fmt.Fprintf(os.Stdout, "actual: %v\n", res)
		if res != test.expected {
			t.Errorf("func(%v, %v) = %v", test.input, test.m, res)
		}
	}

}

func getStringHash(input string, m mode) string {
	reader := bufio.NewReader(strings.NewReader(input))
	in, isPrefix, err := reader.ReadLine()
	if isPrefix || err != nil {
		fmt.Errorf("error")
		return ""
	}

	switch m {
	case SHA254:
		b := sha256.Sum256(in)
		fmt.Fprintf(os.Stdout, "in getStringHash: %x\n", b)
		return fmt.Sprintf("%x", b)
	case SHA384:
		b := sha512.Sum384(in)
		fmt.Fprintf(os.Stdout, "in getStringHash: %x\n", b)
		return fmt.Sprintf("%x", b)
	case SHA512:
		b := sha512.Sum512(in)
		fmt.Fprintf(os.Stdout, "in getStringHash: %x\n", b)
		return fmt.Sprintf("%x", b)
	default:
		fmt.Errorf("invalid mode param: %v", m)
	}
	return ""
}
