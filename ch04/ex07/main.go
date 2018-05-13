package rev

import (
	"fmt"
	"os"
)

// この実装にすると以前のテストケースが通らなくなる
// テストケースでは引数のbyte自体の検証をしているため
// 参照先の配列自体を変更しないと失敗する
// この関数では別のstringのスライスの変更をしているため引数のbyteスライス自体には変更が入らない
//func reverse(b []byte) {
//	s := strings.Split(string(b), "")
//	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
//		s[i], s[j] = s[j], s[i]
//	}
//	fmt.Fprintf(os.Stdout, "input: %v, after: %v\n", b, s)
//	b = []byte(strings.Join(s, ""))
//}

type CharMap struct {
	headIndex int
	byteNum   int
}

func reverse(b []byte) {
	orig := make([]byte, len(b))
	copy(orig, b)

	err, charMap := createMap(b)
	if err != nil {
		fmt.Errorf("invalid utf8 chars")
		os.Exit(1)
	}

	for i, j := 0, len(charMap)-1; i < j; i, j = i+1, j-1 {
		former := charMap[i]
		latter := charMap[j]

		// former
		for num := 0; num < former.byteNum; num++ {
			b[former.headIndex+num] = orig[latter.headIndex+num]
		}
		// latter
		for num := 0; num < latter.byteNum; num++ {
			b[latter.headIndex+num] = orig[former.headIndex+num]
		}
	}
}

func createMap(b []byte) (error, []CharMap) {
	cMap := []CharMap{}
	// 頭から見ていく
	// 1111 1100 でマスクする
	// マスク結果を比較して, byteNumをいれる
	for i := 0; i < len(b); {
		bitN := b[i]
		switch {
		// TODO 共通化
		case bitN >= 252: // 6 bytes
			cMap = append(cMap, CharMap{i, 6})
			i = i + 6
		case bitN >= 248: // 5 bytes
			cMap = append(cMap, CharMap{i, 5})
			i = i + 5
		case bitN >= 240: // 4 bytes
			cMap = append(cMap, CharMap{i, 4})
			i = i + 4
		case bitN >= 224: // 3 bytes
			cMap = append(cMap, CharMap{i, 3})
			i = i + 3
		case bitN >= 192: // 2 bytes
			cMap = append(cMap, CharMap{i, 2})
			i = i + 2
		case bitN < 128: // 1 byte = ascii
			cMap = append(cMap, CharMap{i, 1})
			i++
		default:
			return fmt.Errorf("invalid utf8 bytes"), nil
		}
	}
	return nil, cMap
}
