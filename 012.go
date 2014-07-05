package main

import "fmt"

func num_divisors(n uint64) uint64 {
	count := uint64(8)
	if n%uint64(3) == 0 {
		count += 1
	}
	if n%uint64(7) == 0 {
		count += 1
	}
	for i := uint64(11); i < n/2; i++ {
		if (n % i) == 0 {
			count += 1
		}
	}
	return count
}

func triangles(ch chan uint64) {
	tri := uint64(0)
	for i := uint64(1); ; i++ {
		tri += i
		ch <- tri
	}
}

func filter(in, out chan uint64) {
	for {
		n := <-in
		if n%uint64(2) != 0 {
			continue
		}
		if n%uint64(3) != 0 {
			continue
		}
		if n%uint64(4) != 0 {
			continue
		}
		if n%uint64(5) != 0 {
			continue
		}
		//if n % uint64(7) != 0 { continue }
		//if n % uint64(8) != 0 { continue }
		//if n % uint64(9) != 0 { continue }
		out <- n
	}
}

func main() {
	triangle := uint64(0)
	ch := make(chan uint64)
	out := make(chan uint64)
	go triangles(ch)
	go filter(ch, out)
	for triangle := range out {
		divs := num_divisors(triangle)
		fmt.Printf("%d  (%d)\n", triangle, divs)
		if divs > 500 {
			fmt.Println(triangle)
			break
		}
	}
	fmt.Println(triangle)
}
