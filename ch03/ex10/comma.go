package comma

import (
	"bytes"
)

const digit = 3

// commna は負でない10進表記整数文字列にカンマを挿入します
func comma(s string) string {

	orig := bytes.NewBufferString(s)
	var out bytes.Buffer

	for orig.Len() > 0 {
		b := orig.Next(1)
		out.Write(b)
		if orig.Len() > digit-1 && orig.Len()%digit == 0 {
			out.WriteString(",")
		}
	}
	return out.String()
}
