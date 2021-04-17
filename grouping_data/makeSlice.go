package main

import "fmt"

func main() {
	herFoods := []string{"banana", "apple", "berry", "chorizo"}
	foods := make([]string, 4, 5)
	fmt.Println(foods)
	fmt.Println(len(foods))
	fmt.Println(cap(foods))

	for idx, food := range herFoods {
		foods[idx] = food
	}

	foods = append(foods, "fish")
	foods = append(foods, "muffin") // doubles the current capacity of the array from 5 to 10
	fmt.Println(cap(foods))         // logs 10
	fmt.Println(foods)

}
