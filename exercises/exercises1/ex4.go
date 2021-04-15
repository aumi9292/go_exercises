package main

import "fmt"

type money int

var x money

func main() {
	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Printf("%d\n", x)
}
