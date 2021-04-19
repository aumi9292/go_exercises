package main

import "fmt"

func sayHi() {
	fmt.Println("Hello")
}

func sayBye() {
	fmt.Println("Goodbye")
}

func main() {
	sayHi()
	defer sayBye()
	fmt.Println("conversing conversing conversing")
}
