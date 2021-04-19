package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	(*p).name = "Bob"
}

func main() {
	austin := person{
		name: "Austin",
		age:  29,
	}

	changeMe(&austin)

	fmt.Println(austin)

	sentence := "hello my name is austin"
	fmt.Println(strings.Index(sentence, " "))
}
