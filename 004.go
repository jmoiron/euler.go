package main

import (
	"fmt"
	"strconv"
)

func is_palindrome(i int64) bool {
	s := strconv.FormatInt(i, 10)
	i, j := int64(0), int64(len(s)-1)
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	product := int64(0)
	largest_product := int64(0)
	i, j := int64(999), int64(999)

	for i > 0 && j > 0 {
		product = i * j
		if is_palindrome(product) && product > largest_product {
			largest_product = product
			fmt.Printf("%d from %d x %d\n", product, i, j)
			i--
			j = 999
			if i*j < largest_product {
				break
			}
		} else if j > 1 {
			j -= 1
		} else {
			i--
			j = 999
		}
	}
}
