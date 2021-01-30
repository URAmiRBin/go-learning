package main

import "fmt"

func fib(c, quit chan int) {
	prepre, pre := 0, 1
	for {
		// blocks until one of the cases can run
		// run one randomly if multiple can run
		select {
		case c <- prepre:
			prepre, pre = pre, prepre+pre
		case <-quit:
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib(c, quit)
}
