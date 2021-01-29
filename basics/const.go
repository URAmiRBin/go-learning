package main

import "fmt"

//Constants cannot be declared using the := syntax.

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const Pi = 3.14
	const Truth = true
	fmt.Printf("Pi is %v and the truth is %v\n", Pi, Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))
}
