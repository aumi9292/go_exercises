package main

import "fmt"

func main() {
	money := true
	power := true

	switch {
	case money && !power:
		fmt.Println("I'm rich, like a figurehead")
	case money && power:
		fmt.Println("I have money and power, but maybe too much ...")
	default:
		fmt.Println("Can't we all just get along?!?")
	}
}
