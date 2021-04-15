package main

import "fmt"

const (
	currentYear   = 2021 - iota
	lastYear      = 2021 - iota
	twoYearsAgo   = 2021 - iota
	threeYearsAgo = 2021 - iota
	fourYearsAgo  = 2021 - iota
)

func main() {

	fmt.Println(currentYear, lastYear, twoYearsAgo, threeYearsAgo, fourYearsAgo)

}
