package ex05

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input       string
		expectedLen int
	}{
		{"a:b:c", 3},
		{"a:b:c:d:e", 5},
		{"a", 1},
		{"", 1},
	}
	sep := ":"

	for _, test := range tests {
		if len(strings.Split(test.input, sep)) != test.expectedLen {
			actual := strings.Split(test.input, sep)
			t.Errorf("input = %v expectedLen = %v actual = %v actualLen = %v",
				test.input, test.expectedLen, actual, len(actual))
		}
	}
}
