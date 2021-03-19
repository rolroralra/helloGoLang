package slice

import (
	"fmt"
	"strings"
	"testing"
)

// Reference: https://blog.golang.org/slices-intro
func TestPrintSlice(t *testing.T) {
	PrintSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	PrintSlice([]string{"a", "b", "c", "d", "efg"})

	PrintSlice(make([]int, 0, 5)[:5])

	arr := [...]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(len(arr), cap(arr))
}

func TestMakeCopyAppend(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}

	// func make(t Type, size ...int)
	slice := make([]int, 5, 5)

	// func copy(dst []T, src []T)
	copy(slice, arr[:])

	// func append(s []T, args ...T)
	slice = append(slice, 6, 7)
	slice[0] = 0

	PrintSlice(slice)
	PrintSlice(arr[:])
}

func TestForEach(t *testing.T) {
	ForEach([]int{3, 4, 5}, func(i interface{}) {
		fmt.Println(i)
	})
}

func TestSlice(t *testing.T) {
	// Array Literal
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[1:3])

	a := arr[1:3]
	b := arr[0:2]
	fmt.Println(a, b)

	b[0] = 999

	// Slices are like references to arrays.
	fmt.Println(a, b)
	fmt.Println(arr)

	// Slice Literal
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	PrintSlice(intSlice)
	PrintSlice(intSlice[2:])
	PrintSlice(intSlice[:5])
	PrintSlice(intSlice[2:5])

	structureSlice := []struct {
		i int
		b bool
	}{
		{2, true},
		{1, false},
		{3, true},
	}

	fmt.Println(structureSlice)
}

func TestSliceOfSlice(t *testing.T) {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func TestNilSlice(t *testing.T) {
	var slice []int
	PrintSlice(slice)
	if slice == nil {
		fmt.Printf("nil slice! %v\n", slice)
	}
}

func TestForRange(t *testing.T) {
	slice := []int{1, 2, 3}

	for i, v := range slice {
		fmt.Printf("index=%d, value=%v\n", i, v)
	}
}
