package main

import "fmt"

var x bool // compiler assigns the boolean zero value false

func main() {
	x = true
	fmt.Println(x)

	a := 7
	b := 42

	fmt.Println(a == b, a > b, a <= b, a != b)
}
