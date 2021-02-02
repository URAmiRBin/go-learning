package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// OR
	// os := runtime.GOOS
	// switch os { ... }
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Good! I envy you")
	default:
		fmt.Printf("%s\n", os)
	}

	today := time.Now().Weekday()

	switch today {
	case time.Saturday:
		fmt.Println("Today is saturday")
	case time.Sunday:
		fmt.Println("Today is sunday")
	case time.Monday:
		fmt.Println("Today is monday")
	case time.Tuesday:
		fmt.Println("Today is tue")
	case time.Wednesday:
		fmt.Println("Today is wed")
	case time.Thursday:
		fmt.Println("Today is th")
	case time.Friday:
		fmt.Println("Today is friday")
	}

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is saturday")
	case today + 1:
		fmt.Println("Tomorrow is saturday")
	default:
		fmt.Println("Saturday is long to come")
	}

	// Switch with no conditions for if then else chains
	x := 10
	switch {
	case x < 10:
		fmt.Println("It's less than 10")
	case x > 10:
		fmt.Println("It's more than 10")
	default:
		fmt.Println("It's 10")
	}
}
