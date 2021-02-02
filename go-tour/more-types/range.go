package main

import "fmt"

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64}

	for index, value := range pow {
		fmt.Printf("2^%d = %d\n", index, value)
	}

	fmt.Printf("========\n")

	// For convenience instead of index, _ can use index
	for index := range pow {
		fmt.Printf("index is %d\n", index)
	}

	fmt.Printf("========\n")

	for _, value := range pow {
		fmt.Printf("values is is %d\n", value)
	}

}
