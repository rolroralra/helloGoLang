package primitive

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func PrintVariableInfo(i interface{}) {
	r := reflect.ValueOf(i)

	fmt.Printf("Type: %T, Size: %v, DefaultValue: %v\n", i, r.Type().Size(), i)
}

func TestUnsafeSizeof(t *testing.T) {
	var (
		int8Variable       int8
		int16Variable      int16
		int32Variable      int32
		int64Variable      int64
		intVariable        int
		uint8Variable      uint8
		uint16Variable     uint16
		uint32Variable     uint32
		uint64Variable     uint64
		uintVariable       uint
		float32Variable    float32
		float64Variable    float64
		complex64Variable  complex64
		complex128Variable complex128
		byteVariable       byte
		runeVariable       rune
		stringVariable     string
		arrayVariable      = [...]int{1, 2, 3, 4, 5}
	)

	fmt.Printf("Type: %T, Size: %v, DefaultValue: %v\n", int8Variable, unsafe.Sizeof(int8Variable), int8Variable)
	fmt.Printf("Type: %T, Size: %v, DefaultValue: %v\n", int16Variable, unsafe.Sizeof(int16Variable), int8Variable)
	fmt.Printf("Type: %T, Size: %v, DefaultValue: %v\n", int32Variable, unsafe.Sizeof(int32Variable), int8Variable)
	fmt.Printf("Type: %T, Size: %v, DefaultValue: %v\n", int64Variable, unsafe.Sizeof(int64Variable), int8Variable)
	fmt.Printf("Type: %T, Size: %v, DefaultValue: %v\n", intVariable, unsafe.Sizeof(intVariable), int8Variable)
	PrintVariableInfo(uint8Variable)
	PrintVariableInfo(uint16Variable)
	PrintVariableInfo(uint32Variable)
	PrintVariableInfo(uint64Variable)
	PrintVariableInfo(uintVariable)
	PrintVariableInfo(float32Variable)
	PrintVariableInfo(float64Variable)
	PrintVariableInfo(complex64Variable)
	PrintVariableInfo(complex128Variable)
	PrintVariableInfo(byteVariable)
	PrintVariableInfo(runeVariable)
	PrintVariableInfo(stringVariable)
	PrintVariableInfo(arrayVariable) // Array is also primitive type. But Slice is reference type.
}
