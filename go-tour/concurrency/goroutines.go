package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// Note: Go runs in the same address space
	// so access to shared memory should be synchronized using sync package
	go say("world")
	say("hello")
}
