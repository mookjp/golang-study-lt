package ex01

func countBitDiff(srcBytes *[32]byte, targetBytes *[32]byte) int {
	var diffCount int

	for i, v := range targetBytes {
		for bitN := 0; bitN < 8; bitN++ {
			var bitDiff bool
			if bitN == 0 {
				bitDiff = int(v)&1 == int(srcBytes[i])&1
			} else {
				bitDiff = int(v)&2*bitN == int(srcBytes[i])&2*bitN
			}
			if !bitDiff {
				diffCount++
			}
		}
	}
	return diffCount
}
