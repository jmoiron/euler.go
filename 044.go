package main

import (
	"fmt"
)

var pentas = []int{}
var pentamap = map[int]bool{}

func init() {
	for i := 1; i < 10000; i++ {
		p := (i * ((3 * i) - 1)) / 2
		pentas = append(pentas, p)
		pentamap[p] = true
	}

}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func ptest(i, n int) bool {
	if _, ok := pentamap[i+n]; ok {
		if _, ok := pentamap[abs(i-n)]; ok {
			return true
		}
	}
	return false
}

func main() {
	for i := 0; i < len(pentas); i++ {
		for j := 1; j < len(pentas); j++ {
			pi, pj := pentas[i], pentas[j]
			if ptest(pi, pj) {
				fmt.Println(abs(pi - pj))
				return
			}
		}
	}
}
