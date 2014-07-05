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

func abundant_numbers() []int {
	ans := []int{}
	for i := 10; i < 28123; i++ {
		if divisor_sum(i) > i {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	ans := abundant_numbers()
	ansmap := map[int]int{}
	for _, a := range ans {
		ansmap[a] = 1
	}
	sum := 1
	for i := 2; i < 28124; i++ {
		found := false
		// if i is the sum of two abundant n's, there exists an ans < i such
		// that i-a is an ans;  use the map to check this membership in O(1)
		for _, a := range ans {
			if a > i {
				break
			}
			if ansmap[i-a] != 0 {
				found = true
				break
			}
		}
		if !found {
			sum += i
		}
	}

	fmt.Println(sum)
}
