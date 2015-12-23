package main

import (
	"fmt"
)

func main() {
	int2intmap := map[int]int{
		1: 2,
		2: 3,
		3: 4,
		4: 1,
	}
	int2intmap[5] = 1
	int2intmap[4] = 5
	delete(int2intmap, 1)
	for key, value := range int2intmap {
		fmt.Println(key, value)
	}
	fmt.Println(len(int2intmap))
	if _, exists := int2intmap[9]; exists {
		fmt.Println("It exists!")
	} else {
		fmt.Println("It don't exist.")
	}
}
