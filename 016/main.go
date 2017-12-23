package main

import (
	"fmt"
)

func main() {
	sum := 0.0
	input := 2.0
	power := input
	for i := 2; i <= 50; i++ {
		if power > 1 {
			power *= (input / 10.0)
		} else {
			power *= input
		}
		fmt.Println(power)
	}
	/*n := math.Pow(2, 15)
	fmt.Println(n)
	for n >= 1 {
		r := int(n) % 10
		sum += r
		n = n / 10
	}*/
	fmt.Println(sum)
}
