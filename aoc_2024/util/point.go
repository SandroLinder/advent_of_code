package util

import "fmt"

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}

func (p Point) String() string {
	return fmt.Sprintf("Point: (%d, %d)", p.Y, p.X)
}
