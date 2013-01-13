package main

import (
	"fmt"
	"strconv"
)

func is_palindrome(s string) bool {
	var i, j int
	j = len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func is_dual_palindrome(i int) bool {
	b10 := strconv.FormatInt(int64(i), 10)
	b02 := strconv.FormatInt(int64(i), 2)
	return is_palindrome(b10) && is_palindrome(b02)
}

func main() {
	dual_palindromes := []int{}
	for i := 0; i < 1000000; i++ {
		if is_dual_palindrome(i) {
			dual_palindromes = append(dual_palindromes, i)
		}
	}
	fmt.Println(dual_palindromes)
	sum := 0
	for _, v := range dual_palindromes {
		sum += v
	}
	fmt.Printf("Sum: %d\n", sum)
}
