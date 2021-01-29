package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// The package name is the same as the last element of the import path.

func main() {
	fmt.Println("Time is : ", time.Now())
	fmt.Println("You are number ", rand.Intn(100), " in line")
	fmt.Println("Your password is ", math.Pi)
	fmt.Println("SQRT of 7 is ", math.Sqrt(7))
}
