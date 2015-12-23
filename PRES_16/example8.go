package main

import (
	"fmt"
)

func mapfunc(callback func(int), x ...int) {
	for _, val := range x {
		callback(val)
	}
}

func main() {
	mapfunc(func(x int) { fmt.Println(x) }, 1, 2, 3, 4, 5, 6)
}
