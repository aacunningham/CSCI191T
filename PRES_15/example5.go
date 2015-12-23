package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	person1 := &person{"Aaron", 23}
	person2 := &person{"You", 40}
	fmt.Println(person1.name, " ", person1.age)
	person2.age = 30
	fmt.Println(person2.age)
}
