package main

import "fmt"

func main() {
	a := 1
	b := 10

	for a < b {
		a *= 2
		fmt.Println(a)
	}

	fmt.Println(a)

}
