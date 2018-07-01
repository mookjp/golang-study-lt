package intset

import (
	"bytes"
	"fmt"
)

// 負ではない小さな整数のセット
type IntSet struct {
	words []uint64
}

// 負ではない値xをセットが含んでいるか否かを返却
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// セットに負ではない値xを追加
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// sとtの和集合をsに設定します
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// "{1 2 3}" の形式の文字列としてセットを返します
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// TODO: 値の扱いを見る
// 要素数を返します
func (s *IntSet) Len() int {
	return len(s.words)
}

// セットからxを取り除きます
// TODO
func (S *IntSet) Remove(x int) {

}

//セットからすべての要素を取り除きます
// TODO
func (s *IntSet) Clear() {

}

// セットのコピーを返します
func (s *IntSet) Copy() *IntSet {
	newSet := IntSet{}

	for _, num := range s.words {
		newSet.words = append(newSet.words, num)
	}

	return &newSet
}
