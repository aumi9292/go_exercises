package main

import "fmt"

const a = 10
const b string = "hello"

const (
	d     = 42
	e int = 100
)

func main() {
	fmt.Println(a, b, d, e)
}
