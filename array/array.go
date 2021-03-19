package array

import (
	"fmt"
	"reflect"
)

func PrintArray(arr interface{}) {
	a := reflect.ValueOf(arr)

	if a.Kind() != reflect.Array {
		panic("Parameter must be Array.")
	}

	fmt.Printf("len=%d, cap=%d, %v\n", a.Len(), a.Cap(), a.Interface())
}
