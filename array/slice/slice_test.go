package slice

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrintSlice(t *testing.T) {
	PrintSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	PrintSlice([]string{"a", "b", "c", "d", "efg"})


	PrintSlice(make([]int, 0, 5)[:5])

	arr := [...]int{1,2,3}
	fmt.Println(arr)
	fmt.Println(len(arr), cap(arr))
}

func TestCopy(t *testing.T) {
	arr := [...]int{1,2,3,4,5}
	slice := make([]int, 5, 5)
	copy(slice, arr[:])
	slice = append(slice, 6, 7)
	slice[0] = 0

	PrintSlice(slice)
	PrintSlice(arr[:])

}

func TestForEach(t *testing.T) {
	ForEach([]int{1, 2, 3}, func(i interface{}) {
		fmt.Println(i)
	})
}

func Test(t *testing.T) {
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

func Test2(t *testing.T) {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
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
