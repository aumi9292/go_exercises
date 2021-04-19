package main

import (
	"fmt"
	"strings"
)

func main() {
	func(word string) {
		fmt.Println(strings.ToUpper(word))
	}("hello!")
}
