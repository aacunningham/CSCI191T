package main

import (
	"fmt"
)

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	val, even := half(23)
	fmt.Println(val, even)
}
