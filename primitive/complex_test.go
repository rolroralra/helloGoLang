package primitive

import (
	"fmt"
	"testing"
)

type Complex complex128

func (z Complex) String() string {
	return fmt.Sprintf("[%v+%vi]", real(z), imag(z))
}

func TestComplex(t *testing.T) {
	var z = complex(1, 2)
	var z2 = 3 - 2i

	fmt.Println(z)
	fmt.Println(z2)
	fmt.Println(real(z))
	fmt.Println(imag(z))

	fmt.Println(z == 1+2i)

	var z3 Complex = complex(1, 2)
	fmt.Println(z3)
}
