package main

import "fmt"

const (
	a   = 42
	b   = 42.78
	spy = "James Bond"
)

func main() {
	fmt.Printf("%T\n%T\n%T\n", a, b, spy)
}
