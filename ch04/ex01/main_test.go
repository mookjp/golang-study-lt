package ex01

import (
	"encoding/binary"
	"fmt"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	fmt.Fprintf(os.Stdout, "x: %v\n", binary.BigEndian.Uint64([]byte{'x'}))
}

func TestBitCount(t *testing.T) {
	var tests = []struct {
		input1   [32]byte
		input2   [32]byte
		expected int
	}{
		//{[32]byte{}, [32]byte{}, 0},
		//{makeSequence(1), makeSequence(1), 0},
		//{[32]byte{1}, [32]byte{}, 1},
		{[32]byte{3}, [32]byte{}, 2},
		//{[32]byte{4}, [32]byte{}, 3},
		//{[32]byte{255}, [32]byte{}, 8},
		//{[32]byte{255, 1}, [32]byte{}, 9},
	}
	for _, test := range tests {
		fmt.Fprintf(os.Stdout, "input: %v %v expected: %v\n", test.input1, test.input2, test.expected)
		if got := countBitDiff(&test.input1, &test.input2); got != test.expected {
			t.Errorf("countBitDiff(%v, %v) = %v", test.input1, test.input2, got)
		}
	}

}

func makeSequence(start int) [32]byte {
	var bytes [32]byte
	for i := 0; i < 32; i, start = i+1, start+1 {
		bytes[i] = byte(start)
	}
	return bytes
}
