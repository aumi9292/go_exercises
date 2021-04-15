package main

import "fmt"

func main() {
	var smallInt int
	var largeInt int

	fmt.Println("Please enter a small number")
	fmt.Scan(&smallInt)

	fmt.Println("Please enter a larger number")
	fmt.Scan(&largeInt)

	fmt.Printf("%T", largeInt) // Scan is able to coerce/convert input to int types declared on 6 and 7. This means no conversion is necessary for line 17

	quotient := largeInt % smallInt
	fmt.Printf("%d\n", quotient)
}
