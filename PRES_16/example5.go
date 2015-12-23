package main

import (
	"fmt"
)

func dogold(y int) (int, bool) {
	return y * 7, y > 25
}

func main() {
	dog, old := dogold(26)
	if old {
		fmt.Println("Aaron is", dog, "in dog years and is old")
	} else {
		fmt.Println("Aaron is", dog, "in dog years and is not old")
	}
}
