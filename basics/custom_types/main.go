package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	var a int
	a = 42
	b = 43
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b) // b is of type main.hotdog

	// a = b //raises cannout use a (type hotdog) as type int in assignment

	a = int(b) // int() converts some expression into the type int
	fmt.Printf("%T\n", a)

	// Below, trying these conversions raises exceptions
	// num := 10

	// num = string(num)

	// fmt.Println("%T\n", num)
}
