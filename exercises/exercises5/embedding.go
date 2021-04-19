package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {

	whiteRascal := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}

	gasGuzzler := truck{
		vehicle: vehicle{
			doors: 2,
			color: "grey",
		},
		fourWheel: true,
	}

	fmt.Println(whiteRascal, gasGuzzler)
	fmt.Println(whiteRascal.luxury)
	fmt.Println(gasGuzzler.vehicle.color)
	fmt.Println(gasGuzzler.color) // color is a promoted field name

}
