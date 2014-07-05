package main

import (
	"fmt"
	"sort"
)

func next(nums []int) []int {

	last := len(nums) - 1
	tail := 0
	for i := 1; i <= last; i++ {
		if nums[i] > nums[i-1] {
			break
		}
		if i == last && nums[i] < nums[i-1] {
			return nil
		}
	}
	/* find the longest tail sorted in decreasing order */
	for tail = last; nums[tail-1] > nums[tail] && tail > 0; tail-- {
	}

	//fmt.Println("Tail", tail)

	if tail == last {
		nums[last], nums[last-1] = nums[last-1], nums[last]
		return nums
	}

	if tail >= 1 {
		min := 10
		mini := 0
		token := nums[tail-1]
		for i := tail; i < len(nums); i++ {
			if nums[i] > token && nums[i] < min {
				min = nums[i]
				mini = i
			}
		}
		nums[tail-1], nums[mini] = nums[mini], nums[tail-1]
		sort.Ints(nums[tail:])
		return nums

	}

	return nil
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 1; nums != nil; i++ {
		//fmt.Println(i, nums)
		if i == 1000000 {
			for f := 0; f < len(nums); f++ {
				fmt.Print(nums[f])
			}
			fmt.Println()
			return
		}
		nums = next(nums)
	}
}
