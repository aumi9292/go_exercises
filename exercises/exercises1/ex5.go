package main

import "fmt"

type money int

var x money
var y int

func main() {
	fmt.Printf("%v %T\n", x, x)
	x = 42
	y = int(x)
	fmt.Printf("%d\n", x)

	fmt.Printf("%v %T\n", y, y)
}
