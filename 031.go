package main

import "fmt"

var values = []int{100, 50, 20, 10, 5, 2, 1}

func main() {
	// there's only one way of making it exclusively with one coin of each
	// value, and we can't use £2 more than this one way, so we'll disregard
	counts := []int{0, 0, 0, 0, 0, 0, 200}

	length := len(values)
	pennies := length - 1
	// start at 2, one for the starting position, and 1 for the single £2 coin
	tot := 2

	// the algorithm works like this;  all coins are arranged highest to lowest
	// value, with 200 pennies being the starting position as valid change for
	// 2GBP.  The left-most coin that can be subtracted from the pennies is then
	// incremented, and pennies decremented by that amount.  This position is
	// a "pivot", whose place is kept.

	// When the pennies reach below 2, the position to the left of the pivot
	// is increased, and everything between that and the end is zeroed out.
	// The remainder is set in pennies, and the algorithm starts up again.

	for x := 0; counts[0] < 2; x++ {
		if counts[pennies] < 2 {
			for i := pennies - 1; i > 0; i-- {
				if counts[i] > 0 && (counts[i-1]*values[i-1]) < 200 {
					// if we can increase the next highest coin which can receive more
					counts[i-1]++
					// adjust the pennies based on the value of that coin
					counts[pennies] += (values[i] * counts[i]) - (values[i-1])
					counts[i] = 0
					// because we checked only one coine value and not the
					// total sum, we can get an underflow in the pennies which
					// will be corrected in the next loop, so increase sum if
					// this did not occur
					if counts[pennies] >= 0 {
						tot += 1
					}
					break
				}
			}
		}
		for i := pennies - 1; i >= 0; i-- {
			if values[i] <= counts[pennies] {
				counts[i] += 1
				counts[pennies] -= values[i]
				tot += 1
				break
			}
		}
	}

	fmt.Println(tot)

}
