package ex01

import (
	"testing"
	"fmt"
	"os"
	"../tempconv"
)

func TestKToC(t *testing.T) {
	var tests = []struct {
		input   tempconv.KelvinScale
		expected tempconv.Celsius
	}{
		{tempconv.KelvinScale(-1), tempconv.Celsius(-546.3)},
		{tempconv.KelvinScale(0), tempconv.Celsius(-273.15)},
		{tempconv.KelvinScale(1),tempconv.Celsius(0)},
		{tempconv.KelvinScale(2),tempconv.Celsius(273.15)},
	}

	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v expected: %v\n", test.input, test.expected)
		if got := tempconv.KToC(test.input); got != test.expected {
			t.Errorf("KtoC(%v) = %v", test.input, got)
		}
	}
}
