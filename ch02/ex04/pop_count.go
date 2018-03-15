package ex04

// PopCount は xのポピュレーションカウント（1が設定されているビット数）を返します
func PopCount(x uint64) int {
	var count int
	for i := 0; i < 64; i++ {
		// 最下位ビットの検査
		var last = x&1
		if last == 1 {
			count++
		}
		x = x >> 1
	}
	return count
}