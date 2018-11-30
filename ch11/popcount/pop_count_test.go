package popcount

import (
	"fmt"
	"os"
	"testing"
)

func TestPopCount(t *testing.T) {
	var tests = []struct {
		input    uint64
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
	}

	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v expected: %v\n", test.input, test.expected)
		if got := PopCount(test.input); got != test.expected {
			t.Errorf("KtoC(%v) = %v", test.input, got)
		}
	}
}
