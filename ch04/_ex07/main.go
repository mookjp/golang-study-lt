package rev

import "fmt"

// utf8の文字列を逆順に並び替えます
func reverse(b []byte) (error, []byte) {
	loopMax := len(string(b)) / 2

	for j, loop := 0, 0; j < len(b) && loop < loopMax; {
		bitN := b[j] & 252

		switch {
		// TODO 共通化
		case bitN >= 252: // 6 bytes
			return fmt.Errorf("no impl"), nil
		case bitN >= 248: // 5 bytes
			return fmt.Errorf("no impl"), nil
		case bitN >= 240: // 4 bytes
			return fmt.Errorf("no impl"), nil
		case bitN >= 224: // 3 bytes
			return fmt.Errorf("no impl"), nil
		case bitN >= 192: // 2 bytes
			return fmt.Errorf("no impl"), nil
		case bitN < 128: // 1 byte = ascii
			fmt.Printf("b[j+1:]: %v, b[:j+1]: %v\n", b[j+1:], b[:j+1])
			b = append(b[j+1:], b[:j+1]...)
			fmt.Printf("b: %v\n", b)
			j++
		default:
			return fmt.Errorf("invalid utf8 bytes"), nil
		}
		loop++
		fmt.Printf("loopMax: %v, loop: %v\n", loopMax, loop)
	}
	return nil, b
}

func r(b []byte) (error, []byte) {
	bitN := b[0] & 252

	switch {
	// TODO 共通化できないか
	case bitN >= 252: // 6 bytes
		return nil, b
	case bitN >= 248: // 5 bytes
		return nil, b
	case bitN >= 240: // 4 bytes
		return nil, b
	case bitN >= 224: // 3 bytes
		return nil, b
	case bitN >= 192: // 2 bytes
		return nil, b
	case bitN < 128: // 1 byte = ascii
		return nil, b
	default:
		return fmt.Errorf("invalid utf8 bytes"), nil
	}

}
