package main

import "fmt"

func main() {
	fizzBuzz(100)
}

func fizzBuzz(upperLimit int) {
	for i := 0; i <= upperLimit; i += 1 {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
