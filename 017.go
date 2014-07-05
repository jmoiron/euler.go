package main

var digits = map[int]int{
	1: 3, //"one",
	2: 3, //"two",
	3: 5, //"three",
	4: 4, //"four",
	5: 4, //"five",
	6: 3, //"six",
	7: 5, //"seven",
	8: 5, //"eight",
	9: 4, //"nine",
}

var tens = map[int]int{
	10: 3, //"ten",
	11: 6, //"eleven",
	12: 6, //"twelve",
	13: 8, //"thirteen",
	14: 8, //"fourteen",
	15: 7, //"fifteen",
	16: 7, //"sixteen",
	17: 9, //"seventeen",
	18: 8, //"eighteen",
	19: 8, //"nineteen",
	2:  6, //"twenty",
	3:  6, //"thirty",
	4:  5, //"forty",
	5:  5, //"fifty",
	6:  5, //"sixty",
	7:  7, //"seventy",
	8:  6, //"eighty",
	9:  6, //"ninety",
}

func num_letters(n int) int {
	// catch special cases like single digits or teens
	if digits[n] > 0 {
		return digits[n]
	}
	if tens[n] > 0 {
		return tens[n]
	}
	if n == 1000 {
		return 3 + 8
	}
	// special case "one hundred", "two hundred", etc
	if n%100 == 0 {
		return digits[n/100] + 7
	}
	if n > 100 {
		return digits[n/100] + 10 + num_letters(n%100)
	}
	if n%10 == 0 {
		return tens[n/10]
	}
	return tens[n/10] + digits[n%10]
}

func main() {
	sum := uint64(0)
	for i := 1; i <= 1000; i++ {
		println(i, num_letters(i))
		sum += uint64(num_letters(i))
	}
	println(sum)
}
