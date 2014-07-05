package main

import "fmt"

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

func main() {
	prime_sum := uint64(0)
	ch := make(chan int)
	go sieve(ch, 2000000)
	for prime := range ch {
		prime_sum += uint64(prime)
	}
	fmt.Println(prime_sum)
}
