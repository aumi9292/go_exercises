package main

import "fmt"

func main() {

	denver := struct {
		name       string
		population int
	}{
		name:       "Denver",
		population: 350000000,
	}

	fmt.Println(denver)

}
