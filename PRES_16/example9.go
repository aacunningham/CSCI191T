package main

import (
	"fmt"
)

func fib(x int) int {
	if x <= 1 {
		return 1
	} else {
		return fib(x-1) + fib(x-2)
	}
}

func main() {
	var i int = fib(6)
	fmt.Println(i)
}
