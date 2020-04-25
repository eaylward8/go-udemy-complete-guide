package main

import (
	"fmt"
	"math"
)

type triangle struct {
	b float64
	h float64
}

type square struct {
	l float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{b: 2, h: 3}
	sq := square{l: 5}

	printArea(t)
	printArea(sq)
}

// printArea of something that implements shape interface
func printArea(s shape) {
	fmt.Println(s.getArea())
}

// triangle getArea
func (t triangle) getArea() float64 {
	return (t.b * t.h) * 0.5
}

// square getArea
func (s square) getArea() float64 {
	return math.Pow(s.l, 2)
}
