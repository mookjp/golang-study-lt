package ex02

import "io"

type ByteCountWriter struct {
	origWriter io.Writer
	count      int64
}

func (c *ByteCountWriter) Write(p []byte) (int, error) {
	c.origWriter.Write(p)
	c.count = c.count + int64(len(p))
	return len(p), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	byteCounter := ByteCountWriter{w, 0}
	return &byteCounter, &byteCounter.count
}
