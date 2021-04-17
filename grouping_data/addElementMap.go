package main

import "fmt"

func main() {
	populations := map[string]int{"USA": 350, "Canada": 37, "China": 1400, "Japan": 126}

	populations["Nigeria"] = 201

	fmt.Println(populations)

	// ranging over the map, accessing each k-v pair
	for country, population := range populations {
		fmt.Printf("The country %s has a population of %d\n", country, population*1000000)
	}
}
