package main

import "fmt"

func main() {
	nums := [5]int{5, 7, 9, 1, 3} // this is an array, not a slice

	for idx, num := range nums {
		fmt.Println(idx, num)
	}
	fmt.Printf("Type: %T\n", nums)
}
