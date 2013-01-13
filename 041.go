package main

import (
	"fmt"
	"math"
	"strconv"
)

func notPrime(n int) bool {
	if n&2 == 2 {
		return false
	}
	if (n-1)%6 == 0 {
		return false
	}
	if (n+1)%6 == 0 {
		return false
	}
	return true
}

func isPrime(n int) bool {
	if notPrime(n) {
		return false
	}
	max := int(math.Ceil(math.Sqrt(float64(n))))
	for i := int(3); i < max; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func permute(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	perms := []string{}
	head := string(s[0])
	tail := s[1:]

	for _, perm := range permute(tail) {
		for i := 0; i < len(s); i++ {
			newperm := perm[:i] + head + perm[i:]
			perms = append(perms, newperm)
		}
	}
	return perms
}

func main() {
	// there are less pandigital numbers than there are primes, so we filter based
	// on whether or not numbers are pandigital, starting with 9-digit pandigital
	// numbers

	// once we've found a prime, we know that the largest pandigital prime must have
	// that many digits, so we continue over the rest of the permutations in that
	// digit set and then know the max is the highest

	// each digit set has n! possible pandigital permutations (9! for 9 digits, 8!
	// for 8, and so on).  9! is only 362880, which isn't large for a computer.

	// using a divisibility rule for 3, we note that 9 and 8 digit pandigital
	// have digits adding to 45 (3*15) and 36 (3*12), thus are divisible by 3 and not prime

	max := 0
	digits := "7654321"
	for i := 0; i < len(digits); i++ {
		for _, perm := range permute(digits[i:]) {
			i := atoi(perm)
			if isPrime(i) {
				if i > max {
					max = i
				}
			} else {
			}
		}
		if max > 0 {
			break
		}
	}
	fmt.Printf("%d\n", max)
}
