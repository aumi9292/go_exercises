package main

import "fmt"

var x int // x is has a package level scope

func main() {
	fmt.Println(x)
	printPackageVar()
	{
		x := 5 // code block in a code block, variables have block scope
		fmt.Println(x)
	}

	newId := idGenerator()

	fmt.Println(x)
	fmt.Println(newId())
	fmt.Println(newId())
	fmt.Println(newId())
	fmt.Println(newId())
}

func printPackageVar() {
	fmt.Println(x)
}

func idGenerator() func() int {
	id := 0
	return func() int {
		id += 1
		return id
	}
}
