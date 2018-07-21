package main_test

import (
	"testing"

	"fmt"
	"os"
)

func TestParseAttArg(t *testing.T) {
	var tests = []struct {
		input         string
		expectedName  string
		expectedValue string
		expectedErr   error
	}{
		{"class=\"hoge\"", "class", "hoge", nil},
		{"name=\"fuga-hoge\"", "name", "fuga-hoge", nil},
		{"div", "div", "", nil},
	}
	for _, test := range tests {
		actualName, actualValue, err := main.ParseAttrArg(test.input)
		if err != nil {
			fmt.Fprintf(os.Stdout, "error: %v", err)
		}
		if actualName != test.expectedName || actualValue != test.expectedValue {
			t.Errorf("assertion error. input = %v, actual = %v %v %v, expected = %v %v %v",
				test.input, actualName, actualValue, err, test.expectedName, test.expectedValue, test.expectedErr)
		}
	}
}
