package main

import (
	"fmt"
	"math/big"
)

func sq(i *big.Int) *big.Int {
	j := &big.Int{}
	j.Exp(i, i, nil)
	return j
}

func main() {
	sum := &big.Int{}
	for i := 1; i < 1001; i++ {
		bi := big.NewInt(int64(i))
		sum.Add(sum, sq(bi))
	}
	s := sum.String()
	fmt.Println(s[len(s)-10:])
}
