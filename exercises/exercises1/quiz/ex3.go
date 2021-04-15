package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)            // & is the address_of operator. If it were not included, a copy of the value stored at "name" would be passed, and changes would not be reflected in "name" declared on line 9
	name = strings.Title(name) // string.Title returns a new string with all 1st chars of each word capitalized
	fmt.Printf("Hello %s\n", name)
}
