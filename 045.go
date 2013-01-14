package main

import (
	"fmt"
)

var penmap = map[int64]bool{}
var hexmap = map[int64]bool{}

func main() {
	for i := int64(2); ; i++ {
		t := (i * (i + 1)) / 2
		p := (i * ((3 * i) - 1)) / 2
		h := (i * ((2 * i) - 1))

		penmap[p] = true
		hexmap[h] = true

		if penmap[t] && hexmap[t] && t > 40755 {
			fmt.Println(t)
			return
		}

	}
}
