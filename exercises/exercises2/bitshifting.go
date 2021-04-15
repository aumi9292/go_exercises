package main

import "fmt"

func main() {
	a := 10
	b := a << 1
	fmt.Printf("%d\t%b\t%#X\n", a, a, a)
	fmt.Printf("%d\t%b\t%#X\n", b, b, b)

}
