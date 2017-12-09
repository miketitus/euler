package main

import (
	"fmt"
	"math"
)

type node struct {
	value uint64
	left  *node
	right *node
}

var input uint64

func main() {
	input = 600851475143 // cannot use a const because overflow :(
	root := factorTree(input)
	printTree(root, 0)
}

func isPrimeSqrt(value uint64) bool {
	// copied from www.thepolyglotdeveloper.com/2016/12/determine-number-prime-using-golang/
	if value == 1 {
		return true
	}
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%uint64(i) == 0 {
			return false
		}
	}
	return value > 1
}

func getDivisor(dividend uint64) uint64 {
	var divisor uint64 = 1
	sqrt := math.Floor(math.Sqrt(float64(dividend)))
	for i := uint64(sqrt); i > 0; i-- {
		if isPrimeSqrt(i) && dividend%i == 0 {
			divisor = i
			break
		}
	}
	return divisor
}

func factorTree(i uint64) *node {
	var n *node
	if isPrimeSqrt(i) {
		n = &node{value: i, left: nil, right: nil}
	} else {
		leftDiv := getDivisor(i)
		rightDiv := i / leftDiv
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
