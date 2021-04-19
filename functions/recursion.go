package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func loopFactorial(num int) int {
	factorial := 1
	for ; num > 1; num-- {
		factorial *= num
	}
	return factorial
}

func main() {
	fmt.Println(factorial(4))
	fmt.Println(loopFactorial(4))

}
