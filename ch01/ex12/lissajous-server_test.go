package main

import (
	"bytes"
	"net/http"
	"testing"
)

// TODO: mainをgoroutineで2回呼べない。。
//func TestServerReturnsResponseWithoutParams(t *testing.T) {
//	go func() {
//		// run server
//		main()
//	}()
//
//	_, err := http.Get("http://localhost:8000")
//	if err != nil {
//		t.Errorf("Got error in requesting. %v", err)
//	}
//}

func TestServerReturnsResponseWithParams(t *testing.T) {
	go func() {
		// run server
		main()
	}()

	_, err := http.Get("http://localhost:8000/?size=400&cycles=120")
	if err != nil {
		t.Errorf("Got error in requesting. %v", err)
	}
}

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
