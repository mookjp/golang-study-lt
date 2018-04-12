package comma

import (
	"fmt"
	"os"
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"anagrams", "arsmagna", true},
		{"christmas", "trimscash", true},
		{"hoge", "hogee", false},
		{"aabb", "abbb", false},
	}

	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v %v expected: %v\n", test.input1, test.input2, test.expected)
		if got := anagram(test.input1, test.input2); got != test.expected {
			t.Errorf("anagram(%v, %v) = %v", test.input1, test.input2, got)
		}
	}
}
