package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter a number")
	fmt.Scanln(&number)
	number *= 8
	fmt.Println("Your number times 8: ", number)
}
