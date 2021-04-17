package main

import "fmt"

func main() {
	me := []string{"Austin", "Miller"}
	celia := []string{"Celia", "Horowitz"}
	zack := []string{"Zack", "Baltich"}
	names := [][]string{me, celia, zack}

	fmt.Println(names)
}
