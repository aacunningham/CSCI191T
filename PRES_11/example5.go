package main

import (
	"fmt"
)

func takes_pointer_and_adds_five(x *int) {
	*x += 5
}

func main() {
	var num int = 20
	fmt.Println(num)
	takes_pointer_and_adds_five(&num)
	fmt.Println(num)
}
