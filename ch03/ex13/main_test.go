package ex13

import (
	"fmt"
	"testing"
)

func TestConsts(t *testing.T) {
	expected := c{
		KB: 1000,
		MB: 1000000,
		GB: 1000000000,
		TB: 1000000000000,
		//PB: 1000000000000000,
		//EB: 1000000000000000000,
		//ZB: 1000000000000000000000,
		//YB: 1000000000000000000000000,
	}
	res := consts()
	fmt.Printf("expected: %v\n", expected)
	if res != expected {
		t.Errorf("actual: %v\n", res)
	}
}
