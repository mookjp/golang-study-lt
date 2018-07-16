package ex10_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mookjp/golang-study-lt/ch07/ex10"
)

func Test(t *testing.T) {
	var tests = []struct {
		input    ex10.Sentence
		expected bool
	}{
		{[]string{"g", "o", "d", "o", "g"}, true},
		{[]string{"d", "o", "g"}, false},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n", test.input, test.expected)

		// execute
		actual := ex10.IsPalindrome(test.input)

		// actual

		if actual != test.expected {
			t.Errorf("assertion error. input = %v, actual = %v, expected = %v",
				test.input, actual, test.expected)
		}
	}
}
