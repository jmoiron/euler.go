package main

import (
	"fmt"
	"strconv"
)

func fifth(i uint64) uint64 {
	return i * i * i * i * i
}

func sum_fifth(s string) uint64 {
	sum := uint64(0)
	for _, c := range s {
		n, _ := strconv.Atoi(string(c))
		sum += fifth(uint64(n))
	}
	return sum
}

func main() {
	sum := 0
	for i := 2; i < 2000000; i++ {
		str := strconv.Itoa(i)
		if sum_fifth(str) == uint64(i) {
			sum += i
			fmt.Printf("%d == sum_fifth(%d)\n", i, i)
		}
	}
	fmt.Printf("%d\n", sum)
}
