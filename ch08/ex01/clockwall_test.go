package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParseArgs(t *testing.T) {
	var tests = []struct {
		input    []string
		expected []clockSettings
	}{
		{[]string{"US/NewYork=localhost:8010"},
			[]clockSettings{
				{"US/NewYork", "localhost", 8010}}},
		{[]string{"US/NewYork=localhost:8010", "Asia/Tokyo=localhost:8020"},
			[]clockSettings{
				{"US/NewYork", "localhost", 8010},
				{"Asia/Tokyo", "localhost", 8020},
			},
		},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n",
			test.input, test.expected)

		// Execution
		res := parseArgs(test.input)

		for i, setting := range res {
			if setting != test.expected[i] {
				if setting.address != test.expected[i].address ||
					setting.port != test.expected[i].port ||
					setting.timezone != test.expected[i].timezone {
					t.Errorf("assertion error. input: %v, expected: %v, actual: %v", test.input, test.expected[i], setting)
				}
			}
		}
	}
}
