package main

import (
	"fmt"
)

func main() {
	var current, pos, i int
	positions := [7]int{1, 10, 100, 1000, 10000, 100000, 1000000}
	var digits [7]int
	var s string
	for ; ; i++ {
		s = fmt.Sprint(i)
		if pos <= positions[current] && positions[current] < pos+len(s) {
			digits[current] = int(s[positions[current]-pos]) - 48
			fmt.Printf("Digit %d found in number %d at position %d\n", digits[current], i, positions[current])
			current++
		}
		pos += len(s)
		if current == 7 {
			break
		}
	}
	fmt.Println(digits)
	mul := 1
	for _, d := range digits {
		mul *= d
	}
	fmt.Printf("Product of digits: %d\n", mul)
}
