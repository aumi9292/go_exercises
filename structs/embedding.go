package main

import "fmt"

type person struct {
	first string
	last  string
}

type hero struct {
	person
	alias  string
	powers []string
}

func main() {
	me := person{
		first: "Austin",
		last:  "Miller",
	}

	mr_incredible := hero{
		person: person{
			first: "Bob",
			last:  "Credible",
		},
		alias:  "Mr Incredible",
		powers: []string{"strength", "fitness"},
	}

	fmt.Println(me, mr_incredible.last) // inner types are "promoted" to outer types, so you don't need to type mr_incredible.person.first, you can just type mr_incredible.first. You can specify, which is useful when there are naming collisions, i.e., when the outer type has a field with the same name as an inner type
}
