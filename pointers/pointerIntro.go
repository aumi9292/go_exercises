package main

import "fmt"

func main() {
	num := 10
	num2 := 5
	repeatNum := 10
	// & is the address of operator, "& what is the address?"
	fmt.Println(num, &num, &repeatNum)   // variables that point to the same number point to different locations in memory
	fmt.Printf("%T: %v\n", &num2, &num2) // The type of a &var is *type, "pointer to the type of" *int and int are two distinct types

	addressOfNum := &num
	fmt.Println(*addressOfNum) // de-referencing the variable, this operator is different syntax than *<SomeType>. This operator is used to get the value stored at the address tied to the variable

	// fmt.Println(*num) // raises error, you can only use * with pointer variables

	*addressOfNum = 50
	fmt.Println(num) // num and &num work with the same location in memory, so changes to one reflect the other

}
