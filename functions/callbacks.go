package main

import (
	"fmt"
	"strings"
)

func say(sentence string) string {
	return sentence
}

func shout(cb func(sentence string) string, sentence string) string {
	return strings.ToUpper(cb(sentence))
}

func sum(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func oddOperation(operation func(nums ...int) int, numList ...int) int {
	var odds []int
	for _, val := range numList {
		if val%2 != 0 {
			odds = append(odds, val)
		}
	}
	return operation(odds...)
}

func main() {
	fmt.Println(shout(say, "hello my name is austin"))
	fmt.Println(oddOperation(sum, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
