package maps

import (
	"fmt"
	"testing"
)

type Vertex struct {
	x, y int
}

func Test(t *testing.T) {
	m := make(map[string]Vertex)
	m["a"] = Vertex{x: 1}
	PrintMap(m)

	// Map Literal
	mapLiteral := map[string]Vertex{
		"a": Vertex{x: 1, y: 2},
		"b": Vertex{1, 2},
	}
	mapLiteral2 := map[string]Vertex{
		"a": {x: 1, y: 2},
		"b": {1, 2},
	}

	PrintMap(mapLiteral)
	PrintMap(mapLiteral2)
}

func TestForRange(t *testing.T) {
	m := map[string]Vertex{
		"a": Vertex{x: 1, y: 2},
		"b": Vertex{1, 2},
	}

	for key, value := range m {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}

func TestGetValue(t *testing.T) {
	m := map[string]int{}

	m["a"] = 1
	m["b"] = 2

	// Delete key in map
	delete(m, "b")

	PrintMap(m)

	elem, ok := m["b"]
	fmt.Println(elem, ok)

	elem, ok = m["a"]
	fmt.Println(elem, ok)
}

// map delete : delete(m, index)
// slice delete : s = append(s[:i], s[i+1:])
func TestDeleteKey(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for key, value := range m {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	delete(m, "b")
	PrintMap(m)

}
