package main

import (
	"fmt"
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
	i, j := big.NewInt(1), big.NewInt(1)
	for j.Cmp(big.NewInt(100)) != 0 {
		j.Add(j, big.NewInt(1))
		i.Mul(i, j)
	}
	fmt.Println(sum_digits(i.String()))
}
