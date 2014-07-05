package main

import (
	"math/big"
	"strconv"
)

func sum_digits(s string) uint64 {
	sum := uint64(0)
	for _, c := range s {
		n, _ := strconv.Atoi(string(c))
		sum += uint64(n)
	}
	return sum
}

func main() {
	n := new(big.Int)
	n.Exp(big.NewInt(2), big.NewInt(1000), nil)
	println(sum_digits(n.String()))
}
