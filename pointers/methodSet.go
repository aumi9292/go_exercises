package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) growOlder() {
	p.age += 1
}

func main() {
	austin := person{
		name: "Austin",
		age:  29,
	}

	austin.growOlder()
	fmt.Println(austin)

	austin.growOlder()
	fmt.Println(austin)
}
