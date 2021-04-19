package main

import "fmt"

// pointers are useful to 1) pass large data structures by a reference rather than passing the large data structure each time and 2) alter passed data structures

// go is always pass by value. Pointers/addresses can be passed, which is how go implements pass by reference/mutating methods

func foo(x *int) {
	fmt.Println(*x)
	*x = 43
	fmt.Println(*x)
}

func main() {
	num := 10
	foo(&num)
	fmt.Println(num) //
}
