package main

import (
	"fmt"
	"math"
)

func compute(f func(float64, float64) float64) float64 {
	return f(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	myfunc := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println((compute(math.Pow)))
	fmt.Println((compute(myfunc)))

	f1, f2 := adder(), adder()
	fmt.Println(f1(0), f2(0))
	fmt.Println(f1(1), f2(-2))
	fmt.Println(f1(2), f2(-4))
}
