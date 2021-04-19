package main

import "fmt"

type Being interface {
	walk()
}

type Human struct {
	name string
}

type Lion struct {
	name string
	size string
}

func (h Human) walk() {
	fmt.Println("I'm walking like a human walks")
}

func (l Lion) walk() {
	fmt.Println("Stealthily and pridefully walking in the desert")
}

func growOlder(b Being) {
	switch b.(type) {
	case Human:
		fmt.Printf("I'm a %T and I'm getting one year older\n", b)
	case Lion:
		fmt.Printf("I'm a %T and I'm getting 7 years older\n", b)
	}

}

func main() {
	austin := Human{
		name: "Austin",
	}

	aslan := Lion{
		name: "Aslan",
		size: "large",
	}

	austin.walk()
	aslan.walk()
	growOlder(austin)
	growOlder(aslan)
}
