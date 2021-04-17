package main

import "fmt"

func main() {
	capitals := map[string]string{"Germany": "Berlin", "Mongolia": "Uulan Batar", "Qatar": "Doha", "Nigeria": "Lagos", "China": "Hangzhou"}

	fmt.Println(capitals)

	// comma ok idiom to show that China is a key in capitals
	hangzhou, ok := capitals["China"]
	fmt.Println(hangzhou, ok)

	// ranging over each country, it's not necessary to include the value if you're not using it
	for country := range capitals {
		if capitals[country] == "Hangzhou" {
			delete(capitals, country) // use delete built-in function and pass <map identifier>, <key>
		}
	}

	delete(capitals, "North Korea") // will not raise an error

	fmt.Println(capitals)

	hangzhou, ok = capitals["China"]
	fmt.Println(hangzhou, ok)
}
