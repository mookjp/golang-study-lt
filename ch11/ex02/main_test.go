package intset

import (
	"strconv"
	"testing"
)

func TestIntSetAdd(t *testing.T) {
	var tests = []struct {
		input    map[int]int
		expected string
	}{
		{map[int]int{1: 1}, "{1}"},
		{map[int]int{1: 1, 2: 2, 3: 3}, "{1, 2, 3}"},
	}

	for _, test := range tests {
		t.Logf("input: %v, expected: %v\n", test.input, test.expected)

		inputKeySet := makeKeySetFromMap(test.input)
		set := IntSet{}
		for _, v := range inputKeySet {
			set.Add(v)
		}
		inputKeySetStr := makeArrayString(inputKeySet)
		setStr := set.String()

		if inputKeySetStr != setStr {
			t.Errorf("inputKeySetStr = %v, setStr = %v", inputKeySetStr, setStr)
		}
	}
}

func makeKeySetFromMap(m map[int]int) []int {
	set := make([]int, 0, len(m))
	for _, value := range m {
		set = append(set, value)
	}
	return set
}

func makeArrayString(set []int) string {
	res := "{"
	for i, v := range set {
		if i != 0 {
			res = res + " " + strconv.Itoa(v)
		} else {
			res = res + strconv.Itoa(v)
		}
		if i == len(set)-1 {
			res = res + "}"
		}
	}
	return res
}
