package primitive

import (
	"fmt"
	"testing"
)

func TestByte(t *testing.T) {
	var a uint8 = 0b00100101
	var b uint8 = 0b00111011

	fmt.Printf("%08b & %08b = %08b\n", a, b, a&b)
	fmt.Printf("%08b | %08b = %08b\n", a, b, a|b)
	fmt.Printf("%08b ^ %08b = %08b\n", a, b, a^b)
	fmt.Printf("^%08b = %08b\n", a, ^a)

	fmt.Printf("a = %08b,  b = %08b\n", a, b)
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("a = %08b,  b = %08b\n", a, b)

}
