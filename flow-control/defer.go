package main

import "fmt"

func main() {
	defer fmt.Printf(" World\n")
	fmt.Printf("Hello ")
}
