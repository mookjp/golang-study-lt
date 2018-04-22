package ex01

func main() {

}

func countBitDiff(srcBytes *[32]byte, targetBytes *[32]byte) int {
	var diffCount int
	for i, v := range targetBytes {
		if v != srcBytes[i] {
			diffCount++
		}
	}
	return diffCount
}
