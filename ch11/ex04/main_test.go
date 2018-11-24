package ex04

import (
	"math/rand"
	"testing"
	"time"

	"fmt"

	"gopl.io/ch11/word2"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // 24までのランダムな数字
	runes := make([]rune, n)

	// n = 9 のとき i は 5 以下
	// n = 10 のとき i は 5 以下
	// n = 11 のとき i は 6 以下
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // '\u0999' までのランダムなルーン
		// runes の 頭から r を詰める
		runes[i] = r
		// n = 10 のとき 10-1-0=
		// runes の お尻から r を詰める
		runes[n-1-i] = r
	}
	return addChars(string(runes))
}

func TestRandomPalindromes(t *testing.T) {
	// 疑似乱数生成器を初期化する
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		t.Logf("random palindrome: %s", p)
		if !word.IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func addChars(str string) string {
	return fmt.Sprintf("　 ,. %s ., 　", str)
}
