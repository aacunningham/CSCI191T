package main

import (
	"fmt"
)

func printperson(x string, y int) {
	fmt.Println(x, "is", y, "years old.")
}

func main() {
	printperson("Aaron", 23)
}
