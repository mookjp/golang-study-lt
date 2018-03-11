package tempconv

import (
	"testing"
	"fmt"
	"os"
)

func TestKToC(t *testing.T) {
	var tests = []struct {
		input   KelvinScale
		expected Celsius
	}{
		{KelvinScale(-1), Celsius(-546.3)},
		{KelvinScale(0), Celsius(-273.15)},
		{KelvinScale(1), Celsius(0)},
		{KelvinScale(2), Celsius(273.15)},
	}

	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v expected: %v\n", test.input, test.expected)
		if got := KToC(test.input); got != test.expected {
			t.Errorf("KtoC(%v) = %v", test.input, got)
		}
	}
}
