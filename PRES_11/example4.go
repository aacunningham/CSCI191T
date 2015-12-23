package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter a number")
	fmt.Scan(&number)
	number += 5
	fmt.Println("Your number plus 5: ", number)
}
