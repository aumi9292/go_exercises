package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) calculateArea() float64 {
	return s.side * s.side
}

func (c circle) calculateArea() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	calculateArea() float64
}

func info(s shape) {
	fmt.Println(s.calculateArea())
}

func main() {
	horsePen := square{
		side: 10,
	}

	racetrack := circle{
		radius: 20,
	}

	info(horsePen)
	info(racetrack)

}
