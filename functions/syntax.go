package main

import (
	"fmt"
	"strings"
)

func foo() {
	fmt.Println("I'm from the foo function!")
}

func bar(toPrint string) {
	fmt.Println(toPrint)
}

func add(num1, num2 int) int {
	return num1 + num2 // no implicit return statements
}

func getInitials(name string) (n1, n2 string) {
	lastInitial := strings.Split(name, " ")[1][0]
	return string(name[0]), string(lastInitial)
}

func main() {
	foo()
	bar("I'm an argument passed to bar!")
	fmt.Println(add(10, 100))
	fmt.Println(getInitials("Austin Miller"))
}
