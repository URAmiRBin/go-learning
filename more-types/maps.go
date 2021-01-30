package main

import "fmt"

type Vertex struct {
	X float64
	Y float64
}

// Vertex can be omitted or not
var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		5, 6,
	},
	"Google": {
		6, 7,
	},
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{5.0, 6.0}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m2["Google"])

	// Add new key
	m2["My house"] = Vertex{2, 3}

	// Retrieve with elem = m["My house"]
	elem := m2["My house"]

	fmt.Println(elem)

	// Delete key
	delete(m2, "My house")

	// Printing what does not exists returns 0 value
	fmt.Println(elem)
	fmt.Println(m2["My house"])

	// To make sure if the key exists
	elem, ok := m2["My house"]
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("The given key does not exist")
	}

}
