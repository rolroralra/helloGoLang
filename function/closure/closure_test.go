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

func calc(a int, b int) func(int) int {
	return func(x int) int {
		return a * x + b
	}
}

func TestClosure(t *testing.T) {
	a, b := 1, 2
	functionVariable := calc(a, b)
	fmt.Println(functionVariable(1))
	a, b = 2, 3
	functionVariable = calc(a, b)
	fmt.Println(functionVariable(1))
}
