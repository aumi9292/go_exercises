package main

import "fmt"

//Package level variables, declared variables of types that are given zero values for that type.

// var x int
// var y string
// var z bool

var x int = 42
var y string = "james bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%d %q %t", x, y, z)
	fmt.Println(s)
}
