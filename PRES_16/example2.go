package main

import (
	"fmt"
)

func maxint(x ...int) int {
	var maxval int = 0
	for _, value := range x {
		if value > maxval {
			maxval = value
		}
	}
	return maxval
}

func main() {
	val := maxint(234, 5, 45, 23, 45, 1, 4, 6, 34, 61)
	fmt.Println(val)
}
