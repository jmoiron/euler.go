package main

import "fmt"

func main() {
	sum_sq := uint64(0)
	sq_sum := uint64(0)

	for i := 0; i <= 100; i++ {
		sum_sq += (uint64(i) * uint64(i))
		sq_sum += uint64(i)
	}

	sq_sum *= sq_sum

	fmt.Printf("%d - %d = %d\n", sq_sum, sum_sq, sq_sum-sum_sq)

}
