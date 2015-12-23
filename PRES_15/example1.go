package main

import (
	"fmt"
)

func main() {
	intslice := []int{1, 2, 3, 4, 5}
	fmt.Println(intslice)
	strslice := []string{"hello", " world"}
	fmt.Println(strslice)

	intslice_2 := []int{6, 7, 8, 9, 0}
	new_slice := append(intslice, intslice_2...)
	fmt.Println(new_slice)
}
