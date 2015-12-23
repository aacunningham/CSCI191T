package main

import (
	"fmt"
)

func main() {
	var num_arr [2]int
	for i := 0; i < 2; i++ {
		fmt.Scan(&num_arr[i])
	}
	fmt.Println(num_arr[0] % num_arr[1])
}
