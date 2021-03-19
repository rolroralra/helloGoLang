package closure

import (
	"fmt"
	"testing"
)

func TestSummation(t *testing.T) {
	sum1, sum2 := Summation(), Summation()
	for i := 0; i < 10; i++ {
		fmt.Println(i, sum1(i), sum2(-i))
	}
}

func TestFibonacci(t *testing.T) {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci(%d) = %v\n", i, f())
	}
}
