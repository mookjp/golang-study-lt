package comma

import (
	"bytes"
)

const digit = 3

// commna は負でない10進表記整数文字列にカンマを挿入します
func comma(s string) string {

	orig := bytes.NewBufferString(s)
	var out bytes.Buffer

	former, err := orig.ReadBytes([]byte(".")[0])
	var formerBuf bytes.Buffer
	var latterBuf bytes.Buffer
	if err != nil {
		formerBuf.Write(former)
	} else {
		formerBuf.Write(former[:len(former)-1])
		latterBuf.Write([]byte("."))
		latterBuf.Write(orig.Bytes())
	}

	for formerBuf.Len() > 0 {
		b := formerBuf.Next(1)
		out.Write(b)
		if formerBuf.Len() > digit-1 && formerBuf.Len()%digit == 0 {
			out.WriteString(",")
		}
	}
	for latterBuf.Len() > 0 {
		b := latterBuf.Next(1)
		out.Write(b)
	}

	return out.String()
}
