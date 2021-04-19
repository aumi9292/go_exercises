package main

import "fmt"

// func (r receiver) identifier(params, types) (return/s) {code}

// having variadic paramaters brings them in as a slice of the specified type
func sum(operands ...int) int {
	fmt.Printf("%T, %v\n", operands, operands)
	sum := 0
	for _, num := range operands {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
}
