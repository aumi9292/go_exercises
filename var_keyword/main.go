package main

import "fmt"

// x := 10 // raises "non-declaration statement outside function body" error
var x = 10 // does not raise an error
var z int  // declares a new variable of type int and assigns the zero value of the declared type

func main() {
	// best practice for declaring variables is to use the short declaration operator in the smallest scope possible
	fmt.Println(x)
	fmt.Println(z)

	var zeroInt int
	var zeroFloat float32
	var zeroStr string
	var zeroBool bool
	var zeroFunc func()

	fmt.Println(zeroInt, zeroFloat, zeroStr, zeroBool, zeroFunc)
	// output will be 0, 0.0, "", false, nil
}
