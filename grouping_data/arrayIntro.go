package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println(arr)
	arr[3] = 400
	fmt.Println(arr)
	fmt.Println(len(arr))

}
