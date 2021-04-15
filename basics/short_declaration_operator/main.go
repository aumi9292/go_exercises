package main

import (
	"fmt"
)

// When within a double-quoted string passed to fmt.Printf(), the %T format specifier returns the type of the next argument.

func main() {
	x := 42 // The walrus operator is shorthand syntax for declaring and assigning a variable. The type for a number is int
	var y int32 = 10
	fmt.Printf("%T %T\n", x, y)
}
