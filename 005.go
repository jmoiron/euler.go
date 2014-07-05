package main

import "fmt"

func is_divisible(n uint64) bool {
	primes := []int{3, 7, 9, 11, 13, 17, 19}
	for _, i := range primes {
		if (n % uint64(i)) != 0 {
			return false
		}
	}
	rest := []int{4, 8, 12, 14, 16, 18}
	for _, i := range rest {
		if (n % uint64(i)) != 0 {
			return false
		}
	}
	return true
}

func main() {
	smallest := uint64(2520)

	for i := smallest; ; i += 10 {
		if is_divisible(i) {
			fmt.Printf("%d\n", i)
			return
		}
	}

}
