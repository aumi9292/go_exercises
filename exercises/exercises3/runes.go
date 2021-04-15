package main

import "fmt"

func main() {
	for i := 65; i <= 90; i += 1 {
		for j := 1; j <= 3; j += 1 {
			fmt.Printf("%d.\n\t%#U\t%v\n", j, i, string(i))
		}
	}
}
