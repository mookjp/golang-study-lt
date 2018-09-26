package main

import (
	"fmt"
	"os"
	"testing"
)

func TestBank(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{100, 100, 200}, 400},
		{[]int{1, 2, 3}, 6},
		{[]int{0, 1, 1}, 2},
		{[]int{0, 0, 0}, 0},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n",
			test.input, test.expected)

		// Execution
		Init()
		for _, v := range test.input {
			Deposit(v)
		}
		result := Balance()

		if result != test.expected {
			t.Errorf("assertion error. input: %v, expected: %v, actual: %v", test.input, test.expected, result)
		}
	}
}
