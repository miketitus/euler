package main

import (
	"fmt"
)

func main() {
	var length float64 = 20
	steps := length * 2
	fs := factorial(steps)
	fl := factorial(length)
	fls := fl * fl
	fmt.Printf("fs = %.0f, fl = %.0f, fls = %.0f\n", fs, fl, fls)
	routes := fs / fls
	fmt.Printf("steps = %.0f, length = %.0f, routes = %.0f\n", steps, length, routes)
}

func factorial(i float64) float64 {
	if i == 1 {
		return 1
	}
	return i * factorial(i-1)
}
