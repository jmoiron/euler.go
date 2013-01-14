package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func permute(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	perms := []string{}
	head := string(s[0])
	tail := s[1:]

	for _, perm := range permute(tail) {
		for i := 0; i < len(s); i++ {
			newperm := perm[:i] + head + perm[i:]
			perms = append(perms, newperm)
		}
	}
	return perms
}

func atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func subprime(s string) bool {
	tests := []int{1, 2, 3, 5, 7, 11, 13, 17}
	for i, j := 1, 4; j <= len(s); i++ {
		if atoi(s[i:j])%tests[i] != 0 {
			return false
		}
		j++
	}
	return true
}

func main() {
	// 0 to 9 pandigitals can overflow 32bit ints so..
	matches := []*big.Int{}
	digits := "9876543210"
	for _, perm := range permute(digits) {
		if subprime(perm) {
			num := &big.Int{}
			num.SetString(perm, 10)
			matches = append(matches, num)
		}
	}
	sum := &big.Int{}
	for _, i := range matches {
		sum.Add(sum, i)
	}
	fmt.Println(sum)
}
