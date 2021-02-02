package main

import (
	"fmt"
	"math/cmplx"
)

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte alias for uint8

// rune alias for int32
//      represents a Unicode code point

// float32 float64

// complex64 complex128

var (
	ToBe    bool = true
	NotToBe bool
	MaxInt  uint64     = 1<<64 - 1
	z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", NotToBe, NotToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var i int
	var f float64
	var b bool
	var s string = "sit"
	var c string
	// %q shows the "" mark, so it prints "" if it's default 0(empty) but %v prints nothing
	fmt.Printf("%v %v %v %v %q\n", i, f, b, s, c)
}
