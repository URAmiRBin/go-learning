package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

// By implementing Error() ErrNegativeSqrt becomes an error
// as well as a float64
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Input should be positive, given: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
