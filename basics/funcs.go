package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

// Multiple Results
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(9, 2))
	fmt.Println(multiply(9, 2))
	fmt.Println(swap("World!", "Hello"))
	fmt.Println(split(15))
}
