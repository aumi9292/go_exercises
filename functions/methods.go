package main

import "fmt"

type Person struct {
	first string
	last  string
}

type Professional struct {
	Person
	title  string
	salary int
}

// function syntax: func (r receiver) identifier(params) (returns) {code}

func (person Person) introduce() {
	fmt.Printf("Hello, my name is %s\n", person.first)
}

// this does not actually alter the receiver (I think */& oeprators are needed)
func (prof Professional) getPromoted() {
	prof.salary += 50000
}

func main() {
	austin := Professional{
		Person: Person{
			first: "Austin",
			last:  "Miller",
		},
		title:  "software engineer",
		salary: 100000,
	}

	fmt.Println(austin)
	austin.introduce()
	austin.getPromoted()
	fmt.Println(austin)
}
