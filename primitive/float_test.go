package primitive

import (
	"fmt"
	"math"
	"testing"
)

const (
	EpsilonFloat32 float64 = 1e-5
	EpsilonFloat64 float64 = 1e-14
)

func TestFloat32(t *testing.T) {
	var got float32 = 10.0

	for i := 0; i < 10; i++ {
		got -= float32(0.1)
	}

	var want float32 = 9.0

	fmt.Println(got, want)
	fmt.Println(math.Abs(float64(got - want)))

	if math.Abs(float64(got-want)) > EpsilonFloat32 {
		panic("Float Rounding Error.")
	} else {
		fmt.Printf("%v = %v\n", got, want)
	}
}

func TestFloat64(t *testing.T) {
	var got = 10.0

	for i := 0; i < 10; i++ {
		got -= 0.1
	}

	var want = 9.0

	fmt.Println(got, want)
	fmt.Println(math.Abs(got - want))

	if math.Abs(got-want) > EpsilonFloat64 {
		panic("Float Rounding Error.")
	} else {
		fmt.Printf("%v = %v\n", got, want)
	}
}
