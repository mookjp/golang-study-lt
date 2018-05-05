package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestMainFunc(t *testing.T) {
	var tests = []struct {
		input    []string
		expected []string
	}{
		{[]string{}, []string{}},
		{[]string{"あ"}, []string{"あ"}},
		{strings.Split("あああ", ""), []string{"あ"}},
		{strings.Split("ああああ", ""), []string{"あ"}},
		{strings.Split("あああああ", ""), []string{"あ"}},
		{strings.Split("あいうえおおお", ""), []string{"あ", "い", "う", "え", "お"}},
		{strings.Split("ああいいううええおお", ""), []string{"あ", "い", "う", "え", "お"}},
		{strings.Split("ああああああいうえお", ""), []string{"あ", "い", "う", "え", "お"}},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v, expected: %v\n", test.input, test.expected)
		res := removeDuplicated(test.input)
		fmt.Fprintf(os.Stdout, "res: %v\n", res)
		if len(res) != len(test.expected) {
			t.Errorf("func(%v) = %v", test.input, res)
			fmt.Println("==================================================")
			continue
		}
		for i, v := range res {
			if v != test.expected[i] {
				t.Errorf("result[%v](%v) != expected[%v](%v)", i, v, i, test.expected[i])
			}
		}
		fmt.Println("==================================================")
	}

}
