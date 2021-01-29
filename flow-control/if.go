package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func powWithLim(x, y float64, lim int) float64 {
	if v := math.Pow(x, y); v < float64(lim) {
		return v
	} else {
		return float64(lim)
	}
}

func main() {
	fmt.Println(sqrt(20), sqrt(-20))
	fmt.Println(powWithLim(2, 3, 10))
	fmt.Println(powWithLim(3, 3, 10))
}
