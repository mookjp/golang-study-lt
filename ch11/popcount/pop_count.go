package popcount

// pc[i]はiのポピュレーションカウント
var pc [256]byte

func init() {
	for i := range pc {
		// i = 1のとき
		// pc[1/2] -> pc[0]
		// byte(1&1) -> byte(1)
		// pc[1] = pc[0] + byte(1) -> 0 + 1 = 1

		// i = 2のとき
		// pc[2/2] -> pc[1]
		// byte(2&1) -> byte(0)
		// pc[2] = pc[1] + byte(0) -> 1 + 0 = 1

		// i = 3のとき
		// pc[3/2] -> pc[1]
		// byte(3&1) -> byte(1)
		// pc[3] = pc[1] + byte(1) -> 1 + 1 = 2
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Fprintf(os.Stdout, "pc[%v]: %v\n", i, pc[i])
	}
}

// PopCount は xのポピュレーションカウント（1が設定されているビット数）を返します
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
