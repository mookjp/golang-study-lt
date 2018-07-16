package ex10

import "sort"

type Sentence []string

func (x Sentence) Len() int {
	return len(x)
}
func (x Sentence) Less(i, j int) bool {
	if i < j {
		return true
	}
	return false
}
func (x Sentence) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func IsPalindrome(s Sentence) bool {
	reversed := make(Sentence, len(s))
	copy(reversed, s)
	sort.Sort(sort.Reverse(reversed))

	for i, v := range reversed {
		if v != s[i] {
			return false
		}
	}
	return true

}
