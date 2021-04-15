package main

import "fmt"

// challenge: Print out integers 33 - 122, with their corresponding ascii characters to the right

func main() {
	for asci_int := 33; asci_int <= 122; asci_int += 1 {
		fmt.Printf("%d\t%s\n", asci_int, string(asci_int))
	}
}
