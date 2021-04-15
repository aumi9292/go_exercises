package main

import "fmt"

func main() {
	a := 10 == 10
	b := 100 > 10
	c := 100 >= 10
	d := 10 < 100
	e := 10 <= 100
	f := 10 != 100

	fmt.Println(a, b, c, d, e, f)
}
