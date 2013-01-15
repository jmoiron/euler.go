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

func sq(i int) int { return i * i }

func goldbach(n int) bool {
	var j, prod int

	for i := 0; i < numprimes && primes[i] < n; i++ {
		prod = 0
		for j = 1; prod+primes[i] < n; j++ {
			prod = 2 * sq(j)
		}
		if prod+primes[i] == n {
			//fmt.Printf("%d = %d + 2 x %d^2\n", n, primes[i], j-1)
			return true
		}
	}
	return false
}

func isPrime(n int) bool {
	_, ok := primemap[n]
	return ok
}

func main() {
	i := 9
	for goldbach(i) {
		i += 2
		for ; isPrime(i); i += 2 {

		}
	}
	fmt.Println(i)
}
