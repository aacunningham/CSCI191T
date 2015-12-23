package main

import (
	"fmt"
)

func dog(y int) (x int) {
	x = y * 7
	return
}

func main() {
	dog := dog(26)
	fmt.Println("Aaron is", dog, "in dog years and is old")
}
