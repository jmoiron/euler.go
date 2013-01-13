package main

import (
	"fmt"
	"sort"
	"strconv"
)

/* prime map generator, from 027.go */

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
	// create a map of primes up to 1000000,
	// but it's unlikely we'll reach very high values for n
	in := make(chan int)
	go sieve(in, 1000000)
	for next := range in {
		primes[next] = 1
	}
}

func truncations(s string) []string {
	ss := []string{s}
	for i := 1; i < len(s); i++ {
		ss = append(ss, s[i:])
	}
	for i := len(s) - 1; i > 0; i-- {
		ss = append(ss, s[:i])
	}
	return ss
}

func isPrime(s string) bool {
	i, _ := strconv.Atoi(s)
	_, ok := primes[i]
	return ok
}

func all_prime(ss []string) bool {
	for _, v := range ss {
		if !isPrime(v) {
			return false
		}
	}
	return true
}

func main() {
	matches := []int{}
	for k, _ := range primes {
		ts := truncations(strconv.Itoa(k))
		if len(ts) == 1 {
			continue
		}
		//fmt.Printf("%d truncations: %v\n", k, ts)
		if all_prime(ts) {
			matches = append(matches, k)
		}
		if len(matches) == 11 {
			break
		}
	}
	sort.Ints(matches)
	sum := 0
	for _, v := range matches {
		sum += v
	}
	fmt.Printf("Found %d matches for rotatable primes: %v\n", len(matches), matches)
	fmt.Printf("Sum: %d\n", sum)
}
