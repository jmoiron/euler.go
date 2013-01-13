package main

import (
	"fmt"
	"math"
)

func isSquare(i int) bool {
	// Quickfail
	if i < 1 || ((i & 2) != 0) || ((i & 7) == 5) || ((i & 11) == 8) {
		return false
	}
	test := int(math.Sqrt(float64(i)))
	return test*test == i
}

func main() {
	triplets := [][3]int{}
	for i := 1; i < 500; i++ {
		for j := i; j < 500; j++ {
			k := i*i + j*j
			if !isSquare(k) {
				continue
			}
			k = int(math.Sqrt(float64(k)))
			if i+j+k > 1000 {
				break
			}
			triplets = append(triplets, [3]int{i, j, k})
		}
	}
	counts := map[int]int{}
	for _, t := range triplets {
		sum := t[0] + t[1] + t[2]
		counts[sum] = counts[sum] + 1
	}
	max, pos := 0, 0
	for k, v := range counts {
		if v > max {
			pos = k
			max = v
		}
	}
	fmt.Printf("%d triplets add up to %d\n", max, pos)
}
