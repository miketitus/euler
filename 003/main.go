package main

import (
	"fmt"
	"math"
)

type node struct {
	value int
	left  *node
	right *node
}

const input = 13195

var isPrime = [input + 1]bool{}

func main() {
	computePrimes()
	root := factorTree(input)
	printTree(root, 0)
}

func computePrimes() {
	var x, y, n int
	nsqrt := math.Sqrt(input)
	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= input && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) + y*y
			if n <= input && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= input && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}
	isPrime[1] = true
	isPrime[2] = true
	isPrime[3] = true
	for n = 5; float64(n) <= nsqrt; n++ {
		if isPrime[n] {
			for y = n * n; y < input; y += n * n {
				isPrime[y] = false
			}
		}
	}
}

func getDivisor(dividend int) int {
	divisor := 1
	sqrt := math.Floor(math.Sqrt(float64(dividend)))
	//fmt.Printf("%f is %d\n", sqrt, int(sqrt))
	for i := int(sqrt); i > 0; i-- {
		if isPrime[i] && dividend%i == 0 {
			divisor = i
			break
		}
	}
	return divisor
}

func factorTree(i int) *node {
	//fmt.Println("factorTree: ", i)
	var n *node
	if isPrime[i] {
		// prime
		n = &node{value: i, left: nil, right: nil}
	} else {
		leftDiv := getDivisor(i)
		rightDiv := i / leftDiv
		//fmt.Printf("leftDiv %d, rightDiv %d\n", leftDiv, rightDiv)
		lNode := factorTree(leftDiv)
		rNode := factorTree(rightDiv)
		n = &node{value: i, left: lNode, right: rNode}
	}
	return n
}

func printTree(n *node, indent int) {
	for i := 0; i < indent*2; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%d\n", n.value)
	if n.left != nil {
		printTree(n.left, indent+1)
	}
	if n.right != nil {
		printTree(n.right, indent+1)
	}
}
