package main

import "fmt"

func makeCounter() func() int {
	counter := 0
	return func() int {
		counter += 1
		return counter
	}
}

func main() {
	counter1 := makeCounter()
	counter2 := makeCounter()

	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter2())
	fmt.Println(counter2())

}
