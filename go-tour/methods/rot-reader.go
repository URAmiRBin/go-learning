package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	fmt.Printf("%v\n", n)
	if err != nil {
		return n, err
	}
	size := n
	for i := 0; i < size; i++ {
		if b[i] < 91 && b[i] > 64 {
			if b[i]+13 < 91 {
				b[i] += 13
			} else {
				b[i] -= 13
			}
		} else {
			if b[i]+13 < 122 {
				b[i] += 13
			} else {
				b[i] -= 13
			}
		}

	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
