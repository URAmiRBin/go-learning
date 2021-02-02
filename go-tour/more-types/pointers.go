package main

import "fmt"

func main() {
	// Pointer holds the memory address of a value
	// Pointers zero value is nil

	// Unlike C, go has no pointer arithmetic: safer lang

	var p *int
	i := 42
	// p points to i
	p = &i

	fmt.Println(*p)
	fmt.Println(&i)
	fmt.Println(p)

	*p = *p + 2

	fmt.Println(*p)
	fmt.Println(i)
}
