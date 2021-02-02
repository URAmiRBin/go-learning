package main

import (
	"fmt"
)

type Vertex struct {
	X float64
	Y float64
}

// Methods are functions with receiver
// For a big struct it's best practice to use pointer for methods
// this Mirror for example copies a new Vertex as input
func (v Vertex) Mirror() Vertex {
	return Vertex{-v.X, -v.Y}
}

// Methods that modify the receiver should use its' pointer
func (v *Vertex) Scale(scalar int) {
	v.X *= float64(scalar)
	v.Y *= float64(scalar)
}

// Methods can be defined for non struct types
type MyFloat float64

func (f MyFloat) Abs() MyFloat {
	if f < 0 {
		return -f
	}
	return f
}

// Turn method into function by just giving f MyFloat as input to Abs

func main() {
	v := Vertex{1, 2}
	p := &v
	fmt.Println(v.Mirror())

	v.Scale(4)
	fmt.Println(v)

	// Can use p.Scale(4) instead of (*p)
	(*p).Scale(4)
	fmt.Println(v)

	f := MyFloat(-4.6)
	fmt.Println(f.Abs())
}
