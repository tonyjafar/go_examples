package main

import (
	"fmt"
)

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type getArea interface {
	Shape() float64
}

func main() {

	mySquare := square{sideLength: 33}
	myTriangle := triangle{height: 34.5, base: 78}

	printArea(mySquare)
	printArea(myTriangle)

}

func (s square) Shape() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) Shape() float64 {
	return 0.5 * t.base * t.height
}

func printArea(g getArea) {
	fmt.Println(g.Shape())
}
