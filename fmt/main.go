package main

import (
	"fmt"
	"os"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) // type of the next value
	fmt.Printf("%x\n", y) // the next value formatted as hexadecimal
	fmt.Printf("%b\n", y) // the next value formatted as binary

	var float float32 = 7.0 / 2.0
	fmt.Printf("%.2f\n", float) // the precision is specified as an integer after the ., before the f format verb
	str := fmt.Sprintln("Hello world")
	fmt.Print(str)

	fmt.Fprint(os.Stdout, "Hello there file world") //os.Stdout implements a "writer" interface, so Fprint can be used to write to the console as well as a file
}
