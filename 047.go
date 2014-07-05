package main

import (
	"fmt"
)

var numprimes = 100000
var primes = make([]int, numprimes)
var primemap = map[int]bool{}

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
	i := 0
	in := make(chan int)
	go sieve(in, numprimes)
	for next := range in {
		primes[i] = next
		primemap[next] = true
		i++
	}
}

func setlen(is []int) int {
	m := map[int]bool{}
	for _, v := range is {
		m[v] = true
	}
	length := 0
	for _, _ = range m {
		length++
	}
	return length
}

func factorize(n int) []int {
	// fast fail on primes
	if _, ok := primemap[n]; ok {
		return []int{n}
	}
	factors := []int{}
	var p int
	for i := 0; primes[i] <= n/2; i++ {
		p = primes[i]
		if n%p == 0 {
			n = n / p
			factors = append(factors, p)
			/* i will be ++'d after this and we want to start at the first prime */
			i = -1
		}
	}
	factors = append(factors, n)
	return factors
}

func main() {
	var run, i int
	for i = 10; run < 4; i++ {
		if setlen(factorize(i)) == 4 {
			run++
		} else {
			run = 0
		}
	}
	fmt.Println(i - 4)
}
