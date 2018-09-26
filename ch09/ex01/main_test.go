package main

import (
	"fmt"
	"os"
	"testing"
)

func TestBank(t *testing.T) {
	var tests = []struct {
		deposits  []int
		withdraws []int
		expected  int
	}{
		{[]int{0, 1, 1}, []int{0, 0, 0}, 2},
		{[]int{0, 1, 2}, []int{0, 1, 2}, 0},
		{[]int{1, 1, 2}, []int{0, 1, 2}, 1},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "deposits: %v, withdraws: %v, expected: %v\n",
			test.deposits, test.withdraws, test.expected)

		// Execution
		Init()
		for _, v := range test.deposits {
			Deposit(v)
		}
		for _, v := range test.withdraws {
			result := Withdraw(v)
			if result == false {
				t.Errorf("failed to withdraw")
				return
			}
		}
		result := Balance()

		if result != test.expected {
			t.Errorf("assertion error. deposits: %v, withdraws: %v, expected: %v, actual: %v",
				test.deposits, test.withdraws, test.expected, result)
		}
	}
}

func TestResult(t *testing.T) {
	var tests = []struct {
		deposits  int
		withdraws int
		expected  bool
	}{
		{1, 1, true},
		{1, 2, false},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "deposits: %v, withdraws: %v, expected: %v\n",
			test.deposits, test.withdraws, test.expected)

		// Execution
		Init()
		Deposit(test.deposits)
		result := Withdraw(test.withdraws)

		if result != test.expected {
			t.Errorf("assertion error. deposits: %v, withdraws: %v, expected: %v, actual: %v",
				test.deposits, test.withdraws, test.expected, result)
		}
	}
}
