package main

import "fmt"

func main() {
	// A slice does not store any data, it just describes a section of an underlying array.
	names := [4]string{
		"Majima", "Kiryu", "Saejima", "Akiyama",
	}
	fmt.Println(names)

	x := names[0:2]
	y := names[1:3]
	fmt.Println(x, y)

	y[1] = "Ichiban"
	fmt.Println(names)

	// Slice Literals
	q := []int{2, 3, 4}
	fmt.Println(q)
	r := []struct {
		b bool
		i int
	}{
		{false, 2},
		{true, 3},
		{false, 5},
	}
	fmt.Println(r)

	// Slice defaults
	fmt.Println(names[1:])
	fmt.Println(names[:2])

	// Default Slice
	var e []int
	if e == nil {
		fmt.Println("It's nil")
	}

	// To build dunamically sized arrays make slices
	d := make([]int, 5)
	fmt.Println(d, " : ", "len:", len(d), " cap:", cap(d))

	d2 := make([]int, 0, 5)
	fmt.Println(d2, " : ", "len:", len(d2), " cap:", cap(d2))

	sliceOfEmpty := d2[0:2]
	fmt.Println(sliceOfEmpty, " : ", "len:", len(sliceOfEmpty), " cap:", cap(sliceOfEmpty))

	// 2D Slices
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][1] = "X"

	fmt.Println(board)

}
