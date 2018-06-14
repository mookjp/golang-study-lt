package main

import (
	"fmt"
	"os"
	"testing"
	"golang.org/x/net/html"
	"strings"
)

func Test(t *testing.T) {
	var tests = []struct {
		input    string
		expected []string
	}{
		{`
			<html>
				<body>
					<a href="localhost">aaa</a>
					<p>
						<a href="localhost2">aaa</a>
						<a href="localhost3">aaa</a>
					</p>
					<ul>
						<li><a href="localhost4">aaa</a></li>
					</ul>
				</body>
			</html>
		`,
		[]string{"localhost", "localhost2", "localhost3", "localhost4"}},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v(%v), expected: %v(%v)\n",
			test.input, string(test.input), test.expected, test.expected)
		doc, err := html.Parse(strings.NewReader(test.input))
		if err != nil {
			t.Error(err)
			return
		}

		res := visit(nil, doc)
		fmt.Fprintf(os.Stdout, "res: %v\n", res)
		for i, v := range res {
			if v != test.expected[i] {
				t.Errorf("assertion error. input: input[%v] = %v, expected[%v] = %v",
					i, v, i, test.expected[i])
			}
		}
	}
}
