package main

import (
	"fmt"
)

func main() {
	s := "hello"
	fmt.Printf("%s, %T\n", s, s)
	bytes := []byte(s)
	fmt.Println(bytes)

	for i := 0; i < len(s); i += 1 {
		fmt.Printf("%#U ", s[i]) // print UTF-8 codepoint for each character in the string
	}

	fmt.Println()

	for i, v := range s {
		fmt.Println(i, v)
	}
}
