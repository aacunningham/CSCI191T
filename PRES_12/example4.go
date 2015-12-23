package main

import (
	"fmt"
)

func main() {
	var num_list []int
	for i := 1; i <= 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			num_list = append(num_list, i)
		}
	}
	var total int = 0
	for _, num := range num_list {
		total += num
		fmt.Println(num)
	}

	fmt.Println(total)
}
