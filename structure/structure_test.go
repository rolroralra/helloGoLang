package structure

import (
	"fmt"
	"testing"
)


func Test(t *testing.T) {
	// Structure Literal
	//v := Vertex{1, 2}
	v := Vertex{Y: 3}
	fmt.Println(v)
	p := &v

	p.X = 3
	p.Y = 4
	fmt.Println(p)
	fmt.Println(v)
	fmt.Println(p.Length())
}
