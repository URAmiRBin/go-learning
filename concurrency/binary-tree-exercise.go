package main

import (
	"fmt"
	"sync"

	"golang.org/x/tour/tree"
)

var wg sync.WaitGroup

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Walktree(t, ch)
	close(ch)
}

func Walktree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walktree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walktree(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 == false && ok2 == false {
			break
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	t := &tree.Tree{&tree.Tree{nil, 2, nil}, 3, nil}
	c := make(chan int)
	go Walk(t, c)
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(2)))
}
