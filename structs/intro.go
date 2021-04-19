package main

import "fmt"

type name struct {
	first string
	last  string
}

func main() {

	me := name{
		first: "Austin",
		last:  "Miller",
	}

	celia := name{
		first: "Celia",
		last:  "Horowitz",
	}

	fmt.Println(me.first, celia.last)

}
