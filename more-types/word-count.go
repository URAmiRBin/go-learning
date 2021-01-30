package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	result := make(map[string]int)

	for i := 0; i < len(words); i++ {
		elem, ok := result[words[i]]
		if ok {
			result[words[i]] = elem + 1
		} else {
			result[words[i]] = 1
		}
	}
	return result
}

func main() {
	fmt.Println(WordCount("Salam Khubi Salam merci"))
}
