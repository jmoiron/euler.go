package main

import "fmt"

var primes = map[int]int{}

func sieve(ch chan int, max int) {
	nums := make([]int, max+1)
	ch <- 2
	for i := 3; i < max; i += 2 {
		if nums[i] == 0 {
			ch <- i
			for n := i; n <= max; n += i {
				nums[n] = 1
			}
		}
	}
	close(ch)
}

func init() {
	// create a map of primes up to a value; 1000^2 is 1000000,
	// but it's unlikely we'll reach very high values for n
	in := make(chan int)
	go sieve(in, 5000)
	for next := range in {
		primes[next] = 1
	}
}

func prime_list(max int) []int {
	vals := []int{}
	in := make(chan int)
	go sieve(in, max)
	for p := range in {
		vals = append(vals, p)
	}
	return vals
}

// test how many numbers n^2 + an + b remains prime for
func quad_length(a, b int) int {
	for i := 0; ; i++ {
		v := i*i + a*i + b
		if primes[v] == 0 {
			return i
		}
	}
	return -1
}

// note we can significantly simplify the search by noting that when n=0,
// n^2 + an + b == b and thus b must be a (positive) prime number
func main() {
	bvals := prime_list(1000)
	maxlen := 0
	maxa, maxb := 0, 0

	for a := -999; a < 1000; a += 2 {
		for _, b := range bvals {
			if a > b {
				continue
			}
			ql := quad_length(a, b)
			if ql > maxlen {
				maxlen = ql
				maxa, maxb = a, b
			}
		}
	}
	fmt.Printf("%d x %d = %d\n", maxa, maxb, maxa*maxb)
}
