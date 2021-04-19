package main

import "fmt"

func makeGreeter(greeting string) func(name string) {
	return func(name string) {
		fmt.Printf("%s, %s\n", greeting, name)
	}
}

func main() {
	casualGreeter := makeGreeter("Hey what's up")
	elevatedGreeter := makeGreeter("Dear person it's lovely to see you")

	casualGreeter("Celia")
	casualGreeter("Austin")
	elevatedGreeter("Austin")
}
