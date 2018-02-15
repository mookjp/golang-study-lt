package main

import (
	"testing"
	"bytes"
)

func TestLissajous(t *testing.T) {
	params := lissajousParams{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}

	buf := &bytes.Buffer{}
	lissajous(buf, params)

	if buf.Len() == 0 {
		t.Errorf("It did not output anything. bytes.Buffer.Len() == 0")
	}
}
