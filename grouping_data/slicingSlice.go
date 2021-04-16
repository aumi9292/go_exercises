package main

import "fmt"

func main() {
	nums := []int{4, 100, 154, 21349283}
	fmt.Println(nums)

	for index, num := range nums {
		fmt.Printf("The value at index %d is %v\n", index, num)
	}

	fmt.Println(nums[1:]) // returns a new slice from position 1 to end
	//fmt.Println(nums[1:-1]) // cannot use negative index, will raise error
	fmt.Println(nums[1:3]) // second operand is exclusive ending index to return (up to but not including)
	//fmt.Println(nums[100:]) // raises a panic, not caught in compilation
	fmt.Println(nums) // slicing using : is a non-mutating action

	// use a for loop without the range syntax to iterate over a slice and print elements at each index
	for idx := 0; idx < len(nums); idx += 1 {
		fmt.Printf("The value at %d is %v\n", idx, nums[idx])
	}

	// using the : operator to access slices from 0 to end, 1 to end, 2 to end, etc.
	for idx := 0; idx < len(nums); idx += 1 {
		fmt.Println(nums[idx:])
	}

}
