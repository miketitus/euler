package main

import (
	"fmt"
)

func main() {
	var length uint64 = 16
	steps := length * 2
	fs := factorial(steps)
	fl := factorial(length)
	fls := fl * fl // this overflows :(
	fmt.Printf("fs = %d, fl = %d, fls = %d\n", fs, fl, fls)
	routes := fs / fls
	fmt.Printf("steps = %d, length = %d, routes = %d\n", steps, length, routes)
}

func factorial(i uint64) uint64 {
	if i == 1 {
		return 1
	}
	return i * factorial(i-1)
}
