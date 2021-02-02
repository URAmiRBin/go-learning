package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	x := 0
	// for is while
	for x < 10 {
		x++
	}
	fmt.Println(x)

	// This runs forever
	// for {
	// 	fmt.println("Forever")
	// }
}
