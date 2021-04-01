package fop

import (
	"fmt"
	"testing"
)

type Operation func(int, int) int

func sum(a int, b int) int {
	return a + b
}
func minus(a int, b int) int {
	return a - b
}
func multiply(a int, b int) int {
	return a * b
}
func divide(a int, b int) int {
	return a / b
}
func remainder(a int, b int) int {
	return a % b
}

func TestFunctionType(t *testing.T) {
	var sumFunc func(int, int) int
	sumFunc = sum

	fmt.Println(sumFunc(1, 2) == sum(1, 2))
}

func TestFunctionSlice(t *testing.T) {
	funcSlice := []func(int, int) int{
		sum,
		minus,
		multiply,
		divide,
		remainder,
	}

	for _, function := range funcSlice {
		fmt.Println(function(7, 3))
	}
}

func TestFunctionMap(t *testing.T) {
	funcMap := map[string]Operation{
		"add":       sum,
		"subtract":  minus,
		"multiply":  multiply,
		"divide":    divide,
		"remainder": remainder,
	}

	for key, function := range funcMap {
		fmt.Println(key, function(7, 3))
	}
}
