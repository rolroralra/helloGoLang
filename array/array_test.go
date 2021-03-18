package array

import (
	"fmt"
	"testing"
)

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
