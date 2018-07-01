package ex01

func countBitDiff(srcBytes *[32]byte, targetBytes *[32]byte) int {
	var diffCount int

	// TODO: シフトでできないか？
	for i, v := range targetBytes {
		for bitN := 0; bitN < 8; bitN++ {
			var bitDiff bool
			if bitN == 0 {
				bitDiff = int(v)&1 == int(srcBytes[i])&1
			} else {
				bitDiff = int(v)&power(2, bitN) == int(srcBytes[i])&power(2, bitN)
			}
			if !bitDiff {
				diffCount++
			}
		}
	}
	return diffCount
}

func power(x int, powerNum int) int {
	if x == 0 {
		return 0
	}
	if powerNum == 0 {
		return 1
	}

	var res int
	for i := 0; i < powerNum; i++ {
		if i == 0 {
			res = x
			continue
		}
		res = res * x
	}
	return res
}
