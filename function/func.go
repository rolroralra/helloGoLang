package function

import (
	"fmt"
	"math"
)

func Sum(x, y int) int {
	return x + y
}

func Swap(x, y interface{}) (interface{}, interface{}) {
	return y, x
}

func DivideAndRemainder(x, y int) (int, int) {
	return x / y, x % y
}

func Sqrt(x float64) float64 {
	result := 1.0

	x = math.Abs(x)
	for i := 0; i < 10; i++ {
		result -= (result * result - x) / (2 * result)
	}

	return result
}

// Variable Args implements by Slice
// args might be Slice type.
func VarArgs(args ...interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", args, args)
}

