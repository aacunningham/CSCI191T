package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "35"
	v, _ := strconv.Atoi(str)
	str = strconv.Itoa(v + 10)
	fmt.Println(str)
}
