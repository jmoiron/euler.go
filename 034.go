package main

import "fmt"

var digit_fact = map[rune]int{}

//'0':1, '1':1, '2':2, '3':6, '4':24, '5':120,
//'6': 720, '7': 5040, '8':40320, '9':362880}

func init() {
	digit_fact['0'] = 1
	digit_fact['1'] = 1
	digit_fact['2'] = 2
	fact := 2
	for i := 3; i < 10; i++ {
		digit_fact[rune(i+48)] = i * fact
		fact = i * fact
	}
}

func factsum(n int) int {
	str := fmt.Sprintf("%d", n)
	sum := 0
	for _, r := range str {
		sum += digit_fact[r]
	}
	return sum
}

func main() {
	/* 326880 * 7 is a simple upper bound; better are available  */
	sum := 0
	for i := 13; i < 7*362880; i++ {
		fs := factsum(i)
		if fs == i {
			println(i, fs)
			sum += i
		}
	}
	println("Sum:", sum)
}
