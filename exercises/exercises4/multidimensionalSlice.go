package main

import "fmt"

func main() {
	//bond := []string{"James", "Bond", "Shaken not stirred"} // can be added directly as a slice within the 2d slice

	moneyPenny := []string{"Miss", "MoneyPenny", "Hello James"}

	characters := [][]string{{"James", "Bond", "Shaken not stirred"}, moneyPenny}
	fmt.Println(characters)

	for idx, character := range characters {
		for _, stat := range character {
			fmt.Println(idx, stat)
		}
	}
}
