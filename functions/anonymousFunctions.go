package main

import "fmt"

// func greet() {
// 	fmt.Println("Hi!")
// }

func main() {
	func() {
		fmt.Println("Hi!")
	}()
}
