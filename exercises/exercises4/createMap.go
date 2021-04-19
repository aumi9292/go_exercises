package main

import (
	"fmt"
	"strings"
)

func main() {
	favoriteThings := map[string][]string{
		"Austin": {"biking", "hiking", "walking"},
		"Celia":  {"hiking", "food", "kids"},
		"Zack":   {"drumming", "art", "laughing"},
	}

	favoriteThings["Kaylie"] = []string{"reading", "photography", "hiking"}

	delete(favoriteThings, "Zack")

	for person, activities := range favoriteThings {
		fmt.Printf("%s's favorite activities are %s\n", person, strings.Join(activities, ", "))
	}
}
