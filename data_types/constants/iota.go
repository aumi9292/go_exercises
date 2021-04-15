package main

import "fmt"

const (
	a = iota * 1
	b = iota * 2
	c = iota * 3
)

const (
	A = iota * 2
	B = iota * 4
	C = iota * 6
)

func main() {
	fmt.Println("first iota group", a, b, c, "\nsecond iota group", A, B, C)
}
