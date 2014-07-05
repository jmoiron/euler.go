package main

import "fmt"

/* The go tutorial goes over this parallelized algorithm for a prime sieve
   http://golang.org/doc/go_tutorial.html
*/

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)
	/* this creates 10000 channels, all filtering for divisibility by their prime */
	for i := 0; i < 10000; i++ {
		prime := <-ch
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	prime := <-ch
	fmt.Println(prime)
}
