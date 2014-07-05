package main

import (
	"fmt"
	"math/big"
)

func pow(a, b int) *big.Int {
	mul := big.NewInt(int64(a))
	result := big.NewInt(int64(a))
	for i := 1; i < b; i++ {
		result.Mul(result, mul)
	}
	return result
}

func main() {
	bigmap := map[string]int{}

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			num := pow(a, b)
			bigmap[num.String()] = 1
		}
	}

	fmt.Println(len(bigmap))

}
