package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	austin := person{
		firstName:               "Austin",
		lastName:                "Miller",
		favoriteIceCreamFlavors: []string{"marionberry cheescake", "coffee", "pb brownie"},
	}
	fmt.Println(austin.firstName)
	fmt.Println(austin.lastName)

	for idx, flavor := range austin.favoriteIceCreamFlavors {
		fmt.Println(idx, flavor)
	}
}
