package archive

import (
	"bufio"
	"errors"
	"io"
)

type format struct {
	name, magic string
}

var formats []format

func RegisterFormat(name, magic string) {
	formats = append(formats, format{name, magic})
}

// GetFormat は format の文字列とエラーを返却します
func GetFormat(r io.Reader) (string, error) {
	br := bufio.NewReader(r)
	format := sniff(br)
	if format.name == "" {
		return "", errors.New("this format is not supported")
	}
	return format.name, nil
}

// bytes がワイルドカードか magic で構成されていることを検証する
func match(magic string, bytes []byte) bool {
	if len(magic) != len(bytes) {
		return false
	}
	for i, char := range bytes {
		// TODO: archive の magic もワイルドカードとして?が入るのか？
		if magic[i] != char && magic[i] != '?' {
			return false
		}
	}
	return true
}

func sniff(br *bufio.Reader) format {
	for _, format := range formats {
		bytes, err := br.Peek(len(format.magic))
		if err == nil && match(format.magic, bytes) {
			return format
		}
	}
	return format{}
}
