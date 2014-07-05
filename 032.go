package main

import "fmt"

func is_9_pandigital(a, b, c int) bool {
	full := fmt.Sprintf("%d%d%d", a, b, c)
	if len(full) != 9 {
		return false
	}
	set := map[rune]int{}
	for _, r := range full {
		if r == '0' || set[r] > 0 {
			return false
		}
		set[r] = 1
	}
	return true
}

func main() {
	panprods := map[int]int{}
	for i := 2; i < 99; i++ {
		for j := i; j < 9999/i; j++ {
			if is_9_pandigital(i, j, i*j) {
				panprods[i*j]++
			}
		}
	}
	sum := 0
	for k, _ := range panprods {
		sum += k
	}
	println(sum)
}
