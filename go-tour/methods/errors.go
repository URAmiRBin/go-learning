package main

import (
	"fmt"
	"strconv"
	"time"
)

// built in fmt
// type error interface {
// 	Error() string
// }

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s happended at %v", e.What, e.When)
}

func run() error {
	return &MyError{
		time.Now(),
		"Somebody touched something",
	}
}

func main() {
	i, err := strconv.Atoi("4")
	if err != nil {
		fmt.Println("Couldn't convert number: %v\n", err)
	} else {
		fmt.Println("Converted integer : ", i)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}
}
