package maps

import (
	"fmt"
	"reflect"
)

func PrintMap(m interface{}) {
	v := reflect.ValueOf(m)

	if v.Kind() != reflect.Map {
		panic("Parameter must be map.")
	}

	fmt.Printf("%v\n", m)
}
