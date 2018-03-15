package popcount

// pc[i]はiのポピュレーションカウント
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1) // TODO i&1とは…
	}
}

// PopCount は xのポピュレーションカウント（1が設定されているビット数）を返します
func PopCount(x uint64) int {
	var b byte
	for i := uint64(0); i < 9; i++ {
		b += pc[byte(x>>(i*8))]
	}
	return int(b)
}