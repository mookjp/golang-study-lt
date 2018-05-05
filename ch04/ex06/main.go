package main

import "unicode"

func compressSpace(b []byte) []byte {
	if len(b) < 2 {
		return b
	}

	if unicode.IsSpace(rune(b[0])) && unicode.IsSpace(rune(b[1])) {
		if len(b) < 3 {
			return []byte(" ")
		}
		return compressSpace(append([]byte(" "), b[2:]...))
	}
	return append(b[:1], compressSpace(b[1:])...)
}
