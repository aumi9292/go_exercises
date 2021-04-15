package main

import "fmt"

var x = 2 //best practice to avoid package-level variables

const (
	_  = iota             // zeroith iota is thrown away, next iota is 1
	kb = 1 << (iota * 10) // each iota is multiplied by 10, which is used to shift the bit over an additional 10 zeros
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 3 // variable shadowing overwrites package scoped variable
	fmt.Printf("%d\t%b\n", x, x)

	y := x << 1 // shifting the bit moves both bits to the next places
	fmt.Printf("%d\t%b\n", y, y)

	// kb := 1024
	// mb := 1024 * kb
	// gb := 1024 * mb

	// fmt.Printf("%d\t%b\n", kb, kb)
	// fmt.Printf("%d\t%b\n", mb, mb)
	// fmt.Printf("%d\t%b\n", gb, gb)

	fmt.Printf("%d\t%b\n", kb, kb)
	fmt.Printf("%d\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
