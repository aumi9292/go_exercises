package main

import "fmt"

func multiply(nums ...int) int {
	product := 1
	for _, num := range nums {
		product *= num
	}
	return product
}

func operateOnEvens(operation func(nums ...int) int, numList ...int) int {
	evens := []int{}
	for _, num := range numList {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return operation(evens...)
}

func main() {
	list := []int{2, 4, 1, 3, 5, 7}
	fmt.Println(operateOnEvens(multiply, list...))
}
