package main

import (
	"fmt"
	"strconv"
)

/* a list of small primes for fraction reducing */
var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}

type Fraction struct {
	num, dem int
}

func (f *Fraction) Float() float64 {
	return float64(f.num) / float64(f.dem)
}
func (f *Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.num, f.dem)
}

func (f *Fraction) Reduce() *Fraction {
	n, d := f.num, f.dem
	r := &Fraction{n, d}
	for _, p := range primes {
		if p > n || p > d {
			return r
		}
		if n%p == 0 && d%p == 0 {
			r.num = n / p
			r.dem = d / p
			return r.Reduce()
		}
	}
	return r
}

func Product(f1, f2 *Fraction) *Fraction {
	return &Fraction{f1.num * f2.num, f1.dem * f2.dem}
}

/* returns a fraction created by cancelling the second digit
 * of the parents numerator and the first digit of the parents denominator,
 * or nil
 */
func cancel_digit(f *Fraction) *Fraction {
	n, d := strconv.Itoa(f.num), strconv.Itoa(f.dem)
	if n[1] == d[0] {
		return &Fraction{int(n[0]) - 48, int(d[1]) - 48}
	}
	return nil
}

func main() {
	fmt.Print("Digit Cancelling Fractions.\n")
	fractions := []*Fraction{}
	for i := 11; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			f := &Fraction{i, j}
			d := cancel_digit(f)
			if d != nil && d.Float() == f.Float() {
				r := d.Reduce()
				fmt.Printf("%s == %s (%s)\n", f, d, r)
				fractions = append(fractions, r)
			}
		}
	}
	p := fractions[0]
	for i := 1; i < len(fractions); i++ {
		p = Product(p, fractions[i])
	}
	fmt.Printf("Denominator Product: %s (reduced: %s)\n", p, p.Reduce())
}
