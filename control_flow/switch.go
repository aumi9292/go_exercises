package main

import "fmt"

func main() {

	// All of these are expression switches, there are also switch types
	// switching on no value
	switch {
	case false:
		fmt.Println("shouldn't print")
	case 1 != 1:
		fmt.Println("shouldn't print either")
	case true:
		fmt.Println("should print!")
		// fallthrough // if "fallthrough" keyword is added, execution does fall through
	case true:
		fmt.Println("Will this print?") // this will not print, execution does not "fall through" as it does in JS
		// fallthrough
	case false:
		fmt.Println("This will get printed if all above statements that evaluate to true also include the 'fallthrough' keyword")
	}

	var x int

	switch {
	case false:
		x = 1
	case false:
		x = 2
	case true:
		x = 100
		// fallthrough
	default:
		x = 10
	}

	fmt.Println(x)

	// switch on a value
	name := "Austin"
	switch name {
	case "Bob":
		fmt.Println("Hi, I'm Bob")
	case "Jillian":
		fmt.Println("I'M JILLIAN!")
	// case true: // true cannot be compared to a string, this raises an error
	// 	fmt.Println("Will 'Austin' be evaluated to true?")
	case "Austin":
		fmt.Println("Hi I'm Austin")
	}

}
