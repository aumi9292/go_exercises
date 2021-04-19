package main

import "fmt"

func sayHi() {
	defer func() {
		fmt.Println("short intro after immediate greeting")
	}()
	fmt.Println("Hi!")
}

func sayBye() {
	fmt.Println("BYE!")
}

func main() {
	sayHi()
	defer sayBye()
	fmt.Println("intermediate actions")
}
