package main

import (
	"fmt"
)

func variad_func(x int, y ...int) int {
	var total int = x
	for _, num := range y {
		total += num
	}
	return total
}

func main() {
	num_list := []int{4, 3, 2, 1}
	t := variad_func(5, num_list...)
	fmt.Println(t)
}
