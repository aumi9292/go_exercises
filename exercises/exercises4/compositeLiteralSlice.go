package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for idx, num := range nums {
		fmt.Println(idx, num)
	}

	fmt.Printf("%T", nums)
}
