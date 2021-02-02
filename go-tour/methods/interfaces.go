package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

type Abser interface {
	Abs() float64
}

type MyFloat float64
type MyInt int32

func main() {
	var a Abser
	// a.Abs() gives error here; someone should implement it
	f := MyFloat(-2.3)
	v := Vertex{4, 3}

	a = f // MyFloat implements Abser
	fmt.Println(a.Abs())
	fmt.Printf("TYPE: %T\n", a)

	a = &v // Vertex implements Abser
	fmt.Println(a.Abs())
	fmt.Printf("TYPE: %T\n", a)

	// a = v does not work

	var b Abser = MyInt(-2)
	fmt.Println(b.Abs())
	fmt.Printf("TYPE: %T\n", b)

	var myvertex *Vertex
	b = myvertex
	fmt.Println(b.Abs())
	fmt.Printf("TYPE: %T\n", b)

	fmt.Println("==========")

	var emptyInterface interface{}
	fmt.Printf("%v, %T\n", emptyInterface, emptyInterface)

	emptyInterface = 14
	fmt.Printf("%v, %T\n", emptyInterface, emptyInterface)

	emptyInterface = "hello"
	fmt.Printf("%v, %T\n", emptyInterface, emptyInterface)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// This function works with &vertex
func (v *Vertex) Abs() float64 {
	// Unlike others, Go doesn't throw null ptr exception
	// when being called by a nil receiver
	if v == nil {
		fmt.Println("nil")
		return 0.0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Implicitly implements interface Abser
func (i MyInt) Abs() float64 {
	if i < 0 {
		return float64(-i)
	}
	return float64(i)
}
