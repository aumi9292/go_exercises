package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) speak() {
	fmt.Printf("Hi, my name is %s %s and I am %d years old\n", p.first, p.last, p.age)
}

func main() {
	austin := Person{
		first: "Austin",
		last:  "Miller",
		age:   29,
	}

	austin.speak()
}
