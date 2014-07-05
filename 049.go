package main

import (
	"fmt"
	"strconv"
)

var maxprime = 10000
var primes = []int{}
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
	in := make(chan int)
	go sieve(in, maxprime)
	for next := range in {
		primes = append(primes, next)
		primemap[next] = true
	}
}

func isPermutation(i, j int) bool {
	fti, ftj := map[rune]int{}, map[rune]int{}
	si, sj := strconv.Itoa(i), strconv.Itoa(j)
	if len(si) == len(sj) {
		for _, v := range si {
			fti[v] += 1
		}
		for _, v := range sj {
			ftj[v] += 1
		}
		for k, v := range fti {
			if ftj[k] != v {
				return false
			}
		}
		return true
	}
	return false
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {

	fourdigit := []int{}
	for i := 0; i < len(primes); i++ {
		if primes[i] > 1000 {
			fourdigit = append(fourdigit, primes[i])
		}
	}

	for i := 0; i < len(fourdigit); i++ {
		ip := fourdigit[i]
		for j := i + 1; j < len(fourdigit); j++ {
			jp := fourdigit[j]
			n := jp + (jp - ip)
			if _, ok := primemap[n]; ok {
				if isPermutation(ip, jp) && isPermutation(jp, n) {
					if ip == 1487 {
						continue
					}
					fmt.Printf("%d%d%d\n", ip, jp, n)
				}
			}
		}
	}

}
