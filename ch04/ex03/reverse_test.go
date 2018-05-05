package rev

import (
	"fmt"
	"os"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		input    [4]int
		expected [4]int
	}{
		{[4]int{0, 1, 2, 3}, [4]int{3, 2, 1, 0}},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n", test.input, test.expected)
		reverse(&test.input)
		for i, v := range test.input {
			if v != test.expected[i] {
				t.Errorf("assertion error. input: input[%v] = %v, expected[%v] = %v", i, v, i, test.expected[i])
			}
		}
	}
}
