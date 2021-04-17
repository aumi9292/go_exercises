package main

import "fmt"

// maps are unordered, k-v stores that have fast lookup
func main() {
	characters := map[string]string{ // map[string]int is the type
		"Darth Vader": "Star Wars",
		"Aragorn":     "Lord of the Rings",
		"Dumbledore":  "Harry Potter",
	}

	fmt.Println(characters)
	fmt.Println(characters["Darth Vader"])
	fmt.Println(characters["Austin"]) // if the key does not exist, the return value of bracket access is the zero value for the type

	austin, ok := characters["Austin"]

	fmt.Println(austin, ok)

	// common if logic to execute some code if a map has a particular value, comma ok idiom
	if harryPotter, ok := characters["Dumbledore"]; ok {
		fmt.Printf("%s is a character in %s", "Dumbledore", harryPotter)
	}

}
