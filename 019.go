package main

import "fmt"

var months = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
var month_len = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var leap_len = []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func main() {
	sum := 0
	sunday := 5
	year := 1901
	is_leap := false
	for year < 2001 {
		checker := month_len
		if is_leap {
			checker = leap_len
		}
		tally := sunday
		for _, v := range checker {
			if tally == 0 {
				sum += 1
			}
			if tally < v {
				break
			}
			if tally >= v {
				tally -= v
			}
		}
		sunday += 7
		if is_leap && sunday >= 366 {
			year++
			sunday = sunday % 366
			is_leap = false
		} else if !is_leap && sunday >= 365 {
			year++
			sunday = sunday % 365
			if year%4 == 0 {
				is_leap = true
			}
		}
	}
	fmt.Printf("Year: %d, num sundays: %d\n", year, sum)
}
