package main

import "fmt"

func main() {
	nums := []int{4, 5, 1, 54, 60}
	fmt.Println(len(nums))
	fmt.Println(nums[3]) // access by index

	for index, value := range nums {
		fmt.Println(index, value)
	}
}
