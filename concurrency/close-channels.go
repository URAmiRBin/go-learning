package main

import (
	"fmt"
)

func fib(n int, c chan int) {
	prepre, pre := 0, 1
	for i := 0; i < n; i++ {
		c <- prepre
		prepre, pre = pre, prepre+pre
	}
	// This must be called from sender
	// Indicates to receiver that I send all my things
	// and no longer am sending
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fib(cap(c), c)
	// range c is terminated when sender closes channel
	for i := range c {
		fmt.Println(i)
	}

	// v, ok : ok is false if channel is closed and there is no more values to receive
	c2 := make(chan int, 10)
	go fib(2, c2)
	for i := 0; i < 4; i++ {
		v, ok := <-c2
		fmt.Println(v, ok)
	}

	// Closing channel is only useful when the receiver needs to know when it's finished
	// like terminating a range loop
	// It's not necessary like files

}
