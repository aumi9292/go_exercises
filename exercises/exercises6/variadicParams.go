package main

import "fmt"

func foo(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func bar(toAdd []int) int {
	sum := 0
	for _, num := range toAdd {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(foo(1, 2, 3, 4, 5))
	fmt.Println(bar([]int{1, 2, 3, 4, 5}))
}
