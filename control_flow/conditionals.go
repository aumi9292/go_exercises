package main

import "fmt"

func main() {
	if true {
		fmt.Println("It's true!")
	}

	if false {
		fmt.Println("This won't print!")
	}

	if !false {
		fmt.Println(`This uses the "not" operator (and a raw string literal)!`)
	}

	if !(2 != 2) {
		fmt.Println("Too much logic here")
	}

	// initialization statement, does not evaluate truthiness here, evaluates the next statement, separated by the semicolon

	if veryLocalVar := 10; veryLocalVar == 10 {
		fmt.Println("Can you access this very local variable outside of this conditional?")
	}

	// fmt.Println(veryLocalVar) // raises an error, variable initialized in an if-clause have scope restricted to the conditional block

	// if, else if, else
	money := 0

	if money > 100 {
		fmt.Println("I'm rich!")
	} else if money > 50 {
		fmt.Println("We're doing alright")
	} else if money >= 0 {
		fmt.Println("Well ...")
	} else {
		fmt.Println("I owe some people money")
	}
}
