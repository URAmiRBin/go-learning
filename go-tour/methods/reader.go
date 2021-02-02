package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		// Read gets a b and fills it with data and returns size and err
		n, err := r.Read(b)
		fmt.Printf("n = %v | err = %v | b = %v\n", n, err, b)
		fmt.Printf("This package: %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
