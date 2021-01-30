package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%v with age %v", p.Name, p.Age)
}

type Stringer interface {
	String() string
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Im an int %v\n", v)
	case string:
		fmt.Printf("Im a string %v\n", v)
	default:
		fmt.Printf("I don't know who I am %v\n", v)
	}
}

func main() {
	// TYPE ASSERTION
	var i interface{} = "hello"
	// assigns underlying content of i to s
	s, ok := i.(string)
	fmt.Println(s, ok)

	// f := i.(float64) panics
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// TYPE SWITCH
	do(21)
	do("hello")

	// STRINGER
	// new returns a pointer
	p := new(Person)
	p.Name = "Amir"
	p.Age = 23
	// Automatically calls stringer
	fmt.Println(p)
}
