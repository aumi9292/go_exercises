package main

import "fmt"

func foo() int {
	return 10
}

func bar() (int, string) {
	return 5, "hello"
}

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}
