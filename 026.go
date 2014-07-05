package main

import "fmt"

/* use long division to see how long the cycle is */
func cycle_length(n, d int) int {
	set := map[int]int{}
	length := 0

	for n != 0 {
		if n < d {
			n *= 10
			continue
		}
		n = n % d
		// if the remainder has been encountered before, we've cycled
		if set[n] > 0 {
			return length
		} else {
			set[n] = 1
		}
		length += 1
	}
	// if we've gotten here, it doesn't cycle
	return 0
}

func main() {
	max := 0
	maxi := 0

	for i := 2; i < 1000; i++ {
		length := cycle_length(1, i)
		if length > max {
			max = length
			maxi = i
		}
	}

	fmt.Printf("%d has cycle length %d\n", maxi, max)
}
