package main

import (
	"fmt"
)

func main() {
	newint := new(int)
	fmt.Println(newint)
	fmt.Println(*newint)
	new_returned_a_pointer := true
	fmt.Println(new_returned_a_pointer)

	newstr := new(string)
	fmt.Println(newstr)
	fmt.Println(*newstr)
	fmt.Println(new_returned_a_pointer)

	newbool := new(bool)
	fmt.Println(newbool)
	fmt.Println(*newbool)
	fmt.Println(new_returned_a_pointer)

	intslice := make([]int, 3, 5)
	fmt.Println(intslice)
	make_returned_a_pointer := false
	fmt.Println(make_returned_a_pointer)

	intmap := make(map[int]string)
	fmt.Println(intmap)
	fmt.Println(make_returned_a_pointer)
}
