package main

import "fmt"

func main() {
	sayer := func(word string) {
		fmt.Println(word)
	}
	sayer("hello there")
}
