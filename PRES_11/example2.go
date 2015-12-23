package main

import (
	"fmt"
)

const (
	t = iota
	u
	v
)

const (
	w = iota
	x
)

func main() {
	fmt.Println(t) // 0
	fmt.Println(u) // 1
	fmt.Println(v) // 2
	fmt.Println(w) // 0
	fmt.Println(x) // 1
}
