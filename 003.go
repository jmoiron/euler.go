package main

import "fmt"
import "math"

/* All primes aside from 2 and 3 are of the form 6k+-1 */
func not_prime(n int64) bool {
	if (n-1)%6 == 0 {
		return false
	}
	if (n+1)%6 == 0 {
		return false
	}
	return true
}

func is_prime(n int64) bool {
	if not_prime(n) {
		return false
	}
	max := int64(math.Ceil(math.Sqrt(float64(n))))
	for i := int64(2); i < max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := int64(600851475143)
	largest := int64(0)

	if (n % 2) == 0 {
		largest = 2
	}

	for i := int64(3); i < int64(math.Sqrt(float64(n))); i += 2 {
		if (n % i) == 0 {
			if is_prime(i) {
				fmt.Print(i, "\n")
				largest = i
			}
		}
	}
	fmt.Printf("Largest: %d\n", largest)

}
