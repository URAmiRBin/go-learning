package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	// When given no other arguments, buffer size is 0
	// channel sends after receiving
	mychannel := make(chan int)

	go sum(ints[:3], mychannel)
	go sum(ints[3:], mychannel)

	x, y := <-mychannel, <-mychannel
	fmt.Println(x + y)

	bufferedChannel := make(chan int, 2)
	bufferedChannel <- 1
	bufferedChannel <- 2
	// This returns error
	// bufferedChannel <- 3

	// If for example 1 is read from channel, then 3 can be received
}
