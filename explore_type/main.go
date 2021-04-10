package main

import "fmt"

var y int = 42
var z = "shaken not stirred" // type is inferred

func main() {
	fmt.Printf("%T, %T\n", y, z)
	// z = 10 would raise an exception, can't use 10 as type string in assignment
}
