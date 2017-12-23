package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1.0; i < 16.0; i++ {
		fmt.Println(sumDigits(math.Pow(2, i)))
	}
}

func sumDigits(n float64) float64 {
	sum := 0.0
	for n >= 1 {
		r := float64(int(n) % 10)
		sum += r
		n = n / 10
	}
	return sum
}
