package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	cp.Y = 2
	fmt.Println(cp)

}
