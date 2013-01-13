package main

import (
	"fmt"
	"strconv"
	"strings"
)

func is_pandigital(s string) bool {
	if len(s) != 9 {
		return false
	}
	set := map[rune]int{}
	for _, r := range s {
		if r == '0' || set[r] > 0 {
			return false
		}
		set[r] = 1
	}
	return true
}

func cat_product(i int, set []int) string {
	prods := []string{}
	for _, v := range set {
		prods = append(prods, strconv.Itoa(i*v))
	}
	return strings.Join(prods, "")
}

func main() {
	/* 9 * (1, 2, 3, 4, 5) is the biggest run possible */
	runs := [4][]int{
		[]int{1, 2},
		[]int{1, 2, 3},
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3, 4, 5},
	}

	pandigiprods := []string{}

	for i := 1; i < 50000; i++ {
		for _, r := range runs {
			prod := cat_product(i, r)
			if is_pandigital(prod) {
				fmt.Printf("%d x %v = %s\n", i, r, cat_product(i, r))
				pandigiprods = append(pandigiprods, prod)
			}
		}
	}
	fmt.Printf("Pandigiprods: %v\n", pandigiprods)

	max := 0
	for _, prod := range pandigiprods {
		num, _ := strconv.Atoi(prod)
		if num > max {
			max = num
		}
	}
	fmt.Printf("Max: %d\n", max)
}
