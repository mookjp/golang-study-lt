package intset

import (
	"fmt"
	"os"
	"testing"
)

func TestString(t *testing.T) {
	var tests = []struct {
		input    []int
		expected string
	}{
		{[]int{1, 9, 144}, "{1 9 144}"},
		{[]int{9, 42}, "{9 42}"},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n",
			test.input, test.expected)

		// Execution
		var x IntSet
		for _, num := range test.input {
			x.Add(num)
		}
		res := x.String()

		// Assertion
		if res != test.expected {
			t.Errorf("assertion error. input: %v, expected: %v, actual: %v", test.input, test.expected, res)
		}
	}
}
