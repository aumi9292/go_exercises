package main

import "fmt"

func main() {
	favSport := "skiing"
	switch favSport {
	case "biking":
		fmt.Println("Time to go for a ride!")
	case "running":
		fmt.Println("Watch out for those knees!")
	case "skiing":
		fmt.Println("Watch out for those trees!")
	}
}
