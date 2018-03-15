package ex05

// PopCount は xのポピュレーションカウント（1が設定されているビット数）を返します
func PopCount(x uint64) int {
	var count int
	var x2 = x
	for x2 > 0 {
		count++
		x2 = x2 & (x2 -1)
	}
	return count
}