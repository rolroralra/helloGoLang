package structure

import "math"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) Length() float64 {
	return math.Sqrt(float64(v.X * v.X + v.Y * v.Y))
}

func CreateVertex() Vertex {
	return Vertex{0, 0}
}

