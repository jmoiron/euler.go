package main

import "fmt"

func chain_len(n uint64) uint {
	i := uint(1)
	for ; n > 1; i++ {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return i
}

func main() {
	longest_chain, longest_num := uint(0), uint64(0)
	for i := uint64(2); i < 1000000; i++ {
		length := chain_len(i)
		if length > longest_chain {
			longest_chain = length
			longest_num = i
			fmt.Printf("Leader: %d  len(%d)\n", i, length)
		}
	}
	fmt.Printf("Longest: %d len(%d)\n", longest_num, longest_chain)
}
