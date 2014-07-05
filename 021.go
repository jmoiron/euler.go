package main

import "fmt"

func divisor_sum(n int) int {
	sum := 1
	for i := 2; i < (n/2)+1; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	amicable_sum := 0
	for i := 3; i < 10000; i++ {
		ds := divisor_sum(i)
		pair := divisor_sum(ds)

		if i == pair && i != ds {
			fmt.Printf("%d is amicable to %d\n", i, ds)
			amicable_sum += i
		}
	}
	fmt.Printf("Sum of amicable numbers < 10000: %d\n", amicable_sum)
}
