package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 100}
	fmt.Println(nums)

	appendedNums := append(nums, 3000000000, -1, 223) //append() is non-mutating
	fmt.Println(nums, "\n", appendedNums)

	doubleyAppended := append(nums, appendedNums...) // ... after the argument unfurls the composite type, like JS spread syntax, but affixed rather than prefixed
	fmt.Println(doubleyAppended)
}
