package main

import "fmt"

type Vertex struct {
	X    int // public
	Y    int
	name string // private
}

var (
	v1 = Vertex{1, 2, "v1"}
	v2 = Vertex{X: 1, name: "v2"}
	v3 = Vertex{}
	p  = &Vertex{1, 2, "vp"}
)

func main() {
	v := Vertex{1, 2, "z"}
	v.X = 4
	fmt.Println(v)

	p := &v
	// can use p.X for convenience instead of (*p).X
	p.X = 12
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
