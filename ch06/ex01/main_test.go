package intset

import (
	"fmt"
	"os"
	"testing"
)

func TestLen(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{42}, 1},
		{[]int{1, 9, 144}, 3},
		{[]int{9, 42}, 2},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n",
			test.input, test.expected)

		// Execution
		var x IntSet
		for _, num := range test.input {
			x.Add(num)
		}
		res := x.Len()

		// Assertion
		if res != test.expected {
			t.Errorf("assertion error. input: %v, expected: %v, actual: %v", test.input, test.expected, res)
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		input    IntSet
		expected IntSet
	}{
		{IntSet{}, IntSet{}},
		{IntSet{[]uint64{100}}, IntSet{[]uint64{100}}},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n",
			test.input, test.expected)

		// Execution
		copied := test.input.Copy()

		// Assertion
		for i, copiedNum := range copied.words {
			if copiedNum != test.expected.words[i] {
				t.Errorf("assertion error. input: %v, expected: %v, actual: %v", test.input, test.expected, copied)
			}
		}
	}
}
