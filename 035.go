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

func rotate(s string) string {
	if len(s) == 1 {
		return s
	}
	return string(s[len(s)-1]) + s[:len(s)-1]
}

func isPrime(s string) bool {
	i, _ := strconv.Atoi(s)
	_, ok := primes[i]
	return ok
}

func rotations(s string) []string {
	rs := []string{s}
	for r := rotate(s); r != s; r = rotate(r) {
		rs = append(rs, r)
	}
	return rs
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
	ceiling := 1000000
	for k, _ := range primes {
		if k > ceiling {
			continue
		}
		rs := rotations(strconv.Itoa(k))
		if all_prime(rs) {
			fmt.Printf("%d rotations all prime: %v\n", k, rs)
			matches = append(matches, k)
		}
	}
	sort.Ints(matches)
	fmt.Printf("Found %d rotatable primes under %d: %v\n", len(matches), ceiling, matches)
}
