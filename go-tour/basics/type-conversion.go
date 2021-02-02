package main

import (
	"fmt"
	"math"
)

func main() {
	i := 69
	// Conversions should be explicit to change the real type
	j := float64(i)
	u := uint(j)
	fmt.Println(i, j, u)

	x, y := 3, 5
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)

	// Go also have type inference
	g := 3.02 + 5i
	fmt.Printf("type of %v is %T\n", g, g)
}
