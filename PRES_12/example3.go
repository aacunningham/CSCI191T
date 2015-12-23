package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		var fb string = ""
		if i%3 == 0 {
			fb += "Fizz"
		}
		if i%5 == 0 {
			fb += "Buzz"
		}
		if fb != "" {
			fmt.Println(fb)
		} else {
			fmt.Println(i)
		}
	}
}
