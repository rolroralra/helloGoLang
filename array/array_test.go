package array

import (
	"fmt"
	"testing"
)

func TestPrintArray(t *testing.T) {
	PrintArray([...]int{1, 2, 3})
	PrintArray([0]int{})
}

func TestForRange(t *testing.T) {
	array := [...]int{3, 4, 5}
	for i, v := range array {
		fmt.Printf("index=%d, value=%v\n", i, v)
	}
}

func TestArray(t *testing.T) {
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "go lang"

	// Array Literal
	primes := [5]int{2, 3, 5, 7}

	fmt.Println(arr)
	fmt.Println(arr[0], arr[1])

	fmt.Println(primes)
	fmt.Println(primes[1:3])
}

func TestArrayIsPrimitiveType(t *testing.T) {
	arr := [2][5]int{{1, 2, 3, 4, 5}, {2, 4, 6, 8, 10}}

	temp := arr[0]

	temp[0] = 99

	fmt.Println(arr)    // [[1 2 3 4 5] [2 4 6 8 10]]
	fmt.Println(arr[0]) // [1 2 3 4 5]
	fmt.Println(temp)   // [99 2 3 4 5]
}
