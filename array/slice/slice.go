package slice

/**
 * reference: https://blog.golang.org/slices-intro
 **/

import (
	"fmt"
	"reflect"
)


func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("Parameter must be Slice.")
	}

	if s.IsNil() {
		return nil
	}

	result := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		result[i] = s.Index(i).Interface()
	}

	return result
}

func PrintSlice(slice interface{}) {
	s := InterfaceSlice(slice)
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), slice)
}

// Variable Args implements by Slice
// args might be Slice type.
func VariableArguments(args ...int) {
	fmt.Printf("Type: %T, Value: %v\n", args, args)
}


func ForEach(slice interface{}, function func (interface{})) {
	s := InterfaceSlice(slice)
	for v := range s {
		function(v)
	}
}
