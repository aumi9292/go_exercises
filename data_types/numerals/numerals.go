package main

import "fmt"

func main() {
	string := "H"
	fmt.Println(string)
	byteSlice := []byte(string)
	fmt.Println(byteSlice)
	num := byteSlice[0]
	fmt.Println(num)
	fmt.Printf("%T\n", num)
	fmt.Printf("%b\n", num)
	fmt.Printf("%x\n", num)
	fmt.Println([]byte("hello there you  "))
}
