package main

import "fmt"

func main() {
	for i := 1; i <= 10000; i += 1 {
		fmt.Printf("%d\t%b\t%#X\t%U", i, i, i, i)
	}
}
