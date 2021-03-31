package vargs

import (
	"fmt"
	"strings"
	"testing"
)

// Variable Arguments in Function Parameters
func PrintAll(args ...interface{}) {
	stringSlice := make([]string, len(args))
	for i, v := range args {
		stringSlice[i] = fmt.Sprint(v)
	}

	fmt.Println(strings.Join(stringSlice, " "))
}

func TestVargs(t *testing.T) {
	PrintAll("Hello", "World", "!")

	var interfaceSlice []interface{}
	interfaceSlice = append(interfaceSlice, 1, 2, 3, 4, 5)

	// Variable Arguments in Function Call. Only Silce can be used "..."
	interfaceSlice = append(interfaceSlice, interfaceSlice...)
	PrintAll(interfaceSlice...) // 1 2 3 4 5 1 2 3 4 5

	var interfaceArray = [...]interface{}{1, "Hello", 3.3, "World", '!'}
	//PrintAll(interfaceArray...)	// Compile Error
	PrintAll(interfaceArray[:]...) // 1 Hello 3.3 World 33
}
