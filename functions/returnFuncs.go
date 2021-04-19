package main

import "fmt"

// syntax func (r receiver) identifier(params) return {code}
// return could be func(), func(string) string, etc.

func makeGreeter(greeting string) func(string) {
	return func(name string) {
		fmt.Printf("%s, %s\n", greeting, name)
	}
}

func main() {

	casualGreeting := makeGreeter("Yo")
	casualGreeting("Austin")

	fancyGreeting := makeGreeter("Why hello there fellow world citizen")
	fancyGreeting("Celia")

}
