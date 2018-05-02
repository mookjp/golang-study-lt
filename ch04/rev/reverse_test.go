package rev

import (
	"fmt"
	"os"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n", test.input, test.expected)
		reverse(test.input)
		for i, v := range test.input {
			if v != test.expected[i] {
				t.Errorf("assertion error. input: input[%v] = %v, expected[%v] = %v", i, v, i, test.expected[i])
			}
		}
	}

}
