package comma

import (
	"bytes"
)

// commna は負でない10進表記整数文字列にカンマを挿入します
func anagram(base string, target string) bool {
	if len(base) != len(target) {
		return false
	}

	baseBuf := bytes.NewBufferString(base)
	targetBuf := bytes.NewBufferString(target)
	leftBuf := targetBuf

	for baseBuf.Len() > 0 {
		baseByte, err := baseBuf.ReadByte()
		if err != nil {
			break
		}
		found := false
		var nextBuf bytes.Buffer
		for leftBuf.Len() > 0 {
			targetByte, err := leftBuf.ReadByte()
			if err != nil {
				break
			}
			if baseByte == targetByte {
				found = true
				nextBuf.Write(leftBuf.Bytes())
				break
			} else {
				nextBuf.WriteByte(targetByte)
			}
		}
		if !found {
			return false
		}
		leftBuf.Reset()
		leftBuf.Write(nextBuf.Bytes())
	}

	return true
}
