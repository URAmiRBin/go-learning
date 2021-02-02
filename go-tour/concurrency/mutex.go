package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.mu.Lock()
	s.v[key]++
	s.mu.Unlock()
}

func (s *SafeCounter) Val(key string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	// Removing sleep will give a number less than 1000
	// because all go routines are not done yet
	// and some are waiting in line for lock
	time.Sleep(time.Microsecond)
	fmt.Println(c.Val("somekey"))
}
