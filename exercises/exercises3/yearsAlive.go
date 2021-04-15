package main

import "fmt"

func main() {
	year := 1991
	for {
		fmt.Println(year)
		year += 1
		if year > 2021 {
			break
		}
	}
}
