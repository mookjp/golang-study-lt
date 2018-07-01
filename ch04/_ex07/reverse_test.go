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
		//{[]byte(""), []byte("")},
		//{[]byte("a"), []byte("a")},
		//{[]byte("ab"), []byte("ba")},
		{[]byte("abc"), []byte("cba")},
	}
	for _, test := range tests {
		err, res := reverse(test.input)
		fmt.Fprintf(os.Stdout, "input: %v, res: %v, expected: %v\n", test.input, res, test.expected)
		if err != nil {
			t.Errorf("error: %v, input: %v, expected: %v", err, test.input, test.expected)
		}
		if len(res) != len(test.expected) {
			t.Errorf("len is not matched, input: %v, res: %v", test.input, res)
		}
		for i, v := range res {
			if v != test.expected[i] {
				t.Errorf("assertion error. input: res[%v] = %v, expected[%v] = %v", i, v, i, test.expected[i])
			}
		}
	}
}
