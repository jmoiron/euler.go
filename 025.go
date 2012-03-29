package main

import (
	"fmt"
	"math/big"
)

// generate an endless amt of fibonacci numbers
func fib(ch chan int) {
	send := func(i big.Int) {
		ch <- len(i.String())
	}

	fibseq := [3]big.Int{}
	fibseq[0].Set(big.NewInt(1))
	fibseq[1].Set(big.NewInt(1))
	fibseq[2].Set(big.NewInt(2))
	send(fibseq[0])
	send(fibseq[1])
	send(fibseq[2])
	for {
		fibseq[0].Set(&fibseq[1])
		fibseq[1].Set(&fibseq[2])
		fibseq[2].Add(&fibseq[0], &fibseq[1])
		send(fibseq[2])
	}

}

// find the first fib w/ 1000 digits & print which fib it is
func main() {
	ord := 1
	ch := make(chan int, 100)
	go fib(ch)
	for next := range ch {
		if next >= 1000 {
			break
		}
		ord++
	}
	fmt.Println(ord)
}
