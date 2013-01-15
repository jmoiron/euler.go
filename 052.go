package main

import (
	"fmt"
	"strconv"
)

func isPermutation(i, j int) bool {
	fti, ftj := map[rune]int{}, map[rune]int{}
	si, sj := strconv.Itoa(i), strconv.Itoa(j)
	if len(si) == len(sj) {
		for _, v := range si {
			fti[v] += 1
		}
		for _, v := range sj {
			ftj[v] += 1
		}
		for k, v := range fti {
			if ftj[k] != v {
				return false
			}
		}
		return true
	}
	return false
}

func main() {
	for i := 1; ; i++ {
		n, n2, n3, n4, n5, n6 := i, i*2, i*3, i*4, i*5, i*6
		if isPermutation(n, n2) && isPermutation(n, n3) && isPermutation(n, n4) && isPermutation(n, n5) && isPermutation(n, n6) {
			fmt.Printf("%d\n", n)
			break
		}
	}
}
