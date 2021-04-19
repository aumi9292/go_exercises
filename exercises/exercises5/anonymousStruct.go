package main

import (
	"fmt"
	"strings"
)

func main() {
	denver := struct {
		name    string
		streets []string
		parks   map[string][]string
	}{
		name:    "Denver",
		streets: []string{"Colfax", "Speer", "Colorado", "Tennyson"},
		parks:   map[string][]string{"East": {"City Park", "Lindsley Park"}, "Central": {"Governor's Park", "Cheesman Park", "Wahsington Park"}},
	}
	fmt.Println(denver.parks["East"][1])

	for neighborhood, parks := range denver.parks {
		fmt.Printf("The %s neighborhood has the following parks: %s\n", neighborhood, strings.Join(parks, "\n"))
	}
}
