package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	// Declared variables must be used
	var j, k int = 1, 2
	var golang, scala, kotlin = true, false, "maybe"
	haskell := "we don't do that here"
	fmt.Println(i, c, python, java)
	fmt.Println(golang, scala, kotlin)
	fmt.Println(j, k)
	fmt.Println(haskell)
}
