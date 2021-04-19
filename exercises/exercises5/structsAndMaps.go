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

	celia := person{
		firstName:               "Celia",
		lastName:                "Horowitz",
		favoriteIceCreamFlavors: []string{"caramel coffee", "gelato"},
	}

	people := map[string]person{austin.lastName: austin, celia.lastName: celia}
	fmt.Println(people)

	fmt.Println(people["Miller"].firstName)
	fmt.Println(people["Horowitz"].lastName)

	for k, v := range people {
		fmt.Println(k, v)
	}

	favorites := append(people["Miller"].favoriteIceCreamFlavors, people["Horowitz"].favoriteIceCreamFlavors...)

	for idx, flavor := range favorites {
		fmt.Println(idx, flavor)
	}
}
