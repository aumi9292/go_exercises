package main

import "fmt"

func main() {

	// for loop with init; condition; post;
	for i := 0; i <= 10; i += 1 {
		fmt.Println(i)
	}

	// nested for loop
	for integer := 0.0; integer <= 10.0; integer += 1.0 {
		for decimal := 0.1; decimal <= 1; decimal += 0.1 {
			fmt.Println(integer + decimal)
		}
	}

	// for loop with condition only
	a := 1
	b := 10

	for a < b {
		a *= 2
		fmt.Println(a)
	}

	// for loop with no condition, just break in an if statement within the for loop
	breakForLoopInt := 1

	for {
		if breakForLoopInt >= 10 {
			break
		}
		breakForLoopInt += 1
		fmt.Println(breakForLoopInt)
	}

	// continue with for loop

	continueInt := 1

	for ; continueInt < 10; continueInt += 1 { // for clause w/ condition & post, no init
		if continueInt%2 != 0 {
			continue
		}
		fmt.Println(continueInt)
	}

}
