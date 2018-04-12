package comma

import (
	"fmt"
	"os"
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"10", "10"},
		{"100", "100"},
		{"1000", "1,000"},
		{"10000", "10,000"},
		{"1000000", "1,000,000"},
	}

	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v expected: %v\n", test.input, test.expected)
		if got := comma(test.input); got != test.expected {
			t.Errorf("comma(%v) = %v", test.input, got)
		}
	}
}
