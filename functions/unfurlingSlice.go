package main

import "fmt"

// Variadic parameters must be the final parameter

func sum(nums ...int) int {
	sum := 0
	fmt.Println(cap(nums))
	fmt.Println(len(nums))
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sum(nums...))
	fmt.Println(sum()) // variadic ... parameter means 0 or more arguments
}
