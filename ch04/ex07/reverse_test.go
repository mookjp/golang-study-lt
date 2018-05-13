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
		{[]byte("あ"), []byte("あ")},
		{[]byte("あい"), []byte("いあ")},
		{[]byte("あいうえお"), []byte("おえういあ")},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v(%v), expected: %v(%v)\n",
			test.input, string(test.input), test.expected, string(test.expected))
		reverse(test.input)
		fmt.Fprintf(os.Stdout, "after reverse input: %v\n", test.input)
		for i, v := range test.input {
			if v != test.expected[i] {
				t.Errorf("assertion error. input: input[%v] = %v, expected[%v] = %v",
					i, v, i, test.expected[i])
			}
		}
	}
}

func TestCreateMap(t *testing.T) {
	var tests = []struct {
		input    []byte
		expected []CharMap
	}{
		{[]byte("ab"), []CharMap{
			CharMap{0, 1},
			CharMap{1, 1},
		}},
		{[]byte("あい"), []CharMap{
			CharMap{0, 3},
			CharMap{3, 3},
		}},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v(%v), expected: %v\n",
			test.input, string(test.input), test.expected)
		err, res := createMap(test.input)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		fmt.Fprintf(os.Stdout, "res: %v\n", res)
	}
}
