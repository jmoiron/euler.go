package main

import "fmt"

func ceildiv(n, d int) int {
	if n%d == 0 {
		return n / d
	}
	return (n / d) + 1
}

func diagonal_sum(spiral [][]int) int {
	size := len(spiral)
	sum := 0
	for i, j := 0, 0; i < size; {
		sum += spiral[i][j]
		i++
		j++
	}
	for i, j := 0, size-1; i < size; {
		sum += spiral[i][j]
		i++
		j--
	}
	// adjust for counting the middle twice
	sum -= 1
	return sum
}

func main() {
	size := 1001

	spiral := make([][]int, size)
	for i, _ := range spiral {
		spiral[i] = make([]int, size)
	}

	toggle := 0
	moves := 1
	// 0 = right, 1 = down, 2 = left, 3 = up
	direction := 0
	v, i, j := 1, size/2, size/2

	for v < size*size {
		for n := 0; n < moves; n++ {
			spiral[i][j] = v
			switch direction {
			case 0:
				j++
			case 1:
				i++
			case 2:
				j--
			case 3:
				i--
			}
			v++
		}
		direction = (direction + 1) % 4
		if toggle > 0 {
			moves++
			toggle = 0
		} else {
			toggle++
		}
	}

	fmt.Println(diagonal_sum(spiral))
}
